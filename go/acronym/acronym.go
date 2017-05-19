package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(fullPhrase string) string {
	acroynm := ""

	re := regexp.MustCompile("([a-z])([A-Z])")
	fullPhrase = re.ReplaceAllString(fullPhrase, "$1 $2")

	fullPhrase = strings.Replace(fullPhrase, "-", " ", -1)
	words := strings.Split(fullPhrase, " ")

	for i := 0; i < len(words); i++ {
		firstLetter := strings.Split(words[i], "")[0]
		firstLetter = strings.ToUpper(firstLetter)
		acroynm += firstLetter
	}

	return acroynm
}
