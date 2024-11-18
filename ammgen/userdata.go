package ammgen

import (
	"strings"

	"github.com/mdigger/translit"
	"gitlab.com/stud777/mgen/genent"
)

type EntityStruct struct {
	rawStr []string
	entity []genent.Entity
}

func NewEntityStruct() *EntityStruct {
	return &EntityStruct{}
}

type UD string

type Data string

func (rd EntityStruct) Transform(rawStrs []string) ([]genent.Entity, error) {
	rd.rawStr = rawStrs
	var rawUser string
	for i := range rd.rawStr {
		rawUser = strings.ReplaceAll(strings.ReplaceAll(rd.rawStr[i], " ", ";"), ";;", ";")
		var user, _ = makeUser(rawUser)
		rd.entity = append(rd.entity, Data(user))
	}
	return rd.entity, nil
}

func (d Data) Value() (genent.UserData, error) {
	return UD(d), nil
}

func (ud UD) Csv() (string, error) {
	return string(ud), nil
}

func (ud UD) User() (genent.MoodleUserData, error) {
	var slice = strings.Split(string(ud), ";")
	var result genent.MoodleUserData
	result.Username = slice[0]
	result.Firstname = slice[1]
	result.Lastname = slice[2]
	result.Email = slice[3]
	result.Password = slice[4]
	result.Country = slice[5]
	result.Lang = slice[6]
	result.Cohort = slice[7]
	return result, nil
}

func makeUser(str string) (string, error) {
	var result string
	var slice = strings.Split(str, ";")
	var name = slice[0] + string([]rune(slice[1])[0]) + string([]rune(slice[2])[0])
	var username = translit.Ru(strings.ToLower(name))
	result = username + ";" + slice[0] + ";" + slice[1] + ";" + slice[3] + ";" + slice[4] + ";" + "ru" + ";" + "ru" + ";"
	return result, nil
}
