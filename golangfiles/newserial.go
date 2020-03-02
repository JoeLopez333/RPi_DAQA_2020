package main

import (
		"bufio"
		"io/ioutil"
		"fmt"
		"log"
		"os"
		"path/filepath"
//		"strings"
//		"encoding/json"

		"go.bug.st/serial.v1"
)

func main(){
		absPath, err := filepath.Abs("/home/pi/testfiles/golangfiles")
		files,err := ioutil.ReadDir(absPath)
		if err != nil{log.Println("failed to create file")}
		
		fmt.Println(len(files))
		//filepath := absPath + "/" + "testdata" + string(len(files)) + ".txt"
		filepath := absPath + "/" + "testdata" + ".txt"
		fmt.Println(filepath)
		file, err := os.Create(filepath)
		if err != nil{log.Println("failed to create file")}
		defer file.Close()

		//jsonWriter := json.NewWriter(file)
		//defer jsonWriter.Flush()
		//jsonWriter.Write(string("hello"))
		file.WriteString("hello")

		mode := &serial.Mode{
				BaudRate: 115200,
		}
		port, err := serial.Open("/dev/ttyACM0", mode)
		if err != nil{log.Println("failed to open port")}
		
		defer port.Close()

		serReader := bufio.NewReader(port)
		
		for {
				msgByte, err := serReader.ReadBytes('\n')
				if err != nil{log.Println("serial read failed")}
				//msgStr := strings.TrimRight(string(msgByte), "\r\n")
				file.WriteString(string(msgByte))
				//strsplit := strings.Split(msgStr, " ")
				//fmt.Println("%q", strsplit)
				//fmt.Println(len(strsplit))
		}
		
		//path := absPath + "/" + "

}
