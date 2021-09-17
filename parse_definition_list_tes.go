package md

// import (
// 	"fmt"
// 	"strings"
// )

// var isDefinition bool = false

// func ParseDefinitionList(lines []string) []string {
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

// 	return lines
// }
