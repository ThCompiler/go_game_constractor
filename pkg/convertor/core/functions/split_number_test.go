package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SplitNumberIntoThreesSuite struct {
	suite.Suite
}

func (s *SplitNumberIntoThreesSuite) TestEmptyNumber() {
	res := SplitNumberIntoThrees("")
	assert.Empty(s.T(), res)
}

func (s *SplitNumberIntoThreesSuite) TestNumberWithOneDigit() {
	res := SplitNumberIntoThrees("1")
	assert.EqualValues(s.T(),
		[]objects.RuneDigitTriplet{
			{
				Units:    '1',
				Dozens:   '0',
				Hundreds: '0',
			},
		},
		res,
	)
}

func (s *SplitNumberIntoThreesSuite) TestNumberWithTwoDigit() {
	res := SplitNumberIntoThrees("12")
	assert.EqualValues(s.T(),
		[]objects.RuneDigitTriplet{
			{
				Units:    '2',
				Dozens:   '1',
				Hundreds: '0',
			},
		},
		res,
	)
}
func (s *SplitNumberIntoThreesSuite) TestNumberWithThreeDigit() {
	res := SplitNumberIntoThrees("123")
	assert.EqualValues(s.T(),
		[]objects.RuneDigitTriplet{
			{
				Units:    '3',
				Dozens:   '2',
				Hundreds: '1',
			},
		},
		res,
	)
}

func (s *SplitNumberIntoThreesSuite) TestNumberWithTripletAndTwoDigit() {
	res := SplitNumberIntoThrees("1234")
	assert.EqualValues(s.T(),
		[]objects.RuneDigitTriplet{
			{
				Units:    '1',
				Dozens:   '0',
				Hundreds: '0',
			},
			{
				Units:    '4',
				Dozens:   '3',
				Hundreds: '2',
			},
		},
		res,
	)
}

func (s *SplitNumberIntoThreesSuite) TestAnyChars() {
	assert.NotPanics(s.T(), func() {
		_ = SplitNumberIntoThrees("adsc")
	})
}

func TestSplitNumberIntoThreesSuite(t *testing.T) {
	suite.Run(t, new(SplitNumberIntoThreesSuite))
}
