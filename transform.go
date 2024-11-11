package main

type RawEntityTransformer interface {
	Transform(rawStrs []string) ([]Entity, error)
}

type Entity interface {
	Value() (UserData, error)
}

type UserData interface {
	User() (MoodleUserData, error)
	Csv() (string, error)
}
