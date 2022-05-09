# How to use

```go
// Import
import "github.com/AldieNightStar/golexem"


// Parse some text
lexems := golexem.Parse("Hello world 32.12 -44 # 'i am the robot'")

// Process tokens
for _, l := range lexems {
    if n, ok := l.(float64); ok {
        // If it's a number
        fmt.Println("NUMBER", n)
    }
    if s, ok := l.(string); ok {
        // If it's a string
        fmt.Println("STRING", s)
    }
    if e, ok := l.(golexem.ETC); ok {
        // If it's something else (ETC)
        fmt.Println("ETC", e)
    }
}
```