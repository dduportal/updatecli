package awsami

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

func (a *AMI) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("Target not supported for the plugin AWS/AMI")
}
