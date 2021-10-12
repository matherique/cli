package toolkit_test

import (
	"testing"

	toolkit "github.com/matherique/cli-toolkit"
)

func TestNewCommand(t *testing.T) {
	name := "foo"
	cmd := toolkit.New(name)

	if cmd.Name() != name {
		t.Fatalf("expect %s, got: %s", name, cmd.Name())
	}
}

func TestAddSubcommands(t *testing.T) {
	cmd := toolkit.New("foo")
	sub := toolkit.New("bar")

	cmd.AddSub(sub)
	if len(cmd.Sub()) != 1 {
		t.Fatalf("expect %d subcommands, got: %d", 1, len(cmd.Sub()))
	}
}

func TestAddDescription(t *testing.T) {
	cmd := toolkit.New("foo")

	cmd.AddDesc("foo foo")

	if cmd.Desc() != "foo foo" {
		t.Fatalf("expect '%s' as description, got '%s'", "foo foo", cmd.Desc())
	}
}
