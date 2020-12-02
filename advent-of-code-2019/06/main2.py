
import os 
from collections import defaultdict

def loadFile(filename):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = defaultdict(lambda: [])
    with open(os.path.join(dir_path, filename), "r") as f:
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

def intersection(lst1, lst2): 
    lst3 = [value for value in lst1 if value in lst2] 
    return lst3 
  

def process(loadData):
    s = loadData()
    p = 0
    c = True
    t = 0
    orbits = {}
    for o in s:
        d, i = get_orbit(s, o)
        print("{}: {} direct & {} indirect".format(o, len(d), len(i)))
        orbits[o] = d + i
    
    t = 0
    sanYou = intersection(orbits["SAN"], orbits["YOU"])
    
    t = len(orbits["SAN"]) + len(orbits["YOU"]) - 2 * len(sanYou)
    print("total:", t)

def main():
    # process(lambda: loadFile("sample.txt"))
    process(lambda: loadFile("sample2.txt"))
    process(lambda: loadFile("data.txt"))

main()

