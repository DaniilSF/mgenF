package main

import (
	"gitlab.com/stud777/mgen/ammgen"
	"gitlab.com/stud777/mgen/genent"
	"gitlab.com/stud777/stuff/datagen"
)

func main() {
	var gen genent.RawUserDataGenerator
	gen = datagen.RawUserDataGenerator{}
	// gen := datagen.RawUserDataGenerator{}

	userQty := 10

	var rawStrs []string
	for i := 0; i < userQty; i++ {
		s, err := gen.GetRawUserDataString()
		if err != nil {
			panic(err)
		}
		rawStrs = append(rawStrs, s)
	}

	// === Init Transformer
	// trm := bamgen.NewTransformHandler()
	trm := ammgen.NewEntityStruct()

	ents, err := trm.Transform(rawStrs)
	if err != nil {
		panic(err)
	}

	var datas []genent.UserData
	for _, ent := range ents {
		ud, err := ent.Value()
		if err != nil {
			panic(err)
		}
		datas = append(datas, ud)
	}

	// === Init CsvStorer
	// csvstr := bamgen.NewCsvHandler()
	csvstr := ammgen.NewCsv()

	outFilePath := "./moodleuser.csv"
	err = csvstr.Save(datas, outFilePath)
	if err != nil {
		panic(err)
	}

}
