package buffer

type Replacer interface {
	Evict() FrameId // makes space in buffer pool for new page to be added
	Pin(FrameId)    // pins a frame, indicating it should not be evicted until unpinned
	Unpin(FrameId)  // unpins a frame, indicating it's a candidate for eviction
	Size() uint     // returns number of elements currently in the replacer that can be evicted
}