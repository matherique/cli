package cmd_test

import (
	"testing"

	"github.com/matherique/cmd"
)

func TestCommand(t *testing.T) {
	c := cmd.New("foo")
	sc := cmd.New("bar")
	c.AddSub(sc)

	c.SetHandler(func(arg []string) error {
		// do something
		return nil
	})

}
