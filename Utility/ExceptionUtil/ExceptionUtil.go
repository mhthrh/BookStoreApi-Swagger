package ExceptionUtil

type Exception struct {
	Exceptions map[int]string
}

func New() *Exception {
	return &Exception{Exceptions: map[int]string{
		1001: "[ERROR] invalid Path, path must be /books/[id]",
		1002: "[ERROR] deserializing book",
		1003: "[ERROR] validating book",
		1004: "[DEBUG] Inserting book.",
		1005: "[DEBUG] get all records",
		1006: "[ERROR] serializing book",
		1007: "[DEBUG] get record id",
		1008: "[ERROR] fetching book",
		1009: "[ERROR] serializing book",
		1010: "[DEBUG] updating record id",
		1011: "[ERROR] book not found",
		1012: "[ERROR] serializing book",
		1013: "[ERROR] serializing book",
		1014: "[ERROR] deleting record id",
		1015: "[ERROR] deleting record id does not exist",
	},
	}
}

func (e *Exception) SelectException(i int) string {
	for i2, s := range e.Exceptions {
		if i == i2 {
			return s
		}
	}
	return "Null reference exception, Welcome to GA."
}
