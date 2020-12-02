import copy
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

    def fix_nop_nor(self, i):
        print("Fix %d:" % i)
        if self._data[i][0] == "nop" and self._data[i][1] != 0:
            print("\t nop %d => jmp %d" % (self._data[i][1], self._data[i][1]))
            self._data[i][0] = "jmp"
            return True
        elif self._data[i][0] == "jmp":
            print("\t jmp %d => nop %d" % (self._data[i][1], self._data[i][1]))
            self._data[i][0] = "nop"
            return True
        return False

    def execute(self):
        ik, iv = self.instruction()
        if self._data_exec[self._ptr] > 0:
            return False
        self._data_exec[self._ptr] += 1
        if self._debug:
            print("\t", self._ptr, ik, iv, ":", self._data_exec[self._ptr], self._accumulator)
        if ik == "nop":
            pass
        elif ik == "acc":
            self._accumulator += iv
        elif ik == "jmp":
            self._ptr += iv
            return True
        else:
            print("unknow op: %s" % ik)
        
        self._ptr += 1
        return True

def main():
    d = data()
    for i, dd in enumerate(d):
        p = Program(copy.deepcopy(d))
        if p.fix_nop_nor(i):
            while not p.isend():
                r = p.execute()
                if not r:
                    break
            if r:
                print(p._accumulator)
                break

main()

