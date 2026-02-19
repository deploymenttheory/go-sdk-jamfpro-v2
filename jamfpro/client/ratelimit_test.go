package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewResponseTimeTracker(t *testing.T) {
	tracker := newResponseTimeTracker()
	requireNotNil(t, tracker)
	assert.Equal(t, 0.2, tracker.alpha)
}

func TestResponseTimeTracker_Record_FirstCall(t *testing.T) {
	tracker := newResponseTimeTracker()
	d := tracker.record(50 * time.Millisecond)
	assert.Equal(t, time.Duration(0), d)
}

func TestResponseTimeTracker_Record_AtOrBelowBaseline(t *testing.T) {
	tracker := newResponseTimeTracker()
	tracker.record(10 * time.Millisecond)
	d := tracker.record(10 * time.Millisecond)
	assert.Equal(t, time.Duration(0), d)
}

func TestResponseTimeTracker_Record_ExcessAboveBaseline(t *testing.T) {
	tracker := newResponseTimeTracker()
	tracker.record(10 * time.Millisecond)
	d := tracker.record(100 * time.Millisecond)
	assert.Greater(t, d, time.Duration(0))
	assert.LessOrEqual(t, d, adaptiveDelayMax)
}

func TestResponseTimeTracker_Record_ExcessCapped(t *testing.T) {
	tracker := newResponseTimeTracker()
	tracker.record(time.Millisecond)
	d := tracker.record(1 * time.Minute)
	assert.Equal(t, adaptiveDelayMax, d)
}

func TestResponseTimeTracker_Record_Concurrent(t *testing.T) {
	tracker := newResponseTimeTracker()
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				tracker.record(time.Duration(j) * time.Millisecond)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func requireNotNil(t *testing.T, v interface{}) {
	t.Helper()
	if v == nil {
		t.Fatal("expected non-nil")
	}
}
