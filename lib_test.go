package toolkit_test

import (
	"testing"

	toolkit "github.com/matherique/cli-toolkit"
)

func TestCommand(t *testing.T) {
	cmd := toolkit.New("foo")
	subcmd := toolkit.New("bar")
	cmd.AddSub(subcmd)

	cmd.SetHandler(func(arg []string) error {
		// do something
		return nil
	})

}
