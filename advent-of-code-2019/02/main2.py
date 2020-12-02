
import os 


def loadDataFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    with open(os.path.join(dir_path, "data.txt"), "r") as f:
        data = f.readline()
        return [int(i) for i in data.split(",")]
    return []

def loadData():
    a = [int(i) for i in "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,13,19,1,9,19,23,1,6,23,27,2,27,9,31,2,6,31,35,1,5,35,39,1,10,39,43,1,43,13,47,1,47,9,51,1,51,9,55,1,55,9,59,2,9,59,63,2,9,63,67,1,5,67,71,2,13,71,75,1,6,75,79,1,10,79,83,2,6,83,87,1,87,5,91,1,91,9,95,1,95,10,99,2,9,99,103,1,5,103,107,1,5,107,111,2,111,10,115,1,6,115,119,2,10,119,123,1,6,123,127,1,127,5,131,2,9,131,135,1,5,135,139,1,139,10,143,1,143,2,147,1,147,5,0,99,2,0,14,0".split(",")]
    a[1] = 12
    a[2] = 2
    return a

def loadDataNounVerb(i, j):
    a = [int(i) for i in "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,13,19,1,9,19,23,1,6,23,27,2,27,9,31,2,6,31,35,1,5,35,39,1,10,39,43,1,43,13,47,1,47,9,51,1,51,9,55,1,55,9,59,2,9,59,63,2,9,63,67,1,5,67,71,2,13,71,75,1,6,75,79,1,10,79,83,2,6,83,87,1,87,5,91,1,91,9,95,1,95,10,99,2,9,99,103,1,5,103,107,1,5,107,111,2,111,10,115,1,6,115,119,2,10,119,123,1,6,123,127,1,127,5,131,2,9,131,135,1,5,135,139,1,139,10,143,1,143,2,147,1,147,5,0,99,2,0,14,0".split(",")]
    a[1] = i
    a[2] = j
    return a


def loadSample():
    return [int(i) for i in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]

def loadTest():
    return [int(i) for i in "1,0,0,3,99".split(",")]

def loadStr(s):
    return [int(i) for i in s.split(",")]


def op(s, p):
    if s[p] == 1:
        s[s[p+3]] = s[s[p+1]] + s[s[p+2]]
    elif s[p] == 2:
        s[s[p+3]] = s[s[p+1]] * s[s[p+2]]
    elif s[p] == 99:
        return False
    return True

def res(loadData):
    s = loadData()
    p = 0
    ops = []
    while op(s, p):
        p += 4
    return s[0]

def process(loadData):
    s = loadData()
    p = 0
    c = True
    ops = []
    while c:
        c = op(s, p)
        p += 4
    print("---")
    print(",".join([str(i) for i in s]))

def main():
    process(loadTest)
    process(loadSample)
    process(lambda: loadStr("1,0,0,0,99"))
    process(lambda: loadStr("2,3,0,3,99"))
    process(lambda: loadStr("2,4,4,5,99,0"))
    process(lambda: loadStr("1,1,1,4,99,5,6,0,99"))
    process(lambda: loadStr("1,1,1,4,99,5,6,0,99"))
    process(loadData)

    for i in range(0, 100):
        for j in range(0, 100):
            print(i, j)
            r = res(lambda: loadDataNounVerb(i, j))
            if r == 19690720:
                print("FOUND")
                return


main()

