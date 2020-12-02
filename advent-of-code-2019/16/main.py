
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

INPUT = "59704176224151213770484189932636989396016853707543672704688031159981571127975101449262562108536062222616286393177775420275833561490214618092338108958319534766917790598728831388012618201701341130599267905059417956666371111749252733037090364984971914108277005170417001289652084308389839318318592713462923155468396822247189750655575623017333088246364350280299985979331660143758996484413769438651303748536351772868104792161361952505811489060546839032499706132682563962136170941039904873411038529684473891392104152677551989278815089949043159200373061921992851799948057507078358356630228490883482290389217471790233756775862302710944760078623023456856105493"


def get_pattern(pattern, pos):
    if pos == 0:
        return pattern
    
    res = []
    for p in pattern:
        res += [p] * (pos + 1)
    return res

def calculate_phase(inputt, pattern):
    res = ""
    for s in range(len(inputt)):
        p = get_pattern(pattern, s)
        s = 0
        i_pattern = s + 1
        for d in inputt:
            s += int(d) * p[i_pattern]
            i_pattern = (i_pattern + 1) % len(p)
        res += str(s)[-1]
    return res

def get_phase(c, nb):
    for i in range(0, nb):
        c = calculate_phase(c, [0, 1, 0, -1])
    return c


def main():
    print("---")
    c = "80871224585914546619083218645595"
    a = get_phase(c, 100)
    print(c, a, a[0:8])

    print("---")
    c = "19617804207202209144916044189917"
    a = get_phase(c, 100)
    print(c, a, a[0:8])

    print("---")
    c = "69317163492948606335995924319873"
    a = get_phase(c, 1)
    print(c, a, a[0:8])
    a = get_phase(c, 10)
    print(c, a, a[0:8])
    a = get_phase(c, 100)
    print(c, a, a[0:8])
    a = get_phase(c, 1000)
    print(c, a, a[0:8])
    a = get_phase(c, 10000)
    print(c, a, a[0:8])
    # print(get_phase(INPUT, 10000)[0:8])

# o    : 6335995924319873
# 1    :                3
# 10   :               73
# 100  :             9873
# 1000 :         24319873
# 10000: 6335995924319873

main()
