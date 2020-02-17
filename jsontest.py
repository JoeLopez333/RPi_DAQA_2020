import json
import os
import datetime
import time
import serial, string

ser = serial.Serial('/dev/ttyACM0', 115200, 8, 'N', 1, timeout=1)
directory = "./json_logs"

json_parsed_dict = []
can_msg = []

#format time into filename
time = datetime.datetime.now().strftime("%m_%d_%H_%M")
output = " "
filename = time

with open("./json_logs/" + filename + '.json', 'w') as outfile:
    print("openinging " + str(outfile))
    while True:
        #dict_entry = {"time" : datetime.datetime.now().strftime("%H_%M_%S"),
        #        "text" : "hello"}
        # time.sleep(1)
        # json_parsed_dict.append(dict_entry)
        while output != "":
            output = ser.read(19)
            msg = output[1:13]
            time = output[13:18]
            strlength = len(output)
            if strlength > 3:
                print("Message: " + str(msg))
                # print(output[2:strlength+1])
                # print("fuck")
                dict_entry = {"msg" : str(msg), "time" : str(time)}
                json.dump(dict_entry, outfile)
                outfile.write('\n')
            else:
                print("Something went wrong")
            print(output[18])
        output = " "
        print("hello")
        json.dump(dict_entry, outfile)




