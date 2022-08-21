package stringUtility

import (
	"regexp"
	"strings"
	"unicode"
)

func IsStringInSlice(slice []string, stringValue string) bool {
	for _, stringInSlice := range slice {
		if strings.ToLower(stringInSlice) == strings.ToLower(stringValue) {
			return true
		}
	}

	return false
}

func StringSliceByIndex(text string, startIndex, stopIndex int) string {
	splitString := strings.Split(text, "")
	var desiredStringSlice []string
	for i, char := range splitString {
		if i >= startIndex && i <= stopIndex {
			desiredStringSlice = append(desiredStringSlice, char)
		} else {
			continue
		}
	}

	desiredString := strings.Join(desiredStringSlice, "")
	return desiredString
}

func TrimALLSpace(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}, text)
}

func SpecialCharacterFilter(input string) string {
	splittedStringChar := []rune(input)

	newCharSlice := make([]string, 0)

	for _, char := range splittedStringChar {
		charCode := int(char)

		emojiRegex := regexp.MustCompile(`[\x{1F300}-\x{1F6FF}]`)

		isCharNumber := charCode >= 48 && charCode <= 57
		isCharUpperCaseLetter := charCode >= 65 && charCode <= 90
		isCharLowerCaseLetter := charCode >= 97 && charCode <= 122
		isCharSpace := string(char) == " "
		isCharEmoji := emojiRegex.MatchString(string(char))

		if isCharNumber || isCharUpperCaseLetter || isCharLowerCaseLetter || isCharSpace || isCharEmoji {
			newCharSlice = append(newCharSlice, string(char))
		}

	}

	return strings.Join(newCharSlice, "")

}
