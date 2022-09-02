package timeUtil

import "time"

const (
	EarliestTime = "1970-01-01 08:00:00"
	DateFormat   = "2006-01-02"
	MonthFormat  = "200601"
	TimeFormat   = "2006-01-02 15:04:05"
	ZeroTimeStr  = "0000-00-00 00:00:00"
)

// Format 解析时间为字符串
func Format(t time.Time) string {
	if t.IsZero() {
		return ZeroTimeStr
	}
	return t.Format(TimeFormat)
}

// Parse 解析字符串时间
func Parse(value string) (time.Time, error) {
	if value == ZeroTimeStr {
		return time.Time{}, nil
	}
	return time.Parse(TimeFormat, value)
}

// Now 查询当前时间的字符串格式：2006-01-02 15:04:05
func Now() string {
	return time.Now().Format(TimeFormat)
}

// Date 查询当前日期的字符串格式：2006-01-02
func Date() string {
	return time.Now().Format(DateFormat)
}
