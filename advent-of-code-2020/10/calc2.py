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

def validate_data(d):
    p = 0
    for v in d:
        if p >= (v+3):
            return False
    return True

def main():
    max = 25
    d = data()
    d.sort()

    l = []
    p = 0
    ldiff = []
    vdiff = defaultdict(lambda: 0)
    for v in d:
        if p < (v+3):
            l.append(v)
            ldiff.append(v - p)
            vdiff[v - p] += 1
            p = v
            print("add ", v)
        else:
            print(l)
            print("fail for ", v)
            break

    ldiff.append(3)
    vdiff[3] += 1

            
    print("Final list", [0]+l)
    print("Final diff list", ldiff)
    print("max:", l[-1] + 3)
    print("diff:", vdiff)
    print("res:", vdiff[1] * vdiff[3])

    cmps = []
    cmp = 0
    for v in ldiff:
        if v == 1:
            cmp += 1
        else:
            if cmp > 1:
                cmps.append(cmp - 1)
            cmp = 0
    if cmp > 1:
        cmps.append(cmp - 1)

    t = 1
    for c in cmps:
        if c < 3:
            t *= 2**(c)
        else:
            t *= 2**(c) - (c - 2)

    print("cmps:", cmps)
    print("result:", t)

main()

