package branch

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

func (g Gitea) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("target not supported for the plugin Gitea branch")
}
