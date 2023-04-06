package buffer

import "github.com/HrvojeCosic/godb/src/storage"

type ClockReplacer struct {
	bufferPoolManager *BufferPoolManager
}

func (cr ClockReplacer) getBufferPoolManager() *BufferPoolManager {
	return cr.bufferPoolManager
}

func (cr *ClockReplacer) evict() bool {
	pages := cr.bufferPoolManager.pages
	hasEvicted := false
	for pageIdx, page := range pages {
		if page.PinCount() == 0 {
			frameIdToEvict := cr.bufferPoolManager.pageTable[page.PageId()]
			availableFrames := cr.bufferPoolManager.availableFrames

			delete(cr.bufferPoolManager.pageTable, page.PageId()) // remove page table pair
 
			// TODO: FIX HACK (PAGE HAS ID OF -1, SO IT'S "EVICTED")
			*(pages[pageIdx]) = *storage.NewPage(-1, [5]byte{0}, false, 1) // evict page from buffer pool

			// mark corresponding frame as available
			for _, frameId := range availableFrames {
				if frameId == frameIdToEvict {
					availableFrames = append(availableFrames, frameIdToEvict)
				}
			}

			hasEvicted = true
			break
		} else {
			page.SetPinCount(0)
		}
	}

	return hasEvicted
}

func (cr *ClockReplacer) pin(pageId storage.PageId) {
	for _, page := range cr.bufferPoolManager.pages {
		if (page.PageId() == pageId) {
			page.SetPinCount(1)
		}
	}
}

func (cr *ClockReplacer) unpin(pageId storage.PageId) {
	for _, page := range cr.bufferPoolManager.pages {
		if (page.PageId() == pageId) {
			page.SetPinCount(0)
		}
	}
}