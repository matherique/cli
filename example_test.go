package toolkit_test

import (
	"fmt"

	toolkit "github.com/matherique/cli-toolkit"
)

func ExampleNewCommand() {
	cmd := toolkit.New("foo")

	fmt.Println(cmd.Name())

	// Output:
	// foo
}

func ExampleAddSubcommands() {
	cmd := toolkit.New("foo")
	sub := toolkit.New("bar")

	cmd.AddSub(sub)
	fmt.Println(len(cmd.Sub()))

	// Output:
	// 1
}

func ExampleAddDescription() {
	cmd := toolkit.New("foo")

	cmd.AddDesc("foo foo")

	fmt.Println(cmd.Desc())
	// Output:
	// foo foo
}

func ExampleAddLongDescription() {
	cmd := toolkit.New("foo")
	cmd.AddLongDesc("long foo")

	fmt.Println(cmd.LongDesc())

	// Output:
	// long foo
}

func ExampleAddHandler() {
	cmd := toolkit.New("foo")
	handler := func(args []string) error {
		fmt.Println("foo handler")
		return nil
	}

	cmd.AddHandler(handler)

	h := cmd.Handler()
	arg := make([]string, 0)
	h(arg)

	// Output:
	// foo handler
}
