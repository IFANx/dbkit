package util

func CheckStrExists(item string, arr []string) bool {
	for _, elem := range arr {
		if item == elem {
			return true
		}
	}
	return false
}
