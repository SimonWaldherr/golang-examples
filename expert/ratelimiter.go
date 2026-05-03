// Description: Token-bucket rate limiter using time.Ticker and channels
// Tags: rate limiter, ticker, channel, goroutine, throttle, concurrency
package main

import (
	"fmt"
	"time"
)

// RateLimiter controls how often events are allowed.
type RateLimiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

// NewRateLimiter creates a limiter that allows rate events per second.
// burst is the maximum number of tokens that can be pre-filled.
func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, burst),
		stop:   make(chan struct{}),
	}

	// Pre-fill the bucket up to burst capacity.
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}

	// Refill tokens at the given rate.
	interval := time.Second / time.Duration(rate)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case rl.tokens <- struct{}{}: // add a token if there is space
				default: // bucket full, discard
				}
			case <-rl.stop:
				return
			}
		}
	}()

	return rl
}

// Allow blocks until a token is available, then consumes it.
func (rl *RateLimiter) Allow() {
	<-rl.tokens
}

// TryAllow returns true and consumes a token if one is available immediately.
func (rl *RateLimiter) TryAllow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

// Stop shuts down the background refill goroutine.
func (rl *RateLimiter) Stop() {
	close(rl.stop)
}

func main() {
	// Allow 5 requests per second, burst of 3.
	limiter := NewRateLimiter(5, 3)
	defer limiter.Stop()

	fmt.Println("=== Blocking rate limiter (5 req/s, burst 3) ===")
	for i := 1; i <= 8; i++ {
		start := time.Now()
		limiter.Allow()
		fmt.Printf("request %d allowed after %v\n", i, time.Since(start).Round(time.Millisecond))
	}

	fmt.Println("\n=== Non-blocking TryAllow ===")
	// Refill a bit before trying non-blocking calls.
	time.Sleep(300 * time.Millisecond)
	for i := 1; i <= 5; i++ {
		if limiter.TryAllow() {
			fmt.Printf("request %d: token available\n", i)
		} else {
			fmt.Printf("request %d: rate limit exceeded, dropping\n", i)
		}
	}
}
