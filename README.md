# 🧰 CLI toolbox

![example workflow](https://github.com/matherique/cli-toolbox/actions/workflows/test.yml/badge.svg)


 > toolbox to create CLI in go

## Usage 

```go 
package main

import (
  "fmt"
  "os"

  "github.com/matherique/cli-toolbox"
)

func main() {
  cmd := toolkit.New("foo")
  cmd.AddDesc("foot command description")

  subcmd := toolkit.New("bar")
  subls.AddDesc("bar subcommand description")
  
  cmd.AddSub(subcmd)
  
  if err := cmd.Run(os.Args); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
} 

```

## Author

👤 **Matheus Henrique**

- Email: matherique@gmail.com
- Github: [@matherique](https://github.com/matherique)
- LinkedIn: [@matherique](https://linkedin.com/in/matherique)


