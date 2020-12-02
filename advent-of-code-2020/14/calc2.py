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


def get_addrs(v):
    r = []
    try:
        i = v.index('X')
        r += get_addrs(v[0:i] + "0" + v[i+1:])
        r += get_addrs(v[0:i] + "1" + v[i+1:])
    except Exception:
        r.append(v)
    return r

def get_masks(mask):
    masks = []
    mask = mask.replace("0", "Y")
    for m in get_addrs(mask):
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
        masks.append(mask)
    return masks

def mem_set(masks, mem, k, v):
    for m in masks:
        a = k & m["0"] | m["1"]
        # print("set val {0:b} mem {0:b}".format(v, a))
        mem[a] = v

def main():
    dat = data()
    print(dat)

    masks = {}
    mem = {}
    for d in dat:
        if d[0] == "mask":
            masks = get_masks(d[1])
        else:
            mem_set(masks, mem, d[1], d[2])

    s = 0
    for p, m in mem.items():
        s += m
    print("Sum:", s)


main()

