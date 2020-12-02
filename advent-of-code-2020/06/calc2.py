import math
from pprint import pprint

def init_dict():
    r = {}
    for o in range(97,124):
        r[chr(o)]=False
    return r

def init_user(u):
    r = {
        'qf': init_dict(),
        'q': {},
    }
    for c in u:
        r['qf'][c]=True
        r['q'][c]=True
    return r

def data():
    t = []
    g = {
        'qf': init_dict(),
        'q': {},
        'p': [],
    }
    with open("list.txt", "r") as f:
        for line in f:
            if line == "\n":
                if g['p']:
                    t.append(g)
                g = {
                    'qf': init_dict(),
                    'q': {},
                    'p': [],
                }
            elif line.strip():
                g['p'].append(init_user(line.strip()))
                for c in line.strip():
                    g['qf'][c]=True
                    g['q'][c]=True
    
    if g['p']:
        t.append(g)
    return t

def main():
    c = 0
    for g in data():
        q = {}
        for k, v in g['q'].items():
            r = True
            for p in g['p']:
                if k not in p['q']:
                    r = False
                    break
            if r:
                q[k] = True

        c += len(q)
    print(c)

main()

