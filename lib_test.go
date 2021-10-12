package toolkit_test

import (
	"testing"

	toolkit "github.com/matherique/cli-toolkit"
)

func TestNewCommand(t *testing.T) {
	name := "foo"
	desc := "foo bar"
	cmd := toolkit.New(name, desc)

	if cmd.Name() != name {
		t.Fatalf("expect %s, got: %s", name, cmd.Name())
	}
	if cmd.Desc() != desc {
		t.Fatalf("expect %s, got: %s", desc, cmd.Desc())
	}
}

func TestAddSubcommands(t *testing.T) {
	cmd := toolkit.New("foo", "foo foo")
	sub := toolkit.New("bar", "bar bar")

	cmd.AddSub(sub)
	if len(cmd.Subs()) != 1 {
		t.Fatalf("expect %d subcommands, got: %d", 1, len(cmd.Subs()))
	}
}
