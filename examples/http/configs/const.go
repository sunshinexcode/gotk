package configs

import "github.com/sunshinexcode/gotk/vapi"

var (
	BasicAuth = vapi.Accounts{
		"test": "test",
	}

	CronTraceIdPrefix = "cron-"

	// TestOpen = true
	TestOpen = false

	TestUrl = "http://localhost:8080"
)
