package main

const NO_OF_PROVINCES = 200;
func initPath(from int,to int,paths [][]bool)[][]bool{
	paths[to][from] = true;
	paths[from][to] = true;
	return paths
}
func doesPathExist(from int , to int,paths [][]bool,visited []bool) bool{
	var aux = make([]bool,NO_OF_PROVINCES)
	if (from >to){
		aux = searchForPath(to, from, paths, visited)
	} else{
		aux=searchForPath(from,to,paths,visited);
	}
	return aux[to];
}
func returnPath(from int , to int,paths [][]bool,visited []bool) []bool{
	var aux = make([]bool,NO_OF_PROVINCES)
	if (from >to){
		aux = searchForPath(to, from, paths, visited)
	} else{
		aux=searchForPath(from,to,paths,visited);
	}
	return aux;
}
func moveToProvince(from int,to int, paths [][]bool) []bool {
	visited := make([]bool,NO_OF_PROVINCES);
	visited[from]=true;
	if (paths[from][to] == true) {
		visited[to] = true;
		return visited;
	}else{
		return searchForPath(from,to,paths,visited);
	}
}

func searchForPath(from int, to int, paths [][]bool, visited []bool) []bool {
	visited[from] = true;
	for i:=0;i<len(paths);i++{
		if (paths[from][i] == true){
			if (paths[from][to] == true){
				visited[to] = true;
				return visited;
			}
			visited := searchForPath(from+1,to,paths,visited);
			return visited;
		}
	}
	return visited;
}
func initDefaultPathVals(paths [][]bool)[][]bool{
	for i:=0;i<cap(paths[0]);i++{
		for j:=0;j<cap(paths);j++{
			paths[i][j]=false;
		}
	}
	return paths;
}