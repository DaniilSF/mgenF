package bamgen

import (
	"encoding/csv"
	"os"

	"gitlab.com/stud777/mgen/genent"
)

type CsvHandler struct{}

func NewCsvHandler() *CsvHandler {
	return &CsvHandler{}
}

func (c *CsvHandler) Save(ud []genent.UserData, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Username", "Firstname", "Lastname", "Email", "Password", "Country", "Lang", "Cohort"})
	for _, user := range ud {
		moodleUser, err := user.User()
		if err != nil {
			return err
		}
		line := []string{
			moodleUser.Username, moodleUser.Firstname, moodleUser.Lastname,
			moodleUser.Email, moodleUser.Password, moodleUser.Country,
			moodleUser.Lang, moodleUser.Cohort,
		}
		writer.Write(line)
	}

	return nil
}
