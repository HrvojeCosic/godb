package buffer

import (
	"sync"

	"github.com/HrvojeCosic/godb/src/storage"
)

const MaxPoolSize = 4

type FrameId int

type BufferPoolManager struct {
	pages 			[MaxPoolSize]*storage.Page  // pages is current collection of pages held by buffer pool
	availableFrames []FrameId         		    // availableFrames is available spots by frame's id for new pages to come into buffer pool 
	latch 	        sync.Mutex                  // latch protects the shared "pages"
	pageTable 		map[storage.PageId]FrameId
	diskManager 	*storage.DiskManager
	replacer 		Replacer
}

func NewBufferPoolManager(replacer Replacer) *BufferPoolManager {
	availableFrames := make([]FrameId, MaxPoolSize)
	for i := 0; i < MaxPoolSize; i++ {
		availableFrames = append(availableFrames, FrameId(i))
	}

	return &BufferPoolManager{
		pages: [MaxPoolSize]*storage.Page{},
		availableFrames: availableFrames,
		diskManager: new(storage.DiskManager),
		replacer: NewClockReplacer(MaxPoolSize),
		pageTable: make(map[storage.PageId]FrameId),
		latch: sync.Mutex{},
	}
}
