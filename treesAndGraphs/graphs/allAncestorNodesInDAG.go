package graphs

import "fmt"

func getAncestors(n int, edges [][]int) [][]int {
	// make map int to arr creating adjacency list
	m := map[int][]int{}
	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
	}
	fmt.Println(m)
	// map[0:[3 4] 1:[3] 2:[4 7] 3:[5 6 7] 4:[6]]
	ans := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for idx, arr := range m {
			for _, val := range arr {
				if i == val {
					temp = append(temp, idx)
					// need another loop to look back through map for idx val?
					// if m[]
				}
			}
		}
		// sort temp
		// if since 3 goes to 5, we also need to look at what goes to 3, which would be 0, 1
		ans = append(ans, temp)
	}
	fmt.Println(ans)
	return edges
}
