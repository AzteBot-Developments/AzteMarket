package utils

func StringInSlice(src string, slice []string) bool {
	for _, str := range slice {
		if str == src {
			return true
		}
	}
	return false
}
