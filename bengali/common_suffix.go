package bengali

var CommonSufixList = [][]rune{
	{BN_CONSONANTS["ga"], BN_VOWEL_SIGNS["u"], BN_CONSONANTS["la"], BN_VOWEL_SIGNS["o"]},
}

func GetCommonSuffixList() []string{
	var wl []string

	for _ , item := range CommonSufixList{
		wl = append(wl, string(item))
	}

	return wl
}
