package utility

import (
	"bufio"
	"os"
	"path"
)

func GetWordlistFromFile(filename string) ([]string, error) {

	fpath := filename
	var result []string
	if !path.IsAbs(filename) {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		fpath = path.Join(cwd, fpath)
	}

	file, err := os.Open(fpath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		result = append(result, fscanner.Text())
	}

	if err := fscanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetWordlistFromFileAsRune(p string) [][]rune{
	var result [][]rune
	rawWl , _ := GetWordlistFromFile(p)
	
	for _ , item := range rawWl{
		result = append(result, []rune(item))
	}

	return result

}
