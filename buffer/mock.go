package buffer

import (
	"sync"

	"github.com/HrvojeCosic/godb/storage"
)

func mockBufferPoolManager(replacer Replacer) *BufferPoolManager {
	pages := [MaxPoolSize]*storage.Page{}
	pageTable := map[storage.PageId]FrameId{}
	frames := make([]FrameId, MaxPoolSize)

	// make page and frame indices reverse of each other
	for i := 0; i < MaxPoolSize; i++ {
		pages[i] = storage.NewPage(storage.PageId(i))
		frames[i] = FrameId(MaxPoolSize - 1 - i)
		pageTable[pages[i].PageId()] = FrameId(MaxPoolSize - 1 - i)
	}

	return &BufferPoolManager{
		pages: pages,
		diskManager: storage.NewDiskManagerMock(),
		latch:  &sync.Mutex{},
		availableFrames: frames,
		pageTable: pageTable,
		replacer: replacer,
	}
}