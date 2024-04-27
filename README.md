# PassGEN
```
     ____                      _____   _____  _     __
    /    \ ____   ___ ___     //  //  / ___/ / \   / /
   /  ___//   /| //_ //_  == //  __  / __/  / / \ / /  Password Generator  
  /__/    \___\|__//__//    //___// /____/ /_/   \_/

```
PassGEN is a tool for character-based password generation.

Generate your own password dictionaries based on character sets!

### Command 
###### Installation
```shell
go install github.com/yael-castro/passgen/cmd/passgen@latest
```
###### See how to use
```shell
passgen 
```

### Library
###### Installation
```shell
go get github.com/yael-castro/passgen/pkg/password@latest
```
###### See how to use

```go
package main

import (
	"fmt"
	"github.com/yael-castro/passgen/pkg/password"
)

func main() {
	p := password.New('A', 'B', 'C')

	for p.Next() {
            pw := p.Generate()
            fmt.Println(string(pw))
	}
}

```