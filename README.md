# ProjRoot

Just find go project root absolute path.

## Install

```bash
go get -u github.com/tMinamiii/projroot
```

## Example

```go
package main

import "github.com/tMinamiii/projroot"

func main() {
    projRoot := projroot.Find()
    fmt.Println(projRoot)
}
```
