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

func TestAddLongDescription(t *testing.T) {
	ld := "long long foo"

	cmd := toolkit.New("foo")
	cmd.AddLongDesc(ld)

	if cmd.LongDesc() != ld {
		t.Fatalf("expect '%s', got '%s'", ld, cmd.LongDesc())
	}
}

func TestAddHandler(t *testing.T) {
	cmd := toolkit.New("foo")
	handler := func(args []string) error {
		// do something
		return nil
	}

	cmd.AddHandler(handler)

	if cmd.Handler() == nil {
		t.Fatalf("expect func to be assigned, but got <nil>")
	}
}
