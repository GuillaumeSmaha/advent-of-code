
import os
import copy

def loadDataFile(file):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    data = []
    with open(os.path.join(dir_path, file), "r") as f:
        for line in f:
            if line:
                data.append([1 if x == "#" else 0 for x in line])
    return lambda: data

# def loadDataFile():
#     dir_path = os.path.dirname(os.path.realpath(__file__))
#     with open(os.path.join(dir_path, "data.txt"), "r") as f:
#         data = f.readline()
#         return [int(i) for i in str(data.strip())]
#     return []

def loadSample():
    return [int(i) for i in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]

def loadTest():
    return [int(i) for i in "1,0,0,3,99".split(",")]

def loadStr(s):
    return [int(i) for i in s]


def get_layers(s, w, t):
    layers = []
    nb = len(s) // (w * t)
    for i in range(nb):
        pos = i * (w * t)
        layers.append({
            'data': [],
            'xy': [],
        })
        layers[i]['data'] = s[pos:pos+(w * t)]
    return layers

def find_fewest_zero(layers):
    m = layers[0]['data'].count(0)
    mL = 0
    for i, l in enumerate(layers):
        c = l['data'].count(0)
        print(c)
        if m > c:
            m = c
            mL = i
    if mL == 100000000:
        mL = 0
    return mL



def manathan(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1])

prems = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31]



def process(loadData):
    s = loadData()
    print(s)

    Y = len(s)
    X = len(s[0])
    print("Y: {}".format(Y))
    print("X: {}".format(X))

    ast = []
    ast_coordx_idx = [0]*Y
    for y in range(Y):
        ast_coordx_idx[y] = [0]*X
        for x in range(X):
            if s[y][x] == 1:
                ast_coordx_idx[y][x] = len(ast)
                ast.append((y, x))
    
    ast_sight = [0] * len(ast)
    for a in ast:
        print("Ast: {}".format(a))
        ast_sight_visible = []
        ast_sight_markers = [0]*Y
        for y in range(Y):
            ast_sight_markers[y] = [0]*X
        for b in ast:
            print("\tCheck with Ast: {}".format(b))
            if a[0] == b[0] and a[1] == b[1]:
                print("\t\tPass same".format())
                continue
        
            # if ast_sight_markers[b[1]][b[0]] != 0:
            #     print("\t\tAlready".format())
            #     continue
            
            dy = b[0] - a[0]
            dx = b[1] - a[1]
            print("\t\tdy = {}, dx={}".format(dy, dx))
            dxyabs = abs(dx) + abs(dy)
            # if dxyabs > ((X+Y)//2 + 1):
            #     print("\t\t\t Pass too big abs(dx,dy)={}".format(dxyabs))
            #     continue

            dyI = 0
            dxI = 0
            while dyI != dy or dxI != dx:
                dyI = dy
                dxI = dx
                for p in reversed(prems):
                    if dx//p == dx/p and dy//p == dy/p:
                        dy = dy//p
                        dx = dx//p
            
            print("\t\tMin dy = {}, dx={}".format(dy, dx))
            cy = a[0] + dy
            cx = a[1] + dx
            marked = 0
            while cx >= 0 and cy >= 0 and cx < X and cy < Y:
                print("\t\t\tLook at ({},{})".format(cy, cx))
                if s[cy][cx] == 1:
                    print("\t\t\t\tFound an ast at ({},{})".format(cy, cx))
                    if marked == 0 and ast_sight_markers[cy][cx] == 0:
                        print("\t\t\t\t\tNearest".format(cy, cx))
                        marked = 1
                        ast_sight_visible.append((cy,cx))
                        ast_sight_markers[cy][cx] = 2
                    elif ast_sight_markers[cy][cx] == 0:
                        print("\t\t\t\t\tHidden".format(cy, cx))
                        ast_sight_markers[cy][cx] = 1
                    else:
                        print("\t\t\t\t\tAlready Hidden".format(cy, cx))

                cy = cy + dy
                cx = cx + dx
        
        idx = ast_coordx_idx[a[0]][a[1]]
        ast_sight[idx] = len(ast_sight_visible)


    print()
    print("Map:")
    a = 0
    for y in range(Y):
        for x in range(X):
            if len(ast) != a and ast[a][1] == x and ast[a][0] == y:
                print("#", end="")
                a += 1
            else:
                print(".", end="")
        print()
    print()
    print("Map NB:")
    a = 0

    aa = 1
    if max(ast_sight) > 9:
        aa = 2
    elif max(ast_sight) > 99:
        aa = 3
    for y in range(Y):
        for x in range(X):
            if len(ast) != a and ast[a][1] == x and ast[a][0] == y:
                print(str(ast_sight[a]).zfill(aa), end="")
                a += 1
            else:
                print(".".ljust(aa, "."), end="")
        print()
    print()

    print(ast_sight)
    mL = 0
    mA = 0
    for a, l in enumerate(ast_sight):
        if l > mL:
            mL = l
            mA = a

    print("Max: {}".format(mL))
    print("Pos: {}".format(ast[mA]))


def main():
    # process(loadDataFile("test1.txt"))
    # process(loadDataFile("test2.txt"))
    # process(loadDataFile("test3.txt"))
    # process(loadDataFile("test4.txt"))
    # process(loadDataFile("test5.txt"))
    process(loadDataFile("data.txt"))

main()

