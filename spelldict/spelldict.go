package spelldict

import (
	"bonmanush/utility"
	"bytes"
	"encoding/gob"
	"os"
	"path"
)

type AkademiSpellDB struct {
	Count int
	Db    []string
}

func ReadAkademiSpellGob(filename string) (AkademiSpellDB, error) {

	fpath := filename
	result := new(AkademiSpellDB)
	if !path.IsAbs(filename) {
		cwd, err := os.Getwd()
		if err != nil {
			return *result, err
		}

		fpath = path.Join(cwd, fpath)
	}

	fdata, err := os.ReadFile(filename)

	if err != nil {
		return *result, err
	}

	buf := bytes.NewBuffer(fdata)
	gDec := gob.NewDecoder(buf)
	if err := gDec.Decode(result); err != nil {
		return *result, err
	}

	return *result, nil

}

func ReadAkademiSpellTxt(filename string) (AkademiSpellDB, error) {
	wl, err := utility.GetWordlistFromFile(filename)
	if err != nil {
		return AkademiSpellDB{}, err
	}

	return NewAkademiSpellDB(wl), nil
}

func NewAkademiSpellDB(db []string) AkademiSpellDB {

	return AkademiSpellDB{
		Count: len(db),
		Db:    db,
	}

}

func (a *AkademiSpellDB) DumbGob(filepath string) error {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	if err := enc.Encode(a); err != nil {
		return err
	}

	return os.WriteFile(filepath, buf.Bytes(), 0644)

}

func (a *AkademiSpellDB) GobEncode() ([]byte, error) {
	buf := bytes.Buffer{}

	gEnc := gob.NewEncoder(&buf)

	if err := gEnc.Encode(a.Count); err != nil {
		return nil, err
	}

	if err := gEnc.Encode(a.Db); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func (a *AkademiSpellDB) GobDecode(b []byte) error {
	buf := bytes.NewBuffer(b)
	gDec := gob.NewDecoder(buf)

	if err := gDec.Decode(&a.Count); err != nil {
		return err
	}

	if err := gDec.Decode(&a.Db); err != nil {
		return err
	}

	return nil
}
