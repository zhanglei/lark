package utils

import "strconv"

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int32ToStr(val int32) string {
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

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func StrToInt64(str string) (val int64) {
	val, _ = strconv.ParseInt(str, 10, 64)
	return
}
