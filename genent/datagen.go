package genent

type RawUserDataGenerator interface {
	// GetUserDataString return user data string - surname name midname (email@mail.vom, NULL) (non-generated-password, NULL)  [prep].
	// (e.g. Иванов Андрей Васильевич realemail@mail.com  some-password  prep)
	GetRawUserDataString() (string, error)
}
