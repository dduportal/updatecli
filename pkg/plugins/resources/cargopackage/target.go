package cargopackage

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

func (cp *CargoPackage) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("Target not supported for the plugin Cargo Package")
}
