// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"math"
)

type Booster struct {
	BoostType  BoostTypeEnum
	HorsePower uint8
}

func (b *Booster) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := b.BoostType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, b.HorsePower); err != nil {
		return err
	}
	return nil
}

func (b *Booster) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if b.BoostTypeInActingVersion(actingVersion) {
		if err := b.BoostType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !b.HorsePowerInActingVersion(actingVersion) {
		b.HorsePower = b.HorsePowerNullValue()
	} else {
		if err := _m.ReadUint8(_r, &b.HorsePower); err != nil {
			return err
		}
	}
	return nil
}

func (b *Booster) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.HorsePowerInActingVersion(actingVersion) {
		if b.HorsePower < b.HorsePowerMinValue() || b.HorsePower > b.HorsePowerMaxValue() {
			return fmt.Errorf("Range check failed on b.HorsePower (%v < %v > %v)", b.HorsePowerMinValue(), b.HorsePower, b.HorsePowerMaxValue())
		}
	}
	return nil
}

func BoosterInit(b *Booster) {
	return
}

func (*Booster) EncodedLength() int64 {
	return 2
}

func (*Booster) BoostTypeSinceVersion() uint16 {
	return 0
}

func (b *Booster) BoostTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.BoostTypeSinceVersion()
}

func (*Booster) BoostTypeDeprecated() uint16 {
	return 0
}

func (*Booster) HorsePowerMinValue() uint8 {
	return 0
}

func (*Booster) HorsePowerMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Booster) HorsePowerNullValue() uint8 {
	return math.MaxUint8
}

func (*Booster) HorsePowerSinceVersion() uint16 {
	return 0
}

func (b *Booster) HorsePowerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.HorsePowerSinceVersion()
}

func (*Booster) HorsePowerDeprecated() uint16 {
	return 0
}
