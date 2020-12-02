
import os
import copy


def loadDataFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    with open(os.path.join(dir_path, "data.txt"), "r") as f:
        data = f.readline()
        return [int(i) for i in data.split(",")]
    return []

def loadSample():
    return [int(i) for i in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]

def loadTest():
    return [int(i) for i in "1,0,0,3,99".split(",")]

def loadStr(s):
    return [int(i) for i in s.split(",")]

def get_val(s, p, immediate):
    if immediate:
        return s[p]
    return s[s[p]]

def exec_op(s, op, p, immediate=[False, False], intputAmp=[1]):
    output = None

    if op > 99:
        val = str(op).zfill(5)
        op = int(val[-2:])
        immediate=[val[-3] == "1", val[-4] == "1"]

    if op == 1:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) + get_val(s, p+2, immediate[1])
        p += 4
    elif op == 2:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) * get_val(s, p+2, immediate[1])
        p += 4
    elif op == 3:
        if intputAmp[0] >= len(intputAmp):
            return True, output
        s[s[p+1]] = intputAmp[intputAmp[0]]
        intputAmp[0] += 1
        p += 2
    elif op == 4:
        output = s[s[p+1]]
        p += 2
    elif op == 5:
        if get_val(s, p+1, immediate[0]) != 0:
            p = get_val(s, p+2, immediate[1])
        else:
            p += 3
    elif op == 6:
        if get_val(s, p+1, immediate[0]) == 0:
            p = get_val(s, p+2, immediate[1])
        else:
            p += 3
    elif op == 7:
        if get_val(s, p+1, immediate[0]) < get_val(s, p+2, immediate[1]):
            s[s[p+3]] = 1
        else:
            s[s[p+3]] = 0
        p += 4
    elif op == 8:
        if get_val(s, p+1, immediate[0]) == get_val(s, p+2, immediate[1]):
            s[s[p+3]] = 1
        else:
            s[s[p+3]] = 0
        p += 4
    elif op == 99:
        return (False, output)
    return (p, output)


def call_amp(s, intputAmp, p=0):
    output = []
    oldP = p
    # until stop or waiting for input
    while p is not False and p is not True:
        oldP = p
        p, o = exec_op(s, s[p], p, intputAmp=intputAmp)
        if o is not None:
            output.append(o)
    return output, p, oldP

def call_chaim_amp(sList, conf, inputInit):
    a,b,c,d,e = conf
    inputs = [
        [1, a] + inputInit,
        [1, b],
        [1, c],
        [1, d],
        [1, e],
    ]
    waitings = [True] * len(sList)
    poses = [0] * len(sList)
    while sum(waitings) != 0:
        for i in range(len(sList)):
            o, p, oldP = call_amp(sList[i], inputs[i], poses[i])
            poses[i] = oldP
            waitings[i] = p
            # pass output to next input
            if i == len(sList) - 1:
                inputs[0] += o
            else:
                inputs[i+1] += o
    return o

def find_max_thrusters_loopback(s, inputInit):
    m = 0
    mConfig = ()
    for a in range(5, 10):
        for b in range(5, 10):
            for c in range(5, 10):
                for d in range(5, 10):
                    for e in range(5, 10):
                        if (a != b and a != c and a != d and a != e
                            and b != c and b != d and b != e
                            and c != d and c != e
                            and d != e):
                            conf = (a,b,c,d,e)
                            sList = (copy.copy(s), copy.copy(s), copy.copy(s), copy.copy(s), copy.copy(s))
                            o = call_chaim_amp(sList, conf, inputInit)
                            if o[-1] > m:
                                m = o[-1]
                                mConfig = conf
    return (m, mConfig)


def process(loadData):
    s = loadData()
    p = 0
    print("")
    print("---")
    print("process:")
    m, mConfig = find_max_thrusters_loopback(s, [0])
    print("Max: {}".format(m))
    print("Config: {}".format(mConfig))

def main():
    process(lambda: loadStr("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"))
    process(lambda: loadStr("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"))
    process(loadDataFile)

main()

