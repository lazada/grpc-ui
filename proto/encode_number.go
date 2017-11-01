/*
Encoders for number types
*/

package proto

import (
	"fmt"
	"math"
	"strconv"
	"errors"
)

func _interfaceFloat64ToFloat64(value interface{}) (res float64, err error) {
	// try to convert interface with float64 inside
	_value, ok := value.(float64)
	if !ok {
		err = fmt.Errorf("Conversion error (%T) %#v to float64", value, value)
		// try to convert interface with string inside
		var vv string
		vv, ok = value.(string)
		if !ok {
			//err = fmt.Errorf("Conversion error (%T) %#v to string", value, value)
			return
		}
		// convert string to float
		return strconv.ParseFloat(vv, 64)
	}
	res = float64(_value)
	return
}

func _interfaceFloat64ArrayToIntArray(value interface{}) (res []int, err error) {
	_value, ok := value.([]interface{})
	if !ok {
		err = fmt.Errorf("Conversion error (%T) %#v to []interface{}", value, value)
		return
	}
	res = make([]int, len(_value))
	for i, v := range _value {
		var vv float64
		vv, err = _interfaceFloat64ToFloat64(v)
		if err != nil {
			return
		}
		res[i] = int(vv)
	}
	return
}

func _interfaceFloat64ArrayToFloat64Array(value interface{}) (res []float64, err error) {
	_value, ok := value.([]interface{})
	if !ok {
		err = fmt.Errorf("Conversion error (%T) %#v to []interface{}", value, value)
		return
	}
	res = make([]float64, len(_value))
	for i, v := range _value {
		var vv float64
		vv, err = _interfaceFloat64ToFloat64(v)
		if err != nil {
			return
		}
		res[i] = vv
	}
	return
}

func encodeFloat(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []float64
		_value, err = _interfaceFloat64ArrayToFloat64Array(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeFixed32((math.Float32bits(float32(v))))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeFixed32((math.Float32bits(float32(_value))))...)
	}
	return
}

func encodeDouble(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []float64
		_value, err = _interfaceFloat64ArrayToFloat64Array(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeFixed64((math.Float64bits(float64(v))))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeFixed64((math.Float64bits(float64(_value))))...)
	}
	return
}

func encodeBool(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		_value, ok := value.([]interface{})
		if !ok {
			err = fmt.Errorf("encodeBool: Conversion error (%T) %#v to []interface{}", value, value)
			return
		}
		var repBuf []byte
		for _, v := range _value {
			vv, ok := v.(string)
			if !ok {
				err = fmt.Errorf("encodeBool: Conversion error (%T) %#v to string", v, v)
			}
			if vv  == "true" {
				repBuf = append(repBuf, EncodeVarint(uint64(1))...)
			} else if vv == "false" {
				repBuf = append(repBuf, EncodeVarint(uint64(0))...)
			} else {
				err = fmt.Errorf("encodeBool: Conversion error (%v) %#v to bool", vv, v)
			}

		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		_value, ok := value.(string)
		if !ok {
			err = fmt.Errorf("encodeBool: Conversion error (%T) %#v to string", value, value)
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)

		if _value  == "true" {
			resultBuf = append(resultBuf, EncodeVarint(uint64(1))...)
		} else if _value == "false" {
			resultBuf = append(resultBuf, EncodeVarint(uint64(0))...)
		} else {
			err = fmt.Errorf("encodeBool: Conversion error (%v) to bool", _value)
		}
	}
	return
}

func encodeSInt32(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []int
		_value, err = _interfaceFloat64ArrayToIntArray(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeZigzag32(uint64(v))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeZigzag32(uint64(_value))...)
	}
	return
}

func encodeSInt64(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []int
		_value, err = _interfaceFloat64ArrayToIntArray(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeZigzag64(uint64(v))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeZigzag64(uint64(_value))...)
	}
	return
}

func encodeFixed32(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []int
		_value, err = _interfaceFloat64ArrayToIntArray(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeFixed32(uint32(v))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeFixed32(uint32(_value))...)
	}
	return
}

func encodeFixed64(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []int
		_value, err = _interfaceFloat64ArrayToIntArray(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeFixed64(uint64(v))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		var _value float64
		_value, err = _interfaceFloat64ToFloat64(value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeFixed64(uint64(_value))...)
	}
	return
}

func encodeInt(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []int
		_value, err = _interfaceFloat64ArrayToIntArray(value)
		if err != nil {
			return
		}
		var repBuf []byte
		for _, v := range _value {
			repBuf = append(repBuf, EncodeVarint(uint64(v))...)
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(repBuf)...)
	} else {
		valStr, ok := value.(string)
		if !ok {
			return nil, errors.New("Should be a string")
		}

		_value, err := strconv.ParseInt(valStr, 10, 64)
		if err != nil {
			return nil, err
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeVarint(uint64(_value))...)
	}
	return
}
