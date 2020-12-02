import math
import copy
import time
from pprint import pprint
from collections import defaultdict

def data():
    r = []
    with open("list.txt", "r") as f:
        r.append(int(f.readline().strip()))
        r.append([int(x) for x in  f.readline().strip().split(",") if x != "x"])
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

    start = d[0]
    dmin = 999999999999999999999
    dmin_bus = -1
    for b in d[1]:
        d = b * ( (start // b) + 1) - start
        if d < dmin:
            dmin_bus = b
            dmin = d
    
    print("Bus:", dmin_bus)
    print("Ealier depart:", dmin)
    print("Res:", dmin * dmin_bus)

main()

