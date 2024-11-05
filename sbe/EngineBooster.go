// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"math"
)

type EngineBooster struct {
	BoostType  BoostTypeEnum
	HorsePower uint8
}

func (e *EngineBooster) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := e.BoostType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.HorsePower); err != nil {
		return err
	}
	return nil
}

func (e *EngineBooster) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if e.BoostTypeInActingVersion(actingVersion) {
		if err := e.BoostType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.HorsePowerInActingVersion(actingVersion) {
		e.HorsePower = e.HorsePowerNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.HorsePower); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngineBooster) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.HorsePowerInActingVersion(actingVersion) {
		if e.HorsePower < e.HorsePowerMinValue() || e.HorsePower > e.HorsePowerMaxValue() {
			return fmt.Errorf("Range check failed on e.HorsePower (%v < %v > %v)", e.HorsePowerMinValue(), e.HorsePower, e.HorsePowerMaxValue())
		}
	}
	return nil
}

func EngineBoosterInit(e *EngineBooster) {
	return
}

func (*EngineBooster) EncodedLength() int64 {
	return 2
}

func (*EngineBooster) BoostTypeSinceVersion() uint16 {
	return 0
}

func (e *EngineBooster) BoostTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BoostTypeSinceVersion()
}

func (*EngineBooster) BoostTypeDeprecated() uint16 {
	return 0
}

func (*EngineBooster) HorsePowerMinValue() uint8 {
	return 0
}

func (*EngineBooster) HorsePowerMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*EngineBooster) HorsePowerNullValue() uint8 {
	return math.MaxUint8
}

func (*EngineBooster) HorsePowerSinceVersion() uint16 {
	return 0
}

func (e *EngineBooster) HorsePowerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HorsePowerSinceVersion()
}

func (*EngineBooster) HorsePowerDeprecated() uint16 {
	return 0
}
