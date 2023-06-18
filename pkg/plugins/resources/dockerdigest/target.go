package dockerdigest

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/result"
)

func (ds *DockerDigest) Target(source string, dryRun bool, resultTarget *result.Target) error {
	return fmt.Errorf("Target not supported for the plugin Docker Digest")
}
