package main

import "log"
import "fmt"
import "bufio"
import "os"
import "strings"
import "os/exec"
import "path/filepath"
import "os/user"



func main() {
fmt.Println("Shell started. Enter a command to run(use 'help' for valid commands):")
inputReader := bufio.NewReader(os.Stdin)
activeDir, err := os.Getwd() //Directory getter for later functions.
if err!=nil {
	os.Exit(1)}

for {
	fmt.Print(activeDir + " $ ")
	input, err := inputReader.ReadString('\n')

	if err != nil { 
		fmt.Print("Error: Read Input failed")}
	input = strings.TrimSpace(input) //Rewrite input with correct format(w/o spacing)
	if input == "" {
		fmt.Print("Error: Input Empty")
		continue}

	inputArr := strings.Fields(input)
	commandVar := inputArr[0] // Set command used equal to a variable not including the arguments.
	argumentsVar := inputArr[1:] //set all arguments equal to a variable not including the intial command.
 
switch commandVar {
case "help":
fmt.Println("Avaliable Commands include: 'ls', 'whoami', 'cp', 'help', 'mv', 'mkdir', 'wc', 'exit'")

case "ls":
runCommand("ls", argumentsVar)

case "wc":
argumentsVar[0] = activeDir + "/" + argumentsVar[0]
runCommand("wc", argumentsVar)

case "cp":
argumentsVar[0] = activeDir + "/" + argumentsVar[0]
argumentsVar[1] = activeDir + "/" + argumentsVar[1]
runCommand("cp", argumentsVar)

case "mkdir":
argumentsVar[0] = activeDir + "/" + argumentsVar[0]
runCommand("mkdir", argumentsVar)

case "mv":
argumentsVar[0] = activeDir + "/" + argumentsVar[0]
argumentsVar[1] = activeDir + "/" + argumentsVar[1]
runCommand("mv", argumentsVar)

case "cd":
activeDir = commandCD(activeDir, argumentsVar)

case "exit":
	os.Exit(0)

case "whoami":
currentUser, err := user.Current()
   if err != nil {
      log.Fatalf(err.Error())
   }
   username := currentUser.Username
   id := currentUser.Uid
   fmt.Println(username, id)
case "default":
runCommand(commandVar, argumentsVar)
}
}
}

func runCommand(command string, args []string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error: Command not Available", err)
    }
}
func commandCD(currentDirectory string, newDirectory []string) string { //Function to use CD by returning the new directory to move to if it exits, if not it returns back the currentdirectory the user is in.
if len(newDirectory) == 0 { //Checks if the directory to go to is empty or not, if empty prints error and returns the current active directory the user is in.
fmt.Println("Error: Directory Input Blank")
return currentDirectory}

filePathDirectory := filepath.Join(currentDirectory, newDirectory[0]) // Used code @ https://pkg.go.dev/path/filepath
if _, err := os.Stat(filePathDirectory)
err != nil {
fmt.Println("Error: Directory not found - ", err)
return currentDirectory}
return filePathDirectory}
