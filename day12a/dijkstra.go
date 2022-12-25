package main

import "container/heap"

type Move[N any] struct {
	Node N   // The node to move to
	Cost int // The cost of the move
}

type smallestHolder[N any] struct {
	Node      N   // The node
	TotalCost int // The sum of the moves to the node
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[N any] []smallestHolder[N]

func (pq PriorityQueue[N]) Len() int { return len(pq) }

func (pq PriorityQueue[N]) Less(i, j int) bool {
	return pq[i].TotalCost < pq[j].TotalCost
}

func (pq PriorityQueue[N]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[N]) Push(x any) {
	item := x.(smallestHolder[N])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[N]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func ShortestPath[Node comparable](source Node, neighbours func(Node) []Move[Node], isFinished func(Node) bool) int {
	dist := map[Node]int{source: 0}
	prev := make(map[Node]Node)

	toSee := PriorityQueue[Node]{smallestHolder[Node]{Node: source, TotalCost: 0}}
	heap.Init(&toSee)

	for 0 < toSee.Len() {
		smallest := heap.Pop(&toSee).(smallestHolder[Node])

		if isFinished(smallest.Node) {
			// We found the node we are searching for
			return dist[smallest.Node]
		}

		if dist[smallest.Node] < smallest.TotalCost {
			continue
		}

		for _, move := range neighbours(smallest.Node) {
			d, ok := dist[move.Node]
			alt := dist[smallest.Node] + move.Cost

			if !ok || alt < d {
				dist[move.Node] = alt
				prev[move.Node] = smallest.Node
				heap.Push(&toSee, smallestHolder[Node]{Node: move.Node, TotalCost: alt})
			}
		}
	}

	return -1
}
