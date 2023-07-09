from base import * 


days = 236057086 + 4728
year = 1
while (days > daysYear(year)):
    days -= daysYear(year)
    year += 1

print("year ", year, " day", days)

# days2date right
# date2days error