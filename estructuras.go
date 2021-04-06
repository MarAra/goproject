package main

import "fmt"

type part struct {
    description string
    count int
}

type car struct {
    name string
    topSpeed float64
}

type house struct {
    address string
    rooms int
    bathrooms int
    garage bool
    price float64
}

func main() {
    var myStruct struct {
        number float64
        word   string
        toggle bool
    }

    myStruct.number = 12.34
    myStruct.word = "Mario Araoz"
    myStruct.toggle = true

    fmt.Println(myStruct.number)
    fmt.Println(myStruct.word)
    fmt.Println(myStruct.toggle)

    var porche car
    porche.name = "porche 911"
    porche.topSpeed = 323

    fmt.Printf("Name: %s \n", porche.name)
    fmt.Printf("Top Speed: %f \n", porche.topSpeed)

    var bolts part
    bolts.description = "Hex bolt"
    bolts.count = 24

    showInfo(bolts)

    var myHouse = getDefaultHouse()
    fmt.Println(myHouse)
}

func showInfo(myPart part) {
    fmt.Printf("Description: %s \n", myPart.description)
    fmt.Printf("Count: %d \n", myPart.count)
}

func getDefaultHouse() house {
    var defaultHouse house
    defaultHouse.address = "Av evergreen 123"
    defaultHouse.rooms = 3
    defaultHouse.bathrooms = 2
    defaultHouse.garage = true
    defaultHouse.price = 30000.2

    return defaultHouse
}