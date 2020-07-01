package main

func abolishLaw(i int, laws []Law) []Law {
	//[TESTED]
	var laws2 = append(laws[:i], laws[i+1:]...)
	return laws2
}
