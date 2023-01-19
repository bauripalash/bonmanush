#!/bin/env python3
from types import NoneType
from typing import List, OrderedDict

def process_item(item : str) -> List[str]:
    item = item.strip()
    x = "".join(list(map( lambda i : i.strip(), item.split("-"))))
    z : List[str] = x.split(" ")
    return list(OrderedDict.fromkeys(z))

def get_dict_as_string() -> List[str]:
    dict_txt_path : str = "../dict/dict.txt"
    raw_result : List[str] = []
    with open(dict_txt_path , "r") as f:
        raw_result = f.readlines()

    result : List[str] = []

    print(f"Words in Original Dict : {len(raw_result)}")
    for item in raw_result:
        result.extend(process_item(item))
    
    return list(OrderedDict.fromkeys(result))


def write_dict(path : str , data : List[str]) -> NoneType:
    #print(data) 
    print(f"Words in New Dict : {len(data)}")
    with open(path , "w") as f:
        f.write("\n".join(data))

#print(clean_ws(get_dict_as_string())[101])
#print(get_dict_as_string())
write_dict("samsad_dict.txt" , get_dict_as_string())

