# Hyper-Text Template Language(HTTL) Golang port

## Documentation

## Install

```shell
go get github.com/xiaoma20082008/httl-go.git httl
```

## Usage

### import

```go
package main

import (
	"fmt"
	"github.com/xiaoma20082008/httl"
)

func main() {
	template, err := httl.NewEngine().GetTemplate("about", "", "UTF-8")
	if err != nil {
		return
	}
	args := make(map[string]any)
	result, err := template.Evaluate(args)
	if err != nil {
		return
	}
	fmt.Println(result)
}
```

### configuration

```shell
touch httl.properties
```

## Credit

* [HTTL](https://github.com/httl/httl.git) Written by Java
