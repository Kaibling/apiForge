package status

import "sync/atomic"

var IsReady atomic.Value

func init() {
	// Initialize readiness to false until setup is complete
	IsReady.Store(false)
}
