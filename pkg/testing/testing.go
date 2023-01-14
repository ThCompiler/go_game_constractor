package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestExpected struct {
	CheckError      bool // mean only check error without compare
	HaveError       bool // in returning values end of values is error
	ExpectedErr     error
	ExpectedReturns []interface{}
}

type TestFunction func(args ...interface{}) []interface{}

type TestCase struct {
	Name     string
	Args     []interface{}
	Expected TestExpected
}

type TestCasesSuite struct {
	suite.Suite
}

func (s *TestCasesSuite) RunTest(fun TestFunction, cases ...TestCase) {
	for _, cs := range cases {
		cs := cs
		s.T().Run(cs.Name, func(t *testing.T) {
			runTestCase(t, cs, fun)
		})
	}
}

func checkExpected(t *testing.T, res []interface{}, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HaveError {
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

	if expected.CheckError {
		assert.Error(t, gottenError, "Testcase with name: %s", caseName)
	} else {
		checkErrorCorrect(t, gottenError, expected, caseName)
	}

	*res = (*res)[:size-1]
}

func checkErrorCorrect(t *testing.T, gottenError error, expected TestExpected, caseName string) {
	t.Helper()

	if expected.ExpectedErr == nil {
		assert.NoError(t, gottenError,
			"Testcase with name: %s", caseName)
	} else {
		assert.EqualError(t, gottenError, expected.ExpectedErr.Error(),
			"Testcase with name: %s", caseName)
	}
}

func runTestCase(t *testing.T, test TestCase, fun TestFunction) {
	t.Helper()

	defer func(t *testing.T) {
		t.Helper()

		if r := recover(); r != nil {
			if err, is := r.(error); is {
				assert.Failf(t, "Error testcase: ", "%s %s", test.Name, err)
			}

			assert.Failf(t, "Error testcase: ", "%s %v", test.Name, r)
		}
	}(t)

	res := fun(test.Args...)

	checkExpected(t, res, test.Expected, test.Name)
}

func RunTest(t *testing.T, fun TestFunction, cases ...TestCase) {
	t.Helper()

	for _, cs := range cases {
		cs := cs
		t.Run(cs.Name, func(t *testing.T) {
			runTestCase(t, cs, fun)
		})
	}
}

func ToTestArgs(args ...interface{}) []interface{} {
	return args
}

func ToTestValuesExpected(expedites ...interface{}) TestExpected {
	return TestExpected{
		ExpectedReturns: expedites,
	}
}

func ToTestErrorExpected(err error) TestExpected {
	return TestExpected{
		ExpectedErr: err,
		HaveError:   true,
	}
}

func ToTestCheckErrorExpected() TestExpected {
	return TestExpected{
		HaveError:  true,
		CheckError: true,
	}
}

func ToTestExpected(checkError bool, err error, expedites ...interface{}) TestExpected {
	if !checkError && err != nil {
		return ToTestValuesExpected(expedites...)
	}

	if checkError {
		return ToTestCheckErrorExpected()
	}

	return ToTestErrorExpected(err)
}
