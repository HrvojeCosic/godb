package buffer

import (
	"testing"
)

func TestPinUnpin(t *testing.T) {
	bpm := mockBufferPoolManager(NewClockReplacer(10))
	for _, page := range bpm.pages {
		correspondingFrame := bpm.pageTable[page.PageId()]
		bpm.replacer.Unpin(correspondingFrame)
	}
	if (bpm.replacer.Size() != uint(len(bpm.pages))) {
		t.Error("Expected replacer size to be 0 after unpinning all pages")
	}

	bpm.replacer.Pin(2)
	bpm.replacer.Pin(-1)
	bpm.replacer.Pin(3)
	if (bpm.replacer.Size() != uint(len(bpm.pages)) - 2) {
		t.Error("Expected replacer size to be two less than pages length after unpinning one valid and two invalid pages")
	}
}

// TODO: TEST THREAD-SAFETY
func TestEvict(t *testing.T) {
	bpm := mockBufferPoolManager(NewClockReplacer(10))
	for _, page := range bpm.pages {
		correspondingFrame := bpm.pageTable[page.PageId()]
		bpm.replacer.Unpin(correspondingFrame)
	}

	bpm.replacer.Pin(0)
	bpm.replacer.Pin(2)
	expectedFrameId := FrameId(1)
	actualFrameId := bpm.replacer.Evict()
	if (actualFrameId != expectedFrameId) {
		t.Errorf("Expected frame with id of %d to be evicted, evicted %d instead", expectedFrameId, actualFrameId)
	}
	
	expectedFrameId2 := FrameId(3)
	actualFrameId2 := bpm.replacer.Evict()
	if (actualFrameId2 != expectedFrameId2) {
		t.Errorf("Expected frame with id of %d to be evicted, evicted %d instead", expectedFrameId2, actualFrameId2)
	}
}