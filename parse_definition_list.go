package md

// import (
// 	"fmt"
// 	"strings"
// )

// var isDefinition bool = false

// var lines []string = []string{
// 	" Hello cutie",
// 	" Hello world ",
// 	"  : I am a bad ass",
// 	"Hi world",
// 	": Are you a bad ass",
// 	": Yes I am",
// 	"Hello girls",
// 	": I am smart",
// 	"## ahelo",
// 	"Hello cutie",
// 	"Hello world ",
// 	": I am a bad ass",
// 	"Hi world",
// 	": Are you a bad ass",
// 	": Yes I am",
// 	"Hello girls",
// 	": I am smart",
// 	"## ahelo",
// }

// func main() {
// 	recurseString(lines)
// }

// func recurseString(lines []string) []string {
// 	for i, line := range lines {

// 		if string(strings.Trim(line, " ")[0]) == string(":") {
// 			// cater for possibility of it being a start text
// 			if !isDefinition && string(strings.Trim(lines[i-1], " ")[0]) != string(":") {
// 				isDefinition = true
// 				lines[i-1] = fmt.Sprintf("<dl><dt>%s</dt>", lines[i-1]) // first definition line
// 			} else if isDefinition && string(strings.Trim(lines[i-1], " ")[0]) != string(":") && string(strings.Trim(lines[i-1], " ")[0]) != "<" { // definition term
// 				lines[i-1] = fmt.Sprintf("<dt>%s</dt>", lines[i-1]) // first definition line
// 			}

// 			if isDefinition && string(strings.Trim(lines[i], " ")[0]) == string(":") {
// 				lines[i] = fmt.Sprintf("<dd>%s</dd>", strings.Trim(line, " ")[1:]) // first definition line
// 			}
// 		} else {
// 			if i-1 < 0 {
// 				continue
// 			}
// 			if string(strings.Trim(lines[i-1], " ")[0]) != string("<") && isDefinition {
// 				isDefinition = false
// 				lines[i-2] += "</dl>" //last line for this batch of definition list

// 			}
// 		}

// 		//fmt.Println("line is", lines[i])
// 	}

// 	fmt.Println("lines are", lines)

// 	//fmt.Println("results are", lines)

// 	return lines
// }

// func recurseString(lines []string) []string {
// 	results := make([]string, 0)
// 	for i, line := range lines {
// 		if strings.Trim(string(line[0]), " ") == string(":") {
// 			if isDefinition == false {
// 				isDefinition = true
// 				startingIndex = i - 1
// 			}

// 			if strings.Trim(string(lines[i-1][0]), " ") != string(":") {

// 				if len(results) == 0 { // at first sign of definition list
// 					lines[i] = fmt.Sprintf("<dl><dt>%s</dt>", string(lines[i-1]))
// 					results = append(results, "<dl>")
// 				} else {
// 					results = append(results, fmt.Sprintf("<dt>%s</dt>", string(lines[i-1])))
// 					lines[i] = fmt.Sprintf("<dt>%s</dt>", string(lines[i-1]))
// 				}
// 			}

// 			results = append(results, fmt.Sprintf("<dd>%s</dd>", string(line[1:])))
// 			lines[i] = fmt.Sprintf("<dd>%s</dd>", string(line[1:]))
// 		} else {
// 			previous := i - 1

// 			if isDefinition && previous > -1 && strings.Trim(string(lines[i-1][0]), " ") != string(":") {
// 				results = append(results, "</dl>")
// 				lines[previous-1] += "</dl>"
// 				isDefinition = false
// 				results = []string{}
// 			}
// 		}

// 	}

// 	//fmt.Println("results are", lines)

// 	return results
// }
