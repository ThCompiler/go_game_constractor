package stringutilits

import "regexp"

func ClearStringFromPunctuation(str string) string {
	var re = regexp.MustCompile(`[[:punct:]]`)
	return re.ReplaceAllString(str, "")
}
