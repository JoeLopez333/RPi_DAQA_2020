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
    #gps vid is 5446
    if (port.vid == 5446):
        ser = serial.Serial(port.device, 115200, 8, 'N', timeout=1)
        print("found gps")

CS = DigitalInOut(board.CE1)
RESET = DigitalInOut(board.D25)
spi = busio.SPI(board.SCK, MOSI=board.MOSI, MISO=board.MISO)

rfm9x = adafruit_rfm9x.RFM9x(spi, CS, RESET, 915.0, baudrate=1000000)
rfm9x.tx_power = 23
packet = None
packet_count = 0
time.sleep(6)
files = os.listdir("/home/pi/testingdata")
print(len(files))
now = datetime.now()
timestr = now.strftime("%m-%d_%I-%M")
gpslog = open(r"/home/pi/testingdata/gps" +timestr+".txt" , "a")
#gpslog.write("Hello \n")
print("hello")
comfile = open(r"/home/pi/testingdata/communication.txt", "r+")
str8 = comfile.readline()
logcount = 0
while str8:
    gpslog.write(str8)
    rfm9x.send(bytes(str8, "utf-8"))
    str8 = comfile.readline()
    logcount = logcount+1

os.remove("/home/pi/testingdata/communication.txt")

while True:
    #rfm9x.send(bytes("logging" + str(logcount) + " things\n", "utf-8"))
    packet = rfm9x.receive(timeout=.01)
    if packet != None: 
        packet_count = packet_count + 1
        ser.write(packet)
        print(packet)
    packet = None
    gpsmsg = ser.read_until()
    gpsstring = str(gpsmsg)
    #print(gpsstring)
    now = datetime.now()
    timestr = now.strftime("%H-%M-%S-%f")[:-3]
    #print(timestr)
    # print(gpsstring)
    gpslog.write(gpsstring[:len(gpsstring)-5] +'\'' + " " + timestr + "\n")
