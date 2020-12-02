import math
from pprint import pprint
from collections import defaultdict


def data():
    text = ""
    with open("list.txt", "r") as f:
        for line in f:
            text += line.strip()
    
    r = {}
    t = [x for x in text.split(".") if len(x) > 0]
    for v in t:
        p = v.split(" contain ")
        pp = p[0].split(" ")
        k = " ".join(pp[0:-1])
        r[k] = defaultdict(lambda: 0)
        if p[1] != "no other bags":
            for vv in p[1].split(", "):
                vvv = vv.split(" ")
                r[k][" ".join(vvv[1:-1])] += int(vvv[0])
    return r

def data_expand_key(d, k, m):
    r = {
        k: m
    }
    if k in d:
        for kk, v in d[k].items():
            for rk, rv in data_expand_key(d, kk, m*v).items():
                if rk not in r:
                    r[rk] = 0
                r[rk] += rv
    return r

def data_expand():
    r = {}
    d = data()
    for k, v in d.items():
        r[k] = data_expand_key(d, k, 1)
        del r[k][k]
    return r

def main():
    c = 0
    d = data_expand()
    for _, v in d["shiny gold"].items():
        c += v
    print(c)

main()

