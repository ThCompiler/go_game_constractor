package converter

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/option"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/currency"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConvertSuite struct {
	ts.TestCasesSuite
	rs *ru.Russian
}

func (s *ConvertSuite) SetupTest() {
	s.rs = ru.LoadRussian()
}

func (s *ConvertSuite) TestNegativeNumber() {
	s.RunTest(
		Convert[string],
		ts.TestCase{
			Name:     "-2",
			Args:     ts.TTA("-2", option.Default(s.rs)),
			Expected: ts.TTVEWNE("минус две целых"),
		},
	)
}

func (s *ConvertSuite) TestPositiveNumber() {
	s.RunTest(
		Convert[string],
		ts.TestCase{
			Name:     "2",
			Args:     ts.TTA("2", option.Default(s.rs)),
			Expected: ts.TTVEWNE("две целых"),
		},
		ts.TestCase{
			Name:     "2000",
			Args:     ts.TTA("2000", option.Default(s.rs)),
			Expected: ts.TTVEWNE("две тысячи целых"),
		},
		ts.TestCase{
			Name:     "1000",
			Args:     ts.TTA("1000", option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна тысяча целых"),
		},
		ts.TestCase{
			Name:     "7000",
			Args:     ts.TTA("7000", option.Default(s.rs)),
			Expected: ts.TTVEWNE("семь тысяч целых"),
		},
	)
}

func (s *ConvertSuite) TestDecimalNumber() {
	s.RunTest(
		Convert[float64],
		ts.TestCase{
			Name:     "1.0001",
			Args:     ts.TTA(1.0001, option.NewOptionsBuilder(s.rs).WithoutRound().Build()),
			Expected: ts.TTVEWNE("одна целая одна десятитысячная"),
		},
		ts.TestCase{
			Name:     "1.12",
			Args:     ts.TTA(1.12, option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна целая двенадцать сотых"),
		},
		ts.TestCase{
			Name:     "0.01",
			Args:     ts.TTA(0.01, option.Default(s.rs)),
			Expected: ts.TTVEWNE("ноль целых одна сотая"),
		},
		ts.TestCase{
			Name:     "1.00",
			Args:     ts.TTA(1.00, option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна целая"),
		},
	)
}

func (s *ConvertSuite) TestFractionalNumber() {
	s.RunTest(
		Convert[string],
		ts.TestCase{
			Name:     "1/0",
			Args:     ts.TTA("1/0", option.NewOptionsBuilder(s.rs).Build()),
			Expected: ts.TTVEWNE("одна нулевая"),
		},
		ts.TestCase{
			Name:     "1/",
			Args:     ts.TTA("1/", option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна нулевая"),
		},
		ts.TestCase{
			Name:     "0/2000",
			Args:     ts.TTA("0/2000", option.Default(s.rs)),
			Expected: ts.TTVEWNE("ноль двухтысячных"),
		},
		ts.TestCase{
			Name:     "1/2",
			Args:     ts.TTA("1/2", option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна вторая"),
		},
		ts.TestCase{
			Name:     "1/2000",
			Args:     ts.TTA("1/2000", option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна двухтысячная"),
		},
		ts.TestCase{
			Name:     "1/2001",
			Args:     ts.TTA("1/2001", option.Default(s.rs)),
			Expected: ts.TTVEWNE("одна две тысячи первая"),
		},
	)
}

func (s *ConvertSuite) TestNumberWithCurrency() {
	s.RunTest(
		Convert[string],
		ts.TestCase{
			Name:     "1/2",
			Args:     ts.TTA("1/2", option.NewOptionsBuilder(s.rs.WithOptions(ru.WithCurrency(currency.RUB))).Build()),
			Expected: ts.TTVEWNE("одна вторая рубля"),
		},
		ts.TestCase{
			Name:     "1000",
			Args:     ts.TTA("1000", option.NewOptionsBuilder(s.rs.WithOptions(ru.WithCurrency(currency.RUB))).Build()),
			Expected: ts.TTVEWNE("одна тысяча рублей"),
		},
		ts.TestCase{
			Name:     "1.5",
			Args:     ts.TTA("1.5", option.NewOptionsBuilder(s.rs.WithOptions(ru.WithCurrency(currency.RUB))).Build()),
			Expected: ts.TTVEWNE("один рубль пять копеек"),
		},
	)
}

func TestConvert(t *testing.T) {
	suite.Run(t, new(ConvertSuite))
}
