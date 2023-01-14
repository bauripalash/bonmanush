package spelldict


func containsString(arr []string, target string) bool {
	for _ , item := range arr{
		if item == target{
			return true
		}
	}

	return false
}

func (a *AkademiSpellDB) Contains(target string) bool {
	return containsString(a.Db, target)
}
