
import os
import copy
import math


def loadDataFile(file):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = {}
    with open(os.path.join(dir_path, file), "r") as f:
        for line in f:
            if line:
                l = line.strip().split(' => ')
                l[1] = l[1].split(' ')
                l[1][0] = int(l[1][0])

                l[0] = l[0].split(', ')

                data[l[1][1]] = {
                    'type': l[1][0],
                    'produce': l[1][0],
                    'needs': {},
                }

                for i, v in enumerate(l[0]):
                    l[0][i] = l[0][i].split(',')
                    for j in l[0][i]:
                        k = j.split(' ')
                        data[l[1][1]]['needs'][k[1]] = {
                            'type': k[1],
                            'produce': int(k[0]),
                        }

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

 
def pgcd(a,b):
    while b!=0:
        r=a%b
        a,b=b,r
    return a

def ppcm(a,b):
    if (a==0) or (b==0):
        return 0
    else:
        return (a*b)//pgcd(a,b)

def get_list(s, kind):
    print("-----")
    print("kind", kind)
    remain = 0
    if kind not in s:
        print("@@@@@")
        return ({
            kind: {
                'type': kind,
                'produce': 1,
                'needs': {},
            }
        }, remain)
    print("s[kind]", s[kind])
    d = copy.deepcopy(s[kind])
    res = {}
    for t, n in s[kind]['needs'].items():
        print("t", t)
        print("n", n)
        l, r = get_list(s, t)
        print("l", l)
        remain += r
        for _, nn in l.items():
            print("nn", nn)
            # nn['produce'] *= n['produce']
            # nn['produce'] *= n['produce']
            if nn['type'] in res:
                needs = dict(res[nn['type']]['needs'])
                needs.update(nn['needs'])
                res[nn['type']] = {
                    'type': nn['type'],
                    'produce': res[nn['type']]['produce'] + nn['produce'],
                    'needs': needs,
                }
            else:
                res[nn['type']] = nn

    m = s[kind]['produce']
    # remain = m - 
    print("m", m)
    # if res:
    #     gcd = math.gcd(m, list(res.values())[0]['produce'])
    #     for k, v in res.items(): 
    #         gcd = math.gcd(gcd, v['produce'])

    #     if gcd:
    #         for k, v in res.items(): 
    #             res[k]['produce'] =  v['produce'] // gcd
    print("res")
    print(res)
    print("@@@@@")
    return res, remain


def process(loadData):
    s = loadData()

    # a = get_list(s, 'A')
    a = get_list(s, 'FUEL')
    print('Result:')
    print(a)

def main():
    process(loadDataFile("test1.txt"))
    # process(loadDataFile("test2.txt"))
    # process(loadDataFile("test3.txt"))
    # process(loadDataFile("test4.txt"))
    # process(loadDataFile("test5.txt"))
    # process(loadDataFile("data.txt"))

main()

