package ammgen

import (
	"fmt"
	"os"

	"gitlab.com/stud777/mgen/genent"
)

type CsvStorer interface {
	Save(ud []genent.UserData, filePath string) error
}

type Csv struct{}

func NewCsv() *Csv {
	return &Csv{}
}

func (Csv Csv) Save(uds []genent.UserData, filePath string) error {
	/*err := os.Chdir(filePath)
	if err != nil {
		panic(err)
	}*/
	file, err := os.Create("./moodleuser.csv")
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
	defer file.Close()

	file.WriteString("username;firstname;lastname;email;password;country;lang;\n")
	for i := range uds {
		var str string
		str, _ = uds[i].Csv()
		str = str + "\n"
		file.WriteString(str)
	}

	return nil
}
