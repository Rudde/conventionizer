package conventionizer

import (
	"reflect"
	"strings"
	"testing"
)

const (
	snake    = "this_is_snake_case"
	macro    = "THIS_IS_MACRO_CASE"
	train    = "This-Is-Train-Case" // Used in HTTP Headers
	cobol    = "THIS-IS-COBOL-CASE"
	camel    = "thisIsCamelCase"
	pascal   = "ThisIsPascalCase"
	dashCase = "this-is-dash-case"
	lower    = "lower"
	upper    = "UPPER"
	title    = "Title"

	prefixedSnake = "_prefixed_snake_case"
)

func TestToStudly(t *testing.T) {
	testCase := ToStudly("stUdlY-cAps-StUdLy-CaPs")
	if testCase != "sTUdLy CAPS StUdlY cAPs" {
		t.Error("ToStudly formula changed")
	}
}

func TestToCobol(t *testing.T) {
	testCase := ToCobol(cobol)
	if testCase != cobol {
		t.Error("ToCobol malformed cobol input")
	}
	testCase = ToCobol(camel)
	if testCase != "THIS-IS-CAMEL-CASE" {
		t.Error("ToCobol failed changing format of camelCase")
	}
}

func TestToTrain(t *testing.T) {
	testCase := ToTrain(train)
	if testCase != train {
		t.Error("ToTrain malformed train input")
	}
	testCase = ToTrain(dashCase)
	if testCase != "This-Is-Dash-Case" {
		t.Error("ToTrain failed formatting dash-case input")
	}
	testCase = ToTrain(camel)
	if testCase != "This-Is-Camel-Case" {
		t.Error("ToTrain failed formatting camelCase input")
	}
}

func TestToMacro(t *testing.T) {
	testCase := ToMacro(macro)
	if testCase != macro {
		t.Error("ToMacro malformed macro input")
	}
	testCase = ToMacro(snake)
	if testCase != strings.ToUpper(snake) {
		t.Error("ToMacro failed converting snake_case")
	}
}

func TestToDash(t *testing.T) {
	testCase := ToDash(dashCase)
	if testCase != dashCase {
		t.Error("ToDash malformed dash input")
	}
	testCase = ToDash(snake)
	if testCase != "this-is-snake-case" {
		t.Error("ToDash did not convert snake_case to dash-case")
	}
	testCase = ToDash(camel)
	if testCase != "this-is-camel-case" {
		t.Error("ToDash did not format camelCase input as dash-case")
	}
}

func TestToPascal(t *testing.T) {
	testCase := ToPascal(pascal)
	if testCase != pascal {
		t.Error("ToPascal malformed pascal input")
	}
	testCase = ToPascal(camel)
	if testCase != capitalizeFirstChar(camel) {
		t.Error("ToPascal failed formatting camelCase as PascalCase")
	}
}

func TestToCamel(t *testing.T) {
	testCase := ToCamel(camel)
	if testCase != camel {
		t.Error("ToCamel malformed camel input")
	}
	testCase = ToCamel(pascal)
	if testCase != "thisIsPascalCase" {
		t.Errorf("%s could not be formated to CamelCase", pascal)
	}
	testCase = ToCamel(snake)
	if testCase != "thisIsSnakeCase" {
		t.Errorf("%s could not be formated as camel", snake)
	}
	testCase = ToCamel(dashCase)
	if testCase != "thisIsDashCase" {
		t.Errorf("%s could not be formated as camel", dashCase)
	}
}

func TestToSnake(t *testing.T) {
	testCase := ToSnake(dashCase)
	if testCase != "this_is_dash_case" {
		t.Errorf("%s did not successfully get converted to snake_case", dashCase)
	}
	testCase = ToSnake(upper)
	if testCase != "upper" {
		t.Error("ToSnake failed on upper test")
	}
	testCase = ToSnake(lower)
	if testCase != lower {
		t.Error("ToSnake should not affect lower")
	}
	testCase = ToSnake(snake)
	if testCase != snake {
		t.Error("ToSnake does not handle snake_case input")
	}
	testCase = ToSnake(pascal)
	if testCase != "this_is_pascal_case" {
		t.Error("ToSnake failed converting pascalCase")
	}
	testCase = ToSnake(macro)
	if testCase != strings.ToLower(macro) {
		t.Errorf("ToSnake failed converting MACRO_CASE")
	}
}

func TestCapitalizeFirstChar(t *testing.T) {
	testCase := capitalizeFirstChar(title)
	if testCase != title {
		t.Error("CapitalizeFirstChar malformed Title input")
	}
	testCase = capitalizeFirstChar(upper)
	if testCase != upper {
		t.Error("CapitalizeFirstChar malformed uppercase input")
	}
	testCase = capitalizeFirstChar(lower)
	if testCase != "Lower" {
		t.Error("CapitalizeFirstChar didn't capitalize lower input")
	}
	testCase = capitalizeFirstChar(pascal)
	if testCase != pascal {
		t.Error("CapitalizeFirstChar malformed PascalCase input")
	}
	testCase = capitalizeFirstChar(camel)
	if testCase != "ThisIsCamelCase" {
		t.Error("CapitalizeFirstChar did not capitalize first letter in camelCase input")
	}
}

func TestFormatWords(t *testing.T) {
	formattedWords := formatWords(split(camel), toTitle)
	pascalWords := split(ToPascal(camel))
	if !reflect.DeepEqual(formattedWords, pascalWords) {
		t.Error("formatWords don't return the correct format")
	}
}

func TestToTitle(t *testing.T) {
	testCase := toTitle(title)
	if testCase != title {
		t.Error("toTitle should produce the same Title")
	}
	testCase = toTitle(upper)
	if testCase != "Upper" {
		t.Error("Failed converting upper case to title case")
	}
	testCase = toTitle(lower)
	if testCase != "Lower" {
		t.Error("Failed converting lower case to title case")
	}
}

func TestSplit(t *testing.T) {
	testCase := split(snake)
	if !reflect.DeepEqual(testCase, []string{"this", "is", "snake", "case"}) {
		t.Error("Did not successfully split on _")
	}
	testCase = split(dashCase)
	if !reflect.DeepEqual(testCase, []string{"this", "is", "dash", "case"}) {
		t.Error("Did not successfully split on -")
	}
	testCase = split(lower)
	if !reflect.DeepEqual(testCase, []string{lower}) {
		t.Errorf("Unexcpected result splitting %s", lower)
	}
	testCase = split(camel)
	if !reflect.DeepEqual(testCase, []string{"this", "Is", "Camel", "Case"}) {
		t.Errorf("%s failed splitting on case", camel)
	}
}

func TestSplitOnCase(t *testing.T) {
	testCase := splitOnCase(pascal)
	if !reflect.DeepEqual(testCase, []string{"This", "Is", "Pascal", "Case"}) {
		t.Errorf("%s failed splitting on case", pascal)
	}
	testCase = splitOnCase(camel)
	if !reflect.DeepEqual(testCase, []string{"this", "Is", "Camel", "Case"}) {
		t.Errorf("%s failed splitting on case", camel)
	}
}

func TestSplitOnSeparator(t *testing.T) {
	testCase := splitOnSeparator(snake, '_')
	if !reflect.DeepEqual(testCase, []string{"this", "is", "snake", "case"}) {
		t.Error("Did not successfully split on _")
	}
	testCase = splitOnSeparator(prefixedSnake, '_')
	if !reflect.DeepEqual(testCase, []string{"_prefixed", "snake", "case"}) {
		t.Error("Did not successfully split on _ with prefix")
	}
}

func TestIsMixCase(t *testing.T) {
	if isMixCase(lower) || isMixCase(upper) {
		t.Error("Is not mixed case")
	}

	if !isMixCase(pascal) || !isMixCase(camel) {
		t.Error("Mix case not detected")
	}

	if isMixCase(dashCase) {
		t.Errorf("%s is not mix case", dashCase)
	}
}

func TestIsLetter(t *testing.T) {
	if !isLetter('a') || !isLetter('A') || !isLetter('Æ') {
		t.Error("All letters are letters")
	}

	if isLetter('0') || isLetter('9') {
		t.Error("Number marked as a letter")
	}
}

func TestHasSeparator(t *testing.T) {
	if !hasSeparator(snake, '_') {
		t.Errorf("%s has _ as seperator", snake)
	}

	if hasSeparator(camel, '_') {
		t.Errorf("%s does not have seperator _", camel)
	}

	if !hasSeparator(dashCase, '-') {
		t.Errorf("%s has - as seperator", dashCase)
	}
}
