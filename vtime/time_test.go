package vtime_test

import (
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtime"
)

type TestUtcTime struct {
	CreateTime vtime.UtcTime
	UpdateTime time.Time
}

func init() {
	time.Local = time.UTC
}

func TestUtcTimeMarshalJSON(t *testing.T) {
	date := time.Date(2025, 3, 20, 20, 10, 0, 0, time.UTC)
	testUtcTime := &TestUtcTime{CreateTime: vtime.UtcTime(date), UpdateTime: date}
	testUtcTimeStr, err := vjson.Encode(testUtcTime)

	vtest.Nil(t, err)
	vtest.Equal(t, `{"CreateTime":"2025-03-20 20:10:00","UpdateTime":"2025-03-20T20:10:00Z"}`, testUtcTimeStr)
	vtest.Equal(t, "2025-03-20T20:10:00", time.Time(testUtcTime.CreateTime).Format("2006-01-02T15:04:05"))
	vtest.Equal(t, "2025-03-20 20:10:00", time.Time(testUtcTime.CreateTime).Format("2006-01-02 15:04:05"))
	vtest.Equal(t, "2025-03-20 20:10:00", time.Time(testUtcTime.CreateTime).UTC().Format("2006-01-02 15:04:05"))
	vtest.Equal(t, "2025-03-20 20:10:00", testUtcTime.UpdateTime.Format("2006-01-02 15:04:05"))
	vtest.Equal(t, "2025-03-20 20:10:00", testUtcTime.UpdateTime.UTC().Format("2006-01-02 15:04:05"))
}

func TestAddDayOfNow(t *testing.T) {
	vlog.Debug("TestAddDayOfNow", "vtime.AddDayOfNow(1)", vtime.AddDayOfNow(1))
	vlog.Debug("TestAddDayOfNow", "vtime.AddDayOfNowUtc(1)", vtime.AddDayOfNowUtc(1))
	vlog.Debug("TestAddDayOfNow", "vtime.AddDayOfNow(-1)", vtime.AddDayOfNow(-1))
	vlog.Debug("TestAddDayOfNow", "vtime.AddDayOfNowUtc(-1)", vtime.AddDayOfNowUtc(-1))
}

func TestAddDayOfNowUtc(t *testing.T) {
	vlog.Debug("TestAddDayOfNowUtc", "vtime.AddDayOfNowUtc(-1)", vtime.AddDayOfNowUtc(-1))
	vlog.Debug("TestAddDayOfNowUtc", "vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), 1)", vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), 1))
	vlog.Debug("TestAddDayOfNowUtc", "vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), 0)", vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), 0))
	vlog.Debug("TestAddDayOfNowUtc", "vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), -1)", vtime.AddDayUtc(vtime.AddDayOfNowUtc(-1), -1))
}

func TestGetEpochTime(t *testing.T) {
	vtest.Equal(t, "1970-01-01 00:00:00 +0000 UTC", vtime.GetEpochTime().UTC().String())
}

func TestGetEpochTimeUtc(t *testing.T) {
	vtest.Equal(t, "1970-01-01 00:00:00 +0000 UTC", vtime.GetEpochTimeUtc().String())
	vtest.Equal(t, "1970-01-01 00:00:00 +0000 UTC", vtime.GetEpochTime().UTC().String())
}

func TestGetNextMonthFirstDay(t *testing.T) {
	vlog.Debug("TestGetNextMonthFirstDay", "vtime.GetNextMonthFirstDay(vtime.GetNow())", vtime.GetNextMonthFirstDay(vtime.GetNow()))
	vlog.Debug("TestGetNextMonthFirstDay", "vtime.GetNextMonthFirstDay(vtime.GetNowUtc())", vtime.GetNextMonthFirstDay(vtime.GetNowUtc()))

	date := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)

	vtest.Equal(t, "2023-08-01 00:00:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 8, 1, 0, 0, 1, 0, time.UTC)

	vtest.Equal(t, "2023-08-01 00:00:01 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 8, 2, 0, 0, 0, 0, time.UTC)

	vtest.Equal(t, "2023-08-02 00:00:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 8, 29, 16, 10, 0, 0, time.UTC)

	vtest.Equal(t, "2023-08-29 16:10:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 8, 31, 23, 10, 0, 0, time.UTC)

	vtest.Equal(t, "2023-08-31 23:10:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 8, 31, 23, 59, 59, 0, time.UTC)

	vtest.Equal(t, "2023-08-31 23:59:59 +0000 UTC", date.String())
	vtest.Equal(t, "2023-09-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 2, 28, 23, 00, 0, 0, time.UTC)

	vtest.Equal(t, "2023-02-28 23:00:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-03-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 2, 28, 23, 59, 59, 0, time.UTC)

	vtest.Equal(t, "2023-02-28 23:59:59 +0000 UTC", date.String())
	vtest.Equal(t, "2023-03-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)

	vtest.Equal(t, "2023-03-01 00:00:00 +0000 UTC", date.String())
	vtest.Equal(t, "2023-04-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	vtest.Equal(t, "2023-12-31 23:59:59 +0000 UTC", date.String())
	vtest.Equal(t, "2024-01-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())

	date = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	vtest.Equal(t, "2024-01-01 00:00:00 +0000 UTC", date.String())
	vtest.Equal(t, "2024-02-01 00:00:00 +0000 UTC", vtime.GetNextMonthFirstDay(date).String())
}

func TestGetNextMonthFirstDayNow(t *testing.T) {
	vlog.Debug("TestGetNextMonthFirstDayNow", "vtime.GetNextMonthFirstDayNow()", vtime.GetNextMonthFirstDayNow())
}

func TestGetNextMonthFirstDayNowUtc(t *testing.T) {
	vlog.Debug("TestGetNextMonthFirstDayNowUtc", "vtime.GetNextMonthFirstDayNowUtc()", vtime.GetNextMonthFirstDayNowUtc())
	vlog.Debug("TestGetNextMonthFirstDayNowUtc", "vtime.GetNextMonthFirstDayNow()", vtime.GetNextMonthFirstDayNow())
}

func TestGetNextMonthFirstDayNowUtcUnix(t *testing.T) {
	vlog.Debug("TestGetNextMonthFirstDayNowUtcUnix", "vtime.GetNextMonthFirstDayNow()", vtime.GetNextMonthFirstDayNow())
	vlog.Debug("TestGetNextMonthFirstDayNowUtcUnix", "vtime.GetNextMonthFirstDayNowUtc()", vtime.GetNextMonthFirstDayNowUtc())
	vlog.Debug("TestGetNextMonthFirstDayNowUtcUnix", "vtime.GetNextMonthFirstDayNowUtc().Unix()", vtime.GetNextMonthFirstDayNowUtc().Unix())
	vlog.Debug("TestGetNextMonthFirstDayNowUtcUnix", "vtime.GetNextMonthFirstDayNowUtc().UnixMilli()", vtime.GetNextMonthFirstDayNowUtc().UnixMilli())
	vlog.Debug("TestGetNextMonthFirstDayNowUtcUnix", "vtime.GetNextMonthFirstDayNowUtcUnix()", vtime.GetNextMonthFirstDayNowUtcUnix())
}

func TestGetNowUtc(t *testing.T) {
	vlog.Debug("TestGetNowUtc", "vtime.GetNowUtc()", vtime.GetNowUtc())
	vlog.Debug("TestGetNowUtc", "vtime.GetNow()", vtime.GetNow())
}

func TestGetNowUtcUnix(t *testing.T) {
	vlog.Debug("TestGetNowUtcUnix", "vtime.GetNowUtcUnix()", vtime.GetNowUtcUnix())
}

func TestGetSecondsFromNowToNextMonthFirstDay(t *testing.T) {
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDay", "vtime.GetSecondsFromNowToNextMonthFirstDay()", vtime.GetSecondsFromNowToNextMonthFirstDay())
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDay", "vtime.GetNextMonthFirstDayNowUtcUnix()", vtime.GetNextMonthFirstDayNowUtcUnix())
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDay", "vtime.GetNowUtcUnix()", vtime.GetNowUtcUnix())
}

func TestGetSecondsFromNowToNextMonthFirstDayDuration(t *testing.T) {
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDayDuration", "vtime.GetSecondsFromNowToNextMonthFirstDayDuration()", vtime.GetSecondsFromNowToNextMonthFirstDayDuration())
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDayDuration", "vtime.GetSecondsFromNowToNextMonthFirstDay()", vtime.GetSecondsFromNowToNextMonthFirstDay())
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDayDuration", "vtime.GetNextMonthFirstDayNowUtcUnix()", vtime.GetNextMonthFirstDayNowUtcUnix())
	vlog.Debug("TestGetSecondsFromNowToNextMonthFirstDayDuration", "vtime.GetNowUtcUnix()", vtime.GetNowUtcUnix())

	vtest.Equal(t, vtime.GetSecondsFromNowToNextMonthFirstDay(), int64(vtime.GetSecondsFromNowToNextMonthFirstDayDuration().Seconds()))
}

func TestNew(t *testing.T) {
	vtest.Equal(t, "2025-03-20 15:30:20", vtime.New("2025-03-20 15:30:20").String())
	vtest.Equal(t, int64(1742484620), vtime.New("2025-03-20 15:30:20").Unix())
	vtest.Equal(t, "2025/03/20 15:30:20", vtime.New("2025-03-20 15:30:20").Format("Y/m/d H:i:s"))
	vtest.Equal(t, "2025-03-20 15:30:20", vtime.New("2025-03-20 15:30:20").UTC().String())
	vtest.Equal(t, "2025/03/20 15:30:20", vtime.New("2025-03-20 15:30:20").UTC().Format("Y/m/d H:i:s"))
	vtest.Equal(t, "2025/03/20 14:30:20", vtime.New("2025-03-20 15:30:20").Add(-time.Hour).Format("Y/m/d H:i:s"))
	vtest.Equal(t, "2025/05/20 15:30:20", vtime.New("2025-03-20 15:30:20").AddDate(0, 2, 0).Format("Y/m/d H:i:s"))

	vtest.Equal(t, time.Now().Format("2006-01-02 15:04:05"), vtime.New(vtime.Now()).String())
	vtest.Equal(t, time.Now().Format("2006-01-02 15:04:05"), vtime.New(time.Now()).String())

	vtest.Equal(t, "2025-03-20 15:30:20", vtime.New(1742484620).String())
	vtest.Equal(t, int64(1742484620), vtime.New(1742484620).Unix())
	vtest.Equal(t, "2025-03-20 15:30:20", vtime.New(1742484620).Format("Y-m-d H:i:s"))
	vtest.Equal(t, "2025/03/20 15:30:20", vtime.New(1742484620).Format("Y/m/d H:i:s"))
}

func TestNow(t *testing.T) {
	vtest.Equal(t, time.Now().Format("2006-01-02 15:04:05"), vtime.Now().String())
}

func TestStrToTime(t *testing.T) {
	tm, err := vtime.StrToTime("2025-03-20 15:30:20")

	vtest.Nil(t, err)
	vtest.Equal(t, "2025-03-20 15:30:20", tm.String())
	vtest.Equal(t, "2025-03-20 15:30:20", tm.FormatTo("Y/m/d H:i:s").String())
	vtest.Equal(t, "2025/03/20 15:30:20", tm.Format("Y/m/d H:i:s"))
	vtest.Equal(t, "2025-03-20 15:30:20", tm.UTC().String())
	vtest.Equal(t, int64(1742484620), tm.Unix())
	vtest.Equal(t, int64(1742484620000), tm.UnixMilli())

	tm, err = vtime.StrToTime("2025/03/20 15:30:20", "Y/m/d H:i:s")

	vtest.Nil(t, err)
	vtest.Equal(t, "2025-03-20 15:30:20", tm.String())
	vtest.Equal(t, "2025-03-20 15:30:20", tm.UTC().String())
	vtest.Equal(t, int64(1742484620), tm.Unix())
	vtest.Equal(t, int64(1742484620000), tm.UnixMilli())
}

func TestTimestamp(t *testing.T) {
	vtest.Equal(t, time.Now().Unix(), vtime.Timestamp())
	vtest.Equal(t, time.Now().Unix(), vtime.Now().Timestamp())
	vtest.Equal(t, time.Now().Unix(), vtime.Now().Unix())
}

func TestTimestampMilli(t *testing.T) {
	vtest.Equal(t, time.Now().UnixMilli(), vtime.TimestampMilli())
	vtest.Equal(t, time.Now().UnixMilli(), vtime.Now().TimestampMilli())
	vtest.Equal(t, time.Now().UnixMilli(), vtime.Now().UnixMilli())
}
