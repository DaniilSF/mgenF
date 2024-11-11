package main

import (
	"gitlab.com/stud777/stuff/datagen"
)

func main() {
	gen := datagen.RawUserDataGenerator{}

	userQty := 10

	var rawStrs []string
	for i := 0; i < userQty; i++ {
		s, err := gen.GetRawUserDataString()
		if err != nil {
			panic(err)
		}
		rawStrs = append(rawStrs, s)
	}

	var trm RawEntityTransformer // TODO add inititialisation by implementation
	ents, err := trm.Transform(rawStrs)
	if err != nil {
		panic(err)
	}

	var datas []UserData
	for _, ent := range ents {
		ud, err := ent.Value()
		if err != nil {
			panic(err)
		}
		datas = append(datas, ud)
	}

	var csvstr CsvStorer // TODO add inititialisation by implementation
	outFilePath := "./moodleuser.csv"
	err = csvstr.Save(datas, outFilePath)
	if err != nil {
		panic(err)
	}

}
