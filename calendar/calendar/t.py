def bitCount(n):
    count = 0
    while (n):
        count += n & 1
        n >>= 1;
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
def daysYear(year):
    if (leap(year)):
        leapMonth = (abs(year-bitCount(year))%9)+1
        return 360 + 39 + (abs(10*year-leapMonth)%3)
    else:
        return 360
def to_c_array(values, ctype="float", name="table", formatter=str, colcount=10):
    # apply formatting to each element
    values = [formatter(v) for v in values]

    # split into rows with up to `colcount` elements per row
    rows = [values[i:i+colcount] for i in range(0, len(values), colcount)]

    # separate elements with commas, separate rows with newlines
    body = ',\n    '.join([', '.join(r) for r in rows])

    # assemble components into the complete string
    return '{} {}[] = {{\n    {}}};'.format(ctype, name, body)
4000
l = []
days = 0
for i in range(1,1014560):
    days+=monthsYear(i)
    if (i%4000 == 0):
        l.append(days)
        days = 0
print(len(l))
c = to_c_array(l)
print(c)
