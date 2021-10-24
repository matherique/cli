package cmd_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/matherique/cmd"
)

func TestCommand(t *testing.T) {
	c := cmd.New("foo")
	c.SetHandler(func(arg []string) error {
		return nil
	})

	bar := cmd.New("bar")
	bar.SetHandler(func(args []string) error {
		return fmt.Errorf("custom error")
	})

	baz := cmd.New("baz")
	baz.SetHandler(func(args []string) error {
		return nil
	})

	c.AddSub(bar)
	c.AddSub(baz)

	ts := []struct {
		args []string
		err  error
	}{
		{[]string{"baz"}, nil},
		{[]string{"bar"}, fmt.Errorf("custom error")},
	}

	for _, tt := range ts {
		err := c.Run(tt.args)

		if tt.err != nil {
			if errors.Is(err, tt.err) {
				t.Fatalf("expect %v, got %v", tt.err, err)
			}
		}

	}
}
