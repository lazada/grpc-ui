/*
Decoders for map types
*/

package proto

import (
	"fmt"
	"reflect"

	"github.com/lazada/grpc-ui/reflection"
)

func decodeMessage(buf []byte, n int, typeName string, typeInfo map[string]*reflection.TypeInfo, fieldPath []int) (x interface{}, i int, err error) {
	var msgBuf []byte
	msgBuf, i = DecodeBytes(buf)
	x, err = Decode(typeInfo, typeName, msgBuf, append(fieldPath, n))
	return
}

func appendValue(result *map[string]interface{}, key string, value interface{}, mapField bool) (err error) {

	msgValue := (*result)[key]

	if mapField {

		_value, ok := value.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("Conversion error %#v to map[string]interface{}", value)
			return
		}

		var _msgValue map[string]interface{}

		if msgValue != nil {
			var ok bool
			_msgValue, ok = msgValue.(map[string]interface{})
			if !ok {
				err = fmt.Errorf("Conversion error %#v to map[string]interface{}", msgValue)
				return
			}
		} else {
			_msgValue = make(map[string]interface{})
		}

		var _key string
		if kind := reflect.TypeOf(_value["key"]).Kind(); kind >= reflect.Int && kind <= reflect.Float64 {
			_key = fmt.Sprintf("%d", _value["key"])
		} else {
			var ok bool
			_key, ok = _value["key"].(string)
			if !ok {
				err = fmt.Errorf("Conversion error %#v to string", _value["key"])
				return
			}
		}

		_msgValue[_key] = _value["value"]

		(*result)[key] = _msgValue

	} else {

		var _msgValue []interface{}

		if msgValue != nil {
			var ok bool
			_msgValue, ok = msgValue.([]interface{})
			if !ok {
				err = fmt.Errorf("Conversion error %#v to []map[string]interface{}", value)
				return
			}
		} else {
			_msgValue = make([]interface{}, 0)
		}

		_msgValue = append(_msgValue, value)

		(*result)[key] = _msgValue

	}

	return
}
