package Helper

import (
	"os"
	"time"
)

func Read(path string) (string, error) {
	dat, err := os.ReadFile(path)
	return string(dat), err
}

func Write(s, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}
func GetDate(format string) string {

	if format == "date" {
		return time.Now().Format("02-01-2006")
	}
	if format == "time" {
		time.Now().Format("15:04:05")
	}
	return time.Now().Format("02-01-2006 15:04:05.000")
}

func GetPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, err
}
