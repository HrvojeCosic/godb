package buffer

import "github.com/HrvojeCosic/godb/src/storage"

func mockBufferPoolManager() *BufferPoolManager {
	pages := [MaxPoolSize]*storage.Page{}
	pageTable := map[storage.PageId]FrameId{}
	frames := make([]FrameId, MaxPoolSize)

	// make page and frame indices reverse of each other
	for i := 0; i < MaxPoolSize; i++ {
		pages[i] = storage.NewPage(storage.PageId(i), [5]byte{0}, false, 1)
		frames[i] = FrameId(MaxPoolSize - 1 - i)
		pageTable[pages[i].PageId()] = FrameId(MaxPoolSize - 1 - i)
	}

	return &BufferPoolManager{&storage.DiskManager{}, pages, frames, pageTable}
}