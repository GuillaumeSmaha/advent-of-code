
import os 
import copy

#!/usr/bin/python

class Point:
	def __init__(self,p):
		self.x = p[0]
		self.y = p[1]

	def __str__(self):
		print(self.x, self.y)

	def __repr__(self):
		return "{}, {}".format(self.x, self.y)

def ccw(A,B,C):
	return (C.y-A.y)*(B.x-A.x) > (B.y-A.y)*(C.x-A.x)

def intersect(A,B,C,D):
	return ccw(A,C,D) != ccw(B,C,D) and ccw(A,B,C) != ccw(A,B,D)



def loadDataFile():
    dir_path = os.path.dirname(os.path.realpath(__file__))
    with open(os.path.join(dir_path, "data.txt"), "r") as f:
        d1 = f.readline().split(",")
        d2 = f.readline().split(",")
        return [d1, d2]
    return []

def loadStr(s):
    d1, d2 = s.split(";")
    return [d1.split(","), d2.split(",")]


def op(pos, o):
    v = int(o[1:])
    if o[0] == "U":
        pos[1] += v
    elif o[0] == "D":
        pos[1] -= v
    elif o[0] == "R":
        pos[0] += v
    elif o[0] == "L":
        pos[0] -= v
    return pos, abs(v)

def manathan(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1])

def get_edges_poses_v2(d):
    poses = []
    edges = []
    steps = []
    pos = [0, 0]
    for o in d:
        e = [Point(pos)]
        pos, s = op(pos, o)
        p = pos[:]
        poses.append(p)
        e.append(Point(p))
        edges.append(e)
        steps.append(s)

    return edges, poses, steps

def get_top_right_edge_poses(e):
    if e[1][1] > e[0][1] or e[1][0] > e[0][0]:
        return [e[1], e[0]]

def collided_v2(e1, e2):
    return intersect(e1[0], e1[1], e2[0], e2[1])


def process(loadData):
    print("---")
    s = loadData()

    edges1, poses1, steps1 = get_edges_poses_v2(s[0])
    edges2, poses2, steps2 = get_edges_poses_v2(s[1])
    from pprint import pprint
    # pprint(edges1)
    # pprint(edges2)

    # print(steps1)
    # print(steps2)
    stx = 0
    collisions = []
    for s in range(len(edges1)):
        st1 = 0
        for s1, e1 in enumerate(edges1[:s]):
            st2 = 0
            for s2, e2 in enumerate(edges2[:s]):
                if collided_v2(e1, e2):
                    Cx = e1[0].x
                    if e1[0].x != e1[1].x:
                        Cx = e2[0].x
                    Cy = e1[0].y
                    if e1[0].y != e1[1].y:
                        Cy = e2[0].y
                    print("pos")
                    print(e1)
                    print(e2)
                    print((Cx, Cy))
                    
                    t = st1+st2
                    t += manathan([Cx, Cy], [e1[0].x, e1[0].y])
                    t += manathan([Cx, Cy], [e2[0].x, e2[0].y])
                    print(t)
                    return
                st2 += steps2[s2]
            st1 += steps1[s1]
    
                


        
def main():
    process(lambda: loadStr("R8,U5,L5,D3;U7,R6,D4,L4"))
    process(lambda: loadStr("R75,D30,R83,U83,L12,D49,R71,U7,L72;U62,R66,U55,R34,D71,R55,D58,R83"))
    process(lambda: loadStr("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51;U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))
    process(loadDataFile)

main()

