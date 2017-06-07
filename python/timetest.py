import datetime,time

#now = time.strftime("%Y-%m-%d %H:%M:%S")
#today = time.strftime("%Y-%m-%d 00:00:00")

now = datetime.datetime.now()
delta = datetime.timedelta(days = -365)
daysAfter = now + delta

today = now.strftime("%Y-%m-%d 00:00:00")
yearAgo = daysAfter.strftime("%Y-%m-%d 00:00:00")

print now
print daysAfter

print today
print yearAgo
daysAfter = (now + delta).strftime("%Y-%m-%d 00:00:00")
print daysAfter
