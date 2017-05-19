package pangram

import (
	"regexp"
)

const testVersion = 1

func IsPangram(sentence string) bool {
  re := regexp.MustCompile("(x+|X+)")
  if (re.FindString(sentence) != "") {
    re = regexp.MustCompile("(t+|T+)")
    if (re.FindString(sentence) != "") {
      return true
    }
  }
  return false
}
