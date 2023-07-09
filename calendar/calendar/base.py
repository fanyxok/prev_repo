def bitCount(n):
    count = 0
    while (n):
        count += n & 1
        n >>= 1
    return count
def leap(year):
    if (year + bitCount(year))%8 == 0:
        return True
    else:
        return False
def monthsYear(year):
    if (leap(year)):
        return 10
    else:
        return 9

def leapMonthIdx(year):
    leapMonth = (abs(year-bitCount(year))%9)
    return leapMonth
def daysYear(year):
    if (leap(year)):
        leapMonth = (abs(year-bitCount(year))%9)+1
        return 360 + 39 + (abs(10*year-leapMonth)%3)
    else:
        return 360

def daysMonth(year, month):
    return (39 + (10*year - month)%3)

MONTH = ["Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"]
