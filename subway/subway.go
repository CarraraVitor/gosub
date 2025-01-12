package subway


type pair struct {
	n Node
	p []Node
}

type Step struct {
    Lane Lane
	Path []Node
	Dir  string
}


func ParsePath(path []Node) []*Step {
	n := len(path)
	lanes := MapAllLanes()

	var steps []*Step

    step := &Step{
        Lane: lanes[path[0].Lane],
        Path: []Node{},
        Dir: "",
    }
	for i := range n - 1 {
		curr := path[i]
		next := path[i+1]
		step.Path = append(step.Path, curr)

		if curr.Lane != next.Lane {
            steps = append(steps, step)
            step = new(Step)
            step.Lane = lanes[next.Lane]
            step.Path = []Node{}
            step.Dir = ""
		}
        
        if i == n - 2 {
            if len(step.Path) == 0 {
                step.Path = append(step.Path, next)
            }

            if curr.Lane == next.Lane {
                step.Path = append(step.Path, next)
            }
            steps = append(steps, step)
        }
	}

	adj  := MakeAdjacencyList()
    for _, s := range steps {
        dir := findstepdir(*s, adj)
        s.Dir = dir
    }
    
    return steps
}

func findstepdir(s Step, adj map[int][]Node )string {
    if len(s.Path) < 2 {
        return ""
    }
    first := s.Path[0]
    second := s.Path[1]
	visited := make(map[int]bool)
    visited[first.Id] = true
    visited[second.Id] = true
    var queue []Node
    queue = append(queue, second)
    for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		all_nbors := adj[cur.Id]
        var valid []Node
        for _, n := range all_nbors {
            if n.Lane == first.Lane && !visited[n.Id] {
                valid = append(valid, n) 
                visited[n.Id] = true
            }
        }
        if len(valid) == 0 {
            return cur.Name
        }
        queue = append(queue, valid...)
    }
    return ""
}

func FindPaths(source int, dest int) [][]Node {
	var res [][]Node

	n_map := MapAllNodes()
	n_adj := MakeAdjacencyList()

	visited := make(map[int]bool)
	var queue []pair
	queue = append(queue, pair{n_map[source], []Node{n_map[source]}})
	visited[source] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		neighbors := n_adj[cur.n.Id]
		for _, neighbor := range neighbors {
			path := make([]Node, len(cur.p))
			copy(path, cur.p)
			path = append(path, neighbor)

			if neighbor.Id == dest {
				res = append(res, path)
			} else if visited[neighbor.Id] == false {
				queue = append(queue, pair{n: neighbor, p: path})
			}

			visited[neighbor.Id] = true
		}
	}

	return res
}
