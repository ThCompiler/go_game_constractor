package matchers

import (
	"regexp"
	"strings"
)

type RegexMatcher struct {
	pattern     *regexp.Regexp
	nameMatched string
}

func NewRegexMather(pattern string, nameMatched string) *RegexMatcher {
	return &RegexMatcher{
		pattern:     regexp.MustCompile(pattern),
		nameMatched: nameMatched,
	}
}

func (rm *RegexMatcher) Match(message string) (bool, string) {
	res := rm.pattern.FindString(message)
	
	return res != "", res
}

func (rm *RegexMatcher) GetMatchedName() string {
	return rm.nameMatched
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

func (sm *SelectorMatcher) Match(message string) (bool, string) {
	found := ""

	for _, variant := range sm.variants {
		if strings.EqualFold(variant, message) {
			if sm.replaceMessage == "" {
				found = message
			} else {
				found = sm.replaceMessage
			}

			break
		}
	}

	return found != "", found
}

func (sm *SelectorMatcher) GetMatchedName() string {
	return sm.replaceMessage
}
