package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Fetcher fetches HTTP/HTTP(s) resources.
// The type is safe for concurrent use. Multiple go-routines
// can safely use the same instance of Fetcher
type Fetcher struct {
	client     http.Client
	maxRetries int
}

// NewFetcher creates an instance of the Fetcher
func NewFetcher(c *http.Client, retries int) *Fetcher {
	f := &Fetcher{
		maxRetries: retries,
	}
	if c != nil {
		f.client = *c
	}
	return f
}

// Fetch fetches a given URL and returns the response.
// The function optimizes the fetching and only downloads
// the underlying resource if the MIME type of the resource
// is text/html
func (f *Fetcher) Fetch(u *url.URL) (*http.Response, error) {
	// sanity check
	if u == nil {
		return nil, errors.New("url must not be nil")
	}
	if !u.IsAbs() {
		return nil, fmt.Errorf("target url is not absolute %s", u.String())
	}

	// Get the entity only if resource claims to be an HTML
	req, err := http.NewRequest("HEAD", u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := f.try(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", resp.StatusCode)
	}
	if isHTML(resp) {
		req, err = http.NewRequest("HEAD", u.String(), nil)
		if err != nil {
			return nil, err
		}
		return f.try(req)
	}
	return resp, err
}

// a utility method that retries the HTTP request in case of server
// and retrybable network errors.
func (f *Fetcher) try(req *http.Request) (*http.Response, error) {
	retryable := func(resp *http.Response, err error) bool {
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				return true
			}
			return false
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			return true
		}
		return false
	}

	for i := 0; i < f.maxRetries; i++ {
		resp, err := f.client.Do(req)
		if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
			if retryable(resp, err) {
				time.Sleep(defaultBackoff.duration(i))
				continue
			}
		}
		return resp, err
	}
	return nil, nil
}
