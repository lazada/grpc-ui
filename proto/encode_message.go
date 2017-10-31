/*
Encoders for map types
*/

package proto

import (
)
//
//func encodeMap(value interface{}, n int, wt int, typeName string, typeInfo map[string]*reflection.TypeInfo, fieldPath []int) (resultBuf []byte, err error) {
//
//	rv := reflect.ValueOf(value)
//	for _, k := range rv.MapKeys() {
//		v := rv.MapIndex(k)
//
//		_value := map[string]interface{}{
//			"key":   k.Interface(),
//			"value": v.Interface(),
//		}
//
//		var msgBuf []byte
//		msgBuf, err = Encode(typeInfo, typeName, _value, append(fieldPath, n))
//		if err != nil {
//			return
//		}
//
//		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
//		resultBuf = append(resultBuf, EncodeBytes(msgBuf)...)
//	}
//
//	return
//}
//
//func encodeMessage(value interface{}, n int, wt int, typeInfo map[string]*reflection.TypeInfo, typeName string, fieldPath []int, repeated bool) (resultBuf []byte, err error) {
//
//	if repeated {
//
//		_value, ok := value.([]interface{})
//		if !ok {
//			err = fmt.Errorf("encodeMessage: Conversion error %#v to []interface{}", value)
//			return
//		}
//
//		for _, v := range _value {
//			var msgBuf []byte
//			msgBuf, err = encodeMessage(v, n, wt, typeInfo, typeName, fieldPath, false)
//			resultBuf = append(resultBuf, msgBuf...)
//		}
//
//	} else {
//
//		_value, ok := value.(map[string]interface{})
//		if !ok {
//			err = fmt.Errorf("encodeMessage: Conversion error %#v to map[string]interface{}", value)
//			return
//		}
//
//		var msgBuf []byte
//		msgBuf, err = Encode(typeInfo, typeName, _value, append(fieldPath, n))
//		if err != nil {
//			return
//		}
//
//		resultBuf = append(resultBuf, EncodeTag(n, wt)...)
//		resultBuf = append(resultBuf, EncodeBytes(msgBuf)...)
//	}
//
//	return
//}
