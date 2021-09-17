package md

import (
	"fmt"
	"regexp"
	"strings"
)

// lists parsing

var isBlockQuote bool = false

func isQuote(line string) bool {
	re, _ := regexp.Compile(`^>`)

	return re.MatchString(strings.Trim(line, " "))
}

func handleChildBlockQuote(line string) string {

	if isQuote(line) {
		return fmt.Sprintf(`<blockquote>%s</blockquote>`, strings.Trim(line, " ")[1:])
	}

	return line
}

func ParseBlockQuotes(lines []string) []string {

	for i, line := range lines {

		if isQuote(line) {
			// first bq item
			if !isBlockQuote {
				isBlockQuote = true
				lines[i] = fmt.Sprintf("<blockquote>%s", handleChildBlockQuote(strings.Trim(line, " ")[1:])) // first definition line
			} else {
				lines[i] = handleChildBlockQuote(strings.Trim(line, " ")[1:]) // block quote text
			}
		} else {
			if i-1 < 0 {
				continue
			}
			if !isQuote(strings.Trim(line, " ")) && isBlockQuote {
				isBlockQuote = false
				lines[i-1] += "</blockquote>" //last line for this batch of definition list

			}
		}

	}

	return lines
}
