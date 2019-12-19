package utils

import (
	"time"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HumanDate(date time.Time) string {
	return date.Format("02.01.2006 15:04:05")
}

func Float64ToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func Int64ToString(input_num int64) string {
	return strconv.FormatInt(input_num,36)
}

func IntToString(input_num int64) string {
	return strconv.FormatInt(input_num,10)
}

func StringToInt(str string) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}
func StringToFloat(str string) float64 {
	res, err := strconv.ParseFloat(str, 64);
	if err != nil {
		panic(err)
	}
	return res
}