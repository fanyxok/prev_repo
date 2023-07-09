from base import *
import random

sy = random.randrange(1, 100000)
sm = 0
if leap(sy):
    sm = random.randrange(0,10)
else:
    sm = random.randrange(0,9)

m = random.randrange(0,1000)
mstr = ""
l = ["Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"]
if leap(sy):
    k = leapMonthIdx(sy)
    l.insert(k, l[k])
    l[k+1] = l[k+1].upper()
    mstr = l[sm]
print(sy, mstr)

y = sy
while (m > 0):
    l = ["Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"]
    if leap(y):
        k = leapMonthIdx(y)
        l.insert(k, l[k])
        l[k+1] = l[k+1].upper()
    for i in range(kk, len(l)):
        m -= 1

