package strcase

import (
	"bytes"
	"strings"
)

var (
	// ToCapitalCamel is an alias for ToCapital
	ToCapitalCamel = ToCapital

	// ToCaterpillar is an alias for ToLowerKebab
	ToCaterpillar = ToLowerKebab

	// ToCobol is an alias for ToUpperKebab
	ToCobol = ToUpperKebab

	// ToCSS is an alias for ToLowerKebab
	ToCSS = ToLowerKebab

	// ToDash is an alias for ToLowerKebab
	ToDash = ToLowerKebab

	// ToHyphen is an alias for ToLowerKebab
	ToHyphen = ToLowerKebab

	// ToLazy is an alias for ToFlat
	ToLazy = ToFlat

	// ToLisp is an alias for ToLowerKebab
	ToLisp = ToLowerKebab

	// ToMacro is an alias for ToUpperSnake
	ToMacro = ToUpperSnake

	// ToPascal is an alias for ToCapital
	ToPascal = ToCapital
)

// ToCamel converts the input string `input` into `camelCase`..
//
// Useful for: variable names in JavaScript/TypeScriptS/Go/C/C++/Java,
// internal method names in Go
func ToCamel(input string) string {
	tokens := parse(input)
	for i := 1; i < len(tokens); i++ {
		tokens[i] = ToCapital(tokens[i])
	}
	return strings.Join(tokens, "")
}

// ToCapital converts the input string `input` into `CapitalCase`.
//
// Useful for: external method names in Go, classes and modules in Ruby,
// classes in JavaScript/TypeScript
func ToCapital(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		upperFirstChar := bytes.ToUpper([]byte{tokens[i][0]})[0]
		tokens[i] = string(append([]byte{upperFirstChar}, tokens[i][1:]...))
	}
	return strings.Join(tokens, "")
}

// ToFlat converts the input string `input` into `aflatcasestring`.
//
// Useful for: Java package names, Go package names
func ToFlat(input string) string {
	tokens := parse(input)
	return strings.Join(tokens, "")
}

// ToLowerKebab converts the input string `input` into `lower-kebab-case`.
//
// Useful for: Lisp names
func ToLowerKebab(input string) string {
	tokens := parse(input)
	return strings.Join(tokens, "-")
}

// ToLowerSnake converts the input string `input` into `lower_snake_case`.
//
// Useful for: variable names in Python
func ToLowerSnake(input string) string {
	tokens := parse(input)
	return strings.Join(tokens, "_")
}

// ToUpperKebab converts the input string `input` into `UPPER-KEBAB-CASE`.
//
// Useful for: COBOL
func ToUpperKebab(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return strings.Join(tokens, "-")
}

// ToUpperSnake converts the input string `input` into `UPPER_SNAKE_CASE`.
//
// Useful for: constants in JavaScript/TypeScript/Python, Shell variables
func ToUpperSnake(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return strings.Join(tokens, "_")
}
