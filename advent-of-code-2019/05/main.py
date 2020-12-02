
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


def loadSample():
    return [int(i) for i in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]

def loadTest():
    return [int(i) for i in "1,0,0,3,99".split(",")]

def loadStr(s):
    return [int(i) for i in s.split(",")]

Input = 1

def get_val(s, p, immediate):
    if immediate:
        return s[p]
    return s[s[p]]

def exec_op(s, op, p, immediate=[False, False]):
    if op == 1:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) + get_val(s, p+2, immediate[1])
        p += 4
    elif op == 2:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) * get_val(s, p+2, immediate[1])
        p += 4
    elif op == 3:
        s[s[p+1]] = Input
        p += 2
    elif op == 4:
        print(s[s[p+1]])
        # if immediate[0]:
        #     print(s[s[p+1]])
        # else:
        #     print(s[p+1])
        p += 2
    elif op == 99:
        return False
    elif op > 99:
        val = str(op).zfill(5)
        code_op = int(val[-2:])
        p = exec_op(s, code_op, p, immediate=[val[-3] == "1", val[-4] == "1", val[-5] == "1"])
    return p


def process(loadData):
    s = loadData()
    p = 0
    c = True
    print("")
    print("---")
    print("process:")
    while p is not False:
        p = exec_op(s, s[p], p)
    print("final:" + ",".join([str(i) for i in s]))

def main():
    process(lambda: loadStr("3,0,4,0,99"))
    process(lambda: loadStr("1002,4,3,4,33"))
    process(lambda: loadStr("1101,100,-1,4,0"))
    process(loadDataFile)

main()

