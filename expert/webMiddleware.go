package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func main() {
	// Define route handlers
	homeHandler := http.HandlerFunc(home)
	profileHandler := http.HandlerFunc(profile)

	// Apply the combined middleware to your route handlers
	http.Handle("/", applyMiddleware(homeHandler))
	http.Handle("/profile", applyMiddleware(profileHandler))

	http.ListenAndServe(":8080", nil)
}

// home and profile are our route handlers
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the profile page!")
}

// Combine multiple middlewares into a single applyMiddleware function
func applyMiddleware(handler http.Handler) http.Handler {
	return timingMiddleware(
		loggingMiddleware(
			corsMiddleware(
				gzipMiddleware(handler),
			),
		),
	)
}

// various middleware implementations:

// timingMiddleware logs the time taken to process a request
func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("Elapsed time: %v\n", elapsed)
	})
}

// loggingMiddleware logs the request method and URL
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// gzipMiddleware compresses the response body
func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		grw := gzipResponseWriter{ResponseWriter: w, Writer: gw}
		next.ServeHTTP(grw, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}

// corsMiddleware adds the Access-Control-Allow-Origin header
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// rateLimitMiddleware limits the number of requests per minute
func rateLimitMiddleware(next http.Handler) http.Handler {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiterCtx, err := instance.Get(r.Context(), r.RemoteAddr)
		if err != nil {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiterCtx.Limit))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiterCtx.Remaining))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", limiterCtx.Reset))

		if limiterCtx.Reached {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// authenticationMiddleware checks for a valid Authorization header
func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer my-secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// contentTypeJsonMiddleware sets the Content-Type header to application/json
func contentTypeJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// panicRecoveryMiddleware recovers from panics and returns a 500 error
func panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// cacheItem stores a response and its timestamp
type cacheItem struct {
	response  []byte
	timestamp time.Time
}

// cacheStorage is a simple in-memory cache
type cacheStorage struct {
	data  map[string]*cacheItem
	mutex sync.RWMutex
}

// newCacheStorage creates a new cacheStorage
func newCacheStorage() *cacheStorage {
	return &cacheStorage{data: make(map[string]*cacheItem)}
}

func (cs *cacheStorage) get(key string) (value []byte, found bool) {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()

	item, exists := cs.data[key]
	if exists && !isExpired(item.timestamp) {
		return item.response, true
	}
	return nil, false
}

func (cs *cacheStorage) set(key string, value []byte) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.data[key] = &cacheItem{response: value, timestamp: time.Now()}
}

func isExpired(timestamp time.Time) bool {
	return time.Since(timestamp) > 10*time.Minute
}

var cache = newCacheStorage()

// cacheMiddleware is a middleware that caches responses
func cacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cacheKey := generateCacheKey(r)
		if response, found := cache.get(cacheKey); found {
			w.Write(response)
			return
		}

		recorder := newResponseRecorder()
		next.ServeHTTP(recorder, r)
		response := recorder.Bytes()

		cache.set(cacheKey, response)
		w.Write(response)
	}
}

// generateCacheKey generates a cache key from a request
func generateCacheKey(r *http.Request) string {
	hasher := sha1.New()
	io.WriteString(hasher, r.URL.String())
	io.WriteString(hasher, r.Method)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

type responseRecorder struct {
	http.ResponseWriter
	buffer bytes.Buffer
}

// newResponseRecorder creates a new responseRecorder
func newResponseRecorder() *responseRecorder {
	return &responseRecorder{}
}

func (rr *responseRecorder) Write(data []byte) (int, error) {
	rr.buffer.Write(data)
	return rr.ResponseWriter.Write(data)
}

func (rr *responseRecorder) Bytes() []byte {
	return rr.buffer.Bytes()
}
