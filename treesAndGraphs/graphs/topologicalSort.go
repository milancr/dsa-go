package graphs

/* topological ordering is an ordering of each of the nodes in a directed
graph where for each directed edge from node A to node B, node A appears before
Node B in the ordering

Kahns algorithm is a simple topo sort algo that can find topo ordering in O(V+E) time

NOTE: TOPO ORDERINGS ARE NOT UNIQUE, there can be many valid orderings
Only Directed Acyclic graphs (DAG's) have topological sorts

Kahn's Algorithm:

Intuition: repeatedly remove nodes without any dependencies from the graph and
add them to the topological ordering

as nodes without deps and their outgoing edges are removed from the graph
new nodes without deps should become free

repeat until all nodes are process or cycle is discovered

diagonals all point from left to right
	0-->1
 / \ /^
2   3 |
 \ /  |
	4-->5

start at node 2 bc its the only node wtout deps,
add node two to topo order
now 0 and 4 dont have deps
select either and add to topo order
select other node wtout deps and add to topo order
continue with 3, 5, 1

ans = [2, 0, 4, 3, 5, 1]

In more complex cases, we need to keep track of the degree of each node

13

  2<--0
 / \ /
9   6   all diagonals point to 6
 \ /    6 has a degree of 3
  10

create map of node values:degrees
create a queue used to store all nodes with no incoming edges
add all nodes with 0 degree to queue,
decrement the degrees of each nodes connected to the nodes added to the queue
repeat steps enqueueing all new nodes with 0 degrees


*/

func topoSort(graph [][]int, n int) {

}
