package Helper

import (
	"encoding/json"
	"io"
	"os"
)

// Read string from file
func Read(path string) (string, error) {
	dat, err := os.ReadFile(path)
	return string(dat), err
}

// Write string to file
func Write(i interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b, _ := json.Marshal(i)
	_, err = f.WriteString(string(b))
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

// GetPath Return current directory on executable file
func GetPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, err
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
