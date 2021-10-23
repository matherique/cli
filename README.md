# 🧰 CMD toolbox

![example workflow](https://github.com/matherique/cmd/actions/workflows/test.yml/badge.svg)


 > a set of tools to create a CLI in go

## Usage 

```go 
package main

import (
  "fmt"
  "os"

  "github.com/matherique/cmd"
)

func main() {
  c := cmd.New("foo")
  c.AddDesc("foot command description")

  sc := cmd.New("bar")
  sc.AddDesc("bar subcommand description")
  
  c.AddSub(sc)
  
  if err := c.Run(os.Args); err != nil {
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


