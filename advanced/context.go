// Description: context package - cancellation, deadlines, and value propagation
// Tags: context, cancel, timeout, deadline, WithCancel, WithTimeout, WithValue
package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation simulates work that respects context cancellation.
func slowOperation(ctx context.Context, name string) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("%s: finished\n", name)
		return nil
	case <-ctx.Done():
		fmt.Printf("%s: cancelled – %v\n", name, ctx.Err())
		return ctx.Err()
	}
}

// fetchWithTimeout demonstrates context.WithTimeout.
func fetchWithTimeout(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // always release resources

	if err := slowOperation(ctx, "fetch "+url); err != nil {
		fmt.Println("fetch error:", err)
	}
}

// userIDKey is an unexported type used as a context key to avoid collisions.
type userIDKey struct{}

// withUserID returns a new context that carries the user ID.
func withUserID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIDKey{}, id)
}

// getUserID retrieves the user ID stored in the context.
func getUserID(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey{}).(int)
	return id, ok
}

func main() {
	// 1. WithCancel: manual cancellation
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("cancelling context...")
		cancel()
	}()

	if err := slowOperation(ctx, "job-1"); err != nil {
		fmt.Println("job-1 error:", err)
	}

	// 2. WithTimeout: automatic cancellation after a duration
	fmt.Println("\n--- timeout example ---")
	fetchWithTimeout("https://example.com")

	// 3. WithDeadline: cancellation at an absolute point in time
	fmt.Println("\n--- deadline example ---")
	deadline := time.Now().Add(400 * time.Millisecond)
	dctx, dcancel := context.WithDeadline(context.Background(), deadline)
	defer dcancel()

	if err := slowOperation(dctx, "job-2"); err != nil {
		fmt.Println("job-2 error:", err)
	}

	// 4. WithValue: attach request-scoped data
	fmt.Println("\n--- value example ---")
	vctx := withUserID(context.Background(), 42)
	if id, ok := getUserID(vctx); ok {
		fmt.Printf("user ID from context: %d\n", id)
	}
}
