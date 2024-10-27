# CLI Utils

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

To use this package, install it via `go get`:
```shell
go get github.com/mishankov/go-utlz/cliutils
```

and simply import it in your Go program:
```go
import "github.com/mishankov/go-utlz/cliutils"
```


You can then use the functions provided by the package to read user input, format output, and perform other tasks.

## Example

### `UserInput`

```go
package main

import (
	"fmt"
	"github.com/mishankov/go-utlz/cliutils"
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

### `FormatListWithSeparator`

Formats a list of strings or runes with a specified separator.

```go
list := []string{"apple", "banana", "cherry"}
formattedList := cliutils.FormatListWithSeparator(list, ", ")
fmt.Println(formattedList) // Output: "apple, banana, cherry"

list = []string{"only one"}
formattedList = cliutils.FormatListWithSeparator(list, ", ")
fmt.Println(formattedList) // Output: "only one"

list = []string{}
formattedList = cliutils.FormatListWithSeparator(list, ", ")
fmt.Println(formattedList) // Output: ""

listRunes := []rune{'a', 'p', 'p', 'l', 'e'}
formattedList = FormatListWithSeparator(listRunes, ", ")
fmt.Println(formattedList) // Output: "a, p, p, l, e"
```

### GetEnvOrDefault

Returns the value of an environment variable, or a default value if the variable is not set.

```go
value := cliutils.GetEnvOrDefault("MY_VAR", "default_value")

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.