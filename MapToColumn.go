package utils

// MapToColumn Returns a column from map
func MapToColumn(inputMap []map[string]string, keyName string) []string {
	columnEntries := make([]string, 0)
	for _, element := range inputMap {
		columnEntries = append(columnEntries, element[keyName])
	}
	return columnEntries
}
