import math
import copy
import time
from pprint import pprint
from collections import defaultdict

def data():
    r = []
    with open("list.txt", "r") as f:
        for l in f:
            l = l.strip()
            if l[0:4] == "mask":
                r.append(("mask", l[7:]))
            else:
                p = l[4:].split("]", 1)
                r.append(("mem", int(p[0]), int(p[1][3:])))
    return r


def mask_set(m):
    mask = {
        "m": m,
        "0": 2**(len(m)) - 1,
        "1": 0,
    }
    for i, v in enumerate(m[::-1]):
        if v == "0":
            mask["0"] -= 2**i
        elif v == "1":
            mask["1"] += 2**i

    return mask

def mem_set(mask, mem, k, v):
    mem[k] = v & mask["0"] | mask["1"]

def main():
    dat = data()
    print(dat)

    mask = {}
    mem = {}
    for d in dat:
        if d[0] == "mask":
            mask = mask_set(d[1])
        else:
            mem_set(mask, mem, d[1], d[2])

    s = 0
    for p, m in mem.items():
        s += m
    print("Sum:", s)


main()

