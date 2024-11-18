package bamgen

import "gitlab.com/stud777/mgen/genent"

type CsvStorer interface {
	Save(ud []genent.UserData, filePath string) error
}
