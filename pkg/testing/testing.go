package testing

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MockTestFunction func(ctrl *gomock.Controller) []interface{}

type TestFunction func(args ...interface{}) []interface{}

type TestCase struct {
	Name      string
	Args      []interface{}
	Expected  TestExpected
	InitMocks MockTestFunction
}

type TestCasesSuite struct {
	suite.Suite
}

func (s *TestCasesSuite) RunTest(fun TestFunction, cases ...TestCase) {
	RunTest(s.T(), fun, cases...)
}

func checkExpected(t *testing.T, res []interface{}, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HaveError() {
		checkWithError(t, &res, expected, caseName)
	} else {
		require.Equalf(t, len(res), len(expected.ExpectedReturns),
			"Testcase with name: %s, different len of expected and gotten return values", caseName)
	}

	for i, expected := range expected.ExpectedReturns {
		assert.EqualValuesf(t, expected, res[i], "Testcase with name: %s", caseName)
	}
}

func checkWithError(t *testing.T, res *[]interface{}, expected TestExpected, caseName string) {
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
		checkErrorCorrect(t, gottenError, expected, caseName)
	}

	*res = (*res)[:size-1]
}

func checkErrorCorrect(t *testing.T, gottenError error, expected TestExpected, caseName string) {
	t.Helper()

	if expected.MustErrorExpected().Error == nil {
		assert.NoError(t, gottenError,
			"Testcase with name: %s", caseName)
	} else {
		assert.EqualError(t, gottenError, expected.MustErrorExpected().Error.Error(),
			"Testcase with name: %s", caseName)
	}
}

func checkPanicErrorCorrect(t *testing.T, msg any, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HavePanicError() {
		assert.EqualValuesf(t, msg, expected.MustPanicErrorExpected().Msg, "Testcase with name: %s", caseName)

		return
	}

	if err, is := msg.(error); is {
		assert.Failf(t, "Panic error testcase: ", "%s %s", caseName, err)
	}

	assert.Failf(t, "Panic error testcase: ", "%s %v", caseName, msg)
}

func runTestCase(t *testing.T, test TestCase, fun TestFunction, ctrl *gomock.Controller) {
	t.Helper()

	defer func(t *testing.T, test TestCase) {
		t.Helper()

		if r := recover(); r != nil {
			checkPanicErrorCorrect(t, r, test.Expected, test.Name)
		}
	}(t, test)

	args := test.Args
	if test.InitMocks != nil {
		mocks := test.InitMocks(ctrl)

		args = append(mocks, args...)
	}
	res := fun(args...)

	checkExpected(t, res, test.Expected, test.Name)
}

func RunTest(t *testing.T, fun TestFunction, cases ...TestCase) {
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
