package utils

// InterfaceToStringArray converts an interface to String array
func InterfaceToStringArray(v interface{}) []string {
	stringArrayInterface := v.([]interface{})
	stringArray := make([]string, len(stringArrayInterface))
	for i, v := range stringArrayInterface {
		stringArray[i] = v.(string)
	}
	return stringArray
}
