// Generated SBE (Simple Binary Encoding) message codec

package types

import (
	"io"

	"github.com/KyberNetwork/binance-sbe/sbe"
)

type ServerTimeSbe struct {
	ServerTime uint64
}

func (s *ServerTimeSbe) Encode(_m *sbe.SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if err := _m.WriteUint64(_w, s.ServerTime); err != nil {
		return err
	}

	return nil
}

func (s *ServerTimeSbe) Decode(_m *sbe.SbeGoMarshaller, _r io.Reader) error {
	if err := _m.ReadUint64(_r, &s.ServerTime); err != nil {
		return err
	}

	return nil
}
