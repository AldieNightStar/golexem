# How to use

# Import
```go
import "github.com/AldieNightStar/golexem"
```

# Usage
* Parse the text
```go
tokens := golexem.Parse("123.33 -12 etcLabel 'text' # comment and all\n'test' # Another comment\n")
```
* Walk over the tokens
    * Here you can:
        * Get line number
        * Get type/value of the token
    * To check the type, simply use: `val, ok := token.(golexem.TYPE)`
    * Each token has:
        * `Value` - string value. Could be `STRING`, `COMMENT` or `ETC`
        * `ValueNumber` - number value
        * `LineNumber` - line number of the token
```go
for _, token := range tokens {
    // Get token line number
    golexem.GetTokenLine(token)

    // Check token types. There are: STRING, NUMBER, ETC, COMMENT
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

# Text it parses
```r
This is the etc part
"Some string stuff" `and this too` 'and this one is string'
numbers is 123 -12 -44 22.13
# This is simple comment. You can ignore while parsing it
```