import math
import copy
import time
from pprint import pprint
from collections import defaultdict

def data():
    # return [int(x) for x in "0,3,6".split(",")]
    # return [int(x) for x in "1,3,2".split(",")]
    # return [int(x) for x in "1,2,3".split(",")]
    return [int(x) for x in "18,8,0,5,4,1,20".split(",")]

def init_numbers(mem, dat):
    for idx, v in enumerate(dat):
        mem[v] = {
            "last_idx": idx,
            "first_idx": idx,
            # "cnt": 1,
            # "list": [idx],
        }


def say_number(mem, idx, v):
    if v not in mem:
        mem[v] = {
            "last_idx": idx,
            "first_idx": idx,
            # "cnt": 1,
            # "list": [idx],
        }
        return 0
    
    r = idx - mem[v]["last_idx"]
    mem[v]["last_idx"] = idx
    # mem[v]["cnt"] += 1
    # mem[v]["list"].append(idx)
    return r

def main():
    dat = data()
    print(dat)

    mem = {}
    init_numbers(mem, dat)
    idx = len(dat) - 1
    last_r = dat[-1]
    while idx != 2020:
        v = last_r
        r = say_number(mem, idx, v)
        if idx == 2020-1:
            print(idx+1, "\t", v, "\t", r)
        idx += 1
        last_r = r
    
    print("---")

    mem = {}
    init_numbers(mem, dat)
    idx = len(dat) - 1
    last_r = dat[-1]
    while idx != 30000000:
        v = last_r
        r = say_number(mem, idx, v)
        if (idx % 1000000) == 1000000-1:
            print(idx+1, "\t", v, "\t", r)
        idx += 1
        last_r = r


main()

