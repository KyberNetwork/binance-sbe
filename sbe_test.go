package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/KyberNetwork/binance-sbe/sbe"
	"github.com/KyberNetwork/binance-sbe/types"
	"github.com/stretchr/testify/assert"
)

var (
	sbeMsg  = []byte{3, 0, 50, 0, 2, 0, 0, 0, 0, 200, 0, 19, 0, 1, 0, 2, 1, 1, 112, 23, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 36, 56, 53, 57, 51, 48, 98, 55, 50, 45, 52, 101, 101, 102, 45, 52, 56, 57, 100, 45, 56, 51, 53, 50, 45, 48, 101, 55, 51, 57, 100, 97, 51, 98, 49, 55, 100, 50, 0, 0, 0, 34, 0, 211, 0, 2, 0, 0, 0, 248, 248, 192, 97, 241, 51, 68, 6, 0, 0, 176, 47, 248, 4, 0, 0, 0, 0, 0, 164, 0, 52, 68, 6, 0, 0, 56, 179, 79, 36, 0, 0, 0, 0, 7, 66, 84, 67, 85, 83, 68, 84}
	restMsg = []byte{123, 34, 115, 121, 109, 98, 111, 108, 34, 58, 34, 66, 84, 67, 85, 83, 68, 84, 34, 44, 34, 98, 105, 100, 80, 114, 105, 99, 101, 34, 58, 34, 54, 56, 56, 57, 56, 46, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 98, 105, 100, 81, 116, 121, 34, 58, 34, 52, 46, 53, 52, 49, 51, 57, 48, 48, 48, 34, 44, 34, 97, 115, 107, 80, 114, 105, 99, 101, 34, 58, 34, 54, 56, 56, 57, 56, 46, 48, 49, 48, 48, 48, 48, 48, 48, 34, 44, 34, 97, 115, 107, 81, 116, 121, 34, 58, 34, 54, 46, 51, 55, 57, 57, 49, 48, 48, 48, 34, 125}

	marshaller = sbe.NewSbeGoMarshaller()
)

type BookTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

func decodeSbe() (*types.WsWrapperSbe, error) {
	reader := bytes.NewBuffer(sbeMsg)

	var header sbe.MessageHeader
	if err := header.Decode(marshaller, reader, 2); err != nil {
		return nil, err
	}

	wsData := new(types.WsWrapperSbe)
	if err := wsData.Decode(marshaller, reader); err != nil {
		return nil, err
	}

	return wsData, nil
}

func decodeJson() (*BookTicker, error) {
	var res BookTicker

	if err := json.Unmarshal(restMsg, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func TestMarshaler(t *testing.T) {
	sbeData, err := decodeSbe()
	assert.NoError(t, err)
	t.Log("decodeSbe", "data", sbeData, "err", err)

	jsonData, err := decodeJson()
	assert.NoError(t, err)
	t.Log("decodeJson", "jsonData", jsonData, "err", err)
}

func BenchmarkSbe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeSbe()
	}
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeJson()
	}
}
