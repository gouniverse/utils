package utils

import "testing"

func TestTimeWithTimezone(t *testing.T) {
	testCases := []struct {
		timeString string
		timezone   string
		expected   string
	}{
		// Valid cases
		// {"12:34", "America/Los_Angeles", "05:34", false},
		{"00:00", "Europe/London", "00:00"},
		{"23:59", "Asia/Tokyo", "08:59"},

		// // Invalid time format
		{"25:60", "America/New_York", ""},
		{"abc", "Europe/Berlin", ""},

		// // Invalid timezone
		{"10:00", "InvalidTimezone", ""},

		// // Edge cases
		{"00:00", "America/Los_Angeles", "16:00"}, // Daylight saving time transition
		{"23:59", "America/Los_Angeles", "15:59"}, // Daylight saving time transition
	}

	for _, tc := range testCases {
		t.Run(tc.timeString+"-"+tc.timezone, func(t *testing.T) {
			result := TimeWithTimezone(tc.timeString, tc.timezone)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
