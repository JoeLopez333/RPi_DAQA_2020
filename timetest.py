import datetime
import time

var = datetime.datetime.now().time()
count = 0;

while (count < 5):
    count = count + 1
    time.sleep(1)
    print(var)
    var = datetime.datetime.now().time()


