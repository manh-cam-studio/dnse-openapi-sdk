package websocket

import (
	"encoding/json"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

// FlexInt64 allows unmarshaling ints from both numbers and strings natively
type FlexInt64 int64

func (fi *FlexInt64) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		if s == "" {
			*fi = 0
			return nil
		}
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*fi = FlexInt64(val)
		return nil
	}

	var val int64
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}
	*fi = FlexInt64(val)
	return nil
}

func (fi *FlexInt64) UnmarshalMsgpack(b []byte) error {
	// MsgPack decoding natively handles strict types because the protocol dictates typings.
	// We'll decode using msgpack library natively. Because msgpack might encode it as 
	// integer directly, we can use decoder.
	var val int64
	err := msgpack.Unmarshal(b, &val)
	if err == nil {
		*fi = FlexInt64(val)
		return nil
	}
	// Fallback to string
	var s string
	if err := msgpack.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*fi = 0
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*fi = FlexInt64(v)
	return nil
}

// FlexFloat64 allows unmarshaling floats from both numbers and strings natively
type FlexFloat64 float64

func (ff *FlexFloat64) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		if s == "" {
			*ff = 0
			return nil
		}
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*ff = FlexFloat64(val)
		return nil
	}

	var val float64
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}
	*ff = FlexFloat64(val)
	return nil
}

func (ff *FlexFloat64) UnmarshalMsgpack(b []byte) error {
	var val float64
	err := msgpack.Unmarshal(b, &val)
	if err == nil {
		*ff = FlexFloat64(val)
		return nil
	}
	// Fallback to string
	var s string
	if err := msgpack.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*ff = 0
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*ff = FlexFloat64(v)
	return nil
}
