package buffer

import (
	"errors"
	"sync"

	"github.com/HrvojeCosic/godb/storage"
)

const MaxPoolSize = 4

type FrameId int

type BufferPoolManager struct {
	pages 			[MaxPoolSize]*storage.Page  // pages is current collection of pages held by buffer pool
	availableFrames []FrameId         		    // availableFrames is available spots by frame's id for new pages to come into buffer pool 
	latch 	        *sync.Mutex                 // latch protects the shared "pages"
	pageTable 		map[storage.PageId]FrameId
	diskManager 	storage.DiskManager
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
		diskManager: storage.NewDiskManagerMock(),
		replacer: NewClockReplacer(MaxPoolSize),
		pageTable: make(map[storage.PageId]FrameId),
		latch: &sync.Mutex{},
	}
}

// Fetch the requested page from the buffer pool.
func (bpm BufferPoolManager) FetchPage(pageId storage.PageId) (*storage.Page, error) {
	bpm.latch.Lock()
	defer bpm.latch.Unlock()

	frameId, ok := bpm.pageTable[pageId]
	if (ok) {
		page := bpm.pages[frameId] 
		page.SetPinCount(page.PinCount() + 1)
		return page, nil
	}

	if (len(bpm.availableFrames) == 0) {
		readPage, err := bpm.diskManager.ReadPage(pageId)
		if (err != nil) {
			return nil, err
		}
		frameId := bpm.replacer.Evict()
		bpm.pages[frameId] = readPage
		bpm.pageTable[pageId] = frameId 
		return readPage, nil
	}

	return nil, errors.New("error fetching page")
}