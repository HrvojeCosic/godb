package buffer

import "github.com/HrvojeCosic/godb/src/storage"

type Replacer interface {
	evict() bool          // evict makes space in buffer pool for new page to be added
	pin(storage.PageId)   // pin pins a frame, indicating it should not be evicted until unpinned
	unpin(storage.PageId) // unpin unpins a frame, indicating it's a candidate for eviction
	getBufferPoolManager() *BufferPoolManager
}