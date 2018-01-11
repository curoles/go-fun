package main

import (
    //go get golang.org/x/crypto/ssh/terminal
    "golang.org/x/crypto/ssh/terminal"
    "fmt"
    //"os"
)

func main() {
    /*oldState, err := terminal.MakeRaw(0)
    if err != nil {
            panic(err)
    }
    defer terminal.Restore(0, oldState)*/

    width, height, _ := terminal.GetSize(0)

    /*
    fmt.Println("Is terminal: ", terminal.IsTerminal(1))

    term := terminal.NewTerminal(os.Stdout, "")

    ec := term.Escape
    fmt.Printf("|%s| width=%d height=%d |%s|\n", ec.Red, width, height, ec.Reset)
    */

    fmt.Printf("%d:%d\n", width, height)
}
