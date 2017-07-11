# -*- coding: utf-8 -*-
class Driver:
    level = 5
    create_time = '2015-01-01 00:00:00'
if __name__ =="__main__":
    driverRecordsList = []
    for i in range(100):
        x = Driver()
        driverRecordsList.append(x)
        if i == 10:
            x.create_time = '2013-02-02 00:00:00'
            x.level = 1
        if i == 36:
            x.create_time = '2015-02-02 00:00:00'
            x.level = 2

    driverRecordsList = sorted(driverRecordsList, key=lambda record : record.create_time)
#    driverRecordsList[37].level = 4

    oldestIndex = -1
    # 先找出最早的非五星评价
    for i in range(len(driverRecordsList)):
        print i, driverRecordsList[i], driverRecordsList[i].level,driverRecordsList[i].create_time
        if driverRecordsList[i].level != 5:
            oldestIndex = i
            break
    print oldestIndex
