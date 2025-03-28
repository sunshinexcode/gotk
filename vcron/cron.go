package vcron

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
)

// Add adds a timed task to default cron object.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
func Add(ctx context.Context, pattern string, job gcron.JobFunc, name ...string) (*gcron.Entry, error) {
	return gcron.Add(ctx, pattern, job, name...)
}

// AddSingleton adds a singleton timed task, to default cron object.
// A singleton timed task is that can only be running one single instance at the same time.
// A unique `name` can be bound with the timed task.
// It returns and error if the `name` is already used.
func AddSingleton(ctx context.Context, pattern string, job gcron.JobFunc, name ...string) (*gcron.Entry, error) {
	return gcron.AddSingleton(ctx, pattern, job, name...)
}

// Start starts running the specified timed task named `name`.
// If no`name` specified, it starts the entire cron.
func Start(name ...string) {
	gcron.Start(name...)
}

// Stop stops running the specified timed task named `name`.
// If no`name` specified, it stops the entire cron.
func Stop(name ...string) {
	gcron.Stop(name...)
}
