from base import *

year = 673660
i = 1
days = 0
while ( i < year):
    days += daysYear(i)
    i +=1
print(days)
print("leap" if leap(year) else "non leap")
month = MONTH.index("Ims")
for i in range(month):
    print(daysMonth(year, i + 1))
    days += daysMonth(year, i + 1)
print(days)
days += 4
print(days)
