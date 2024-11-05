// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"io"
)

type OptionalExtras [8]bool
type OptionalExtrasChoiceValue uint8
type OptionalExtrasChoiceValues struct {
	SunRoof       OptionalExtrasChoiceValue
	SportsPack    OptionalExtrasChoiceValue
	CruiseControl OptionalExtrasChoiceValue
}

var OptionalExtrasChoice = OptionalExtrasChoiceValues{0, 1, 2}

func (o *OptionalExtras) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range o {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (o *OptionalExtras) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		o[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (OptionalExtras) EncodedLength() int64 {
	return 1
}

func (*OptionalExtras) SunRoofSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) SunRoofInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SunRoofSinceVersion()
}

func (*OptionalExtras) SunRoofDeprecated() uint16 {
	return 0
}

func (*OptionalExtras) SportsPackSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) SportsPackInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SportsPackSinceVersion()
}

func (*OptionalExtras) SportsPackDeprecated() uint16 {
	return 0
}

func (*OptionalExtras) CruiseControlSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) CruiseControlInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CruiseControlSinceVersion()
}

func (*OptionalExtras) CruiseControlDeprecated() uint16 {
	return 0
}
