import math
from pprint import pprint
from collections import defaultdict


def data():
    lines = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                lines.append(int(line.strip()))
    
    return lines

def chunks(lst, n):
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]



def window(d, i, max=25):
    if max > len(d):
        max = len(d)
    m = i - max
    M = i
    if (len(d)-m) < max:
        m = len(d) - max
        M = len(d)
    return d[m:M]

def main():
    max = 25
    d = data()
    print(d)
    print(d)
    for di,dv in enumerate(d):
        if di < max:
            print("pass", di, dv)
            continue
        w = window(d, di, max)
        print(di, dv, w)
        f = False
        for i in w:
            for j in w:
                if i != w and (i+j) == dv:
                    f = True
                    break
            if f:
                break

        if not f:
            print("break", di, dv)
            break

main()

