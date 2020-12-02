
import os
import copy


def loadDataFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    with open(os.path.join(dir_path, "data.txt"), "r") as f:
        data = f.readline()
        return [int(i) for i in str(data.strip())]
    return []

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
        if m > c:
            m = c
            mL = i
    return mL

def process(loadData, w, t):
    s = loadData()
    print("")
    print("---")
    print("process:")
    layers = get_layers(s, w, t)

    res = [2] * len(layers[0]['data'])
    for p in range(len(layers[0]['data'])):
        for l in layers:
            if l['data'][p] != 2:
                res[p] = l['data'][p]
                break
    
    print(res)
    print("".join([str(i) for i in res]))

    for tt in range(t):
        a = [str(i) for i in res[tt*w:(tt+1)*w]]
        a = ["0" if i == "1" else " " for i in a]
        print("".join(a))

def main():
    process(lambda: loadStr("0222112222120000"), 2, 2)
    process(loadDataFile, 25, 6)

main()

