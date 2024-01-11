package declension

import "github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"

const (
	NOMINATIVE    = words.Declension(`nominative`)    // именительный
	GENITIVE      = words.Declension(`genitive`)      // родительный
	DATIVE        = words.Declension(`dative`)        // дательный
	ACCUSATIVE    = words.Declension(`accusative`)    // винительный
	INSTRUMENTAL  = words.Declension(`instrumental`)  // творительный
	PREPOSITIONAL = words.Declension(`prepositional`) // предложный
)
