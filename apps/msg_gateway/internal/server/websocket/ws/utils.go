package ws

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func intToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func int32ToStr(val int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(val)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func decode(msgCode int, buf []byte) interface{} {
	return nil
}

func encode(msgCode int, body []byte) (buf []byte) {
	var (
		buyCode []byte
		buffer  bytes.Buffer
	)
	buyCode = intToBytes(msgCode)
	buffer.Write(buyCode)
	if body != nil {
		buffer.Write(body)
	}
	buf = buffer.Bytes()
	return
}

func clientKey(uid int64, platform int32) (key string) {
	return int32ToStr(platform) + "-" + int64ToStr(uid)
}
