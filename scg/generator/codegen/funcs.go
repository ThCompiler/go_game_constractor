package codegen

// Code based on goa generator: https://github.com/goadesign/goa

import (
	"bytes"
	"os"
	"strings"
	"unicode"
)

const (
	GeneratorCommandName = "scg"
	maxCommentLen        = 77
)

// TemplateFuncs lists common template helper functions.
func TemplateFuncs() map[string]interface{} {
	return map[string]interface{}{
		"commandLine": CommandLine,
		"comment":     Comment,
	}
}

// CommandLine return the command used to run this process.
func CommandLine() string {
	cmdl := GeneratorCommandName

	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--cmd=") {
			cmdl = arg[6:]

			break
		}
	}

	return cmdl
}

// Comment produces line comments by concatenating the given strings and
// producing 80 characters long lines starting with "//".
func Comment(elems ...string) string {
	var lines []string

	for _, e := range elems {
		lines = append(lines, strings.Split(e, "\n")...)
	}

	trimmed := make([]string, len(lines))

	for i, l := range lines {
		trimmed[i] = strings.TrimLeft(l, " \t")
	}

	t := strings.Join(trimmed, "\n")

	return Indent(WrapText(t, maxCommentLen), "// ")
}

func ToTitle(s string) string {
	return CamelCase(s, true, true)
}

// Indent inserts prefix at the beginning of each non-empty line of s. The
// end-of-line marker is NL.
func Indent(s, prefix string) string {
	var (
		res []byte
		b   = []byte(s)
		p   = []byte(prefix)
		bol = true
	)

	for _, c := range b {
		if bol && c != '\n' {
			res = append(res, p...)
		}

		res = append(res, c)

		bol = c == '\n'
	}

	return string(res)
}

// CopyStringMap create copy of map with strings kay and any values.
func CopyStringMap[Value any](mp map[string]Value) map[string]Value {
	newMap := make(map[string]Value)
	for key, value := range mp {
		newMap[key] = value
	}

	return newMap
}

// Casing exceptions.
var toLower = map[string]string{"OAuth": "oauth"}

func LowerCamelCase(name string) string {
	return CamelCase(name, false, true)
}

// CamelCase produces the CamelCase version of the given string. It removes any
// non letter and non digit character.
//
// If firstUpper is true the first letter of the string is capitalized else
// the first letter is in lowercase.
//
// If acronym is true and a part of the string is a common acronym
// then it keeps the part capitalized (firstUpper = true)
// (e.g. APIVersion) or lowercase (firstUpper = false) (e.g. apiVersion).
func CamelCase(name string, firstUpper, acronym bool) string {
	if name == "" {
		return ""
	}

	runes := []rune(name)
	// remove trailing invalid identifiers (makes code below simpler)
	runes = removeTrailingInvalid(runes)

	// all characters are invalid
	if len(runes) == 0 {
		return ""
	}

	wordStart, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		var correct bool
		// remove leading invalid identifiers
		runes = removeInvalidAtIndex(i, runes)

		runes, correct = checkAndCorrectRune(i, runes)

		i++

		if correct {
			continue
		}

		runes = updateWordInRuneWithCamelCase(wordStart, i, runes, firstUpper, acronym)
		wordStart = i
	}

	return string(runes)
}

func checkAndCorrectRune(i int, runes []rune) ([]rune, bool) {
	correct := true

	switch {
	case i+1 == len(runes):
		correct = false

	case !validIdentifier(runes[i]):
		runes = append(runes[:i], runes[i+1:]...)

	case runes[i+1] == '_':
		// underscore; shift the remainder forward over any run of underscores
		correct = false
		j := 1

		for i+j+1 < len(runes) && runes[i+j+1] == '_' {
			j++
		}

		copy(runes[i+1:], runes[i+j+1:])
		runes = runes[:len(runes)-j]

	case isLower(runes[i]) && !isLower(runes[i+1]):
		correct = false
	}

	return runes, correct
}

func updateWordInRuneWithCamelCase(wordStart, i int, runes []rune, firstUpper, acronym bool) []rune {
	word := string(runes[wordStart:i])
	// is it one of our initialisms?

	acronyms := strings.ToUpper(word)

	switch {
	case commonInitialisms[acronyms]:
		{
			acronyms = prepareAcronyms(acronyms, wordStart, firstUpper, acronym)

			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[wordStart:], []rune(acronyms))
		}
	case wordStart > 0 && strings.ToLower(word) == word:
		// already all lowercase, and not the first word, so uppercase the first character.
		runes[wordStart] = unicode.ToUpper(runes[wordStart])
	case wordStart == 0 && strings.ToLower(word) == word && firstUpper:
		runes[wordStart] = unicode.ToUpper(runes[wordStart])
	}

	if wordStart == 0 && !firstUpper {
		runes[wordStart] = unicode.ToLower(runes[wordStart])
	}

	return runes
}

func prepareAcronyms(acronyms string, wordStart int, firstUpper, acronym bool) string {
	switch {
	case firstUpper && acronym:
		// u is already in upper case. Nothing to do here.
	case firstUpper && !acronym:
		acronyms = strings.ToLower(acronyms)
	case wordStart > 0 && !acronym:
		acronyms = strings.ToLower(acronyms)
	case wordStart == 0:
		acronyms = strings.ToLower(acronyms)
	}

	return acronyms
}

// SnakeCase produces the snake_case version of the given CamelCase string.
// News    => news
// OldNews => old_news
// CNNNews => cnn_news
func SnakeCase(name string) string {
	// Special handling for single "words" starting with multiple upper case letters
	for u, l := range toLower {
		name = strings.ReplaceAll(name, u, l)
	}

	// Remove leading and trailing blank spaces and replace any blank spaces in
	// between with a single underscore
	name = strings.Join(strings.Fields(name), "_")

	// Special handling for dashes to convert them into underscores
	name = strings.ReplaceAll(name, "-", "_")

	var b bytes.Buffer

	ln := len(name)
	if ln == 0 {
		return ""
	}

	n := rune(name[0])
	b.WriteRune(unicode.ToLower(n))

	var isLower, isUnder bool

	lastLower, lastUnder := false, false

	for i := 1; i < ln; i++ {
		r := rune(name[i])
		isLower = unicode.IsLower(r) && unicode.IsLetter(r) || unicode.IsDigit(r)
		isUnder = r == '_'

		addUnderliningToStr(isLower, lastLower, isUnder, lastUnder, &b, name, i, ln <= i+1)

		b.WriteRune(unicode.ToLower(r))

		lastLower = isLower
		lastUnder = isUnder
	}

	return b.String()
}

func addUnderliningToStr(isLower, lastLower, isUnder, lastUnder bool, b *bytes.Buffer, name string, i int,
	isLast bool,
) {
	if isLower || isUnder {
		return
	}

	if lastLower && !lastUnder {
		b.WriteRune('_')
	} else if !isLast {
		rn := rune(name[i+1])
		if unicode.IsLower(rn) && rn != '_' && !lastUnder {
			b.WriteRune('_')
		}
	}
}

// KebabCase produces the kebab-case version of the given CamelCase string.
func KebabCase(name string) string {
	name = SnakeCase(name)
	ln := len(name)

	if name[ln-1] == '_' {
		name = name[:ln-1]
	}

	return strings.ReplaceAll(name, "_", "-")
}

// WrapText produces lines with text capped at maxChars
// it will keep words intact and respects newlines.
func WrapText(text string, maxChars int) string {
	res := ""
	lines := strings.Split(text, "\n")

	for _, v := range lines {
		runes := []rune(strings.TrimSpace(v))

		for l := len(runes); l >= 0; l = len(runes) {
			if maxChars >= l {
				res = res + string(runes) + "\n"

				break
			}

			i := runeSpacePosRev(runes[:maxChars])
			if i == 0 {
				i = runeSpacePos(runes)
			}

			res = res + string(runes[:i]) + "\n"

			if l == i {
				break
			}

			runes = runes[i+1:]
		}
	}

	return res[:len(res)-1]
}

func runeSpacePosRev(r []rune) int {
	for i := len(r) - 1; i > 0; i-- {
		if unicode.IsSpace(r[i]) {
			return i
		}
	}

	return 0
}

func runeSpacePos(r []rune) int {
	for i := 0; i < len(r); i++ {
		if unicode.IsSpace(r[i]) {
			return i
		}
	}

	return len(r)
}

// isLower returns true if the character is considered a lower case character
// when transforming word into CamelCase.
func isLower(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLower(r)
}

// validIdentifier returns true if the rune is a letter or number.
func validIdentifier(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// removeTrailingInvalid removes trailing invalid identifiers from runes.
func removeTrailingInvalid(runes []rune) []rune {
	valid := len(runes) - 1
	for ; valid >= 0 && !validIdentifier(runes[valid]); valid-- {
	}

	return runes[0 : valid+1]
}

// removeInvalidAtIndex removes consecutive invalid identifiers from runes starting at index i.
func removeInvalidAtIndex(i int, runes []rune) []rune {
	valid := i
	for ; valid < len(runes) && !validIdentifier(runes[valid]); valid++ {
	}

	return append(runes[:i], runes[valid:]...)
}

// common words who need to keep their
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JMES":  true,
	"JSON":  true,
	"JWT":   true,
	"LHS":   true,
	"OK":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
