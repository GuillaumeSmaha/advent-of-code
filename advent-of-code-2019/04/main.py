
import os 
import copy

#!/usr/bin/python

        
def check(i):
    c = False
    s = str(i)
    for a in range(1, 6):
        if s[a] == s[a-1]:
            c = True
        elif s[a-1] > s[a]:
            return False
    if not c:
        return False
    return True
    

def main():
    sols = 0
    for i in range(264360, 746326):
        if check(i):
            sols +=1
            print(i)

    print("sols = {}".format(sols))
main()

