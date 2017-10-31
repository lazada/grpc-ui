/*
Encoders for byte types
*/

package proto

import (
	"encoding/base64"
	"fmt"
)

func _interfaceStringArrayToStringArray(value interface{}) (res []string, err error) {
	_value, ok := value.([]interface{})
	if !ok {
		err = fmt.Errorf("Conversion error (%T) %#v to []interface{}", value, value)
		return
	}
	res = make([]string, len(_value))
	for i, v := range _value {
		vv, ok := v.(string)
		if !ok {
			err = fmt.Errorf("Conversion error (%T) %#v to string", v, v)
			return
		}
		res[i] = vv
	}
	return
}

func encodeBytes(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	// Decode base64 to raw bytes
	if repeated {
		var _value []string
		_value, err = _interfaceStringArrayToStringArray(value)
		if err != nil {
			return
		}
		var vv []byte
		for _, v := range _value {
			vv, err = base64.StdEncoding.DecodeString(v)
			if err != nil {
				return
			}
			resultBuf = append(resultBuf, EncodeTag(n, wt)...)
			resultBuf = append(resultBuf, EncodeBytes(vv)...)
		}
	} else {
		_value, ok := value.(string)
		if !ok {
			err = fmt.Errorf("encodeBytes: Conversion error %#v to string", value)
			return
		}
		var vv []byte
		vv, err = base64.StdEncoding.DecodeString(_value)
		if err != nil {
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes(vv)...)
	}
	return
}

func encodeString(value interface{}, n int, wt int, repeated bool) (resultBuf []byte, err error) {
	if repeated {
		var _value []string
		_value, err = _interfaceStringArrayToStringArray(value)
		if err != nil {
			return
		}
		for _, v := range _value {
			resultBuf = append(resultBuf, EncodeTag(n, wt)...)
			resultBuf = append(resultBuf, EncodeBytes([]byte(v))...)
		}
	} else {
		_value, ok := value.(string)
		if !ok {
			err = fmt.Errorf("encodeString: Conversion error %#v to string", value)
			return
		}
		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
		resultBuf = append(resultBuf, EncodeBytes([]byte(_value))...)
	}
	return
}
