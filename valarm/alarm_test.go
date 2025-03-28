package valarm_test

import (
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/valarm"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestGetMsg(t *testing.T) {
	alarm, err := valarm.New(nil)

	vtest.Nil(t, err)

	dataM := vmap.M{"msgtype": "text", "text": vmap.M{"content": "gotk test"}}
	dataJson, err := vjson.Encode(dataM)

	vtest.Nil(t, err)
	vtest.Nil(t, alarm.Send(dataJson))
}

func TestProcess(t *testing.T) {
	alarm, err := valarm.New(vmap.M{"QueueMaxSize": 100, "SendInterval": 2})

	vtest.Nil(t, err)

	alarm.Push("gotk_title_test", "service_test", "trace_test", "content_test")

	vtest.Equal(t, 0, alarm.Process().Size())

	time.Sleep(1 * time.Second)

	vtest.Equal(t, 1, alarm.Process().Size())
}

func TestSend(t *testing.T) {
	alarm, err := valarm.New(nil)

	vtest.Nil(t, err)

	dataM := vmap.M{"msgtype": "text", "text": vmap.M{"content": "gotk test"}}
	dataJson, err := vjson.Encode(dataM)

	vtest.Nil(t, err)
	vtest.Nil(t, alarm.Send(dataJson))

	dataM = vmap.M{"msgtype": "markdown", "markdown": vmap.M{"content": "gotk test\n" +
		">service: gotk"}}
	dataJson, err = vjson.Encode(dataM)

	vtest.Nil(t, err)
	vtest.Nil(t, alarm.Send(dataJson))
}

func TestRunCron(t *testing.T) {
	alarm, err := valarm.New(vmap.M{"QueueMaxSize": 110, "SendInterval": 2})

	vtest.Nil(t, err)

	for i := 0; i < 120; i++ {
		alarm.Push("gotk_title_test", "service_test", "trace_test", "content_test")
	}

	vtest.Nil(t, alarm.RunCron())

	time.Sleep(5 * time.Second)
	err = alarm.RunCron()

	vtest.NotNil(t, err)
	vtest.Equal(t, "duplicated cron job name \"alarmCron\", already exists", err.Error())
}

func TestSetConfig(t *testing.T) {
	alarm, err := valarm.New(nil)

	vtest.Nil(t, err)

	err = alarm.SetConfig(vmap.M{"Test": "test"})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
}
