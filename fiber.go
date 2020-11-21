package utils

import (
	"github.com/gofiber/fiber/v2"
)

// FiberGetValueArray2D return a 2D array of a post body
// Defaults to nil if the form key doesn't exist.
// If a default value is given, it will return that value if the form key does not exist.
// Returned value is only valid within the handler. Do not store any references.
// Make copies or use the Immutable setting instead.
func FiberGetValueArray2D(c *fiber.Ctx, key string, defaultValue ...[]string) [][]string {
	// log.Println(postBody)
	postBody := string(c.Body())
	m, _ := url.ParseQuery(postBody)
	rows := [][]string{}

	for i := 0; i < len(m); i++ {
		index := key + "[" + strconv.Itoa(i) + "][]"
		// log.Println(index)
		keyValue, keyExists := m[index]
		if keyExists == false {
			if i == 0 {
				return defaultValue
			}
			continue
		}
		// log.Println(row)
		rows = append(rows, keyValue)
	}

	return rows
}
