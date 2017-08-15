package utility

import "os"

func DoesFileExist(path string) bool{
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return  false
}