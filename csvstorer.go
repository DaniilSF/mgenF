package main

type CsvStorer interface {
	Save(ud []UserData, filePath string) error
}
