package valarm

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gqueue"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcron"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vqueue"
	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vstr"
)

type (
	Alarm struct {
		Options *Options
		queue   *gqueue.Queue
	}

	Options struct {
		QueueMaxSize int
		SendInterval int // second
		Url          string
	}
)

var (
	defaultOptions = map[string]any{
		"QueueMaxSize": 10000,
		"SendInterval": 30,
		"Url":          "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=",
	}
)

// New create new alarm
func New(options map[string]any) (alarm *Alarm, err error) {
	alarm = &Alarm{Options: &Options{}}
	err = alarm.SetConfig(options)
	return
}

// GetMsg get message
func (alarm *Alarm) GetMsg(msgM vmap.M) (dataJson string) {
	totalColor := "warning"
	if msgM["total"].(int) > 100 {
		totalColor = "red"
	}

	dataM := vmap.M{"msgtype": "markdown", "markdown": vmap.M{"content": vstr.S("###### %s"+
		"\n>**service:** %s"+
		"\n>**trace:** %s"+
		"\n>**total: <font color='%s'>%d</font>**"+
		"\n>**time:** %s"+
		"\n>**content:** <font color='comment'>%s</font>", msgM["title"], msgM["service"], msgM["trace"], totalColor,
		msgM["total"], msgM["time"].(time.Time).UTC().String(), msgM["content"])}}
	dataJson, _ = vjson.Encode(dataM)

	return
}

// Process processes message
func (alarm *Alarm) Process() (data *gmap.Map) {
	data = vmap.New(true)
	qLen := alarm.queue.Len()

	for i := 0; i < int(qLen); i++ {
		msg := alarm.queue.Pop().(vmap.M)
		// Check time
		if time.Since(msg["time"].(time.Time)).Seconds() < float64(alarm.Options.SendInterval/2) {
			alarm.queue.Push(msg)
			continue
		}

		key := vstr.S("%s-%s-%s", msg["title"], msg["service"], msg["trace"])
		if data.Contains(key) {
			dataOld := data.Get(key).(vmap.M)
			msg["total"] = dataOld["total"].(int) + 1
		}
		data.Set(key, msg)
	}

	return
}

// SetConfig set config
func (alarm *Alarm) SetConfig(options map[string]any) (err error) {
	if err = vreflect.SetAttrs(alarm.Options, vmap.Merge(defaultOptions, options)); err != nil {
		return
	}

	alarm.queue = vqueue.New(alarm.Options.QueueMaxSize)
	return
}

// Push pushes the data into the queue
// data deduplication rule: service + title + trace
func (alarm *Alarm) Push(title string, service string, trace string, content string) {
	if int(alarm.queue.Size()) >= alarm.Options.QueueMaxSize {
		return
	}

	alarm.queue.Push(vmap.M{"title": title, "service": service, "trace": trace, "content": content, "time": time.Now(), "total": 1})
}

// Run send msg
func (alarm *Alarm) Run() {
	data := alarm.Process()
	data.Iterator(func(k any, v any) bool {
		_ = alarm.Send(alarm.GetMsg(v.(vmap.M)))
		return true
	})
}

// RunCron send msg by cron
func (alarm *Alarm) RunCron() (err error) {
	if _, err = vcron.AddSingleton(context.TODO(), vstr.S("@every %ds", alarm.Options.SendInterval), func(ctx context.Context) {
		alarm.Run()
	}, "alarmCron"); err != nil {
		return
	}

	vcron.Start()
	return
}

// Send sends message
func (alarm *Alarm) Send(msg string) (err error) {
	_, err = vhttp.C.R().SetHeader("Content-Type", vapi.MimeJson).
		SetBody(msg).
		Post(alarm.Options.Url)

	return
}
