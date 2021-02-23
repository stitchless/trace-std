package main
import (
	"log"
	"os"
	"os/exec"
)
func main() {
	arg := os.Args[1]
	command := exec.Command(arg)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}