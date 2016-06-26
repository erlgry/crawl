package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func isHTML(r *http.Response) bool {
	ct := strings.ToLower(r.Header.Get("Content-Type"))
	if strings.HasPrefix(ct, "text/html") || strings.HasPrefix(ct, "application/xhtml+xml") {
		return true
	}
	return false
}

// backoff implements a backoff policy, randomizing its delays and
// saturating at its last value.
type backoff struct {
	millis []int
}

// defaultBackoff is a backoff policy ranging up to 5s.
var defaultBackoff = backoff{
	[]int{0, 10, 10, 100, 100, 500, 500, 3000, 3000, 5000},
}

// duration returns the time duration of the n'th wait cycle in its
// backoff policy. This is backoff.millis[n], randomized to avoid
// thundering herds.
func (b backoff) duration(n int) time.Duration {
	if n >= len(b.millis) {
		n = len(b.millis) - 1
	}
	return time.Duration(jitter(b.millis[n])) * time.Millisecond
}

// jitter returns a random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func jitter(millis int) int {
	if millis == 0 {
		return 0
	}
	return millis/2 + rand.Intn(millis)
}
