package main

import
(
    "fmt"
    "time"
)

var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
//una coma despues del 10 es porque para declarar arreglos multilinea siempre se necesita una coma
var mynumbers [5]int = [5]int{
                                1,
                                3,
                                4,
                                6,
                                10,
                                }

var daysOfWeek [7]string = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func main() {
    var dates [3]time.Time
    dates[0] = time.Unix(273846376, 0)
    fmt.Println(dates[0])
    dates[1] = time.Unix(173846376, 0)
    fmt.Println(dates[1])
    dates[2] = time.Unix(177846376, 0)
    fmt.Println(dates[2])

    fmt.Println(mynumbers)
    fmt.Printf("%#v\n", notes)

    for i := 0; i < len(dates); i++ {
        fmt.Println("Date", i, "is", dates[i])
    }

    for index, value := range(daysOfWeek) {
        fmt.Printf("position %d value %s \n", index+1, value)
    }
}