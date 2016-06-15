package statement

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb/stress/v2/stressClient"
)

// WaitStatement is a Statement Implementation to prevent the test from returning to early when running GoStatements
type WaitStatement struct {
	StatementID string

	runtime time.Duration
}

// SetID statisfies the Statement Interface
func (w *WaitStatement) SetID(s string) {
	w.StatementID = s
}

// Run statisfies the Statement Interface
func (w *WaitStatement) Run(s *stressClient.StoreFront) {
	runtime := time.Now()
	s.Wait()
	w.runtime = time.Since(runtime)
}

// Report statisfies the Statement Interface
func (w *WaitStatement) Report(s *stressClient.StoreFront) string {
	return fmt.Sprintf("WAIT -> %v", w.runtime)
}
