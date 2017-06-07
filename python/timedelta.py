import datetime

yearAgoStr = (datetime.datetime.now() + datetime.timedelta(days = -365)).strftime("%Y-%m-%d 00:00:00")
todayStr = datetime.datetime.now().strftime("%Y-%m-%d 00:00:00")

cond = "where create_time >= '%s' and create_time < '%s'" %(yearAgoStr, todayStr)

print cond
