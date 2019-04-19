package main

import (
        "fmt"
        "github.com/goburrow/modbus"
)



func main() {
        fmt.Println("Hello World!")


        // Modbus TCP
        client := modbus.TCPClient("10.54.52.180:502")
        //client := modbus.TCPClient("localhost:502")


        results, err := client.ReadHoldingRegisters(35, 12)

        fmt.Println("Result:", results)
        fmt.Println("Error:", err)
}

