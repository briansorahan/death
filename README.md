# death

Package death provides some functions that I find myself writing over and over again in Go programs.

## Install

```
go get -u github.com/briansorahan/death
```

## Usage

### In func main

This is not working code, but should serve to illustrate how I expect `Main` to be used:

```
package main

import (
	"github.com/briansorahan/death"
)

func main() {
	cli, err := NewCLI()

	death.Main(err)

	death.Main(cli.Run(context.Background()))
}
```

### In testing

```
package mypkg

import (
	"testing"

	"github.com/briansorahan/death"
)

func TestMyFunc(t *testing.T) {
	die := death.By(t)

	die(MyFunc())
}
```
