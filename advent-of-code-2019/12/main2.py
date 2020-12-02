
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

def unique(list1): 
  
    # intilize a null list 
    unique_list = [] 
      
    # traverse for all elements 
    for x in list1: 
        # check if exists in unique_list or not 
        if x not in unique_list: 
            unique_list.append(x) 
    return unique_list

combs=[]
combsX=[]
combsY=[]
combsZ=[]
combsV=[]
import copy
tt = 0
eee=[]

def process(loadData):
    moons = loadData()

    cx = []
    cy = []
    cz = []
    # v = []
    for m in moons:
        cx.append(str(m['pos']['x']))
        cy.append(str(m['pos']['y']))
        cz.append(str(m['pos']['z']))
    
    cx = "|".join(cx)
    cy = "|".join(cy)
    cz = "|".join(cz)

    combsX.append(cx)
    combsY.append(cy)
    combsZ.append(cz)

    fnd = [None] * 3
    steps = [0] * 3

    for s in range(1, 100000000):
        # if s % 1000 == 0:
        #     print(s)

        if s != 0 and s % 1000000 == 0:
            # print(s)
            # u = unique(combs)
            # print(u)
            # print(len(u))

        #     u = unique(combsV)
        #     print(u)
        #     print(len(u))
            # print(eee)
            # u = unique(eee)
            # print(len(eee))
            # print(len(u))
            break

        # copie velm
        for m in moons:
            m['velbak'] = copy.copy(m['vel'])


        # gravity
        if s != 0:
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
        c = []
        e = 0
        cx = []
        cy = []
        cz = []
        # v = []
        for m in moons:
            m['veldiff'] = {}
            for p in ['x', 'y', 'z']:
                m['pos'][p] += m['vel'][p]
                m['veldiff'][p] = m['vel'][p] - m['velbak'][p]

            # c.append("{}/{}/{}".format(m['veldiff']['x'], m['veldiff']['y'], m['veldiff']['z']))
            # c.append(str(kinetic_energy(m)))
            # c.append("{}/{}/{}".format(m['pos']['x'], m['pos']['y'], m['pos']['z']))
            # v.append("{}/{}/{}".format(m['vel']['x'], m['vel']['y'], m['vel']['z']))
            # c.append("{}/{}/{},{}/{}/{}".format(m['veldiff']['x'], m['veldiff']['y'], m['veldiff']['z'], m['vel']['x'], m['vel']['y'], m['vel']['z']))
            cx.append(str(m['pos']['x']))
            cy.append(str(m['pos']['y']))
            cz.append(str(m['pos']['z']))
        
        cx = "|".join(cx)
        cy = "|".join(cy)
        cz = "|".join(cz)
        # c = "|".join(c)
        # v = "|".join(v)

        # if v in combsV:
        #     print("CCCC")
        #     print(s)
        #     break
        # combsV.append(v)
        # if cx in combsX:
        if not fnd[0] and len(combsX) > 0 and cx == combsX[0]:
            print("---\nCCCC X")
            print("s=",s)
            print("c=",cx)
            fnd[0] = True
            steps[0] = s
            
        
        if not fnd[1] and len(combsY) > 0 and cy == combsY[0]:
            print("---\nCCCC Y")
            print("s=",s)
            print("c=",cy)
            fnd[1] = True
            steps[1] = s
        
        if not fnd[2] and len(combsZ) > 0 and cz == combsZ[0]:
            print("---\nCCCC Z")
            print("s=",s)
            print("c=",cz)
            fnd[2] = True
            steps[2] = s

        
        # if all(fnd):
        #     break
    
    print(steps)
    print(steps[0] * steps[1] * steps[2])
            # eee.append(s)
            # eee.append(combs.index(c))
            # break
        # if s < 1001:
        #     combs.append(c)
        # combs.append(c)

        # 28854
        # t = 0
        # print("Step {}".format(s))
        # for m in moons:
            # e = energy(m)
            # t += e
            # print("id={} pos=<x={}, y=  {}, z= {}>, vel=<x= {}, y= {}, z= {}>, e={}".format(m["id"], m['pos']['x'], m['pos']['y'], m['pos']['z'], m['vel']['x'], m['vel']['y'], m['vel']['z'], e))
        # print(c)
        # print("total energy: {}".format(t))

    cyclex = 161428
    cycley = 231614
    cyclez = 102356
    print(cyclex * cycley * cyclez)


    # 3826986927369952
    # 478373365921244
def main():
    # process(loadDataFile("test1.txt"))
    # process(loadDataFile("test2.txt"))
    # process(loadDataFile("test3.txt"))
    # process(loadDataFile("test4.txt"))
    # process(loadDataFile("test5.txt"))
    process(loadDataFile("data.txt"))

main()

