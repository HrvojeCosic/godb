package buffer

import (
	"testing"

	"github.com/HrvojeCosic/godb/src/storage"
)

func TestPinUnpin(t *testing.T) {
	bufferPoolManager := mockBufferPoolManager()
	clockReplacer := ClockReplacer{bufferPoolManager}
	var replacer Replacer = &clockReplacer

	testPage := bufferPoolManager.pages[0].PageId()
	replacer.unpin(storage.PageId(testPage))
	
	for _, page := range replacer.getBufferPoolManager().pages {
		if (page.PageId() == testPage) {
			if (page.PinCount() != 0) {
				t.Errorf("Expected page's pin count to be 0, got %d", page.PageId())
			} 
			
			replacer.pin(storage.PageId(testPage)) // test pin on the same page
			if (page.PinCount() != 1) {
				t.Errorf("Expected page's pin count to be 1, got %d", page.PageId())
			}
		}
	}
}

func TestEvictWhenAllPagesArePinned(t *testing.T) {
	bufferPoolManager := mockBufferPoolManager()
	clockReplacer := ClockReplacer{bufferPoolManager}
	var replacer Replacer = &clockReplacer
	hasEvicted := false

	hasEvicted = replacer.evict()
	if (hasEvicted == true) {
		t.Errorf("Expected clock replacer to return false, but returned %t", hasEvicted)
	}

	for _, page := range replacer.getBufferPoolManager().pages {
		if (page.PinCount() != 0) {
			t.Error("Expected clock replacer to set pin count to 0 after checking the frame")
		}
	}
}

// // TODO: TEST THREAD-SAFETY
func TestEvictWithSomePagesUnpined(t *testing.T) {
	bufferPoolManager := mockBufferPoolManager()
	clockReplacer := ClockReplacer{bufferPoolManager}
	var replacer Replacer = &clockReplacer
	hasEvicted := false
	
	pageIdToUnpin := bufferPoolManager.pages[1].PageId()
	bufferPoolManager.pages[pageIdToUnpin].SetPinCount(0)
	frameIdToEvict := bufferPoolManager.pageTable[pageIdToUnpin]
	hasEvicted = replacer.evict()
	
	if (hasEvicted == false) {
		t.Errorf("Expected clock replacer to return true, but returned %t", hasEvicted)
	}

	for _, page := range bufferPoolManager.pages {
		if (page.PageId() == pageIdToUnpin) {
			t.Log(pageIdToUnpin)
			t.Error("Expected evicted page not to take up frames")
		}
	}

	isFreed := false
	for _, frame := range bufferPoolManager.availableFrames {
		if (frame == frameIdToEvict) {
			isFreed = true
		}
	}
	if (!isFreed) {
		t.Error("Expected frame corresponding to the evicted page to be made available")
	}

	_, exists := bufferPoolManager.pageTable[pageIdToUnpin]
	if (exists) {
		t.Log(bufferPoolManager.pageTable)
		t.Error("Expected evicted page not to be in page table")
	}
}