package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSemaphore_AcquireRelease(t *testing.T) {
	sem := newSemaphore(2)
	ctx := context.Background()

	err := sem.acquire(ctx)
	require.NoError(t, err)
	err = sem.acquire(ctx)
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		sem.release()
		sem.release()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("release blocked")
	}
}

func TestSemaphore_AcquireBlocksWhenFull(t *testing.T) {
	sem := newSemaphore(1)
	ctx := context.Background()
	require.NoError(t, sem.acquire(ctx))

	acquired := make(chan struct{})
	go func() {
		_ = sem.acquire(ctx)
		close(acquired)
	}()

	select {
	case <-acquired:
		t.Fatal("acquire should block")
	case <-time.After(50 * time.Millisecond):
	}
	sem.release()
	<-acquired
}

func TestSemaphore_AcquireRespectsContextCancel(t *testing.T) {
	sem := newSemaphore(1)
	require.NoError(t, sem.acquire(context.Background()))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := sem.acquire(ctx)
	require.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}
