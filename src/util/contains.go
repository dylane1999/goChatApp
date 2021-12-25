package util

func DoesIdExist(listOfIds []string, target string) bool {
	for _, item := range listOfIds {
		if item == target {
			return true
		}
	}
	return false
}
