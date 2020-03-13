package main

import (
		"bufio"
		"io/ioutil"
		"fmt"
		"log"
		"os"
		"os/exec"
		"path/filepath"
		"time"
		"strings"
//		"encoding/json"
		"go.bug.st/serial.v1"
		"go.bug.st/serial.v1/enumerator"
)

func main(){
		absPath, err := filepath.Abs("/home/pi/testingdata")
		files,err := ioutil.ReadDir(absPath)
		if err != nil{log.Println("failed to read path")}

		comfile,err := os.Create(absPath + "/communication.txt")
		if err != nil{fmt.Println(err)}
		comfile.Chmod(0777)
		comfile.WriteString("hello\n")
		//defer comfile.Close()

		mode := &serial.Mode{
				BaudRate: 115200,
		}

		var port serial.Port
		ports, err := enumerator.GetDetailedPortsList()
		fmt.Printf("Found %d ports\n", len(ports))
		for _, portcheck := range ports{
			//fmt.Printf("found port: %s\n", portcheck.VID)
			if (portcheck.VID == "16c0"){
				port, err = serial.Open(portcheck.Name, mode)
				//fmt.Println("found teensy")
			}
		}


		
		//port, err := serial.Open("/dev/ttyACM0", mode)
		//fmt.Print("port is ")



		if err != nil{log.Println("failed to open port")}
		defer port.Close()

		serReader := bufio.NewReader(port)

		//Code to sync time
		bufferedbytes := serReader.Buffered()
		bufferedbytes, err = serReader.Discard(bufferedbytes)
		
		_,err = port.Write([]byte("1")) //send handshake
		time.Sleep(time.Millisecond*100)
		msgBytes, err := serReader.ReadBytes('\n') //read RTC time
		msgStr := strings.TrimRight(string(msgBytes), "\r\n")
		//file.WriteString("Received " + msgStr + "\n")
		t,err := time.Parse("2006-01-02T03:04:05", msgStr) //parse for syscall
		dateString := t.Format("2 Jan 2006 03:04:05")
		if (err != nil){log.Println(err)}
		args := []string{"--set",dateString} //add set command
		//execute syscall to set time
		err = exec.Command("date", args...).Run()
		if err != nil{log.Print("Error setting time\n")}
		//file.WriteString("tried to set time to " + dateString + "\n")
	
		filepath := absPath + "/fakefile" + ".txt"
		if err == nil{
				filepath = absPath + "/" + "nongps" + t.Format("01-02_03-04") + ".txt"
		} else{ filepath = absPath + "/" + "nongps" + string(len(files) + 65) + ".txt"}
		fmt.Println(filepath)
		file, err := os.Create(filepath)
		if err != nil{log.Println("failed to create file")}
		defer file.Close()
		
		var can_log bool = false
		var accel_log bool = false //are 
		var cool_log = false //are thermistors logging
		
		for {
				msgBytes, err = serReader.ReadBytes('\n')	
				if (err != nil){
					log.Println("serial read failed")
				}else {
					if (!can_log && string(msgBytes[0]) == "0"){
						comfile.WriteString("CAN logging\n")
						can_log = true
						fmt.Println("can logging")
					}else if (!accel_log && string(msgBytes[0]) == "3"){
						comfile.WriteString("Accel logging\n")
						accel_log = true
						fmt.Println("Accel logging")
					}else if (!cool_log && string(msgBytes[0]) == "5"){
						comfile.WriteString("Cooling loop logging\n")
						cool_log = true
						fmt.Println("cool logging")
					}

					msgStr := strings.TrimRight(string(msgBytes), "\r\n")
					now := time.Now()
					dateString = now.Format(" 03:04:05.000\n")
					file.WriteString(msgStr + dateString)
					}
				}
		
}
