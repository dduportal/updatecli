package gomodule

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

// Target is not support for gomodule
func (g *GoModule) Target(source string, dryRun bool, releaseTarget *result.Target) error {
	return fmt.Errorf("Target not supported for the plugin GO module")
}
