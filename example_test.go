package toolkit_test

import (
	"fmt"

	toolkit "github.com/matherique/cli-toolkit"
)

func ExampleNew() {
	cmd := toolkit.New("foo")

	fmt.Println(cmd.Name())

	// Output:
	// foo
}

func ExampleAddSub() {
	cmd := toolkit.New("foo")
	sub := toolkit.New("bar")

	cmd.AddSub(sub)
	fmt.Println(len(cmd.Sub()))

	// Output:
	// 1
}

func ExampleSetDesc() {
	cmd := toolkit.New("foo")

	cmd.SetDesc("foo foo")

	fmt.Println(cmd.Desc())
	// Output:
	// foo foo
}

func ExampleSetLongDesc() {
	cmd := toolkit.New("foo")
	cmd.SetLongDesc("long foo")

	fmt.Println(cmd.LongDesc())

	// Output:
	// long foo
}

func ExampleSetHandler() {
	cmd := toolkit.New("foo")
	handler := func(args []string) error {
		fmt.Println("foo handler")
		return nil
	}

	cmd.SetHandler(handler)

	h := cmd.Handler()
	arg := make([]string, 0)
	h(arg)

	// Output:
	// foo handler
}
