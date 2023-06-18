package language

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

// Target is not supported for the Golang resource
func (l *Language) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("Target not supported for the plugin Go")
}
