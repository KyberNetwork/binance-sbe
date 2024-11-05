package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	httpSbe "github.com/KyberNetwork/binance-sbe/http-sbe"
	wsSbe "github.com/KyberNetwork/binance-sbe/ws-sbe"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

const (
	ticker24h           = "https://www.binance.com/api/v3/ticker/24hr"
	historicalTradesETH = "https://www.binance.com/api/v3/historicalTrades?symbol=ETHUSDT&limit=1000"
	klineETH            = "https://www.binance.com/api/v3/klines?symbol=ETHUSDT&limit=1000&interval=1h"
	depthBTC            = "https://www.binance.com/api/v3/depth?symbol=BTCUSDT&limit=5000"
	serverTime          = "https://www.binance.com/api/v3/time"
	exchangeInfo        = "https://www.binance.com/api/v3/exchangeInfo"
)

func main() {
	SetupLogger()
	TestWsSbe()
	TestSbeVsJson()
}

func SetupLogger() *zap.SugaredLogger {
	pConf := zap.NewProductionEncoderConfig()
	pConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(pConf)
	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	l := zap.New(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level), zap.AddCaller())
	zap.ReplaceGlobals(l)
	return zap.S()
}

func TestWsSbe() {
	wsClient, err := wsSbe.NewClientWs("", "")
	if err != nil {
		panic(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	wsReq := wsSbe.WsApiRequest{
		Id:     id.String(),
		Method: "ticker.book",
		Params: map[string]interface{}{
			"symbol": "BTCUSDT",
		},
	}

	rawData, err := json.Marshal(wsReq)
	if err != nil {
		panic(err)
	}

	zap.S().Infow("Ready to write", "data", wsReq)

	waiter, err := wsClient.Write(wsReq.Id, rawData)
	if err != nil {
		panic(err)
	}

	zap.S().Infow("Write successfully")

	res, err := waiter.Wait(context.Background())
	if err != nil {
		panic(err)
	}

	zap.S().Infow("Receive response successfully", "data", res)
}

func TestSbeVsJson() {
	sbeC := httpSbe.NewClient(&http.Client{})
	jsonC := NewClient(&http.Client{})

	sbeC.DoRequest(context.Background(), http.MethodGet, serverTime, nil)
	jsonC.DoRequest(context.Background(), http.MethodGet, serverTime, nil)

	for _, endpoint := range []string{
		ticker24h,
		historicalTradesETH,
		klineETH,
		depthBTC,
		serverTime,
		exchangeInfo,
	} {
		zap.S().Infow("Test sbe vs json", "endpoint", endpoint)
		time.Sleep(200 * time.Millisecond)
		now := time.Now()

		eg := errgroup.Group{}
		eg.Go(func() error {
			data, err := jsonC.DoRequest(context.Background(), http.MethodGet, endpoint, nil)
			zap.S().Infow(
				"Received json response",
				"latency", time.Since(now).Milliseconds(),
				"data_size", len(data)/1024,
			)
			return err
		})

		eg.Go(func() error {
			data, err := sbeC.DoRequest(context.Background(), http.MethodGet, endpoint, nil)
			zap.S().Infow(
				"Received sbe response",
				"latency", time.Since(now).Milliseconds(),
				"data_size", len(data)/1024,
			)
			return err
		})

		if err := eg.Wait(); err != nil {
			panic(err)
		}
	}
}
