package main
func abolishLaw (i int, laws []law) []law {
	//[TESTED]
	var laws2 = append(laws[:i], laws[i+1:]...)
	return laws2
}