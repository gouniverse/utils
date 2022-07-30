package utils

// MapToKeyValue returns a key-value array an array of maps
func MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string {
	keyValueMap := map[string]string{}
	for _, element := range inputMap {
		keyValueMap[element[keyName]] = element[valueName]
	}
	return keyValueMap
}
