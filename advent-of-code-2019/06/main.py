
import os 
from collections import defaultdict


def loadDataFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = defaultdict(lambda: [])
    with open(os.path.join(dir_path, "data.txt"), "r") as f:
        for line in f:
            if line:
                s = line.split(")")
                src = s[1].strip("\n")
                dest = s[0].strip("\n")
                data[dest] # will add root
                data[src].append(dest)
    return data

def loadSampleFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = defaultdict(lambda: [])
    with open(os.path.join(dir_path, "sample.txt"), "r") as f:
        for line in f:
            if line:
                s = line.split(")")
                src = s[1].strip("\n")
                dest = s[0].strip("\n")
                data[dest] # will add root
                data[src].append(dest)
    return data


def get_orbit(s, a):
    direct = s[a]
    indirect = []
    for o in direct:
        d, i = get_orbit(s, o)
        indirect += d + i
    return (direct, indirect)


def process(loadData):
    s = loadData()
    p = 0
    c = True
    t = 0
    for o in s:
        d, i = get_orbit(s, o)
        print("{}: {} direct & {} indirect".format(o, len(d), len(i)))
        t += len(d) + len(i)
    print("total:", t)

def main():
    process(loadSampleFile)
    process(loadDataFile)

main()

