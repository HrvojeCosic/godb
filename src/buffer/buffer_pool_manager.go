package buffer

import "github.com/HrvojeCosic/godb/src/storage"

const MaxPoolSize = 4

type FrameId int

type BufferPoolManager struct {
	diskManager *storage.DiskManager
	ClockReplacer *ClockReplacer // ClockReplacer is replacment policy (could be any other RP)
	pages [MaxPoolSize] *storage.Page // pages is current collection of pages held by buffer pool
	availableFrames []FrameId // availableFrames is available spots by frame's id for new pages to come into buffer pool 
	pageTable map[storage.PageId]FrameId
}

func NewBufferPoolManager() *BufferPoolManager {
	diskManager := new(storage.DiskManager)
	ClockReplacer := new(ClockReplacer)
	pageTable := make(map[storage.PageId]FrameId)
	var availableFrames []FrameId
	var pages [MaxPoolSize]*storage.Page

	for i := 0; i < MaxPoolSize; i++ {
		availableFrames = append(availableFrames, FrameId(i))
	}

	return &BufferPoolManager{diskManager, ClockReplacer, pages, availableFrames, pageTable}
}
