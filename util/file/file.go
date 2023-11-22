package file

func Exists(f string) bool {
	_, err := os.Stat(f)
	if err == nil {
		return true
	}
	return false
}
