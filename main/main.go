package main

import (
	"fmt"

	"github.com/AldieNightStar/golexem"
)

func main() {
	lexems := golexem.Parse("123.33 -12 etcLabel 'text' # comment and all\n'test' # Another comment\n1\n22\n33\n's' e123 -1\n22")

	for _, token := range lexems {
		fmt.Println("LINE: ", golexem.GetTokenLine(token))

		if str, ok := token.(golexem.STRING); ok {
			fmt.Println("STRING", str.Value)
		} else if num, ok := token.(golexem.NUMBER); ok {
			fmt.Println("NUMBER", num.ValueNumber)
		} else if etc, ok := token.(golexem.ETC); ok {
			fmt.Println("ETC", etc.Value)
		} else if cmt, ok := token.(golexem.COMMENT); ok {
			fmt.Println("COMMENT", cmt.Value)
		} else {
			fmt.Println("???")
		}
	}
}
