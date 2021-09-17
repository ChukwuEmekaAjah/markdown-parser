package md

import (
	"fmt"
	"regexp"
	"strings"
)

var isBlockQuote bool = false

func isQuote(line string) bool {
	re, _ := regexp.Compile(`^>`)

	return re.MatchString(strings.TrimLeft(line, " "))
}

func handleChildBlockQuote(line string) string {

	if isQuote(line) {
		return fmt.Sprintf(`<blockquote>%s</blockquote>`, strings.TrimLeft(line, " ")[1:])
	}

	return line
}

func ParseBlockQuotes(lines []string) []string {

	for i, line := range lines {

		if isQuote(line) {
			// first bq item
			if !isBlockQuote {
				isBlockQuote = true
				lines[i] = fmt.Sprintf("<blockquote>%s", handleChildBlockQuote(strings.TrimLeft(line, " ")[1:])) // first definition line
			} else {
				lines[i] = handleChildBlockQuote(strings.TrimLeft(line, " ")[1:]) // block quote text
			}
		} else {
			if i-1 < 0 {
				continue
			}
			if !isQuote(strings.TrimLeft(line, " ")) && isBlockQuote {
				isBlockQuote = false
				lines[i-1] += "</blockquote>" //last line for this batch of definition list

			}
		}

	}

	return lines
}
