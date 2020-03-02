"""
Wiring Check, Pi Radio w/RFM9x
 
Learn Guide: https://learn.adafruit.com/lora-and-lorawan-for-raspberry-pi
Author: Brent Rubell for Adafruit Industries
"""
import time
import busio
from digitalio import DigitalInOut, Direction, Pull
import board
# Import the SSD1306 module.
#import adafruit_ssd1306
# Import the RFM9x radio module.
import adafruit_rfm9x

CS = DigitalInOut(board.CE1)
RESET = DigitalInOut(board.D25)
spi = busio.SPI(board.SCK, MOSI=board.MOSI, MISO=board.MISO)
# spi.max_speed_hz=5000000
# while not spi.try_lock():
#     pass
# spi.try_lock()
# spi.configure(baudrate=433000000, phase=0, polarity=0)
rfm9x = adafruit_rfm9x.RFM9x(spi, CS, RESET, 915.0, baudrate=1000000)
rfm9x.tx_power = 20
prev_packet = None

msg_count = 0;

while True:
    # packet = rfm9x.receive()
    # ife_string = str(formula_electric)
    data = bytes("ife_string", "utf-8")
    print ("send " + str(msg_count))
    rfm9x.send(b'\x34\x34\x34\x34ABC TEST !!!')     
    msg_count = msg_count + 1
    time.sleep(1)


#while True:
 
    # Attempt to set up the RFM9x Module
#    try:
#        rfm9x = adafruit_rfm9x.RFM9x(spi, CS, RESET, 915.0,baudrate=1000000)
#        print("hello")
#    except RuntimeError as error:
        # Thrown on version mismatch
        # display.text('RFM9x: ERROR', 0, 0, 1)
#        print('RFM9x Error: ', error)

    #  display.show()
   #  time.sleep(1)



