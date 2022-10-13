package disjointset

type node struct {
	value  interface{}
	parent *node
	size   int
}

// DSU is the type used to the Disjoint Set data structure
type DSU struct {
	nodes map[interface{}]*node
}

// New returns a pointer to an empty initialized DSU instance
func New() *DSU {
	return &DSU{map[interface{}]*node{}}
}

// Contains checks if a given element is present in the disjoint set
func (d *DSU) Contains(x interface{}) bool {
	_, ok := d.nodes[x]
	return ok
}

// Add adds a new element, placed into a set of only the new element, if the element already exists nothing is done
func (d *DSU) Add(x interface{}) {
	if !d.Contains(x) {
		d.nodes[x] = &node{value: x, parent: nil, size: 1}
	}
}

// Find returns the root element that represents the set to which x belongs to
// if the element does not exist in the data structure it returns the nil value
func (d *DSU) Find(x interface{}) interface{} {
	if !d.Contains(x) {
		return nil
	}
	node := d.nodes[x]
	root := node

	for root.parent != nil {
		root = root.parent
	}

	for node.parent != nil {
		node.parent, node = root, node.parent
	}

	return root.value
}

// Union replaces the set containing x and the set containing y with their union. It first determines the roots of the sets containing x and y.
// If the roots are the same or one of the elements does not exist in the data structure, there is nothing more to do. Otherwise, the two sets get be merged.
//
// The root of the new set is the root of the set with bigger size. In case both sets have the same size, the root is the root of set y.
//
// It returns true if the merge happened
func (d *DSU) Union(x, y interface{}) bool {
	rootx := d.Find(x)
	rooty := d.Find(y)

	if rootx == nil || rooty == nil || rootx == rooty {
		return false
	}

	nodex := d.nodes[rootx]
	nodey := d.nodes[rooty]

	if nodex.size <= nodey.size {
		nodex, nodey = nodey, nodex
	}

	nodey.parent = nodex
	nodex.size += nodey.size
	return true
}

// famous algorithm that uses a disjoint set is kruskals algorithm
// disjoint sets have find and union methods / functions

/*

two non connected non directed sets

	 2--3
  /    \
 1      4

 5      8
  \    /
   6--7

s1 = [1,2,3,4]
s2 = [5,6,7,8]

s1 n s2 = 0


find operation returns to us which set a node is a part of
ex find(7) -> s2

union operation adds an edge
ex union(4, 8) -> joins the sets


	 2--3
  /    \
 1      4
				| <-- union(4,8) <-- union(1,5)
 5      8
  \    /
   6--7


s1 u s2 = [1,2,3,4,5,6,7,8] = s3

find(1) in s3 == true
find(5) in s3 == true
! We have a cycle

if we take an edge and both of the vertices are in the same set, there is a cycle in the graph




How to detect cycle in graph:

u = [1,2,3,4,5,6,7,8]
   7
1-----3
|     |
|1    |2
|     |
2-----4
|	 5
6|
|  9
5-----7
|     |
|3    |4
|     |
6-----8
   8

first edge: 1,2 -> find(1), find(2) --> in universal set
remove from u, add to set  s1
s1=[1.2]
s2=[3,4]
s3=[5,6]
s4=[7,8]

s5 = s1 u s2 == [1,2,3,4], remove s1, s2
edge 6 is (2,5), find(2), find(5) and unite their sets
s6 = [1,2,3,4,5,6]
including edge 7 will form a cycle, so we ignore it
edge 8 == (6, 8) find and union
s7 = [1,2,3,4,5,6,7,8]
!lastly is edge 9, but since 5 and 7 belong to the same set they form a cycle

DETECT A CYCLE GRAPHICAL REPRESENTATION:

first edge is (1,2), we dont need to store set names, only a representation
so vertex 1 is a parent of vertex 2

1     3    ...
 \     \
  2     4   ...
Can be vice verse, doesnt matter

union s1, s2

   1  <--   represents [1,2,3,4,5,6] aka s6
 /   \    \
2     3    5
       \4    \     now we look at 7th edge, (1,3) both nodes parents are 1, so both belong to same set aka theres a cycle
			        6



SETS IN ARRAY:
parent = [-1,-1,-1,-1,-1,-1,-1, -1]
					1, 2, 3, 4, 5, 6, 7, 8
if parent[i] == negative number its a parent
step 1, find(1) find(2) -- both are -1, determined in constant time and belong to two different sets
step 2, now perform union(1, 2) 2 is child of 1 --> parent = [-2, 1,-1,-1,-1,-1,-1, -1] negative shows it a parent, 2 shows theres two nodes
step 3, find(3) find(4) -- both are parents so we perform union, 3 is the parent of 4, 3 remains negative but comes 2
because theres two nodes in the set
parent = [-2, 1, -2, 3,-1,-1,-1, -1]
continues -> parent = [-2, 1,-2, 3,-2, 5, -2, 7]
edge 5 -> (2,4) find(2), find(4) both run into different parents indicating different sets, perform union
parent = [-4, 1, 1, 3,-2, 5, -2, 7]
edge 6 -> (2, 5)
run into two different parents, perform union
because 1 has a heavier weight of -4, it becomes the parent for (5, 6)
parent = [-6, 1,-2, 3, 1, 5, -2, 7]
edge 7 -> (1, 3)
! perform find(1) find(3), they have the SAME PARENT indicating a CYCLE
parent = [-8, 1,-2, 3, 1, 5, 1, 7]
graphically all nodes point to 1

another term is collapsing find, when we follow the nodes on a large data set it will take a lot of steps
another way to build the array is, if 5's parent is 1, and 6 points to 5, then 6's parent is also 1.
so 1 can be stored as the value for 6
reduces the time complexity on repeat finds
parent = [-8, 1, 1, 1, 1, 1, 1, 1]

linked lists can be done as well
*/
