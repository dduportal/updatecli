package git

import (
	"fmt"
	"os"

	git "github.com/olblak/updateCli/pkg/plugins/git/generic"
)

// Add run `git add`.
func (g *Git) Add(files []string) error {
	err := git.Add(files, g.GetDirectory())
	if err != nil {
		return err
	}
	return nil
}

// Checkout create and then uses a temporary git branch.
func (g *Git) Checkout() error {
	err := git.Checkout(g.Branch, g.remoteBranch, g.GetDirectory())
	if err != nil {
		return err
	}
	return nil
}

// GetDirectory returns the working git directory.
func (g *Git) GetDirectory() (directory string) {
	return g.Directory
}

// Clean removes the current git repository from local storage.
func (g *Git) Clean() error {
	err := os.RemoveAll(g.Directory) // clean up
	if err != nil {
		return err
	}
	return nil
}

// Clone run `git clone`.
func (g *Git) Clone() (string, error) {

	g.setDirectory()

	err := git.Clone(
		g.Username,
		g.Password,
		g.URL,
		g.GetDirectory())

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = g.Checkout()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return g.Directory, nil
}

// Commit run `git commit`.
func (g *Git) Commit(message string) error {
	err := git.Commit(
		g.User,
		g.Email,
		message,
		g.GetDirectory())

	if err != nil {
		return err
	}
	return nil
}

// Init set Git parameters if needed.
func (g *Git) Init(source string, name string) error {
	g.Version = source
	g.setDirectory()
	g.remoteBranch = git.SanitizeBranchName(g.Branch)
	return nil
}

// Push run `git push`.
func (g *Git) Push() error {
	err := git.Push(
		g.Username,
		g.Password,
		g.GetDirectory())

	if err != nil {
		return err
	}

	fmt.Printf("\n")
	return nil

}
