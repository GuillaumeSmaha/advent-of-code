import math
from pprint import pprint
from collections import defaultdict


def data():
    lines = []
    # with open("list.test.txt", "r") as f:
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

def find_invalid(d, start, max):
    if start < max:
        start = max
    for di,dv in enumerate(d):
        if di < start:
            continue
        w = window(d, di, max)
        f = False
        for i in w:
            for j in w:
                if i != w and (i+j) == dv:
                    f = True
                    break
            if f:
                break

        if not f:
            return [di, dv]
    return None

def main():
    max = 25
    # max = 5
    d = data()
    invalid_i, invalid_v = find_invalid(d, 0, max)
    print("find_invalid", invalid_i, invalid_v)
    for s in range(2, len(d)+1):
        print("check window of size:", s)
        for di,dv in enumerate(d):
            if di < s:
                # print("pass", di, s)
                continue
            w = window(d, di, s)
            # print(di, s, w)
            ss = 0
            for j in w:
                ss += j
            if ss == invalid_v:
                print("found", di, dv, w)
                wm = w[0]
                wM = w[0]
                for v in w:
                    if v < wm:
                        wm = v
                    if v > wM:
                        wM = v
                print("min max:", wm, wM)
                print("res:", (wm + wM))
                return

main()

