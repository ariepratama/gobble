package preprocessors

import "strings"

func StandardPreProcessors(input string) []string {
	doc := strings.ToLower(input)
	terms := strings.Split(doc, " ")
	return terms
}
