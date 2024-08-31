// Package tm assistant methods related to time reporting
package tm

import "time"

// SpendFn Get the program spend time for any format.
func SpendFn() func() time.Duration {
	now := time.Now()
	return func() time.Duration {
		return time.Since(now)
	}
}
