package main

import (
    "fmt"
    "log"
    "util"
)

func main() {
    fmt.Print("Please, enter a temperature in Fahrenheit: ")
    fahrenheit, err := util.GetFloat()
    if err != nil {
        log.Fatal(err) //report the error and stop the program
    }
    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Printf("%.2f degrees Celsius", celsius)
}
