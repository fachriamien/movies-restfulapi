package helper

import (
	"database/sql"
	"time"
)

func ParserNull(value sql.NullString) string {
	var result string
	if value.Valid {
		result = value.String
	} else {
		result = "" // Or set it to any default value you prefer
	}
	return result
}

func ParserTime(value string) string {
	result := ""

	if value != "" {
		// Parse the input time string into a time.Time object
		parsedTime, err := time.Parse(time.RFC3339, value)
		PanicIfError(err)
		// Format the time as desired (without T and Z)
		result = parsedTime.Format("2006-01-02 15:04:05")
	}

	return result
}
