package graphs

/*
Dijkstras algo is a single source shortest path algo for graphs with
NON NEGATIVE edge weights

Depending on how the algo is implemented and what DS's are used the TC is usually
O(E*log(v)) which is competitive against other shortest path algos

The constraint for dijkstras algo is that the graph must only contain non negative
edge weights, this is imposed to ensure that once a node has been visited its
optimal distance cannot be improved

This property allows dijkstras algo to act ina greedy manner by always selecting
the next most promising node

maintain an array called 'dist' where the distance to every node is positive infinity
mark the distance to the start node 's' to be 0
PQ = priority queue
maintain a PQ of key-value pairs of (node index, distance) pairs which tell you
which node to visit next based on sorted min value

Insert (s,0) into PQ and loop while PQ is not empty, pulling out the next most promising (node index, distance) pair

Iterate over all edges outwards from the current node and relax each edge appending a new (node index, distance) key-val
pair to the PQ for every iteration
   Values in between slashes and lines represent distances
    1
 4 /| \  1
 /  |  \   3
0   |   3 --- 4
 \  |  /
 1\ | /  5
    2

		each node initialized to +infinity
						0		 1		2			3			4
dist arr = [inf, inf, inf, inf, inf]
Priority queue is a heap
Lazy implementation because we lazily delete outdated pairs
Hashmap contains:
(index, dist)
			K:V
		(0:0) <- initialize with (0:0)
		(1:4) <-- best distance from node 0 to node 1 is dist[0] + edge.cost = 0 + 4 = 4
		(2:1) <-- best distance from node 0 to node 2 is dist[0] + edge.cost = 0 + 1 = 1
		Concludes visiting edges for node 0, djisktras always selects the next most promising node (greedy)
		Pull next best node from PQ
		_________
		Node 2 is the next most promising node
		(1:3) Best distance from node 2 to node 1 is dist[2] + edge cost 1 + 2 = 3   path is 0 -> 2 -> 1 == 1+2 = 3
		! (1:3) is better than (1:4), so (1:4) is replaced
		(3:6) This is because 0->2 edge distance is 1, 2 -> 3 edge distance is 5 || 1 + 5 = 6
		(3:4) This is because 0-> 2 -> 1 -> 3 = 1 + 2 + 1 = 4 Replace 3:6
		(4:7) Bc 3 -> 4 edge distance is 3 || 4 + 3 = 7
		FIN
Distance array keeps track of best route so far, and priority queue which tells us what to visit next

*/

func dijkstras() {

}
