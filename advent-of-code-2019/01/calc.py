


def data():
    t = 0
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t += calc(int(line, 10))
    print("total: {}".format(t))

def calc(mass):
    return mass // 3 - 2

def main():
    print(calc(12))
    print(calc(14))
    print(calc(1969))
    print(calc(100756))
    data()

main()

