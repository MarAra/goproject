package main

import (
    "fmt"
    "log"
    "util"
)

func main() {
    fmt.Print("Please, enter a new grade: ")
    grade, err := util.GetFloat()
    if err != nil {
        log.Fatal(err) //report the error and stop the program
    } else {
        if grade > 60.0 {
            fmt.Println(" you approved!")
        } else {
            fmt.Println(" you failed!")
        }
    }
}
