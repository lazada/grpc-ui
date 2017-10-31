/*
Decoders for byte types
*/

package proto

import (
	"encoding/base64"
)

func decodeBytes(buf []byte) (x interface{}, i int, err error) {
	// Encode raw bytes as base64 string
	var v []byte
	v, i = DecodeBytes(buf)
	x = base64.StdEncoding.EncodeToString([]byte(v))
	return
}

func decodeString(buf []byte) (x interface{}, i int, err error) {
	var v []byte
	v, i = DecodeBytes(buf)
	x = string(v)
	return
}
