package types

import (
	"fmt"
	"io"

	"github.com/KyberNetwork/binance-sbe/sbe"
)

// id=50
type WsWrapperSbe struct {
	SbeSchemaIdVersionDeprecated sbe.BooleanTypeEnum
	Status                       uint16
	RateLimits                   []RateLimit
	Id                           string
	Result                       interface{}
}

type RateLimit struct {
	RateLimitType uint8
	Interval      uint8
	IntervalNum   uint8
	RateLimit     int64
	Current       int64
}

func (s *RateLimit) Decode(_m *sbe.SbeGoMarshaller, _r io.Reader) error {
	if err := _m.ReadUint8(_r, &s.RateLimitType); err != nil {
		return err
	}
	if err := _m.ReadUint8(_r, &s.Interval); err != nil {
		return err
	}
	if err := _m.ReadUint8(_r, &s.IntervalNum); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.RateLimit); err != nil {
		return err
	}
	if err := _m.ReadInt64(_r, &s.Current); err != nil {
		return err
	}
	return nil
}

func (s *WsWrapperSbe) Decode(_m *sbe.SbeGoMarshaller, _r io.Reader) error {
	if err := s.SbeSchemaIdVersionDeprecated.Decode(_m, _r, 2); err != nil {
		return err
	}

	if err := _m.ReadUint16(_r, &s.Status); err != nil {
		return err
	}

	// Rate Limit
	var RateLimitBlockLength uint16
	if err := _m.ReadUint16(_r, &RateLimitBlockLength); err != nil {
		return err
	}
	var RateLimitNumInGroup uint16
	if err := _m.ReadUint16(_r, &RateLimitNumInGroup); err != nil {
		return err
	}
	if cap(s.RateLimits) < int(RateLimitNumInGroup) {
		s.RateLimits = make([]RateLimit, RateLimitNumInGroup)
	}
	s.RateLimits = s.RateLimits[:RateLimitNumInGroup]

	for i := range s.RateLimits {
		if err := s.RateLimits[i].Decode(_m, _r); err != nil {
			return err
		}
	}

	// request ID
	var requestIdLength uint8
	if err := _m.ReadUint8(_r, &requestIdLength); err != nil {
		return err
	}
	idData := make([]uint8, requestIdLength)
	if err := _m.ReadBytes(_r, idData); err != nil {
		return err
	}
	s.Id = string(idData)

	// Decode data
	var dataLength uint32
	if err := _m.ReadUint32(_r, &dataLength); err != nil {
		return err
	}

	var msgHeader sbe.MessageHeader
	if err := msgHeader.Decode(_m, _r, 2); err != nil {
		return err
	}

	switch msgHeader.TemplateId {
	case 211:
		var bookTicker BookTickerSbe
		if err := bookTicker.Decode(_m, _r); err != nil {
			return err
		}
		s.Result = bookTicker
	default:
		return fmt.Errorf("unsupported template id: %d", msgHeader.TemplateId)
	}

	return nil
}
