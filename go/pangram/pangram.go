package pangram

import (
	"fmt"
  "strings"
)

const testVersion = 1

func IsPangram(sentence string) bool {
  // re := regexp.MustCompile("(x+|X+)")
  // if (re.FindString(sentence) != "") {
  //   re = regexp.MustCompile("(t+|T+)")
  //   if (re.FindString(sentence) != "") {
  //     return true
  //   }
  // }
  // return false
  sentence = strings.ToUpper(sentence)
  
}
