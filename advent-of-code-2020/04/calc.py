
def data():
    t = []
    with open("list.txt", "r") as f:
        for line in f:
            if line.strip():
                t.append(line.strip())
    return t

def passports():
    t = []
    g = {
        "cid": "0",
    }
    with open("list.txt", "r") as f:
        for line in f:
            if line == "\n":
                if g:
                    t.append(g)
                g = {
                  "cid": "0",
                }
            elif line.strip():
                p = line.strip().split(' ')
                for i in p:
                    pp = i.split(':')
                    g[pp[0]]=""
                    if len(pp) > 1:
                        g[pp[0]]=pp[1]
    
    if g:
        t.append(g)
    return t

fields = [
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
    "cid",
]

def check_pp(pp):
    for f in fields:
        if f not in pp or not pp[f].strip():
            return False
    return True

def main():
    c = 0
    for p in passports():
        # print(p)
        if check_pp(p):
            c = c + 1
    print(c)

main()

