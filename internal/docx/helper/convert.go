package helper

import "strconv"

func Str2Int(str string) (val int, err error) {
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return None, err
	}
	return int(parseInt), nil
}
