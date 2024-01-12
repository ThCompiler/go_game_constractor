package declension

type Declension string

const (
	NOMINATIVE    = Declension(`nominative`)    // именительный
	GENITIVE      = Declension(`genitive`)      // родительный
	DATIVE        = Declension(`dative`)        // дательный
	ACCUSATIVE    = Declension(`accusative`)    // винительный
	INSTRUMENTAL  = Declension(`instrumental`)  // творительный
	PREPOSITIONAL = Declension(`prepositional`) // предложный
)
