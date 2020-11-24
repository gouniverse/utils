package utils

// MapToColumn Returns a column from map
func MapToColumn(inputMap []map[string]string, keyName string) []string {
	columnEntries := make([]string, 0)
	for _, element := range inputMap {
		columnEntries = append(columnEntries, element[keyName])
	}
	return columnEntries
}

// MapToKeyValue Returns a column from map
func MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string {
	keyValueMap := map[string]string{}
	for _, element := range inputMap {
		keyValueMap[element[keyName]] = element[valueName]
	}
	return keyValueMap
}
