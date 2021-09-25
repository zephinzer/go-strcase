package strcase

import (
	"bytes"
	"strings"
)

func ToCamel(input string) string {
	tokens := parse(input)
	for i := 1; i < len(tokens); i++ {
		tokens[i] = ToCapital(tokens[i])
	}
	return strings.Join(tokens, "")
}

func ToCapital(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		upperFirstChar := bytes.ToUpper([]byte{tokens[i][0]})[0]
		tokens[i] = string(append([]byte{upperFirstChar}, tokens[i][1:]...))
	}
	return strings.Join(tokens, "")
}

func ToLowerKebab(input string) string {
	tokens := parse(input)
	return strings.Join(tokens, "-")
}

func ToLowerSnake(input string) string {
	tokens := parse(input)
	return strings.Join(tokens, "_")
}

func ToUpperKebab(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return strings.Join(tokens, "-")
}

func ToUpperSnake(input string) string {
	tokens := parse(input)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return strings.Join(tokens, "_")
}
