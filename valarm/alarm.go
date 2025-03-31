// Package valarm provides functionality for sending alarm notifications, particularly to enterprise WeChat webhooks.
// It supports message queuing, deduplication, and scheduled sending.
package valarm

import (
	"context"
	"time"

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
	// Alarm represents an alarm notification system that manages message queuing and sending.
	Alarm struct {
		Options *Options
		queue   *vqueue.Queue
	}

	// Options defines the configuration options for the Alarm system.
	Options struct {
		QueueMaxSize int    // Maximum number of messages in the queue
		SendInterval int    // Interval between message sends in seconds
		Url          string // Webhook URL for sending messages
	}
)

var (
	// defaultOptions provides default configuration values for the Alarm system
	defaultOptions = map[string]any{
		"QueueMaxSize": 10000,
		"SendInterval": 30,
		"Url":          "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=",
	}
)

// New creates a new Alarm instance with the specified options.
// It merges the provided options with default values and initializes the message queue.
func New(options map[string]any) (alarm *Alarm, err error) {
	alarm = &Alarm{Options: &Options{}}
	err = alarm.SetConfig(options)
	return
}

// GetMsg formats the alarm message into a JSON string suitable for sending to the webhook.
// It includes title, service, trace, total count, timestamp, and content in markdown format.
// The total count color changes to red if it exceeds 100.
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

// Process handles message deduplication and aggregation.
// It combines messages with the same title, service, and trace, incrementing their total count.
// Messages that are too recent (less than half the send interval) are requeued.
func (alarm *Alarm) Process() (data *vmap.Map) {
	data = vmap.New(true)
	qLen := alarm.queue.Len()

	for range qLen {
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

// SetConfig updates the Alarm configuration with the provided options.
// It merges the options with default values and initializes the message queue.
func (alarm *Alarm) SetConfig(options map[string]any) (err error) {
	if err = vreflect.SetAttrs(alarm.Options, vmap.Merge(defaultOptions, options)); err != nil {
		return
	}

	alarm.queue = vqueue.New(alarm.Options.QueueMaxSize)
	return
}

// Push adds a new alarm message to the queue.
// Messages are deduplicated based on the combination of service, title, and trace.
// If the queue is full, the message is dropped.
func (alarm *Alarm) Push(title string, service string, trace string, content string) {
	if int(alarm.queue.Size()) >= alarm.Options.QueueMaxSize {
		return
	}

	alarm.queue.Push(vmap.M{"title": title, "service": service, "trace": trace, "content": content, "time": time.Now(), "total": 1})
}

// Run processes and sends all queued messages.
// It handles message deduplication and sends formatted messages to the webhook.
func (alarm *Alarm) Run() {
	data := alarm.Process()
	data.Iterator(func(k any, v any) bool {
		_ = alarm.Send(alarm.GetMsg(v.(vmap.M)))
		return true
	})
}

// RunCron starts a scheduled task to send messages at regular intervals.
// The interval is determined by the SendInterval option.
// It uses the cron package to manage the scheduling.
func (alarm *Alarm) RunCron() (err error) {
	if _, err = vcron.AddSingleton(context.TODO(), vstr.S("@every %ds", alarm.Options.SendInterval), func(ctx context.Context) {
		alarm.Run()
	}, "alarmCron"); err != nil {
		return
	}

	vcron.Start()
	return
}

// Send transmits the formatted message to the configured webhook URL.
// It sets the appropriate content type and sends the message via HTTP POST.
func (alarm *Alarm) Send(msg string) (err error) {
	_, err = vhttp.C.R().SetHeader("Content-Type", vapi.MimeJson).
		SetBody(msg).
		Post(alarm.Options.Url)

	return
}
