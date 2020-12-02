import math
import copy
import time
from pprint import pprint
from collections import defaultdict

CASE_FLOOR = "."
CASE_SEAT = "L"
CASE_SEAT_FULL = "#"

directions = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]]

def data():
    lines = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                lines.append(list(line.strip()))
    
    return lines

def step_cnt_seat_full_yx(d, y, x):
    if d[y][x] == CASE_FLOOR:
        return '-'

    ymin = max(0, y - 1)
    ymax = min(len(d) - 1, y + 1)
    xmin = max(0, x - 1)
    xmax = min(len(d[0]) - 1, x + 1)

    cnt_full = 0
    for yi in range(ymin, ymax + 1):
        for xi in range(xmin, xmax + 1):
            if xi != x or yi != y:
                if d[yi][xi] == CASE_SEAT_FULL:
                    cnt_full += 1
    return cnt_full

def step_next_yx(d, y, x):
    cnt_full = 0
    for dd in directions:
        i = 1
        while True:
            di = [i * x for x in dd]
            yi = y + di[0]
            xi = x + di[1]
            if yi < 0 or yi >= len(d) or xi < 0 or xi >= len(d[0]):
                break

            if d[yi][xi] == CASE_SEAT_FULL:
                # print("\t full", yi, xi)
                cnt_full += 1
                break
            elif d[yi][xi] == CASE_SEAT:
                # print("\t empty", yi, xi)
                break
            else:
                # print("\t none", yi, xi)
                pass
            i += 1

    # print("y, x", y, x, " => ", cnt_full)

    if d[y][x] == CASE_SEAT and cnt_full == 0:
        return CASE_SEAT_FULL, True
    elif d[y][x] == CASE_SEAT_FULL and cnt_full >= 5:
        return CASE_SEAT, True
    return d[y][x], False


def print_data(d):
    for l in d:
        print(''.join([str(x) for x in l]))

def step_next(d):
    cnt_full = 0
    d_init = copy.deepcopy(d)
    is_updated = False
    for y,_ in enumerate(d):
        for x,_ in enumerate(d[0]):
            v, updated = step_next_yx(d_init, y, x)
            is_updated = is_updated or updated
            d[y][x] = v
            if v == CASE_SEAT_FULL:
                cnt_full += 1

    return d, cnt_full, is_updated

def main():
    d = data()
    print("size X Y", len(d[0]), len(d))

    print("\nstep", 0)
    print_data(d)
    p = 0
    is_updated = True
    i = 0
    while is_updated:
        i += 1
        print("\nstep", i)
        d, cnt_full, is_updated = step_next(d)
        print_data(d)
        print("seat", cnt_full)
        p = cnt_full

main()

