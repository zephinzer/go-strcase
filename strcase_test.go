package strcase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type strcaseTest struct {
	suite.Suite
	basicTestCases         []string
	basicExtendedTestCases []string
	edgeCases              []string
	numbersTestCases       []string
}

func Test_strcase(t *testing.T) {
	suite.Run(t, &strcaseTest{
		basicTestCases: []string{
			"hello world",
			"hello-world",
			"hello_world",
			"helloWorld",
			"HelloWorld",
		},
		basicExtendedTestCases: []string{
			"hello a world",
			"hello-a-world",
			"hello_a_world",
			"helloAWorld",
			"HelloAWorld",
		},
		numbersTestCases: []string{
			"catch22 statement",
			"catch22-statement",
			"catch22_statement",
			"catch22statement",
			"catch22Statement",
		},
	})
}

func (s strcaseTest) TestToCamel_basic() {
	expected := "helloWorld"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToCamel(testCase))
	}
}

func (s strcaseTest) TestToCamel_basicExtended() {
	expected := "helloAWorld"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToCamel(testCase))
	}
}

func (s strcaseTest) TestToCamel_numbers() {
	expected := "catch22Statement"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToCamel(testCase))
	}
}

func (s strcaseTest) TestToCapital_basic() {
	expected := "HelloWorld"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToCapital(testCase))
	}
}

func (s strcaseTest) TestToCapital_basicExtended() {
	expected := "HelloAWorld"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToCapital(testCase))
	}
}

func (s strcaseTest) TestToCapital_numbers() {
	expected := "Catch22Statement"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToCapital(testCase))
	}
}

func (s strcaseTest) TestToLowerKebab_basic() {
	expected := "hello-world"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToLowerKebab(testCase))
	}
}

func (s strcaseTest) TestToLowerKebab_basicExtended() {
	expected := "hello-a-world"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToLowerKebab(testCase))
	}
}

func (s strcaseTest) TestToLowerKebab_numbers() {
	expected := "catch22-statement"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToLowerKebab(testCase))
	}
}

func (s strcaseTest) TestToLowerSnake_basic() {
	expected := "hello_world"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToLowerSnake(testCase))
	}
}

func (s strcaseTest) TestToLowerSnake_basicExtended() {
	expected := "hello_a_world"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToLowerSnake(testCase))
	}
}

func (s strcaseTest) TestToLowerSnake_numbers() {
	expected := "catch22_statement"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToLowerSnake(testCase))
	}
}

func (s strcaseTest) TestToUpperKebab_basic() {
	expected := "HELLO-WORLD"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToUpperKebab(testCase))
	}
}

func (s strcaseTest) TestToUpperKebab_basicExtended() {
	expected := "HELLO-A-WORLD"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToUpperKebab(testCase))
	}
}

func (s strcaseTest) TestToUpperKebab_numbers() {
	expected := "CATCH22-STATEMENT"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToUpperKebab(testCase))
	}
}

func (s strcaseTest) TestToUpperSnake_basic() {
	expected := "HELLO_WORLD"
	for _, testCase := range s.basicTestCases {
		s.EqualValues(expected, ToUpperSnake(testCase))
	}
}

func (s strcaseTest) TestToUpperSnake_basicExtended() {
	expected := "HELLO_A_WORLD"
	for _, testCase := range s.basicExtendedTestCases {
		s.EqualValues(expected, ToUpperSnake(testCase))
	}
}

func (s strcaseTest) TestToUpperSnake_numbers() {
	expected := "CATCH22_STATEMENT"
	for _, testCase := range s.numbersTestCases {
		s.EqualValues(expected, ToUpperSnake(testCase))
	}
}
