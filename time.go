package utils

import(
	"time"
)

// StrToTimeUnix converts sting to Unix time
func StrToTimeUnix(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
