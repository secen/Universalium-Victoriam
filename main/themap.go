package main

import "strings"
const (
	entities = "A:E"
	heights = "0:1:2:3"
	forts = "-1:-2:-3:-4"
	sepa = ":"
)
func mapEnemyExists(mp string) bool {
	return strings.Contains(mp,"E");
}
func mapHasMountains(mp string) bool {
	return strings.Contains(mp,"3")
}
func mapDeadExistOnMap(mp string) bool {
	return strings.Contains(mp,"D");
}
func mapHasForests(mp string) bool {
	return strings.Contains(mp,"%");
}
func mapHasWater(mp string) bool {
	contains := strings.Contains(mp, "W");
	return contains;
}
func mapHasForts(mp string) bool{
	return strings.Contains(mp,"-")
}
func getMap(filename string) string{
	return readFromFile(filename)
}
func mapGet2DArr(mp string) []string{
	return strings.Split(mp,"\n");
}
func mapFrom2DArr(mp []string) string{
	var aux string = "";
	for i:=0;i<len(mp);i++{
		aux+=mp[i];
		aux+= "\n";
	}
	return aux;
}
func getEntitiesMap(mp string)string{
	var aux string = mp;
	aux = strings.ReplaceAll(aux,"0","X")
	aux = strings.ReplaceAll(aux,"1","X")
	aux = strings.ReplaceAll(aux,"2","X")
	aux = strings.ReplaceAll(aux,"%","X")
	aux = strings.ReplaceAll(aux,"W","X")
	aux = strings.ReplaceAll(aux,"-1","X")
	return aux;
}
func getHeightMap(mp string)string{
	var aux string = mp;
	aux = strings.ReplaceAll(aux,"A","X")
	aux = strings.ReplaceAll(aux,"E","X")
	aux = strings.ReplaceAll(aux,"D","X")
	aux = strings.ReplaceAll(aux,"%","X")
	aux = strings.ReplaceAll(aux,"W","0")
	aux = strings.ReplaceAll(aux,"-1","X")
	return aux;
}