package md

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// // lists parsing

// var isBlockQuote bool = false

// var lines []string = []string{
// 	"> Hello cutie",
// 	"> Hello world ",
// 	">  : I am a bad ass",
// 	"> >Hi world",
// 	"> : Are you a bad ass",
// 	": Yes I am",
// 	">Hello girls",
// 	": I am smart",
// 	"## ahelo",
// 	"Hello cutie",
// 	"Hello world ",
// 	": I am a bad ass",
// 	"Hi world",
// 	"> : Are you a bad ass",
// 	"> > >1. : Yes I am",
// 	">2. Hello girls",
// 	": I am smart",
// 	"## ahelo",
// }

// func main() {
// 	blockQuotes(lines)
// }

// func isQuote(line string) bool {
// 	re, _ := regexp.Compile(`^>`)

// 	return re.MatchString(strings.Trim(line, " "))
// }

// func handleChildBlockQuote(line string) string {

// 	if isQuote(line) {
// 		return fmt.Sprintf(`<blockquote>%s</blockquote>`, strings.Trim(line, " ")[1:])
// 	}

// 	return line
// }

// func blockQuotes(lines []string) []string {

// 	for i, line := range lines {

// 		if isQuote(line) {
// 			// first bq item
// 			if !isBlockQuote {
// 				isBlockQuote = true
// 				lines[i] = fmt.Sprintf("<blockquote>%s", handleChildBlockQuote(strings.Trim(line, " ")[1:])) // first definition line
// 			} else if isBlockQuote {
// 				lines[i] = handleChildBlockQuote(strings.Trim(line, " ")[1:]) // block quote text
// 			}
// 		} else {
// 			if i-1 < 0 {
// 				continue
// 			}
// 			if !isQuote(strings.Trim(line, " ")) && isBlockQuote {
// 				isBlockQuote = false
// 				lines[i-1] += "</blockquote>" //last line for this batch of definition list

// 			}
// 		}

// 		//fmt.Println("line is", lines[i])
// 	}

// 	fmt.Println("lines are", lines)

// 	//fmt.Println("results are", lines)

// 	return lines
// }
