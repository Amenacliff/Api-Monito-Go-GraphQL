package stringUtility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsStringInSlice(t *testing.T) {
	requires := require.New(t)
	sliceOfString := []string{"bomb", "tag", "beans", "GOLANG"}

	isIn := IsStringInSlice(sliceOfString, "golang")
	requires.True(isIn)
	isIn = IsStringInSlice(sliceOfString, "BEANS")
	requires.True(isIn)
	isIn = IsStringInSlice(sliceOfString, "java")
	requires.False(isIn)
}

func TestStringSliceByIndex(t *testing.T) {
	requires := require.New(t)

	stringTest := "ThisIsAGolangProject"

	sliceResult := "ThisI"

	stringSlice := StringSliceByIndex(stringTest, 0, 4)

	requires.Equal(sliceResult, stringSlice)
}

func TestTrimALLSpace(t *testing.T) {
	requires := require.New(t)

	stringTest := "This Is A Golang Project"

	expectedString := "ThisIsAGolangProject"

	trimmedString := TrimALLSpace(stringTest)

	requires.Equal(expectedString, trimmedString)
}

func TestSpecialCharacterFilter(t *testing.T) {
	requires := require.New(t)

	stringTest := "ThisIsAGolan%()&(&)_*)&&g Project🙏🏽"

	expectedString := "ThisIsAGolang Project🙏🏽"

	finalString := SpecialCharacterFilter(stringTest)

	requires.Equal(expectedString, finalString)
}
