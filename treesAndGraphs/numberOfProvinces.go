package treesandgraphs

/* make use of parent array size N, traverse over all nodes in graph, for each node traversed
we traverse over all the nodes directly connected to it and assign them to a single group represented
by their parent node --> process is called FORMING A UNION, every group has a single parent node
whos parent is given by -1

for every new pair of nodes found, we look for the parents of both nodes, if the parents nodes are the same,
it indicated they have already been united into the same group, if parent nodes differ they have yet to be united
thus for the pair of nodes, (x, y) while forming the union, we assign parent[x] = parent[y] which combines them
into the same group

using a disjoint set is O(n^3) -> better solution is dfs but implementing both for practice
*/
func numberOfProvinces(matrix [][]int) {

}

func union(i, j int, friends []int) {

}

func find(i int, friends []int) int {

}

// https://leetcode.com/explore/featured/card/graph/618/disjoint-set/3845/
