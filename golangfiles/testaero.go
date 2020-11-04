package main

import (
		"bufio"
		//"io/ioutil"
		"fmt"
		//"log"
		//"os"
		//"os/exec"
		//"path/filepath"
		//"time"
		//"strings"
		"go.bug.st/serial.v1"
		"go.bug.st/serial.v1/enumerator"
)
func main(){
		
		mode := &serial.Mode{
				BaudRate: 115200,
		}
//		var port_teensy serial.Port
		var port_aero serial.Port
		
		ports, err := enumerator.GetDetailedPortsList()
		if err != nil{
				fmt.Println("port opening error")
		}
		
		fmt.Printf("found %d ports\n", len(ports))

		for _, portcheck := range ports{
			fmt.Println(portcheck.SerialNumber)
			if (portcheck.SerialNumber == "6423410"){
				fmt.Println("Aero teensy found")
				port_aero, err = serial.Open(portcheck.Name, mode)
				//aero(portcheck.Name)
			}
			//else {
//					if (portcheck.SerialNumber == "6416810"){

			//fmt.Println(portcheck.Name)
			//fmt.Println(portcheck.SerialNumber)
			//port, err = serial.Open(portcheck.Name, mode)
		}

		//_, err = port.Write([]byte("1"))
		//serReaderTeensy := bufio.NewReader(port_teensy)
		serReaderAero := bufio.NewReader(port_aero)
		
		for {
			msgBytes, err := serReaderAero.ReadBytes('\n')
			if (err != nil){
					fmt.Println("read failed")
			}
			fmt.Println(string(msgBytes))
		}

}
