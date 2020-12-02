import math

def data():
    t = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t.append(line.strip())
    return t

def data_seat():
    return [seat_info(t) for t in data()]


def seat_pos(t, max, lower_letter):
    m = 0
    M = max
    for p in t:
        s = (M - m + 1) // 2
        if p == lower_letter:
            M -= s
        else:
            m += s
    return m


def seat_pos_row(t):
    return seat_pos(t[:7], 127, 'F')

def seat_pos_range(t):
    return seat_pos(t[-3:], 7, 'L')

def seat_id(t):
    return seat_pos_row(t) * 8 + seat_pos_range(t)

def seat_info(t):
    row = seat_pos_row(t)
    rang = seat_pos_range(t)
    return [row * 8 + rang, row, rang]


def main():
    d = data_seat()
    d.sort(lambda x: x[0])
    pp = d[0][0] - 1
    for p in d:
        if pp != p[0]-1:
            print("rep = ",pp+1)
        pp = p[0]

main()

