import math
from pprint import pprint
from collections import defaultdict


def data():
    lines = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                lines.append(line.strip())
    
    r = []
    pprint(lines)
    for l in lines:
        v = [x for x in l.split(" ")]
        v[1] = int(v[1])
        r.append(v)

    return r

class Program:
    def __init__(self, data):
        self._ptr = 0
        self._debug = True
        self._data = data
        self._data_exec = [0] * len(data)
        self._accumulator = 0

    def isend(self):
        return self._ptr >= len(self._data)

    def instruction(self):
        return self._data[self._ptr]

    def execute(self):
        ik, iv = self.instruction()
        if self._data_exec[self._ptr] > 0:
            self._ptr = len(self._data)
            return
        self._data_exec[self._ptr] += 1
        if self._debug:
            print(ik, iv, self._data_exec[self._ptr], self._accumulator)
        if ik == "nop":
            pass
        elif ik == "acc":
            self._accumulator += iv
        elif ik == "jmp":
            self._ptr += iv
            return
        else:
            print("unknow op: %s" % ik)
        
        self._ptr += 1

def main():
    p = Program(data())

    while not p.isend():
        p.execute()

main()

