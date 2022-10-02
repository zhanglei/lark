package ws

import (
	"fmt"
	"time"
)

type Timer struct {
	ID       int64
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

func NewTimer(d time.Duration) *Timer {
	t := &Timer{Start: time.Now(), Duration: d}
	t.ID = t.Start.Unix()
	t.End = t.Start
	t.Run()
	return t
}

func (t *Timer) Run() {
	go func() {
		ticker := time.NewTicker(t.Duration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println(t.ID, " 耗时(毫秒):", t.End.Sub(t.Start).Milliseconds())
			}
		}
	}()
}
