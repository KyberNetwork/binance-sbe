package types

import (
	"io"

	"github.com/KyberNetwork/binance-sbe/sbe"
)

// id=211
type BookTickerSbe struct {
	PriceExponent int8
	QtyExponent   int8
	BidPrice      int64
	BidQty        int64
	AskPrice      int64
	AskQty        int64
	Symbol        string
}

func (s *BookTickerSbe) Decode(_m *sbe.SbeGoMarshaller, _r io.Reader) error {
	if err := _m.ReadInt8(_r, &s.PriceExponent); err != nil {
		return err
	}
	if err := _m.ReadInt8(_r, &s.QtyExponent); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.BidPrice); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.BidQty); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.AskPrice); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.AskQty); err != nil {
		return err
	}

	// request ID
	var symbolLength uint8
	if err := _m.ReadUint8(_r, &symbolLength); err != nil {
		return err
	}
	idData := make([]uint8, symbolLength)
	if err := _m.ReadBytes(_r, idData); err != nil {
		return err
	}
	s.Symbol = string(idData)

	return nil
}
