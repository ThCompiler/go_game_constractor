package testing

type TestPanicErrorExpected struct {
	Msg interface{}
}

type TestErrorExpected struct {
	CheckError bool // mean only check error without compare
	Error      error
}

type TestExpected struct {
	PanicError      *TestPanicErrorExpected
	Error           *TestErrorExpected
	ExpectedReturns []interface{}
}

func (te *TestExpected) HaveError() bool {
	return te.Error != nil
}

func (te *TestExpected) HavePanicError() bool {
	return te.PanicError != nil
}

func (te *TestExpected) MustErrorExpected() TestErrorExpected {
	if te.Error == nil {
		panic("Expected error, but not")
	}

	return *te.Error
}

func (te *TestExpected) MustPanicErrorExpected() TestPanicErrorExpected {
	if te.PanicError == nil {
		panic("Expected panic error, but not")
	}

	return *te.PanicError
}

func ToTestValuesExpected(expedites ...interface{}) TestExpected {
	return TestExpected{
		PanicError:      nil,
		Error:           nil,
		ExpectedReturns: expedites,
	}
}

func TTVE(expedites ...interface{}) TestExpected {
	return ToTestValuesExpected(expedites...)
}

func ToTestValuesExpectedWithNilError(expedites ...interface{}) TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			Error: nil,
		},
		ExpectedReturns: expedites,
	}
}

func TTVEWNE(expedites ...interface{}) TestExpected {
	return ToTestValuesExpectedWithNilError(expedites...)
}

func ToTestErrorExpected(err error) TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			Error: err,
		},
		ExpectedReturns: nil,
	}
}

func TTEE(err error) TestExpected {
	return ToTestErrorExpected(err)
}

func ToTestCheckErrorExpected() TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			CheckError: true,
		},
		ExpectedReturns: nil,
	}
}

func TTCEE() TestExpected {
	return ToTestCheckErrorExpected()
}

func ToTestPanicErrorExpected(msg interface{}) TestExpected {
	return TestExpected{
		PanicError: &TestPanicErrorExpected{
			Msg: msg,
		},
		Error:           nil,
		ExpectedReturns: nil,
	}
}

func TTPEE(msg interface{}) TestExpected {
	return ToTestPanicErrorExpected(msg)
}

func ToTestExpected(checkError bool, err error, withPanic bool, panicMsg interface{},
	expedites ...interface{},
) TestExpected {
	if withPanic {
		return ToTestPanicErrorExpected(panicMsg)
	}

	if !checkError && err != nil {
		return ToTestValuesExpected(expedites...)
	}

	if checkError {
		return ToTestCheckErrorExpected()
	}

	return ToTestErrorExpected(err)
}

func TTE(checkError bool, err error, withPanic bool, panicMsg string,
	expedites ...interface{},
) TestExpected {
	return ToTestExpected(checkError, err, withPanic, panicMsg, expedites...)
}
