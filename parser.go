package strcase

import (
	"strings"
)

func parse(input string, tokenSet ...string) []string {
	var currentTokenSet []string
	if tokenSet != nil {
		currentTokenSet = tokenSet
	}
	nextIndex := findNextDelimiter(input)
	if nextIndex == -1 {
		currentTokenSet = append(currentTokenSet, normalizeToken(input))
		return currentTokenSet
	}
	currentTokenSet = append(currentTokenSet, normalizeToken(input[:nextIndex]))
	return parse(input[nextIndex:], currentTokenSet...)
}

func normalizeToken(input string) string {
	return strings.ToLower(strings.Trim(input, string(byteTypeList)))
}

// findNextDelimiter returns the index position of the first
// transition from the left of the input `input`. If no
// transitions are found, -1 is returned
func findNextDelimiter(input string) int {
	previousByteType := typeInit
	for i := range input {
		currentByteType := getByteType(input[i])
		var nextByteType byteType = typeInvalid
		if i+1 < len(input) {
			nextByteType = getByteType(input[i+1])
		}
		if i > 0 && isDelimiter(previousByteType, currentByteType, nextByteType) {
			return i
		}
		previousByteType = currentByteType
	}
	return -1
}

func isDelimiter(previous byteType, current byteType, next byteType) bool {
	if previous&typeAlpha == current&typeAlpha {
		if previous&typeLower == typeLower && current&typeUpper == typeUpper {
			return true
		}
		if current&typeUpper == typeUpper && next&typeLower == typeLower {
			return true
		}
	}
	if previous&typeNumber == typeNumber && current&typeNumber != typeNumber {
		return true
	}
	if previous&typeDelimiter != typeDelimiter && current&typeDelimiter == typeDelimiter {
		return true
	}
	return false
}

func getByteType(input byte) byteType {
	typeOfByte := typeUnknown
	switch true {
	case input >= 'A' && input <= 'Z':
		typeOfByte = typeAlpha | typeUpper
	case input >= 'a' && input <= 'z':
		typeOfByte = typeAlpha | typeLower
	case input >= '0' && input <= '9':
		typeOfByte = typeNumber
	default:
		if predefinedByteType, ok := byteTypeMap[input]; ok {
			typeOfByte = typeDelimiter | predefinedByteType
		}
	}
	return typeOfByte
}
