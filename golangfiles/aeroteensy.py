import time
import busio
from digitalio import DigitalInOut, Direction, Pull
import board
import adafruit_rfm9x
import serial, string
import os
from datetime import datetime
import serial.tools.list_ports

output = " "

#ser = serial.Serial('/dev/ttyACM1', 115200, 8, 'N', 1, timeout=1)
ports = serial.tools.list_ports.comports()
for port in ports:
    print(port.serial_number)
    #look for aero teensy
    if (port.serial_number == "6423410"):
        ser = serial.Serial(port.device, 115200, 8, 'N', timeout=1)
        print("aero teensy")

time.sleep(10)
files = os.listdir("/home/pi/testingdata")
#print(len(files))
now = datetime.now()
timestr = now.strftime("%m-%d_%I-%M")
aerolog = open(r"/home/pi/testingdata/aero" +timestr+".txt" , "a")
ser.flushInput()

while 1:
    msg = ser.read_until()
    aerostr = str(msg)
    now = datetime.now()
    timestr = now.strftime("%H-%M-%S-%f")[:-3]
    aerolog.write(aerostr + " " +timestr + "\n")
    print(msg[1:-2])

#time.sleep(10)

#files = os.listdir("/home/pi/testingdata")



