package status

import "sync/atomic"

var IsReady atomic.Value //nolint: gochecknoglobals

func init() { //nolint: gochecknoinits
	// Initialize readiness to false until setup is complete
	IsReady.Store(false)
}
