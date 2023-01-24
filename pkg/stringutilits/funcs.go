package stringutilits

import "regexp"

func ClearStringFromPunctuation(str string) string {
	re := regexp.MustCompile(`[[:punct:]]`)

	return re.ReplaceAllString(str, "")
}
