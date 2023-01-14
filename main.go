package main

import (
	"bonmanush/spelldict"
	"fmt"
)

func main() {

	sdb, _ := spelldict.ReadAkademiSpellGob("akademi_db.gob")
	fmt.Println(sdb.Count)
	fmt.Println(sdb.Contains(sdb.Db[100]))

}
