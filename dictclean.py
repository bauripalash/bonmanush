
from types import NoneType
from typing import List, OrderedDict

def process_item(item : str) -> List[str]:
    r  : List[str] = []
    r2 : List[str] = []
    result : str = item.strip()
    
    try:
        if result.index("-") >= 1:
            r.extend(result.split("-"))
    except:
        pass

    for i in r:
        r2.append("".join(i.split(" ")))

    return list(OrderedDict.fromkeys(r2))

def get_dict_as_string() -> List[str]:
    dict_txt_path : str = "../dict/dict.txt"
    raw_result : List[str] = []
    with open(dict_txt_path , "r") as f:
        raw_result = f.readlines()

    result : List[str] = []

    print(f"Total Number of Words in Original Dict : {len(raw_result)}")
    for item in raw_result:
        result.extend(process_item(item))
    
    return list(OrderedDict.fromkeys(raw_result))


def write_dict(path : str , data : List[str]) -> NoneType:
    
    print(f"Total Number of Words in New Dict : {len(data)}")
    with open(path , "w") as f:
        f.write("\n".join(data))

        


#print(clean_ws(get_dict_as_string())[101])
#print(get_dict_as_string())
write_dict("dict3.txt" , get_dict_as_string())

