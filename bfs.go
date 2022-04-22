package main

/*

	a -> b -> c -> f
		     |         |
		     d -> i -> j

	var graph = make(map[string][]string)
	graph["a"] = []string{"b"}
	graph["b"] = []string{"a", "c", "d"}
	graph["c"] = []string{"b", "f"}
	graph["d"] = []string{"b", "i"}
	graph["i"] = []string{"d", "j"}
	graph["f"] = []string{"c", "j"}
	graph["j"] = []string{"i", "f"}

	start := "a"
	stop := "i"

	v := search(start, stop, graph)
	curNode := stop
	
	for curNode != start {
		curNode = v[curNode][0]
		fmt.Println(curNode)
	}

*/

func search(start string, stop string, graph map[string][]string) map[string][]string {
	queue := make([]string, 0)
	queue = append(queue, graph[start]...)
	var visited = make(map[string][]string)
	var curNode string

	for len(queue) > 0 {
		curNode = queue[0]
		queue = queue[1:]

		nextNode := graph[curNode]

		if curNode == stop {
			return visited
		}

		for _, next := range nextNode {
			if wasVisited(next, visited) {
				queue = append(queue, next)
				visited[next] = []string{curNode}
			}
		}

	}
	return visited
}

func wasVisited(node string, visited map[string][]string) bool {
	for key, _ := range visited {
		if node == key {
			return false
		}
	}
	return true
}
