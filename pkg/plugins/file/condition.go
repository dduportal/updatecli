package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/olblak/updateCli/pkg/core/scm"
)

// Condition test if a file content match the content provided via configuration.
// If the configuration doesn't specify a value then it fall back to the source output
func (f *File) Condition(source string) (bool, error) {

	// If f.Content is not provided then we use the value returned from source
	if len(f.Content) == 0 {
		f.Content = source
	}

	data, err := Read(f.File, "")
	if err != nil {
		return false, err
	}

	if strings.Compare(f.Content, string(data)) == 0 {
		fmt.Printf("\u2714 Content from file '%v' is correct'\n", filepath.Join(f.File))
		return true, nil
	}

	fmt.Printf("\u2717 Wrong content from file '%v'. \n%s\n",
		f.File, Diff(f.Content, string(data)))

	return false, nil
}

// ConditionFromSCM test if a file content from SCM match the content provided via configuration.
// If the configuration doesn't specify a value then it fall back to the source output
func (f *File) ConditionFromSCM(source string, scm scm.Scm) (bool, error) {

	if len(f.Content) > 0 {
		fmt.Println("Key content defined from updatecli configuration")
	} else {
		f.Content = source
	}

	data, err := Read(f.File, scm.GetDirectory())
	if err != nil {
		return false, err
	}

	if strings.Compare(f.Content, string(data)) == 0 {
		fmt.Printf("\u2714 Content from file '%v' is correct'\n", f.File)
		return true, nil
	}

	fmt.Printf("\u2717 Wrong content from file '%v'. \n%s\n",
		f.File, Diff(f.Content, string(data)))

	return false, nil
}
