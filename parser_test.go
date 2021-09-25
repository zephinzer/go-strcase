package strcase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type parseTest struct {
	suite.Suite
}

func Test_parse(t *testing.T) {
	suite.Run(t, &parseTest{})
}

func (s parseTest) Test_parse_basic() {
	expected := []string{"hello", "world"}
	testCases := []string{
		"hello world",
		"hello-world",
		"hello--world",
		"hello_world",
		"hello__world",
		"helloWorld",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_parse_basic_edge() {
	expected := []string{"h", "w"}
	testCases := []string{
		"h w",
		"h-w",
		"h--w",
		"h_w",
		"h__w",
		"hW",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_parse_basic_edge_extneded() {
	expected := []string{"whack", "a", "mole"}
	testCases := []string{
		"whack a mole",
		"whack-a-mole",
		"whack_a_mole",
		"whackAMole",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_parse_edge() {
	s.EqualValues([]string{"123"}, parse("123"))
	s.EqualValues([]string{"1", "2", "3"}, parse("1-2-3"))
	s.EqualValues([]string{"1", "2", "3"}, parse("1--2--3"))
	s.EqualValues([]string{"1", "2", "3"}, parse("1__2__3"))
	s.EqualValues([]string{"1", "2", "3"}, parse("1_-2_-3"), "wtf?")
	s.EqualValues([]string{"1", "a2", "b3"}, parse("1a2b3"))
	s.EqualValues([]string{"1", "plate", "of", "rice"}, parse("1plateOfRice"))
	s.EqualValues([]string{"1", "of3", "plates"}, parse("1of3plates"))
	s.EqualValues([]string{"1", "of", "3", "plates"}, parse("1of_3plates"))
	s.EqualValues([]string{"1", "for1"}, parse("1for1"))
	s.EqualValues([]string{"1", "for", "1"}, parse("1-for-1"))
	s.EqualValues([]string{"1", "for", "1"}, parse("1-For-1"))
	s.EqualValues([]string{"1", "for", "1"}, parse("1-FOR-1"))
	s.EqualValues([]string{"7", "zip", "archive"}, parse("7zipArchive"))
	s.EqualValues([]string{"7", "zip", "archive"}, parse("7zipArchive"))
}

func (s parseTest) Test_parse_extended() {
	expected := []string{"one", "two", "three"}
	testCases := []string{
		"one two three",
		"one-two-three",
		"one--two--three",
		"one_two_three",
		"one__two__three",
		"oneTwoThree",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_parse_extended_edge() {
	expected := []string{"a", "b", "c"}
	testCases := []string{
		"a b c",
		"a-b-c",
		"a--b--c",
		"a_b_c",
		"a__b__c",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_parse_extensions() {
	expected := []string{"send", "api", "request"}
	testCases := []string{
		"send API request",
		"send-API-Request",
		"send_API_Request",
		"sendAPIRequest",
	}
	for _, testCase := range testCases {
		s.EqualValues(expected, parse(testCase))
	}
}

func (s parseTest) Test_isDelimiter() {
	s.False(isDelimiter(typeLower, typeLower, typeInvalid),
		"should see somethinglikethis as one word")
	s.True(isDelimiter(typeAlpha|typeLower, typeAlpha|typeUpper, typeInvalid),
		"should mark the end of a word like helloWorld")
	s.True(isDelimiter(typeAlpha|typeUpper, typeAlpha|typeUpper, typeAlpha|typeLower),
		"something like whackAMole should be 'whack a mole'")
	s.True(isDelimiter(typeNumber, typeInvalid, typeInvalid),
		"should see numbers as an end since variable names cannot start with a number")
}
