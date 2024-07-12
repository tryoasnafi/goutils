package online

import (
	"context"
	"net/http"
	"time"
)

// Option represents a functional option for configuring a Checker instance.
type Option func(info *Checker)

type Checker struct {
	Method  string        // http.Method
	URL     string        // reachability URL
	Timeout time.Duration // timeout duration
}

// NewChecker creates a new Checker instance with optional configurations provided by opts.
// It initializes a Checker with default values and applies any provided options to customize it.
//
// Examples:
//	// Create new Checker with default value
//	checker := NewChecker()
//	// or create a new Checker with customized method, URL, timeout.
//	checker := NewChecker(
//		WithMethod(http.MethodGet),
//		WithURL("https://example.com/api/health"),
//		WithTimeout(5 * time.Second),
//	)
func NewChecker(opts ...Option) *Checker {
	info := &Checker{
		Method:  http.MethodHead,
		URL:     "https://clients3.google.com/generate_204",
		Timeout: 10 * time.Second,
	}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

// WithTimeout sets the HTTP Method for the Checker.
// Its recommended to use HEAD or GET.
func WithMethod(method string) Option {
	return func(info *Checker) {
		info.Method = method
	}
}

// WithTimeout sets the URL for the Checker.
func WithURL(url string) Option {
	return func(info *Checker) {
		info.URL = url
	}
}

// WithTimeout sets the timeout duration for the Checker.
func WithTimeout(d time.Duration) Option {
	return func(info *Checker) {
		info.Timeout = d
	}
}

// IsOnline checks if the service at the specified URL is reachable. It sends 
// an HTTP request using the configured method and URL with a timeout context.
// Returns true if the service is online and responds without error;
// otherwise, returns false.
func (me *Checker) IsOnline() bool {
	ctx, cancel := context.WithTimeout(context.Background(), me.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, me.Method, me.URL, nil)
	if err != nil {
		return false
	}
	if _, err = http.DefaultClient.Do(req); err != nil {
		return false
	}
	return true
}
