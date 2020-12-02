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

def main():
    max = 25
    d = data()
    d.sort()

    l = []
    p = 0
    vdiff = defaultdict(lambda: 0)
    for v in d:
        if p < (v+3):
            l.append(v)
            vdiff[v - p] += 1
            p = v
            print("add ", v)
        else:
            print(l)
            print("fail for ", v)
            break

    vdiff[3] += 1
    print("Final list", l)
    print("max:", l[-1] + 3)
    print("diff:", vdiff)
    print("res:", vdiff[1] * vdiff[3])

main()

