package test

import (
	"testing"

	"gitlab.com/stud777/mgen/ammgen"
	"gitlab.com/stud777/mgen/bamgen"
	"gitlab.com/stud777/mgen/genent"

	"github.com/stretchr/testify/require"
)

// TODO test names
// TODO test locations

// TODO  github.com/stretchr/testify
// TODO  gen mock

type MockDatageGenEmpty struct{}

func (mdg *MockDatageGenEmpty) GetRawUserDataString() (string, error) {
	return "Долгих Виктор Виталиевич water@broken.wildflower.com  NULL ", nil
}

// trm
func TestAmmgenTransformValid(t *testing.T) {
	// var rawStrs []string
	// rawStrs = append(rawStrs, "Долгих Виктор Виталиевич water@broken.wildflower.com  NULL ")

	//

	mdg := &MockDatageGenEmpty{}

	s, err := mdg.GetRawUserDataString()
	require.NoError(t, err)
	require.NotEmpty(t, s)

	rawStrs := []string{s}

	trm := ammgen.NewEntityStruct()
	ents, err := trm.Transform(rawStrs)

	require.NotEmpty(t, ents)
	require.NoError(t, err)
	// require.Eventually()
}

func TestTrmAmmgenNoValid(t *testing.T) {
	var rawStrs []string
	rawStrs = append(rawStrs, "")
	trm := ammgen.NewEntityStruct()
	ents, err := trm.Transform(rawStrs)
	if err != nil {
		t.Log("Expected error occurred:", err)
		return
	}
	t.Errorf("Expected an error for invalid input, but got result: %v", ents)
}

//csv

func TestCsvAmmgenValid(t *testing.T) {
	input := ammgen.UD("username;firstname;lastname;email")
	expected := "username;firstname;lastname;email"
	output, err := input.Csv()

	if err != nil {
		t.Errorf("Csv() returned an error: %v", err)
	}
	if output != expected {
		t.Errorf("Csv() = %v, expected %v", output, expected)
	}
}
func TestCsvAmmgenNoValid(t *testing.T) {
	input := ammgen.UD("")
	expected := ""
	expectError := false
	output, err := input.Csv()

	if (err != nil) != expectError {
		t.Errorf("Csv() error = %v, expected error = %v", err, expectError)
	}
	if output != expected {
		t.Errorf("Csv() = %v, expected %v", output, expected)
	}
}

type MockTransWithEpmtydata struct{}

func (mtr *MockTransWithEpmtydata) Transform(rawStrs []string) ([]genent.Entity, error) {
	return []genent.Entity{}, nil
}

// save
func TestSaveAmmgenNoValid(t *testing.T) {
	trm := MockTransWithEpmtydata{}
	ents, err := trm.Transform(nil)
	require.NoError(t, err)

	var datas []genent.UserData
	for _, ent := range ents {
		ud, err := ent.Value()
		if err != nil {
			panic(err)
		}
		datas = append(datas, ud)
	}

	filePath := "test_output_invalid.csv"

	storer := ammgen.NewCsv()
	err = storer.Save(datas, filePath)
	require.NoErrorf(t, err, "save failed")

}

// ///////
// //////
// trm
func TestTrmBamgenValid(t *testing.T) {
	var rawStrs []string
	rawStrs = append(rawStrs, "Долгих Виктор Виталиевич water@broken.wildflower.com  NULL ")

	trm := bamgen.NewTransformHandler()
	ents, err := trm.Transform(rawStrs)
	if err != nil {
		panic(err)
	}
	gg := ents[0]
	print(gg)
	t.Log("Tested result", ents[0])
}
func TestTrmBamgenNoValid(t *testing.T) {
	var rawStrs []string
	rawStrs = append(rawStrs, "")
	trm := bamgen.NewTransformHandler()
	ents, err := trm.Transform(rawStrs)
	if err != nil {
		t.Log("Expected error occurred:", err)
		return
	}
	t.Errorf("Expected an error for invalid input, but got result: %v", ents)
}

//save

// func TestSaveValid(t *testing.T) {
//     users := []genent.UserData{
//         genent.MoodleUserData{
//             Username:  "user1",
//             Firstname: "John",
//             Lastname:  "Doe",
//             Email:     "user1@example.com",
//         },
//         genent.MoodleUserData{
//             Username:  "user2",
//             Firstname: "Jane",
//             Lastname:  "Smith",
//             Email:     "user2@example.com",
//         },
//     }

//     csvHandler := bamgen.NewCsvHandler()
//     outputFile := "test_output_valid.csv"

//     err := csvHandler.Save(users, outputFile)
//     if err != nil {
//         t.Fatalf("Expected no error, got: %v", err)
//     }

//     t.Logf("Save valid test passed. Output file: %s", outputFile)
// }

func TestSaveBamgenNoValid(t *testing.T) {
	var users []genent.UserData // Пустой массив
	csvHandler := bamgen.NewCsvHandler()

	filePath := "test_output_invalid_bamgen.csv"
	err := csvHandler.Save(users, filePath)
	if err != nil {
		t.Fatal("Expected error for empty input, got none")
	}

	t.Log("CsvHandler Save invalid test passed")
}

// ALL

// func TestAllBamgenValid(t *testing.T) {
// 	var rawStrs []string
// 	rawStrs = append(rawStrs, "Долгих Виктор Виталиевич water@broken.wildflower.com  NULL ")
// 	var transformer bamgen.RawEntityTransformer =bamgen.NewTransformHandler()
// 	ents, err := transformer.Transform(rawStrs)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var users []genent.MoodleUserData
// 	for _, entity := range ents {
// 		user, err := entity.Value()
// 		if err != nil {
// 			panic(err)
// 		}
// 		users = append(users, user)
// 	}

// 	var csvStorer bamgen.CsvStorer = bamgen.NewCsvHandler()
// 	outputFile := "./moodleuser.csv"
// 	if err := csvStorer.Save(users, outputFile); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Данные успешно сохранены в файл: %s\n", outputFile)
// }

////////////////////////////
///////////////////////////

// func TestAllValid(t *testing.T) {
// 	var rawStrs []string
// 	rawStrs = append(rawStrs, "Долгих Виктор Виталиевич water@broken.wildflower.com NULL")

// 	trm := ammgen.NewEntityStruct()
// 	ents, err := trm.Transform(rawStrs)
// 	if err != nil {
// 		t.Fatalf("Transform() returned an unexpected error: %v", err)
// 	}

// 	if len(ents) == 0 {
// 		t.Fatalf("Transform() returned no entities for valid input")
// 	}

// 	t.Logf("Transform() succeeded, first entity: %+v", ents[0])

// 	var datas []genent.UserData
// 	for _, ent := range ents {
// 		ud, err := ent.Value()
// 		if err != nil {
// 			t.Fatalf("Entity.Value() returned an unexpected error: %v", err)
// 		}
// 		datas = append(datas, ud)
// 	}

// 	if len(datas) != len(ents) {
// 		t.Errorf("Mismatch between number of entities and UserData objects: %d != %d", len(ents), len(datas))
// 	}

// 	t.Logf("Value() succeeded, UserData collected: %+v", datas[0])

// 	csvstr := ammgen.NewCsv()
// 	outFilePath := "./moodleuser.csv"

// 	if _, err := os.Stat(outFilePath); err == nil {
// 		os.Remove(outFilePath)
// 	}

// 	err = csvstr.Save(datas, outFilePath)
// 	if err != nil {
// 		t.Fatalf("Save() returned an unexpected error: %v", err)
// 	}

// 	// Проверка, что файл был создан
// 	if _, err := os.Stat(outFilePath); os.IsNotExist(err) {
// 		t.Fatalf("Save() did not create the output file: %v", outFilePath)
// 	}

// 	t.Logf("Save() succeeded, file created: %s", outFilePath)

// 	defer os.Remove(outFilePath)
// }

// func TestAllInvalid(t *testing.T) {

// 	var rawStrs []string
// 	rawStrs = append(rawStrs, "")

// 	trm := ammgen.NewEntityStruct()
// 	ents, err := trm.Transform(rawStrs)

// 	if err == nil {
// 		t.Fatalf("Transform() should have returned an error for invalid input")
// 	}
// 	t.Logf("Transform() correctly returned an error: %v", err)

// 	if len(ents) != 0 {
// 		t.Fatalf("Transform() returned entities for invalid input: %+v", ents)
// 	}

// 	var datas []genent.UserData
// 	for _, ent := range ents {
// 		ud, err := ent.Value()
// 		if err == nil {
// 			t.Fatalf("Entity.Value() should have returned an error for invalid entity: %+v", ent)
// 		}
// 		t.Logf("Entity.Value() correctly returned an error: %v", err)

// 		datas = append(datas, ud)
// 	}

// 	if len(datas) != 0 {
// 		t.Errorf("Expected no UserData objects, but got: %d", len(datas))
// 	}

// 	csvstr := ammgen.NewCsv()
// 	outFilePath := "./moodleuser_invalid.csv"

// 	if _, err := os.Stat(outFilePath); err == nil {
// 		os.Remove(outFilePath)
// 	}

// 	err = csvstr.Save(datas, outFilePath)
// 	if err == nil {
// 		t.Fatalf("Save() should have returned an error for invalid data")
// 	}
// 	t.Logf("Save() correctly returned an error for invalid data: %v", err)

// 	if _, err := os.Stat(outFilePath); !os.IsNotExist(err) {
// 		t.Fatalf("Save() created a file for invalid data: %s", outFilePath)
// 	}

// 	defer os.Remove(outFilePath)
// }
