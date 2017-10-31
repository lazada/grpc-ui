/*
Decoders for number types
*/

package proto

import (
	"math"
)

func decodeFloat(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []float64
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed32(buf[i:])
			i += shift
			vv = append(vv, float64(math.Float32frombits(v)))
			if i >= max {
				break
			}
		}
		x = []float64(vv)
	} else {
		var v uint32
		v, i = DecodeFixed32(buf)
		x = float64(math.Float32frombits(v))
	}
	return
}

func decodeDouble(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []float64
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed64(buf[i:])
			i += shift
			vv = append(vv, float64(math.Float64frombits(v)))
			if i >= max {
				break
			}
		}
		x = []float64(vv)
	} else {
		var v uint64
		v, i = DecodeFixed64(buf)
		x = float64(math.Float64frombits(v))
	}
	return
}

func decodeBool(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []bool
		var max int = i + int(len)
		for {
			v, shift := DecodeVarint(buf[i:])
			i += shift
			if int(v) > 0 {
				vv = append(vv, true)
			} else {
				vv = append(vv, false)
			}
			if i >= max {
				break
			}
		}
		x = []bool(vv)
	} else {
		var v uint64
		v, i = DecodeVarint(buf)
		if v > 0 {
			x = true
		} else {
			x = false
		}
	}
	return
}

func decodeSInt32(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []int
		var max int = i + int(len)
		for {
			v, shift := DecodeZigzag32(buf[i:])
			i += shift
			vv = append(vv, int(int32(v)))
			if i >= max {
				break
			}
		}
		x = []int(vv)
	} else {
		var v uint64
		v, i = DecodeZigzag32(buf)
		x = int(int32(v))
	}
	return
}

func decodeSInt64(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []int
		var max int = i + int(len)
		for {
			v, shift := DecodeZigzag64(buf[i:])
			i += shift
			vv = append(vv, int(v))
			if i >= max {
				break
			}
		}
		x = []int(vv)
	} else {
		var v uint64
		v, i = DecodeZigzag64(buf)
		x = int(v)
	}
	return
}

func decodeFixed32(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []uint32
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed32(buf[i:])
			i += shift
			vv = append(vv, uint32(int32(v)))
			if i >= max {
				break
			}
		}
		x = []uint32(vv)
	} else {
		var v uint32
		v, i = DecodeFixed32(buf)
		x = uint32(int32(v))
	}
	return
}

func decodeFixed64(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []uint64
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed64(buf[i:])
			i += shift
			vv = append(vv, uint64(v))
			if i >= max {
				break
			}
		}
		x = []uint64(vv)
	} else {
		var v uint64
		v, i = DecodeFixed64(buf)
		x = uint64(v)
	}
	return
}

func decodeSFixed32(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []int32
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed32(buf[i:])
			i += shift
			vv = append(vv, int32(v))
			if i >= max {
				break
			}
		}
		x = []int32(vv)
	} else {
		var v uint32
		v, i = DecodeFixed32(buf)
		x = int32(v)
	}
	return
}

func decodeSFixed64(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []int64
		var max int = i + int(len)
		for {
			v, shift := DecodeFixed64(buf[i:])
			i += shift
			vv = append(vv, int64(v))
			if i >= max {
				break
			}
		}
		x = []int64(vv)
	} else {
		var v uint64
		v, i = DecodeFixed64(buf)
		x = int64(v)
	}
	return
}

func decodeUInt(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []uint
		var max int = i + int(len)
		for {
			v, shift := DecodeVarint(buf[i:])
			i += shift
			vv = append(vv, uint(v))
			if i >= max {
				break
			}
		}
		x = []uint(vv)
	} else {
		var v uint64
		v, i = DecodeVarint(buf)
		x = uint(v)
	}
	return
}

func decodeInt(buf []byte, repeated bool) (x interface{}, i int, err error) {
	if repeated {
		len, shift := DecodeVarint(buf[i:])
		i += shift
		var vv []int
		var max int = i + int(len)
		for {
			v, shift := DecodeVarint(buf[i:])
			i += shift
			vv = append(vv, int(v))
			if i >= max {
				break
			}
		}
		x = []int(vv)
	} else {
		var v uint64
		v, i = DecodeVarint(buf)
		x = int(v)
	}
	return
}
