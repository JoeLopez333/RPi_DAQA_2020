package main


import (
		"github.com/tarm/serial"
		"log"
		"fmt"
		"encoding/json"
		"io/ioutil"
)

type Data struct{
	data []byte
}


func main(){
		fmt.Println("opening")
		c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
		fmt.Println("opened")
		
		s, err:= serial.OpenPort(c)
		if err != nil{
				log.Fatal(err)
		}
		
		n := 8
	//	buf := make([]byte, 8)
		mydata := Data{
				data : make([]byte, 8),
		}


		err = s.Flush()
		mydata.data[0] = 0

		for {
			# log.Println("reading")
			# n, err = s.Read(mydata.data)
			#fmt.Println("%q", mydata.data)
			# log.Println("read?")
			msgByte, err := serial.ReadBytes('\n')
			if err != nil{
					log.Fatal(err)
			}
		}
		log.Println("n is %d ", n)	
		log.Println("%q", mydata.data[:n])
		err = s.Close()

		file, _ := json.MarshalIndent(mydata, ""," ")

		_ = ioutil.WriteFile("test.json", file, 0644)
}
		
