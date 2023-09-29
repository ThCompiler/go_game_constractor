package constants

type (
	NumberForm int8
	WordForm   int8
)

const (
	SINGULAR_WORD = WordForm(0)
	PLURAL_WORD   = WordForm(1)
)

const (
	FIRST_FORM  = NumberForm(0) // when digit is one (1)
	SECOND_FORM = NumberForm(1) // when digit between 2 and 4
	THIRD_FORM  = NumberForm(2) // when digit between 4 and 9
)

const (
	CountWordForms       = 2
	CountNumberNameForms = 3
)
