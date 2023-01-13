package bengali

const (
	BN_URANGE_START       = rune('ঀ') //U+0980
	BN_URANGE_VISIBLE_END = rune('৽') //U+09FD
	BN_URANGE_END         = rune('\u09ff')
)

var (
	BN_URANGE_RESEVERD = []rune{
		'\u0984', '\u098D', '\u098E', '\u0991',
		'\u0992', '\u0992', '\u09A9',
		'\u09B1', '\u09B1', '\u09B3',
		'\u09B4', '\u09B5', '\u09BA',
		'\u09BB', '\u09C5', '\u09C6',
		'\u09C9', '\u09CA', '\u09CF',
		'\u09D0', '\u09D1', '\u09D2',
		'\u09D3', '\u09D4', '\u09D5',
		'\u09D6', '\u09D8', '\u09D9',
		'\u09DA', '\u09DB', '\u09DE',
		'\u09E4', '\u09E5', '\u09FF',
	}

	BN_VOWELS = map[string]rune{
		"a":  '\u0985',
		"aa": '\u0986',
		"i":  '\u0987',
		"ii": '\u0988',
		"u":  '\u0989',
		"uu": '\u098A',
		"vr": '\u098B',
		"vl": '\u098C',
		"e":  '\u098F',
		"ai": '\u0990',
		"o":  '\u0993',
		"au": '\u0994',
	}

	BN_CONSONANTS = map[string]rune{
		"ka":   '\u0995',
		"kha":  '\u0996',
		"ga":   '\u0997',
		"gha":  '\u0998',
		"nga":  '\u0999',
		"ca":   '\u099A',
		"cha":  '\u099B',
		"ja":   '\u099C',
		"jha":  '\u099D',
		"nya":  '\u099E',
		"tta":  '\u099F',
		"ttha": '\u09A0',
		"dda":  '\u09A1',
		"ddha": '\u09A2',
		"nna":  '\u09A3',
		"ta":   '\u09A4',
		"tha":  '\u09A5',
		"da":   '\u09A6',
		"dha":  '\u09A7',
		"na":   '\u09A8',
		"pa":   '\u09AA',
		"pha":  '\u09AB',
		"ba":   '\u09AC',
		"bha":  '\u09AD',
		"ma":   '\u09AE',
		"ya":   '\u09AF',
		"ra":   '\u09B0',
		"la":   '\u09B2',
		"sha":  '\u09B6',
		"ssa":  '\u09B8',
		"sa":   '\u09B8',
		"ha":   '\u09B9',
		"rra":  '\u09DC',
		"rha":  '\u09DD',
		"yya":  '\u09DF',

		// Extra Signs
	}

	BN_SIGNS = map[string]rune{
		"chandrabindu": '\u0981',
		"anusvara":     '\u0982',
		"visarga":      '\u0983',
	}

	BN_VOWEL_SIGNS = map[string]rune{
		"aa":  '\u09BE',
		"i":   '\u09BF',
		"ii":  '\u09C0',
		"u":   '\u09C1',
		"uu":  '\u09C2',
		"vr":  '\u09C3',
		"vrr": '\u09C4',
		"e":   '\u09C7',
		"ai":  '\u09C8',
		"o":   '\u09CB',
		"au":  '\u09CC',
	}

	BN_NUMBERS = map[string]rune{
		"0": '\u09E6',
		"1": '\u09E7',
		"2": '\u09E8',
		"3": '\u09E9',
		"4": '\u09EA',
		"5": '\u09EB',
		"6": '\u09EC',
		"7": '\u09ED',
		"8": '\u09EE',
		"9": '\u09EF',
	}
)

func doesContain(arr []rune, c rune) bool {
	for _, i := range arr {
		if i == c {
			return true
		}
	}

	return false
}

func IsUnresBnChar(c rune) bool {
	if IsBnChar(c) {
		if !doesContain(BN_URANGE_RESEVERD, c) {
			return true
		}
	}
	return false
}

func IsBnChar(c rune) bool {
	if c >= BN_URANGE_START && c <= BN_URANGE_END {
		return true
	}

	return false
}

func IsBnNum(c rune) bool {
	return c >= BN_NUMBERS["0"] && c <= BN_NUMBERS["9"]
}

func IsBnVowel(c rune) bool {
	if IsUnresBnChar(c) {
		if c >= BN_VOWELS["a"] && c <= BN_VOWELS["au"] {
			return true
		}
	}

	return false
}

func IsBnCons(c rune) bool {
	if IsUnresBnChar(c) {
		if c >= BN_CONSONANTS["k"] && c <= BN_CONSONANTS["yya"] {
			return true
		} else if c >= BN_SIGNS["chandrabindu"] && c <= BN_SIGNS["visarga"] {
			return true
		}
	}

	return false
}
