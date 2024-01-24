package ngssp

import (
	"strconv"
	"strings"
)

type StringInt int

func (si *StringInt) UnmarshalText(text []byte) error {
	s64 := StringInt64(0)
	err := s64.UnmarshalText(text)
	if err != nil {
		return err
	}
	*si = StringInt(s64)
	return nil
}

type StringInt64 int64

func (si64 *StringInt64) UnmarshalText(text []byte) error {
	str := strings.TrimSpace(string(text))
	if len(str) == 0 {
		return nil
	}

	s, _ := strconv.ParseInt(string(text), 10, 64)
	// if err != nil {
	// 	return err
	// }
	*si64 = StringInt64(s)
	return nil
}

type StringFloat64 float64

func (sf64 *StringFloat64) UnmarshalText(text []byte) error {
	str := strings.TrimSpace(string(text))
	if len(str) == 0 {
		return nil
	}

	s, _ := strconv.ParseFloat(string(text), 64)
	// if err != nil {
	// 	return err
	// }
	*sf64 = StringFloat64(s)
	return nil
}
