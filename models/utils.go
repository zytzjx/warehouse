package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

/**
 * @note: 返回值为 年月日时分秒格式，加强易读性
 * @auth: t
 * @date 2022年3月29日15:34:03
**/
type LocalTime struct {
	time.Time
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t.Time)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(value interface{}) error {
	// if value, ok := v.(time.Time); ok {
	// 	*t = LocalTime(value)
	// 	return nil
	// }
	// return fmt.Errorf("can not convert %v to timestamp", v)
	switch v := value.(type) {
	case time.Time:
		*t = LocalTime{Time: v}
	case []uint8:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = LocalTime{Time: tt}
	default:
		return fmt.Errorf("cannot convert %v to CustomTime", value)
	}
	return nil
}

func (t *LocalTime) String() string {
	// 如果时间 null 那么我们需要把返回的值进行修改
	if t == nil || t.IsZero() {
		return ""
	}
	return time.Time(t.Time).Format("2006-01-02 15:04:05")
}

func (t *LocalTime) IsZero() bool {
	return time.Time(t.Time).IsZero()
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	*t = LocalTime{Time: t1}
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t.Time)
	// 如果时间值是空或者0值 返回为null 如果写空字符串会报错
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
