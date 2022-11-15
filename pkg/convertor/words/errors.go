package words

import "github.com/pkg/errors"

var (
    UnknownLanguageError = errors.New("There is no resource for such a language")
)
