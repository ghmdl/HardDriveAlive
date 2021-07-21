// Learning Go - next is a GUI version of this utility.

/* Small cross-platform self-contained utility to keep hard drives alive
   by writing a text file to the drive it's launched from based on a
   timer interval specified by the user */

// MdL - July 21, 2021

package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var InitWrite chan bool

func main() {
	InitWrite = make(chan bool)

	var usrTimer uint32
	fmt.Print("How frequently do we write to the file in seconds? (e.g. 10): ")
	fmt.Scanln(&usrTimer)

	if usrTimer < 0 { //TODO: implement a better input checking

		fmt.Println("Invalid input. Only a positive integer is allowed.")
		fmt.Println("Exiting in 3 seconds...") //TODO: loop back program logic on error
		time.Sleep(3 * time.Second)

	} else if usrTimer > 1 {

		for {
			go func() {
				select {
				case <-InitWrite:
					txtToWrite := "This file is being written by HardDriveAlive.\nLast write was on " + time.Now().Format(time.ANSIC)
					writeFile("HardDriveAlive.txt", txtToWrite)
					fmt.Println("Writing to file...", time.Now().Format(time.ANSIC))
				}
			}()

			InitWrite <- true
			time.Sleep(time.Duration(usrTimer) * time.Second)
		}

	} else {
		fmt.Println("Invalid input. Only a positive integer is allowed.")
		fmt.Println("Exiting in 3 seconds...") //TODO
		time.Sleep(3 * time.Second)
	}
}

func writeFile(outFile, txt string) error {

	err := ioutil.WriteFile(outFile, []byte(txt), 0644)
	if err != nil {
		fmt.Println("Error writing to file. Check your permissions in the current directory.\nError: ", err)
		InitWrite <- false
	}

	return nil
}
