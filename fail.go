package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
)

func main() {
    fmt.Print("Please, enter a new grade: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n') //retorna todo lo que el usuario digito hasta que presiono enter
    fmt.Println("Grade entered: ", input)

    input = strings.TrimSpace(input)
    grade,err := strconv.ParseFloat(input, 64)
    if err != nil {
        log.Fatal(err) //report the error and stop the program
    } else {
        if grade > 60.0 {
            fmt.Println(" you approved!")
        } else {
            fmt.Println(" you failed!")
        }
    }


    fmt.Print("Please, enter other grade: ")
    input2, err := reader.ReadString('\n') //retorna todo lo que el usuario digito hasta que presiono enter
    if err != nil {
        log.Fatal(err) //report the error and stop the program
    } else {
        fmt.Print("Grade entered: ", input2)
        grade, err := strconv.ParseInt(input2[:len(input2)-2], 10, 0) //puedo hacer un substring de una cadena[initindex:endindex] y con ParseInt(s string, base int, bitSize int)
        if err != nil {
            log.Fatal(err) //report the error and stop the program
        }

        if grade > 60 {
            fmt.Println(" you approved!")
        } else {
            fmt.Println(" you failed!")
        }
    }

}