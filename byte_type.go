package strcase

type byteType int

type byteTypes []byteType

var byteTypeMap = map[byte]byteType{
	'-':  typeDash,
	'_':  typeUnderscore,
	'.':  typePeriod,
	' ':  typeSpace,
	'\t': typeTab,
	'\n': typeNewline,
}

var byteTypeList []byte

func init() {
	for b := range byteTypeMap {
		byteTypeList = append(byteTypeList, b)
	}
}

const (
	typeInit       byteType = 1 << iota
	typeAlpha      byteType = 1 << iota
	typeUpper      byteType = 1 << iota
	typeLower      byteType = 1 << iota
	typeNumber     byteType = 1 << iota
	typeDelimiter  byteType = 1 << iota
	typePeriod     byteType = 1 << iota
	typeDash       byteType = 1 << iota
	typeUnderscore byteType = 1 << iota
	typeSpace      byteType = 1 << iota
	typeTab        byteType = 1 << iota
	typeNewline    byteType = 1 << iota
	typeUnknown    byteType = 1 << iota
	typeInvalid    byteType = 1 << iota
)
