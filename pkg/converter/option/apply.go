package option

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
)

type OptionsBuilder struct {
	op Options
}

func NewOptionsBuilder(lang words.Language) *OptionsBuilder {
	return &OptionsBuilder{
		op: Default(lang),
	}
}

func (ob *OptionsBuilder) Build() Options {
	return ob.op
}

func (ob *OptionsBuilder) WithRound(round int64) *OptionsBuilder {
	ob.op.RoundNumber = round
	return ob
}

func (ob *OptionsBuilder) WithoutRound() *OptionsBuilder {
	ob.op.RoundNumber = constants.NoRoundIndicator
	return ob
}

func (ob *OptionsBuilder) WithMinusAsWord() *OptionsBuilder {
	ob.op.ConvertMinusSignToWord = true
	return ob
}

func (ob *OptionsBuilder) WithMinusAsSymbol() *OptionsBuilder {
	ob.op.ConvertMinusSignToWord = false
	return ob
}

func (ob *OptionsBuilder) WithShowCurrency(showing NumberPart) *OptionsBuilder {
	ob.op.ShowCurrency = showing
	return ob
}

func (ob *OptionsBuilder) WithShowNumberParts(showing NumberPart) *OptionsBuilder {
	ob.op.ShowNumberParts = showing
	return ob
}

func (ob *OptionsBuilder) WithConvertNumberToWords(showing NumberPart) *OptionsBuilder {
	ob.op.ConvertNumberToWords = showing
	return ob
}

func (ob *OptionsBuilder) WithShowingZeroInDecimalPart() *OptionsBuilder {
	ob.op.ShowZeroInDecimalPart = true
	return ob
}

func (ob *OptionsBuilder) WithoutShowingZeroInDecimalPart() *OptionsBuilder {
	ob.op.ShowZeroInDecimalPart = false
	return ob
}

func (ob *OptionsBuilder) WithUppercaseFirstSymbol() *OptionsBuilder {
	ob.op.AddUppercaseToFirstSymbol = true
	return ob
}

func (ob *OptionsBuilder) WithLowercaseFirstSymbol() *OptionsBuilder {
	ob.op.AddUppercaseToFirstSymbol = false
	return ob
}
