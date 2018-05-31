package hashmap

// hashCode calculates hashcode of string value
// TODO: it can be better while calculating mod
func hashCode(s string) int {
	code := 0
	bs := []byte(s)
	for _, b := range bs {
		code = code*31 + int(b)
	}
	return code
}
