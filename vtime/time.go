package vtime

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type UtcTime time.Time

func (t UtcTime) MarshalJSON() ([]byte, error) {
	var ts = fmt.Sprintf("\"%s\"", time.Time(t).UTC().Format("2006-01-02 15:04:05"))
	return []byte(ts), nil
}

func AddDayOfNow(day int) time.Time {
	return GetNow().AddDate(0, 0, day)
}

func AddDayOfNowUtc(day int) time.Time {
	return GetNowUtc().AddDate(0, 0, day)
}

func AddDayUtc(t time.Time, day int) time.Time {
	return t.UTC().AddDate(0, 0, day)
}

func GetEpochTime() time.Time {
	return time.Unix(0, 0)
}

func GetEpochTimeUtc() time.Time {
	return GetEpochTime().UTC()
}

func GetNextMonthFirstDay(t time.Time) time.Time {
	nextMonth := t.AddDate(0, 1, -t.Day()+1)
	nextMonthFirstDay := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())

	return nextMonthFirstDay
}

func GetNextMonthFirstDayNow() time.Time {
	return GetNextMonthFirstDay(GetNow())
}

func GetNextMonthFirstDayNowUtc() time.Time {
	return GetNextMonthFirstDay(GetNowUtc())
}

func GetNextMonthFirstDayNowUtcUnix() int64 {
	return GetNextMonthFirstDay(GetNowUtc()).Unix()
}

func GetNow() time.Time {
	return time.Now()
}

func GetNowUtc() time.Time {
	return GetNow().UTC()
}

func GetNowUtcUnix() int64 {
	return GetNowUtc().Unix()
}

func GetSecondsFromNowToNextMonthFirstDay() int64 {
	return GetNextMonthFirstDayNowUtcUnix() - GetNowUtcUnix()
}

func GetSecondsFromNowToNextMonthFirstDayDuration() time.Duration {
	return time.Duration(GetSecondsFromNowToNextMonthFirstDay()) * time.Second
}

// New creates and returns a Time object with given parameter.
// The optional parameter can be type of: time.Time/*time.Time, string or integer.
func New(param ...any) *gtime.Time {
	return gtime.New(param...)
}

// Now creates and returns a time object of now.
func Now() *gtime.Time {
	return gtime.Now()
}

// StrToTime converts string to *Time object. It also supports timestamp string.
// The parameter `format` is unnecessary, which specifies the format for converting like "Y-m-d H:i:s".
// If `format` is given, it acts as same as function StrToTimeFormat.
// If `format` is not given, it converts string as a "standard" datetime string.
// Note that, it fails and returns error if there's no date string in `str`.
func StrToTime(str string, format ...string) (*gtime.Time, error) {
	return gtime.StrToTime(str, format...)
}

// Timestamp retrieves and returns the timestamp in seconds.
func Timestamp() int64 {
	return gtime.Timestamp()
}

// TimestampMilli retrieves and returns the timestamp in milliseconds.
func TimestampMilli() int64 {
	return gtime.TimestampMilli()
}
