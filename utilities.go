package main

import (
	"regexp"
	"strings"
)

func getFunctions(val string) [][]string {
	compRegEx := regexp.MustCompile(`.*:\s*function\s*.*`)
	match := compRegEx.FindAllStringSubmatch(val, -1)
	return match
}

func stringInArray(str string, list []string) bool {
	for _, v := range list {
		if strings.Contains(strings.ToLower(str), strings.ToLower(v)) {
			return true
		}
	}
	return false
}
