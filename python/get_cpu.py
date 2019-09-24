import psutil
import time

p = psutil.Process(1)

while 1:
    v = str(p.cpu_percent())
    if "0.0" != v:
        print(111111, v, time.time())
    time.sleep(1)

