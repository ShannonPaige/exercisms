package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(fullPhrase string) string {
	acroynm := ""

  fullPhrase = BreakOnCapitals(fullPhrase)
	fullPhrase = ReplaceHyphens(fullPhrase)
  
	words := strings.Split(fullPhrase, " ")

	for i := 0; i < len(words); i++ {
		firstLetter := strings.Split(words[i], "")[0]
		firstLetter = strings.ToUpper(firstLetter)
		acroynm += firstLetter
	}

	return acroynm
}

func BreakOnCapitals(fullPhrase string) string {
  re := regexp.MustCompile("([a-z])([A-Z])")
  return re.ReplaceAllString(fullPhrase, "$1 $2")
}

func ReplaceHyphens(fullPhrase string) string {
  return strings.Replace(fullPhrase, "-", " ", -1)
}
