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

func main() {

	fmt.Print("How frequently do we write to the file in seconds? (e.g. 10): ")
	var usrTimer int
	fmt.Scanln(&usrTimer)

	initWrite := make(chan bool)

	for {
		go func() {
			select {
			case <-initWrite:
				txtToWrite := "This file is being written by HardDriveAlive.\nLast write was on " + time.Now().Format(time.ANSIC)
				writeFile("HardDriveAlive.txt", txtToWrite)
				fmt.Println("Writing to file...", time.Now().Format(time.ANSIC))
			}
		}()

		initWrite <- true
		time.Sleep(time.Duration(usrTimer) * time.Second)
	}
}

func writeFile(outFile, txt string) error {

	err := ioutil.WriteFile(outFile, []byte(txt), 0644)
	if err != nil {
		fmt.Println("Error writing to file. Check your permissions. ", err)
	}

	return nil
}
