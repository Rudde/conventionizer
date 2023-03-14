package conventionizer

import (
	"encoding/binary"
	"hash/fnv"
	"math/rand"
	"strings"
	"unicode"
)

type separator rune

const (
	underscore separator = '_'
	dash       separator = '-'
	dot        separator = '.'
	space      separator = ' '
)

func generateStudlyWords(words []string) []string {
	result := make([]string, len(words))

	return result
}

func ToStudly(str string) string {
	hasher := fnv.New64a()
	hash := hasher.Sum([]byte(str))
	seed := binary.BigEndian.Uint64(hash)
	rnd := rand.New(rand.NewSource(int64(seed)))
	words := split(str)
	result := make([]string, len(words))
	for wordPos, word := range words {
		lastCap := 0
		var newWord strings.Builder
		for pos, c := range strings.ToLower(word) {
			weight := 50 // Of 100
			if wordPos == 0 && pos == 0 {
				weight -= 20
			} else if pos == 0 {
				weight -= 10
			} else if pos != len(word)-1 {
				weight += 15
			} else if pos == len(word)-1 {
				weight -= 5
			}
			weight += lastCap * 5
			switch c {
			case 'a', 'e', 'i', 'o', 'u', 'æ', 'ø', 'å': // vowel
				weight += 25
			case 'w', 'y': // half-vowel
				weight += 15
			default:
				weight -= 10
			}
			num := rnd.Intn(100-1) + 1
			if num <= weight {
				newWord.WriteRune(unicode.ToUpper(c))
				lastCap = 0
			} else {
				newWord.WriteRune(c)
				lastCap++
			}
		}
		result[wordPos] = newWord.String()
	}
	return strings.Join(result, string(space))
}

func ToCobol(str string) string {
	return strings.ToUpper(ToDash(str))
}

func ToTrain(str string) string {
	return strings.Join(formatWords(split(str), toTitle), string(dash))
}

func ToMacro(str string) string {
	return strings.ToUpper(ToSnake(str))
}

func ToDash(str string) string {
	return strings.ToLower(strings.Join(split(str), string(dash)))
}

func ToSnake(str string) string {
	return strings.ToLower(strings.Join(split(str), string(underscore)))
}

func ToPascal(str string) string {
	return capitalizeFirstChar(ToCamel(str))
}

func ToCamel(str string) string {
	words := split(str)
	var result strings.Builder
	for pos, word := range words {
		if pos == 0 {
			result.WriteString(strings.ToLower(word))
		} else {
			result.WriteString(toTitle(word))
		}
	}
	return result.String()
}

func formatWords(words []string, formater func(string) string) []string {
	var result []string
	for _, word := range words {
		result = append(result, formater(word))
	}
	return result
}

func capitalizeFirstChar(str string) string {
	return string(unicode.ToUpper(rune(str[0]))) + str[1:]
}

func toTitle(str string) string {
	return capitalizeFirstChar(strings.ToLower(str))
}

func splitOnSeparator(str string, on separator) []string {
	var result []string
	var word []rune
	separatorPosition := 0
	for pos, c := range str {
		if separatorPosition != 0 && c != rune(on) {
			result = append(result, string(word))
			word = []rune{}
			separatorPosition = 0
		}
		if c != rune(on) || (c == rune(on) && len(word) == 0) {
			word = append(word, c)
		} else {
			separatorPosition = pos
		}
	}
	return append(result, string(word))
}

func splitOnCase(str string) []string {
	var result []string
	var word []rune
	for pos, c := range str {
		if pos == 0 {
			word = append(word, c)
		} else if isUpperLetter(c) {
			result = append(result, string(word))
			word = []rune{c}
		} else {
			word = append(word, c)
		}
	}
	return append(result, string(word))
}

func split(str string) []string {
	switch {
	case hasSeparator(str, underscore):
		return splitOnSeparator(str, underscore)
	case hasSeparator(str, dash):
		return splitOnSeparator(str, dash)
	case isMixCase(str):
		return splitOnCase(str)
	default:
		return []string{str}
	}
}

func isMixCase(str string) bool {
	hasUpper := false
	hasLower := false
	for _, c := range str {
		if !hasLower && isLowerLetter(c) {
			hasLower = true
		} else if !hasUpper && isUpperLetter(c) {
			hasUpper = true
		}
		if hasLower && hasUpper {
			return true
		}
	}
	return false
}

func hasSeparator(str string, separator separator) bool {
	first := false
	separatorPosition := 0
	for pos, c := range str {
		if !first && isLetter(c) {
			first = true
		} else if first && c == rune(separator) {
			separatorPosition = pos
		} else if first && separatorPosition != 0 && isLetter(c) {
			return true
		}
	}
	return false
}

func isLetter(c rune) bool {
	return isLowerLetter(c) || isUpperLetter(c)
}

func isLowerLetter(c rune) bool {
	return unicode.IsLower(c)
}

func isUpperLetter(c rune) bool {
	return unicode.IsUpper(c)
}
