package main

type Packet struct {
	Priority int
	Content  []byte
}

type PacketHeap []Packet

func (h PacketHeap) Len() int           { return len(h) }
func (h PacketHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h PacketHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PacketHeap) Push(x interface{}) {
	*h = append(*h, x.(Packet))
}

func (h *PacketHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
