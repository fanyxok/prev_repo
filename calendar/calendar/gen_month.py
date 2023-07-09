from base import *
import pandas as pd
MONTH = []
MONTH.append([])
for i in range(1,1000000):
    l = ["Sist", "Spst", "Slst", "Sem", "Sca", "Ims", "Ihuman", "Siais", "Ih"]
    if leap(i):
        k = leapMonthIdx(i)
        print(k)
        l.insert(k, l[k])
        l[k+1] = l[k+1].upper()
    MONTH.append(l)
df = pd.DataFrame(data = MONTH)
df.to_csv("./months.csv",encoding="utf-8")


