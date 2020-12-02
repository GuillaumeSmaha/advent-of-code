


def data():
    t = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t.append(int(line))
    return t

def main():
    t = data()
    for ii, i in enumerate(t):
        for ij, j in enumerate(t):
            for ik, k in enumerate(t):
                if ii != ij and ii != ik:
                    if (i + j +k) == 2020:
                        print(i*j*k)
                        return

main()

