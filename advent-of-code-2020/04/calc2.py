
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

def check_pp_byr(s):
    v = int(s)
    return v >= 1920 and v <= 2002

def check_pp_iyr(s):
    v = int(s)
    return v >= 2010 and v <= 2020

def check_pp_eyr(s):
    v = int(s)
    return v >= 2020 and v <= 2030

def check_pp_hgt(s):
    v = int(s[0:-2])
    if s[-2:] == 'cm':
        return v >= 150 and v <= 193
    elif s[-2:] == 'in':
        return v >= 59 and v <= 76
    return False

def check_pp_hcl(s):
    if s[0] != '#':
        return False

    for c in s[1:]:
        o = ord(c)
        if not (o >= 48 and o <= 57 or o >= 97 and o <= 102):
            return False
    return True

def check_pp_ecl(s):
    return s in [
        'amb',
        'blu',
        'brn',
        'gry',
        'grn',
        'hzl',
        'oth',
    ]

def check_pp_pid(s):
    if len(s) != 9:
        return False
    v = int(s)
    return True

def check_pp_cid(s):
    return True

def check_pp(pp):
    for f in fields:
        try:
            if f not in pp or not pp[f].strip() or not globals()['check_pp_'+f](pp[f]):
                print("failed on "+f)
                return False
        except Exception as _:
            print("failed (exp) on "+f)
            return False
    return True

def main():
    c = 0
    for p in passports():
        print(p)
        if check_pp(p):
            print("=> valid")
            c = c + 1
    print(c)

main()

