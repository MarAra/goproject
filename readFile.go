package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    fileInfo, err := os.Stat("fail.go") //lee el archivo file.go e imprime el tamanio del archivo
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(fileInfo.Size())

    count := 12
    suffix := " minutes of bonus footage"
    format := " DVD"
    languages := append([]string{}, "Espanol")
    fmt.Println(count, suffix, format, languages)
}