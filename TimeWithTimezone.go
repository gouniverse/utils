package utils

import "github.com/golang-module/carbon"

// TimeWithTimezone returns the given timeString formatted according to the provided timezone.
//
// Parameters:
//   - timeString: a string representing the time in "HH:MM" format.
//   - timezone: a string representing the timezone to apply to the timeString.
//
// Returns:
//   - a string representing the formatted time in the provided timezone.
func TimeWithTimezone(timeString string, timezone string) string {
	datetime := "2000-01-01 " + timeString // the year is irrelevant, as we only return the time
	parsed := carbon.Parse(datetime, carbon.UTC)
	return parsed.SetTimezone(timezone).Format("H:i")
}
