package githubrelease

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

func (ghr GitHubRelease) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("target not supported for the plugin GitHub Release")
}
