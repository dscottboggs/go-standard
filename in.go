package standard

// InSlice returns true if the given value is in the given slice
func StringInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}
