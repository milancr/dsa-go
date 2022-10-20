package treesandgraphs

func canReach(arr []int, start int) bool {
	end := len(arr)
	queue := []int{}
	visited := make([]bool, end)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if arr[cur] == 0 {
			return true
		}
		visited[cur] = true

		if cur+arr[cur] < end && !visited[cur+arr[cur]] {
			queue = append(queue, cur+arr[cur])
		}

		if cur-arr[cur] >= 0 && !visited[cur-arr[cur]] {
			queue = append(queue, cur-arr[cur])
		}

	}
	return false
}
