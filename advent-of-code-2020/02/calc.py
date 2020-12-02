
def data():
    t = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t.append(line)
    return t

def policies():
    r = []
    for d in data():
        d = d.strip().split(' ')
        minMax = d[0].split('-')
        minMax = [int(x) for x in minMax]
        r.append([minMax, d[1][0], d[2]])
    return r

def checkPolicy(p):
    c = 0
    for i in p[2]:
        if i == p[1]: 
            c = c + 1
    return c >= p[0][0] and c <= p[0][1]

def main():
    c = 0
    for p in policies():
        print(p)
        if checkPolicy(p):
            c = c + 1
    print(c)

main()

