package toolkit_test

import (
	"fmt"

	toolkit "github.com/matherique/cli-toolkit"
)

func ExampleNewCommand() {
	cmd := toolkit.New("test", "description of test command")
	fmt.Println(cmd.Name())
	fmt.Println(cmd.Desc())

	// Output:
	// test
	// description of test command
}
