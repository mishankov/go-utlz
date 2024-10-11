# CLI Utils
================

A collection of utility functions for building command-line interfaces in Go.

## Overview

This package provides a set of functions to simplify the process of reading user input, formatting output, and performing other common tasks when building CLI tools.

## Functions

### `UserInput`

Reads a line from user input and returns the trimmed result.

* `func UserInput(prompt string) (string, error)`

### `FormatListWithSeparator`

Formats a list of elements into a string, using a separator.

* `func FormatListWithSeparator[T string | rune](l []T, sep string) string`

## Usage

To use this package, simply import it in your Go program:
```go
import "github.com/mishankov/cliutils"
```

You can then use the functions provided by the package to read user input, format output, and perform other tasks.

## Example

Here is an example of how to use the `UserInput` function to read a line of input from the user:
```go
package main

import (
	"fmt"
	"github.com/yourusername/cliutils"
)

func main() {
	prompt := "Enter your name: "
	name, err := cliutils.UserInput(prompt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hello, " + name)
}
```
This code will print "Enter your name: " to the console, read a line of input from the user, and then print "Hello, " followed by the user's input.

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.