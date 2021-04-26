package main

import (
    "bytes"
    "io"
    "log"
    "os"
    "os/exec"
)
func main() {
    // Binary to pall to the function
    arg := os.Args[1]

    // assign binary as child process
    command := exec.Command(arg)

    var stdBuffer bytes.Buffer
    mw := io.MultiWriter(os.Stdout, &stdBuffer)

    command.Stdout = mw
    command.Stderr = mw


    // Run the binary and stream the output.
    if err := command.Run(); err != nil {
        log.Panicln(err)
    }

    // Stream the log output to the terminal window
    log.Println(stdBuffer.String())
}