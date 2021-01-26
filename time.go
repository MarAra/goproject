package main

import (
    "fmt"
    "time"
    "strings"
)

func main() {
    var now time.Time = time.Now()
    var year = now.Year()
    fmt.Println(now)
    fmt.Println(year)
    phrase := "G# r#cks!"
    replacerObj := strings.NewReplacer("#", "o")
    newPhrase := replacerObj.Replace(phrase)
    fmt.Println(newPhrase)
}