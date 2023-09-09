package testing

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MockTestFunction func(ctrl *gomock.Controller) []interface{}

type ArrangeFunction func(args ...interface{}) []interface{}

type ActFunction func(args ...interface{}) []interface{}

type TestCase struct {
	Name        string
	Args        []interface{}
	Expected    TestExpected
	InitMocks   MockTestFunction
	ArrangeCase ArrangeFunction
}

type TestCasesSuite struct {
	suite.Suite
}

func (s *TestCasesSuite) RunTest(fun ActFunction, cases ...TestCase) {
	RunTest(s.T(), fun, cases...)
}

func assertCase(t *testing.T, res []interface{}, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HaveError() {
		assertWithError(t, &res, expected, caseName)
	} else {
		require.Equalf(t, len(expected.ExpectedReturns), len(res),
			"Testcase with name: %s, different len of expected and gotten return values", caseName)
	}

	for i, expected := range expected.ExpectedReturns {
		assert.EqualValuesf(t, expected, res[i], "Testcase with name: %s", caseName)
	}
}

func assertWithError(t *testing.T, res *[]interface{}, expected TestExpected, caseName string) {
	t.Helper()

	size := len(*res)
	require.NotZerof(t, size, "Testcase with name: %s, return nothing, but wait return error", caseName)
	gottenError, ok := (*res)[size-1].(error)

	if !ok && gottenError != nil {
		require.Failf(t, "Last value not error, but expected error",
			"Testcase with name: %s", caseName)
	}

	if expected.MustErrorExpected().CheckError {
		assert.Error(t, gottenError, "Testcase with name: %s", caseName)
	} else {
		checkForCorrectnessError(t, gottenError, expected, caseName)
	}

	*res = (*res)[:size-1]
}

func checkForCorrectnessError(t *testing.T, gottenError error, expected TestExpected, caseName string) {
	t.Helper()

	if expected.MustErrorExpected().Error == nil {
		assert.NoError(t, gottenError,
			"Testcase with name: %s", caseName)
	} else {
		assert.ErrorIs(t, gottenError, expected.MustErrorExpected().Error,
			"Testcase with name: %s", caseName)
	}
}

func checkForCorrectnessPanicError(t *testing.T, msg any, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HavePanicError() {
		assert.EqualValuesf(t, expected.MustPanicErrorExpected().Msg, msg, "Testcase with name: %s", caseName)

		return
	}

	if err, is := msg.(error); is {
		assert.Failf(t, "Panic error testcase: ", "%s %s", caseName, err)
	}

	assert.Failf(t, "Panic error testcase: ", "%s %v", caseName, msg)
}

func runTestCase(t *testing.T, test TestCase, fun ActFunction, ctrl *gomock.Controller) {
	t.Helper()

	defer func(t *testing.T, test TestCase) {
		t.Helper()

		if r := recover(); r != nil {
			checkForCorrectnessPanicError(t, r, test.Expected, test.Name)
		}
	}(t, test)

	args := test.Args

	if test.InitMocks != nil {
		mocks := test.InitMocks(ctrl)
		args = append(mocks, args...)
	}

	if test.ArrangeCase != nil {
		args = test.ArrangeCase(args)
	}

	res := fun(args...)

	assertCase(t, res, test.Expected, test.Name)
}

func RunTest(t *testing.T, fun ActFunction, cases ...TestCase) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, cs := range cases {
		cs := cs
		t.Run(cs.Name, func(t *testing.T) {
			runTestCase(t, cs, fun, ctrl)
		})
	}
}

func ToTestArgs(args ...interface{}) []interface{} {
	return args
}

func TTA(args ...interface{}) []interface{} {
	return ToTestArgs(args...)
}
