package bamgen

import (
	"fmt"
	"strings"

	"gitlab.com/stud777/mgen/genent"
)

type TransformHandler struct{}

func NewTransformHandler() *TransformHandler {
	return &TransformHandler{}
}

func (t *TransformHandler) Transform(rawStrs []string) ([]genent.Entity, error) {
	var entities []genent.Entity
	for i, raw := range rawStrs {
		parts := strings.Fields(raw)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid raw data at index %d: %s", i, raw)
		}
		entity := &MoodleEntity{
			data: genent.MoodleUserData{
				Username:  cleanField(safeGet(parts, 0)),
				Firstname: cleanField(safeGet(parts, 1)),
				Lastname:  cleanField(safeGet(parts, 2)),
				Email:     cleanField(safeGet(parts, 3)),
				Password:  cleanField(safeGet(parts, 4)),
				Country:   cleanField(safeGet(parts, 5)),
				Lang:      cleanField(safeGet(parts, 6)),
				Cohort:    cleanField(safeGet(parts, 7)),
			},
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

type MoodleEntity struct {
	data genent.MoodleUserData
}

func (m *MoodleEntity) Value() (genent.UserData, error) {
	return &MoodleUserAdapter{data: m.data}, nil
}

type MoodleUserAdapter struct {
	data genent.MoodleUserData
}

func (m *MoodleUserAdapter) User() (genent.MoodleUserData, error) {
	return m.data, nil
}

func (m *MoodleUserAdapter) Csv() (string, error) {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s",
		m.data.Username, m.data.Firstname, m.data.Lastname,
		m.data.Email, m.data.Password, m.data.Country,
		m.data.Lang, m.data.Cohort), nil
}
