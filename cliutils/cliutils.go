package cliutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UserInput reads a line from user input and returns the trimmed result.
// It returns an error only if the read fails.
// If the promt is not empty, it is printed to the console.
func UserInput(promt string) (string, error) {
	if len(promt) > 0 {
		fmt.Print(promt)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), err

}

// FormatListWithSeparator formats a list of elements into a string, using a separator.
// Example: FormatListWithSeparator([a, b, c], ", ") returns "a, b, c".
func FormatListWithSeparator[T string | rune](l []T, sep string) string {
	if len(l) == 0 {
		return ""
	}

	r := string(l[0])
	for _, el := range l[1:] {
		r += sep + string(el)
	}
	return r
}

func GetEnvOrDefault(key, def string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return def
	} else {
		return strings.TrimSpace(value)
	}
}
