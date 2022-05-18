# How to use

```go
// Import
import "github.com/AldieNightStar/golexem"

// Parse some random text and get the lexems
lexems := golexem.Parse("123.33 -12 etcLabel 'text' # comment and all\n'test' # Another comment\n")

// Each token has types like: STRING, ETC, NUMBER, COMMENT
// Each token has Value (String) or ValueNumber for number values
// Also by LineNumber you can check which line some token has

// In case of error a lot of tokens could have ETC
// You can check for that and make "invalid token error" in case of that

// Walk over the tokens and print the info
for _, token := range lexems {

    // Print the line number of the token
    // will print -1 if the token is unknown
    fmt.Println("LINE: ", golexem.GetTokenLine(token))

    // Now we checking token type. They should be one of: golexem(STRING, NUMBER, ETC, COMMENT)
    // tok.Value - string value (string, etc, comment)
    // tok.ValueNumber - number value (number)
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
```