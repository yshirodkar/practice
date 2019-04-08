/*
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

Example 1:

Input: 2, [[1,0]] 
Output: true
Explanation: There are a total of 2 courses to take. 
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: 2, [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
Note:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites, numCourses)

	visited := make(map[int]bool)
	curVisited := make(map[int]bool)
	for start, _ := range graph {
		if res := dfs(start, graph, visited, curVisited); !res {
			return false
		}
	}
	return true
}

func dfs(start int, graph map[int][]int, visited map[int]bool, curVisited map[int]bool) bool {
	if val, ok := visited[start]; ok && val {
		return true
	}
	curVisited[start] = true

	neighbors := graph[start]
	for _, neighbor := range neighbors {
		if val, ok := curVisited[neighbor]; ok && val {
			return false
		}
		if res := dfs(neighbor, graph, visited, curVisited); !res {
			return false
		}
	}
	curVisited[start], visited[start] = false, true
	return true
}

// buildGraph returns a graph as {courseID: [requirements courseIDs]} map
// and returns start courses which aren't required by any courses.
func buildGraph(prerequisites [][]int, numCourses int) map[int][]int {
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		if _, ok := graph[p[1]]; ok {
			graph[p[1]] = append(graph[p[1]], p[0])
		} else {
			graph[p[1]] = []int{p[0]}
		}
	}
	return graph
}