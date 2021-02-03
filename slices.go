package main

import
(
    "fmt"
)

func main() {
    daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

    businessDays := daysOfWeek[0:5] //declare a slice
    holidays := daysOfWeek[5:7] //declare another slice

    fmt.Println("Business days: ", businessDays)
    fmt.Println("Holidays: ", holidays)

    //but if I modify an element of the original array
    daysOfWeek[4] = "Jisus"

    fmt.Println("Business days: ", businessDays) //because slice is a pointer to the original elements of the array

    myslice := make([]string, 5) //with this literal I don't have to worry about working with elements from the underlying original array

    fmt.Println(myslice)
}
