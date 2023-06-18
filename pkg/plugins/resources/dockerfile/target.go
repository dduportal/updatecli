package dockerfile

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/updatecli/updatecli/pkg/core/result"
)

// Target updates a targeted Dockerfile from source control management system
func (d *Dockerfile) Target(source string, dryRun bool, resultTarget *result.Target) (err error) {
	dockerfileContent, err := d.contentRetriever.ReadAll(d.spec.File)
	if err != nil {
		return err
	}

	logrus.Debugf("\n🐋 On (Docker)file %q:\n\n", d.spec.File)

	// At the moment, this plugin do not return the currently used value
	// This could be a useful improvement for the source
	resultTarget.OldInformation = "unknown"
	resultTarget.NewInformation = source

	newDockerfileContent, changedLines, err := d.parser.ReplaceInstructions([]byte(dockerfileContent), source)
	if err != nil {
		return err
	}

	if len(changedLines) == 0 {
		logrus.Debugf("no change detected %q, nothing else to do", d.spec.File)
		resultTarget.Changed = false
		resultTarget.Result = result.SUCCESS
		return nil
	}

	lines := []int{}
	for idx := range changedLines {
		lines = append(lines, idx)
	}
	sort.Ints(lines)

	resultTarget.Description = fmt.Sprintf("changed lines %v of file %q", lines, d.spec.File)
	resultTarget.Files = append(resultTarget.Files, d.spec.File)
	resultTarget.Changed = true
	resultTarget.Result = result.ATTENTION

	if !dryRun {
		// Write the new Dockerfile content from buffer to file
		err := d.contentRetriever.WriteToFile(string(newDockerfileContent), d.spec.File)
		if err != nil {
			return err
		}
	}

	return err
}
