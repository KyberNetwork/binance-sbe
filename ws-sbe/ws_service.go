package bws

import (
	"net/http"

	"github.com/gorilla/websocket"

	"time"
)

var (
	BaseWsApiMainURL       = "wss://ws-api.binance.com:443/ws-api/v3?responseFormat=sbe&sbeSchemaId=2&sbeSchemaVersion=0"
	BaseWsMainURL          = "wss://stream.binance.com:9443/ws"
	BaseWsTestnetURL       = "wss://testnet.binance.vision/ws"
	BaseCombinedMainURL    = "wss://stream.binance.com:9443/stream?streams="
	BaseCombinedTestnetURL = "wss://testnet.binance.vision/stream?streams="

	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive                  = false
	ProxyUrl                            = ""
	UseTestnet                          = false
	WebsocketTimeoutReadWriteConnection = time.Second * 10
)

type params map[string]interface{}

type WsApiMethodType string

type WsApiRequest struct {
	Id     string          `json:"id"`
	Method WsApiMethodType `json:"method"`
	Params params          `json:"params"`
}

func getWsProxyUrl() *string {
	if ProxyUrl == "" {
		return nil
	}
	return &ProxyUrl
}

func SetWsProxyUrl(url string) {
	ProxyUrl = url
}

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	if UseTestnet {
		return BaseWsTestnetURL
	}
	return BaseWsMainURL
}

// getWsApiEndpoint return the base endpoint of the API WS according the UseTestnet flag
func getWsApiEndpoint() string {
	return BaseWsApiMainURL
}

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	if UseTestnet {
		return BaseCombinedTestnetURL
	}
	return BaseCombinedMainURL
}

// WsApiInitReadWriteConn create and serve connection
func WsApiInitReadWriteConn() (*websocket.Conn, error) {
	cfg := newWsConfig(getWsApiEndpoint())
	conn, err := WsGetReadWriteConnection(cfg)
	if err != nil {
		return nil, err
	}

	return conn, err
}

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}

var WsGetReadWriteConnection = func(cfg *WsConfig) (*websocket.Conn, error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	if WebsocketKeepalive {
		keepAlive(c, WebsocketTimeoutReadWriteConnection)
	}

	return c, nil
}
