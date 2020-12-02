
def data():
    t = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t.append(line.strip())
    return t

def checkPolicy(p):
    c = 0
    for i in p[2]:
        if i == p[1]: 
            c = c + 1
    return c >= p[0][0] and c <= p[0][1]

def get_cnt(right, down):
    m = data()
    posX = 0
    posY = 0
    paths = [m[0][0]]
    while posY < len(m):
        posX = (posX + right) % len(m[posY])
        posY += down
        if posY < len(m):
            paths.append(m[posY][posX])
    cnt = 0
    for p in paths:
        if p == '#':
            cnt += 1
    return cnt

def main():
    print(get_cnt(3, 1))

main()

