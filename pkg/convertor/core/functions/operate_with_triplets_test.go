package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "RemoveZeroTripletFromBeginning" function
//----------------------------------------------------------------------------------------------------------------------

type RemoveZeroTripletFromBeginningSuite struct {
	ts.TestCasesSuite
}

func (s *RemoveZeroTripletFromBeginningSuite) TestEmptyTriplet() {
	triplet := make([]objects.RuneDigitTriplet, 0)
	res := RemoveZeroTripletFromBeginning(triplet)
	assert.Equal(s.T(), triplet, res)
}

func (s *RemoveZeroTripletFromBeginningSuite) TestNilTriplet() {
	res := RemoveZeroTripletFromBeginning(nil)
	assert.Equal(s.T(), []objects.RuneDigitTriplet(nil), res)
}

func (s *RemoveZeroTripletFromBeginningSuite) TestWithoutZeroTriplet() {
	arg := []objects.RuneDigitTriplet{
		{'0', '1', '2'},
		{'0', '1', '2'},
		{'2', '1', '2'},
		{'4', '1', '2'},
	}
	res := RemoveZeroTripletFromBeginning(arg)
	assert.Equal(s.T(), arg, res)
}

func (s *RemoveZeroTripletFromBeginningSuite) TestWithZeroTripletAtBegin() {
	s.RunTest(
		RemoveZeroTripletFromBeginning,
		ts.TestCase{
			Name:     "Only zero triplet",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'0', '0', '0'}}),
			Expected: ts.TTVE([]objects.RuneDigitTriplet{}),
		},
		ts.TestCase{
			Name:     "Only zero triplets in beginning",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'0', '0', '0'}, {'0', '0', '0'}, {'2', '3', '4'}}),
			Expected: ts.TTVE([]objects.RuneDigitTriplet{{'2', '3', '4'}}),
		},
		ts.TestCase{
			Name: "Many zero triplets",
			Args: ts.TTA([]objects.RuneDigitTriplet{{'0', '0', '0'}, {'0', '0', '0'}, {'2', '3', '4'},
				{'0', '0', '0'}, {'2', '3', '4'}, {'0', '0', '0'}}),
			Expected: ts.TTVE([]objects.RuneDigitTriplet{{'2', '3', '4'}, {'0', '0', '0'},
				{'2', '3', '4'}, {'0', '0', '0'}}),
		},
	)
}

func (s *RemoveZeroTripletFromBeginningSuite) TestWithZeroTripletNotAtBegin() {
	args := []objects.RuneDigitTriplet{{'2', '3', '4'}, {'0', '0', '0'}, {'2', '3', '4'}, {'0', '0', '0'}}
	res := RemoveZeroTripletFromBeginning(args)
	assert.Equal(s.T(), args, res)
}

func TestRemoveZeroTripletFromBeginningSuite(t *testing.T) {
	suite.Run(t, new(RemoveZeroTripletFromBeginningSuite))
}

//----------------------------------------------------------------------------------------------------------------------
// Tests suite for "IndexByEndOfLastNotZeroTriplet" function
//----------------------------------------------------------------------------------------------------------------------

type IndexByEndOfLastNotZeroTripletSuite struct {
	ts.TestCasesSuite
}

func (s *IndexByEndOfLastNotZeroTripletSuite) TestEmptyTriplet() {
	triplet := make([]objects.RuneDigitTriplet, 0)
	res := IndexByEndOfLastNotZeroTriplet(triplet)
	assert.Equal(s.T(), -1, res)
}

func (s *IndexByEndOfLastNotZeroTripletSuite) TestNilTriplet() {
	res := IndexByEndOfLastNotZeroTriplet(nil)
	assert.Equal(s.T(), -1, res)
}

func (s *IndexByEndOfLastNotZeroTripletSuite) TestWithoutZeroTriplet() {
	s.RunTest(
		IndexByEndOfLastNotZeroTriplet,
		ts.TestCase{
			Name:     "One triplet",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'1', '0', '2'}}),
			Expected: ts.TTVE(0),
		},
		ts.TestCase{
			Name:     "Many triplets",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'1', '0', '2'}, {'5', '1', '0'}, {'2', '3', '4'}}),
			Expected: ts.TTVE(0),
		},
	)
}

func (s *IndexByEndOfLastNotZeroTripletSuite) TestWithZeroTripletAtBegin() {
	s.RunTest(
		IndexByEndOfLastNotZeroTriplet,
		ts.TestCase{
			Name:     "Only zero triplet",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'0', '0', '0'}}),
			Expected: ts.TTVE(-1),
		},
		ts.TestCase{
			Name: "Some zero triplet at end",
			Args: ts.TTA([]objects.RuneDigitTriplet{{'2', '3', '4'}, {'2', '3', '4'},
				{'0', '0', '0'}, {'0', '0', '0'}}),
			Expected: ts.TTVE(2),
		},
		ts.TestCase{
			Name:     "Some zero triplet at middle of array",
			Args:     ts.TTA([]objects.RuneDigitTriplet{{'2', '3', '4'}, {'0', '0', '0'}, {'2', '3', '4'}}),
			Expected: ts.TTVE(0),
		},
	)
}

func TestIndexByEndOfLastNotZeroTripletSuite(t *testing.T) {
	suite.Run(t, new(IndexByEndOfLastNotZeroTripletSuite))
}
