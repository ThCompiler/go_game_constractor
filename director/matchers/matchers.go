package matchers

import (
	"regexp"
	"strings"
)

type RegexMatcher struct {
	pattern *regexp.Regexp
}

func NewRegexMather(pattern string) *RegexMatcher {
	return &RegexMatcher{
		pattern: regexp.MustCompile(pattern),
	}
}

func (rm *RegexMatcher) Match(message string) (bool, string) {
	res := rm.pattern.FindString(message)
	return res != "", res
}

type SelectorMatcher struct {
	variants       []string
	replaceMessage string
}

func NewSelectorMatcher(variants []string, replaceMessage string) *SelectorMatcher {
	return &SelectorMatcher{
		variants:       variants,
		replaceMessage: replaceMessage,
	}
}

func (rm *SelectorMatcher) Match(message string) (bool, string) {
	found := ""
	for _, variant := range rm.variants {
		if strings.EqualFold(variant, message) {
			if rm.replaceMessage == "" {
				found = message
			} else {
				found = rm.replaceMessage
			}
			break
		}
	}
	return found != "", found
}
