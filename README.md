# Hyper-Text Template Language(HTTL) Written by Go

## Documentation

## Install

```shell
go get github.com/xiaoma20082008/httl-go.git
```

## Usage

### import

```go
package main

import (
    "fmt"
    httl "github.com/xiaoma20082008/httl-go"
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

```properties
engine=engines.NewDefaultEngine
compiler=Compiler
converter=Converter
filter=Filter
formatter=Formatter
interceptor=Interceptor
listener=Listener
loader=loaders.NewFileLoader
logger=Logger
parser=Parser
resolver=resolvers.NewMultiResolver
resolvers=resolvers.NewGlobalResolver;resolvers.NewSessionResolver;resolvers.NewEngineResolver
translator=Translator
```

## Credit

* [HTTL](https://github.com/httl/httl.git) Written by Java
