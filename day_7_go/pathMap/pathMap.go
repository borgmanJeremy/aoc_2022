package pathMap

type PathMap struct {
	Id   int
	PMap map[string]int
}

func (pathMap *PathMap) AddKey(path string) {

	if _, ok := pathMap.PMap[path]; !ok {
		pathMap.PMap[path] = pathMap.Id
		pathMap.Id++
	}
}

func New() PathMap {
	return PathMap{
		Id:   0,
		PMap: make(map[string]int),
	}
}
