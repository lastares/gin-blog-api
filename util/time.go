package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

//  为 JSONTime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 为 JSONTime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
