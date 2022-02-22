package cmd_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/matherique/cmd"
)

type baz struct{}
func (baz) Handler(a []string) error {
  return fmt.Errorf("custom error")
}

func TestCommand(t *testing.T) {
	c := cmd.New("foo")
	c.SetHandler(foo{})

	br := cmd.New("bar")
	br.SetHandler(bar{})

	bz := cmd.New("baz")
	bz.SetHandler(baz{})

	c.AddSub(br)
	c.AddSub(bz)

	ts := []struct {
		args []string
		err  error
	}{
		{[]string{"bar"}, nil},
		{[]string{"baz"}, fmt.Errorf("custom error")},
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
