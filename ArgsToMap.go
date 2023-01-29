package utils

import "strings"

// ArgsToMap converts command line arguments to a key value map
// supports filled (i.e. --user=12) and unfilled (i.e. --force) arguments
func ArgsToMap(args []string) map[string]string {
	kv := map[string]string{}
	for i := 0; i < len(args); i++ {
		current := args[i]

		if strings.HasPrefix(current, "--") {
			if strings.Contains(current, "=") {
				currentArray := strings.Split(current, "=")
				kv[currentArray[0][2:]] = currentArray[1]
			} else {
				next := ""
				if len(args) > i+1 {
					next = args[i+1]
				}

				if strings.HasPrefix(next, "--") {
					kv[current[2:]] = ""
					continue
				}
				kv[current[2:]] = next
			}
		}
	}
	return kv
}
