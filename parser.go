package md

import (
	"fmt"
	"regexp"
	"strings"
)

var InlineTagsList = []string{"s", "strong i", "strong", "i"} // set in order of precedence
var InlineTags = map[string]string{
	"s":        "~~(.*?)~~",
	"i":        "\\*(.+?)\\*",
	"strong":   "\\*\\*(.+?)\\*\\*",
	"strong i": "\\*\\*\\*(.+?)\\*\\*\\*",
}

var InlineTagsList2 = []string{"strong i", "strong", "i"} // set in order of precedence
var InlineTags2 = map[string]string{
	"i":        "\\_(.+?)\\_",
	"strong":   "\\_\\_(.+?)\\_\\_",
	"strong i": "\\_\\_\\_(.+?)\\_\\_\\_",
}

var LineStartTagsList = []string{"h5", "h4", "h3", "h2", "h1"} // set in order of precendence
var LineStartTags = map[string]string{
	"h1": "#\\s(.*)",
	"h2": "##\\s(.*)",
	"h3": "###\\s(.*)",
	"h4": "####\\s(.*)",
	"h5": "#####\\s(.*)",
}

var MediaTagsList = []string{"img", "a"}
var MediaTags = map[string]string{
	"img": "!\\[(.+?)\\]\\((.+?)\\)",
	"a":   "\\[(.+?)\\]\\((.+?)\\)",
}

var WholeLineTags = map[string]string{
	"hr": "^(\\-\\s?){3,}|(\\*\\s?){3,}|(\\_\\s?){3,}$",
}

var RecursiveTags = map[string]string{
	">": "bq",
	"|": "tr",
}

func Replacer(str string) string {

	// whole line tags
	for tag, rePattern := range WholeLineTags {
		re, _ := regexp.Compile(rePattern)
		str = strings.Trim(str, " ")
		str = re.ReplaceAllString(str, fmt.Sprintf(`<%s />`, tag))
	}

	// inline tags
	for _, tag := range InlineTagsList {
		re, _ := regexp.Compile(InlineTags[tag])

		tags := strings.Split(tag, " ") // for multiple combined tags
		openingTags := ""
		closingTags := ""
		for _, t := range tags {
			openingTags += fmt.Sprintf("<%s>", t)
			closingTags = fmt.Sprintf("</%s>", t) + closingTags
		}

		str = re.ReplaceAllString(str, fmt.Sprintf(`%s%s%s`, openingTags, "$1", closingTags))
	}

	// for repeating inline tag names e.g strong, i
	for _, tag := range InlineTagsList2 {
		re, _ := regexp.Compile(InlineTags2[tag])

		tags := strings.Split(tag, " ") // for multiple combined tags
		openingTags := ""
		closingTags := ""
		for _, t := range tags {
			openingTags += fmt.Sprintf("<%s>", t)
			closingTags = fmt.Sprintf("</%s>", t) + closingTags
		}

		str = re.ReplaceAllString(str, fmt.Sprintf(`%s%s%s`, openingTags, "$1", closingTags))
	}

	// whole line starters
	for _, tag := range LineStartTagsList {
		re, _ := regexp.Compile(LineStartTags[tag])
		previousStringLength := len(str)
		str = re.ReplaceAllString(str, fmt.Sprintf(`<%s>%s</%s>`, tag, "$1", tag))

		if len(str) > previousStringLength { // to prevent multiple matches by other header tags
			break
		}
	}

	// media tags
	for _, tag := range MediaTagsList {
		re, _ := regexp.Compile(MediaTags[tag])

		if tag == "img" {
			str = re.ReplaceAllString(str, fmt.Sprintf(`<%s alt="%s" href="%s" />`, tag, "$1", "$2"))
		}

		if tag == "a" {
			str = re.ReplaceAllString(str, fmt.Sprintf(`<%s title="%s" href="%s">%s</%s>`, tag, "$1", "$2", "$1", tag))
		}
	}

	return str
}
