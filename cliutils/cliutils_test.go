package cliutils

import (
	"testing"
)

type Tests[V string | rune] struct {
	list      []V
	separator string
	result    string
}

func TestFormatListWithSeparator(t *testing.T) {
	testsString := []Tests[string]{
		{
			list:      []string{"onlyOne"},
			separator: ", ",
			result:    "onlyOne",
		},
		{
			list:      []string{"first", "second"},
			separator: ", ",
			result:    "first, second",
		},
		{
			list:      []string{"first", "second", "third"},
			separator: " | ",
			result:    "first | second | third",
		},
		{
			list:      []string{},
			separator: ", ",
			result:    "",
		},
	}

	testsRune := []Tests[rune]{
		{
			list:      []rune{'o'},
			separator: ", ",
			result:    "o",
		},
		{
			list:      []rune{'f', 's'},
			separator: ", ",
			result:    "f, s",
		},
		{
			list:      []rune{'f', 's', 't'},
			separator: " | ",
			result:    "f | s | t",
		},
		{
			list:      []rune{},
			separator: ", ",
			result:    "",
		},
	}

	for _, test := range testsString {
		if actualResult := FormatListWithSeparator(test.list, test.separator); test.result != actualResult {
			t.Fatalf("List: %q. Separator: %q. Want: %q. Got: %q", test.list, test.separator, test.result, actualResult)
		}
	}

	for _, test := range testsRune {
		if actualResult := FormatListWithSeparator(test.list, test.separator); test.result != actualResult {
			t.Fatalf("List: %q. Separator: %q. Want: %q. Got: %q", test.list, test.separator, test.result, actualResult)
		}
	}
}
