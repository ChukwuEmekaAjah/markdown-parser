package md

import (
	"fmt"
	"regexp"
	"strings"
)

// lists parsing

var isList bool = false

func ParseUnorderedList(lines []string) []string {
	for i, line := range lines {

		if string(strings.TrimLeft(line, " ")[0]) == string("-") {
			// first list item
			if !isList {
				isList = true
				lines[i] = fmt.Sprintf("<ul><li>%s</li>", strings.TrimLeft(line, " ")[1:]) // first definition line
			} else if isList { // list item
				lines[i] = fmt.Sprintf("<li>%s</li>", strings.TrimLeft(line, " ")[1:]) // first definition line
			}
		} else {
			if i-1 < 0 {
				continue
			}
			if string(strings.TrimLeft(lines[i-1], " ")[0]) != string("-") && isList {
				isList = false
				lines[i-1] += "</ul>" //last line for this batch of definition list

			}
		}
	}
	return lines
}

func ParseOrderedList(lines []string) []string {
	re, _ := regexp.Compile(`^\d\.`)
	for i, line := range lines {

		if re.MatchString(strings.TrimLeft(line, " ")) {
			// first list item
			if !isList {
				isList = true
				lines[i] = fmt.Sprintf("<ol><li>%s</li>", strings.TrimLeft(line, " ")[2:]) // first ordered list item
			} else if isList { // list item
				lines[i] = fmt.Sprintf("<li>%s</li>", strings.TrimLeft(line, " ")[2:]) // ordered list item
			}
		} else {
			if i-1 < 0 {
				continue
			}
			if !re.MatchString(strings.TrimLeft(line, " ")) && isList {
				isList = false
				lines[i-1] += "</ol>" //last line for this batch of definition list

			}
		}
	}

	return lines
}
