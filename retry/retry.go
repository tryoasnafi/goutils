package retry

import "time"

// Exec execute a function up to maximum retries
func Exec(maxRetries int, delay time.Duration, fn func() error) (bool, []error) {
	var errs []error
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			errs = append(errs, err)
			time.Sleep(delay)
			continue
		}
		return true, nil
	}
	return false, errs
}
