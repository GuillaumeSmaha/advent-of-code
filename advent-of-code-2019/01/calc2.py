
import os 


def data():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    t = 0
    with open(os.path.join(dir_path, "list.txt"), "r") as f:
        for line in f:
            if line.strip():
                t += calcR(int(line, 10))
    print("total: {}".format(t))

def calcR(mass):
    c = calc(mass)
    t = c
    while c >= 0:
        c = calc(c)
        if c >= 0:
            t += c
    return t

def calc(mass):
    return mass // 3 - 2

def main():
    print(calc(12))
    print(calc(14))
    print(calc(1969))
    print(calc(100756))
    data()

main()

