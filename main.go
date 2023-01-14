package main

import (
	"bonmanush/bengali"
	"bonmanush/spelldict"
	"bonmanush/utility"
	"fmt"
)

func main() {

	sdb, _ := spelldict.ReadAkademiSpellGob("akademi_db.gob")
	wl := utility.GetWordlistFromFileAsRune("testcases/gulo.txt")
	sfxList := bengali.CommonSufixList
	dictCheck := true
	var result []string
	for _, item := range wl{
		for _ , sfx := range sfxList{
			iLen := len(item)
			sLen := len(sfx)
			if iLen > sLen{
				tempResult := string(item[:(iLen - sLen)])
				
				if dictCheck{
					if sdb.Contains(tempResult){
						result = append(result, tempResult)
					}
				}else{
					result = append(result, tempResult)

				}

							
			}
		}
	}

	fmt.Printf("%v\n" , result)

}
