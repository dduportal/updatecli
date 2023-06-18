package tag

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

// Target ensure that a specific release exist on GitLab, otherwise creates it
func (g Gitlab) Target(source string, dryRun bool, releaseTarget *result.Target) error {
	return fmt.Errorf("target not supported for the plugin GitLab Tags")
}
