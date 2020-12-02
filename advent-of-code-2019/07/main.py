
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

Input = [1]
InputIdx = 0

def get_val(s, p, immediate):
    if immediate:
        return s[p]
    return s[s[p]]

def exec_op(s, op, p, immediate=[False, False]):
    output = None
    global Input, InputIdx
    if op == 1:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) + get_val(s, p+2, immediate[1])
        p += 4
    elif op == 2:
        s[s[p+3]] = get_val(s, p+1, immediate[0]) * get_val(s, p+2, immediate[1])
        p += 4
    elif op == 3:
        s[s[p+1]] = Input[InputIdx]
        InputIdx += 1
        p += 2
    elif op == 4:
        # print("call output")
        # print(output)
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
    elif op > 99:
        val = str(op).zfill(5)
        code_op = int(val[-2:])
        p, output = exec_op(s, code_op, p, immediate=[val[-3] == "1", val[-4] == "1"])
    return (p, output)


def call_amp(s, intputAmp):
    global Input, InputIdx
    p = 0
    Input = intputAmp
    InputIdx = 0
    output = []
    while p is not False:
        p, o = exec_op(s, s[p], p)
        if o is not None:
            output.append(o)
    return output

def call_chaim_amp(sList, conf, inputInit):
    a,b,c,d,e = conf
    o = inputInit
    for i in range(len(sList)):
        o = call_amp(sList[i], [conf[i]] + o)
    return o

def find_max_thrusters(s, inputInit):
    m = 0
    s = copy.copy(s)
    mConfig = ()
    for a in range(0,5):
        for b in range(0,5):
            for c in range(0,5):
                for d in range(0,5):
                    for e in range(0,5):
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
    global Input, InputIdx
    InputIdx = 0
    s = loadData()
    p = 0
    print("")
    print("---")
    print("process:")
    m, mConfig = find_max_thrusters(s, [0])
    print("Max: {}".format(m))
    print("Config: {}".format(mConfig))

def main():
    process(lambda: loadStr("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"))
    process(lambda: loadStr("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"))
    process(lambda: loadStr("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"))
    process(loadDataFile)

main()

