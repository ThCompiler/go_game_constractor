package objects

import (
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/suite"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "NumericDigitTriplet" functions
//----------------------------------------------------------------------------------------------------------------------

type NumericDigitTripletSuite struct {
	ts.TestCasesSuite
}

func (s *NumericDigitTripletSuite) TestIsZeros() {
	s.RunTest(
		func(ndt NumericDigitTriplet) bool {
			return ndt.IsZeros()
		},
		ts.TestCase{
			Name:     "zero triplet",
			Args:     ts.TTA(NumericDigitTriplet{0, 0, 0}),
			Expected: ts.TTVE(true),
		},
		ts.TestCase{
			Name:     "zero in units of triplet",
			Args:     ts.TTA(NumericDigitTriplet{2, 3, 0}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "zero in dozens of triplet",
			Args:     ts.TTA(NumericDigitTriplet{2, 0, 1}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "zero in hundreds of triplet",
			Args:     ts.TTA(NumericDigitTriplet{0, 2, 1}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "no zero in triplet",
			Args:     ts.TTA(NumericDigitTriplet{3, 2, 1}),
			Expected: ts.TTVE(false),
		},
	)
}

func (s *NumericDigitTripletSuite) TestToRune() {
	s.RunTest(
		func(ndt NumericDigitTriplet) RuneDigitTriplet {
			return ndt.ToRune()
		},
		ts.TestCase{
			Name:     "correct number triplet with zeros",
			Args:     ts.TTA(NumericDigitTriplet{0, 0, 0}),
			Expected: ts.TTVE(RuneDigitTriplet{'0', '0', '0'}),
		},
		ts.TestCase{
			Name:     "correct number triplet with nines",
			Args:     ts.TTA(NumericDigitTriplet{9, 9, 9}),
			Expected: ts.TTVE(RuneDigitTriplet{'9', '9', '9'}),
		},
		ts.TestCase{
			Name:     "incorrect number triplet with tens",
			Args:     ts.TTA(NumericDigitTriplet{10, 10, 10}),
			Expected: ts.TTVE(RuneDigitTriplet{0, 0, 0}),
		},
		ts.TestCase{
			Name:     "incorrect number triplet with big numbers",
			Args:     ts.TTA(NumericDigitTriplet{127, 127, 127}),
			Expected: ts.TTVE(RuneDigitTriplet{0, 0, 0}),
		},
	)
}

func TestNumericDigitTripletSuite(t *testing.T) {
	suite.Run(t, new(NumericDigitTripletSuite))
}

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "RuneDigitTriplet" functions
//----------------------------------------------------------------------------------------------------------------------

type RuneDigitTripletSuite struct {
	ts.TestCasesSuite
}

func (s *RuneDigitTripletSuite) TestIsZeros() {
	s.RunTest(
		func(ndt RuneDigitTriplet) bool {
			return ndt.IsZeros()
		},
		ts.TestCase{
			Name:     "zero triplet",
			Args:     ts.TTA(RuneDigitTriplet{'0', '0', '0'}),
			Expected: ts.TTVE(true),
		},
		ts.TestCase{
			Name:     "zero in units of triplet",
			Args:     ts.TTA(RuneDigitTriplet{'2', '3', '0'}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "zero in dozens of triplet",
			Args:     ts.TTA(RuneDigitTriplet{'2', '0', '1'}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "zero in hundreds of triplet",
			Args:     ts.TTA(RuneDigitTriplet{'0', '2', '1'}),
			Expected: ts.TTVE(false),
		},
		ts.TestCase{
			Name:     "no Zero in triplet",
			Args:     ts.TTA(RuneDigitTriplet{'3', '2', '1'}),
			Expected: ts.TTVE(false),
		},
	)
}

func (s *RuneDigitTripletSuite) TestToNumeric() {
	s.RunTest(
		func(rdt RuneDigitTriplet) NumericDigitTriplet {
			return rdt.ToNumeric()
		},
		ts.TestCase{
			Name:     "correct rune triplet with zeros",
			Args:     ts.TTA(RuneDigitTriplet{'0', '0', '0'}),
			Expected: ts.TTVE(NumericDigitTriplet{0, 0, 0}),
		},
		ts.TestCase{
			Name:     "correct rune triplet with nines",
			Args:     ts.TTA(RuneDigitTriplet{'9', '9', '9'}),
			Expected: ts.TTVE(NumericDigitTriplet{9, 9, 9}),
		},
		ts.TestCase{
			Name:     "incorrect rune triplet with other symbols",
			Args:     ts.TTA(RuneDigitTriplet{'a', 'a', 'a'}),
			Expected: ts.TTVE(NumericDigitTriplet{-1, -1, -1}),
		},
		ts.TestCase{
			Name:     "incorrect rune triplet with big numbers of rune symbol",
			Args:     ts.TTA(RuneDigitTriplet{127, 127, 127}),
			Expected: ts.TTVE(NumericDigitTriplet{-1, -1, -1}),
		},
	)
}

func TestRuneDigitTripletSuite(t *testing.T) {
	suite.Run(t, new(RuneDigitTripletSuite))
}
