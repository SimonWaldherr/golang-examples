// Description: Deterministic concurrency tests with virtual time (Go 1.25)
// Tags: testing, synctest, concurrency, virtual time, timeout, Go 1.25
package synctestexample

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

func waitForWork(ctx context.Context, workTime time.Duration) error {
	select {
	case <-time.After(workTime):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestTimeoutWithoutWaitingTenSeconds(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
		defer cancel()

		err := waitForWork(ctx, time.Hour)

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("got %v, want context deadline exceeded", err)
		}
		if elapsed := time.Since(start); elapsed != 10*time.Second {
			t.Fatalf("virtual elapsed time = %v, want 10s", elapsed)
		}
	})
}
