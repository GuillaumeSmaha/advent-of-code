
import os
import copy

def loadDataFile(file):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = []
    with open(os.path.join(dir_path, file), "r") as f:
        for line in f:
            if line:
                d = line.strip().strip("<>").split(", ")
                s = {
                    "id": len(data),
                    "pos":{
                        "x": int(d[0].split('=')[1]),
                        "y": int(d[1].split('=')[1]),
                        "z": int(d[2].split('=')[1]),
                    },
                    "vel": {
                        "x": 0,
                        "y": 0,
                        "z": 0,
                    }
                }
                data.append(s)
    return lambda: data

# def loadDataFile():
#     dir_path = os.path.dirname(os.path.realpath(__file__))
#     with open(os.path.join(dir_path, "data.txt"), "r") as f:
#         data = f.readline()
#         return [int(i) for i in str(data.strip())]
#     return []

def loadSample():
    return [int(i) for i in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]

def loadTest():
    return [int(i) for i in "1,0,0,3,99".split(",")]

def loadStr(s):
    return [int(i) for i in s]


def get_layers(s, w, t):
    layers = []
    nb = len(s) // (w * t)
    for i in range(nb):
        pos = i * (w * t)
        layers.append({
            'data': [],
            'xy': [],
        })
        layers[i]['data'] = s[pos:pos+(w * t)]
    return layers

def find_fewest_zero(layers):
    m = layers[0]['data'].count(0)
    mL = 0
    for i, l in enumerate(layers):
        c = l['data'].count(0)
        print(c)
        if m > c:
            m = c
            mL = i
    if mL == 100000000:
        mL = 0
    return mL



def manathan(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1])

prems = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31]



def potential_energy(m):
    t = 0
    for p in ['x', 'y', 'z']:
        t += abs(m['pos'][p])
    return t

def kinetic_energy(m):
    t = 0
    for p in ['x', 'y', 'z']:
        t += abs(m['vel'][p])
    return t

def energy(m):
    return kinetic_energy(m) * potential_energy(m)

def process(loadData):
    moons = loadData()


    for s in range(1, 1001):
        # gravity
        for m1 in moons:
            for m2 in moons:
                if m1["id"] >= m2["id"]:
                    continue

                for p in ['x', 'y', 'z']:
                    if m1['pos'][p] < m2['pos'][p]:
                        m1['vel'][p] += 1
                        m2['vel'][p] -= 1
                    elif m1['pos'][p] > m2['pos'][p]:
                        m1['vel'][p] -= 1
                        m2['vel'][p] += 1
        # vel
        for m in moons:
            for p in ['x', 'y', 'z']:
                m['pos'][p] += m['vel'][p]
        
        t = 0
        print("Step {}".format(s))
        for m in moons:
            e = energy(m)
            t += e
            print("id={} pos=<x={}, y=  {}, z= {}>, vel=<x= {}, y= {}, z= {}>, e={}".format(m["id"], m['pos']['x'], m['pos']['y'], m['pos']['z'], m['vel']['x'], m['vel']['y'], m['vel']['z'], e))
        print("total energy: {}".format(t))




def main():
    # process(loadDataFile("test1.txt"))
    # process(loadDataFile("test2.txt"))
    # process(loadDataFile("test3.txt"))
    # process(loadDataFile("test4.txt"))
    # process(loadDataFile("test5.txt"))
    process(loadDataFile("data.txt"))

main()

