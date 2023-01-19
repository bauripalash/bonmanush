#!/bin/env python3

from typing import List, OrderedDict


s_dict_path = 'dict_samsad.txt'
a_dict_path = 'akademi_spellings.txt'
merged_dict_path = "../d.txt"

def read_dict(p : str) -> List[str]:
    result : List[str] = []
    with open(p , "r") as f:
       result = f.readlines() 

    return list(map(lambda x : x.strip() , result))

    
def get_merge_db() -> List[str]:
    mdb = read_dict(s_dict_path) + read_dict(a_dict_path)
    x = list(OrderedDict.fromkeys(mdb))
    #x.sort() 
    return x

def write_db(p : str):
    with open(p , "w") as f:
        f.write("\n".join(get_merge_db()))
        
        

#get_merge_db()
write_db("dict_m.txt")
