import math
import copy
import time
from pprint import pprint
from collections import defaultdict

CASE_FLOOR = "."
CASE_SEAT = "L"
CASE_SEAT_FULL = "#"

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
    if d[y][x] == CASE_FLOOR:
        return d[y][x], False

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

    # print("x, y", x, y, " => X", xmin, xmax, " => Y", ymin, ymax, " => ", cnt_full)

    if d[y][x] == CASE_SEAT and cnt_full == 0:
        return CASE_SEAT_FULL, True
    elif d[y][x] == CASE_SEAT_FULL and cnt_full >= 4:
        return CASE_SEAT, True
    return d[y][x], False


def print_data(d):
    for l in d:
        print(''.join([str(x) for x in l]))

def step_next_cnt_seat(d):
    r = copy.deepcopy(d)
    cnt_full = 0
    for y,_ in enumerate(d):
        for x,_ in enumerate(dd[0]):
            r[y][x] = step_cnt_seat_full_yx(d, y, x)
    return r

def step_next_seat_yx(d, d_seat, y, x):
    if d[y][x] == CASE_FLOOR:
        return d[y][x], False

    s = int(d_seat[y][x])
    if d[y][x] == CASE_SEAT and s == 0:
        return CASE_SEAT_FULL, True
    elif d[y][x] == CASE_SEAT_FULL and s >= 4:
        return CASE_SEAT, True
    return d[y][x], False

def step_next2(d):
    cnt_full = 0
    d_seat = step_next_cnt_seat(d)
    print_data(d_seat)
    print("")
    for y,_ in enumerate(d):
        for x,_ in enumerate(dd[0]):
            v, updated = step_next_seat_yx(d, d_seat, y, x)
            d[y][x] = v
            if v == CASE_SEAT_FULL:
                cnt_full += 1

    return d, cnt_full

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

