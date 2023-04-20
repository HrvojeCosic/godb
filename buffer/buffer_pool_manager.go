package buffer

import (
	"sync"
	"github.com/HrvojeCosic/godb/storage"
)

const MaxPoolSize = 4

type FrameId int

type BufferPoolManager struct {
	pages 		[MaxPoolSize]*storage.Page          // pages is current collection of pages held by buffer pool
	availableFrames []FrameId         		    // availableFrames is available spots by frame's id for new pages to come into buffer pool 
	latch 	        *sync.Mutex                         // latch protects the shared "pages"
	pageTable 	map[storage.PageId]FrameId
	diskManager 	storage.DiskManager
	replacer 	Replacer
}

func NewBufferPoolManager(replacer Replacer) *BufferPoolManager {
	availableFrames := make([]FrameId, MaxPoolSize)
	for i := 0; i < MaxPoolSize; i++ {
		availableFrames[i] = FrameId(i)
	}

	return &BufferPoolManager{
		pages: [MaxPoolSize]*storage.Page{},
		availableFrames: availableFrames,
		diskManager: storage.NewDiskManagerMock(),
		replacer: NewClockReplacer(MaxPoolSize),
		pageTable: make(map[storage.PageId]FrameId),
		latch: &sync.Mutex{},
	}
}

// Fetch the requested page from the buffer pool.
func (bpm *BufferPoolManager) FetchPage(pageId storage.PageId) (*storage.Page, error) {
	bpm.latch.Lock()
	defer bpm.latch.Unlock()

	frameId, ok := bpm.pageTable[pageId]
	if (ok) {
		page := bpm.pages[frameId] 
		page.SetPinCount(page.PinCount() + 1)
		bpm.replacer.Pin(frameId)
		return page, nil
	}

	newFrameId := FrameId(-1)
	if (len(bpm.availableFrames) == 0) {
		newFrameId = bpm.replacer.Evict()
	} else {
		newFrameId = bpm.availableFrames[0]	
		bpm.availableFrames = bpm.availableFrames[1:]
	}
	readPage, err := bpm.diskManager.ReadPage(pageId)
	if (err != nil) {
		return nil, err
	}
	if (bpm.pages[newFrameId] != nil && bpm.pages[newFrameId].IsDirty()) {
		bpm.diskManager.WritePage(readPage)
	}
	bpm.pages[newFrameId] = readPage
	bpm.pageTable[pageId] = newFrameId 
	readPage.SetPinCount(readPage.PinCount() + 1)
	bpm.replacer.Pin(newFrameId)
	return readPage, nil
}

// Unpin the page from the buffer pool. If requested page is not in buffer pool, or it's pin count is already 0, return false, otherwise true
func (bpm *BufferPoolManager) UnpinPage(pageId storage.PageId) bool {
	bpm.latch.Lock()
	defer bpm.latch.Unlock()

	frameId, ok := bpm.pageTable[pageId]
	if (!ok || bpm.pages[frameId].PinCount() == 0) {
		return false
	} else {
		bpm.pages[frameId].SetPinCount(0)
		bpm.replacer.Unpin(frameId)
		return true
	}
}
