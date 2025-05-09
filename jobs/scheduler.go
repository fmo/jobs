package jobs

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Job struct {
	Name     string
	Interval time.Duration
	Timeout  time.Duration
	Command  string
}

func (j *Job) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * j.Interval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				go func() {
					split := strings.Fields(j.Command)
					if len(split) < 2 {
						fmt.Printf("[%s] WARNING: Command or parameter is missing\n", j.Name)
						return
					}

					execCtx := ctx
					if j.Timeout > 0 {
						var cancel context.CancelFunc
						execCtx, cancel = context.WithTimeout(ctx, j.Timeout*time.Millisecond)
						defer cancel()
					}
					cmd := exec.CommandContext(execCtx, split[0], split[1:]...)
					output, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Printf("[%s] ERROR: %v\n", j.Name, err)
					}
					fmt.Println(string(output))
				}()
			case <-ctx.Done():
				return
			}
		}
	}()
}
