// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"reflect"
)

type BoostTypeEnum byte
type BoostTypeValues struct {
	TURBO        BoostTypeEnum
	SUPERCHARGER BoostTypeEnum
	NITROUS      BoostTypeEnum
	KERS         BoostTypeEnum
	NullValue    BoostTypeEnum
}

var BoostType = BoostTypeValues{84, 83, 78, 75, 0}

func (b BoostTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(b)); err != nil {
		return err
	}
	return nil
}

func (b *BoostTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(b)); err != nil {
		return err
	}
	return nil
}

func (b BoostTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(BoostType)
	for idx := 0; idx < value.NumField(); idx++ {
		if b == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on BoostType, unknown enumeration value %d", b)
}

func (*BoostTypeEnum) EncodedLength() int64 {
	return 1
}

func (*BoostTypeEnum) TURBOSinceVersion() uint16 {
	return 0
}

func (b *BoostTypeEnum) TURBOInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TURBOSinceVersion()
}

func (*BoostTypeEnum) TURBODeprecated() uint16 {
	return 0
}

func (*BoostTypeEnum) SUPERCHARGERSinceVersion() uint16 {
	return 0
}

func (b *BoostTypeEnum) SUPERCHARGERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SUPERCHARGERSinceVersion()
}

func (*BoostTypeEnum) SUPERCHARGERDeprecated() uint16 {
	return 0
}

func (*BoostTypeEnum) NITROUSSinceVersion() uint16 {
	return 0
}

func (b *BoostTypeEnum) NITROUSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.NITROUSSinceVersion()
}

func (*BoostTypeEnum) NITROUSDeprecated() uint16 {
	return 0
}

func (*BoostTypeEnum) KERSSinceVersion() uint16 {
	return 0
}

func (b *BoostTypeEnum) KERSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.KERSSinceVersion()
}

func (*BoostTypeEnum) KERSDeprecated() uint16 {
	return 0
}
