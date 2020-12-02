
import os 
import copy
from collections import defaultdict

#!/usr/bin/python

        
def check(i):
    c = False
    s = str(i)
    cnt=defaultdict(lambda: 0)
    for a in range(0, 6):
        cnt[s[a]] += 1
        if a > 0 and s[a-1] > s[a]:
            return False
    if 2 not in cnt.values():
        return False
    return True
    

def main():
    sols = 0
    for i in range(264360, 746325):
        if check(i):
            sols +=1
            print(i)

    print("sols = {}".format(sols))
main()

