package main

import
(
    "fmt"
    "reflect"
)

func main() {
    var myInt int
    fmt.Println(myInt)
    fmt.Println(&myInt)
    fmt.Println(reflect.TypeOf(&myInt), "\n")

    var myFloat = 4.5
    fmt.Println(myFloat)
    fmt.Println(&myFloat)
    fmt.Println(reflect.TypeOf(&myFloat), "\n")

    var otherInt int = 2948
    //declare pointer
    var pointerToInt *int
    //assign pointer
    pointerToInt = &otherInt
    //print the pointer
    fmt.Println(pointerToInt)
    //print the pointer value
    fmt.Println(*pointerToInt, "\n")

    //changing values through pointers
    *pointerToInt = 7878898
    fmt.Println(otherInt)
    fmt.Println(*pointerToInt, "\n")

    var myBool bool
    pointerToBool := &myBool
    fmt.Println(pointerToBool, "\n")

    newInt := 99
    pointRet := usePointers(&newInt)
    fmt.Println(*pointRet)

    truth := true
    negate(&truth)
    fmt.Println(truth)

    lies := false
    negate(&lies)
    fmt.Println(lies)

    greeting := fmt.Sprintf("Thank you! this is a float64 %10.2f %s %d", 2345.45442, "Mario Araoz", 12)
    fmt.Println(greeting)
}

func usePointers(pointerByParam *int) *float64 {
    *pointerByParam = 55
    poin := float64(*pointerByParam) / 9
    return &poin
}

func negate(myBoolean *bool) {
    *myBoolean = !*myBoolean
}