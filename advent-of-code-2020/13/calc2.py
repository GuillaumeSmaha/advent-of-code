import math
import copy
import time
from pprint import pprint
from collections import defaultdict


def data():
    r = []
    with open("list.txt", "r") as f:
        r.append(int(f.readline().strip()))
        g = {i:int(x) for i,x in enumerate(f.readline().strip().split(",")) if x != "x"}
        r.append(g)
    return r

def pgcd(a,b):
    while b!=0:
        r=a%b
        a,b=b,r
    return a

def ppcm(a,b):
    if (a==0) or (b==0):
        return 0
    else:
        return (a*b)//pgcd(a,b)

def main():
    d = data()
    print(d)

    opes = []
    for ix, dx in d[1].items():
        opes.append((ix, dx))
    print(opes)


    a1 = opes[0][1]
    t = 0
    for o in opes:
        i = 0
        while (t+a1*i + o[0]) % o[1] != 0:
                i += 1
        t = t + a1*i
        if o[0] != 0:
            a1 = a1*o[1]
    print("t:", t)
    return

main()

