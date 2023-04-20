package buffer

import (
	"testing"

	"github.com/HrvojeCosic/godb/storage"
)

func TestFetchPage(t *testing.T) {
	replacer := NewClockReplacer(2)
	bpm := NewBufferPoolManager(replacer)
	bpm.diskManager.WritePage(storage.NewPage(1))
	bpm.diskManager.WritePage(storage.NewPage(2))
	bpm.diskManager.WritePage(storage.NewPage(3))

	// Test fetching with no pages in bpm
	fid := storage.PageId(2)
	page, _ := bpm.FetchPage(fid)
	if (page.PageId() != fid) {
		t.Errorf("Expected page id to be %d, but got %d", fid, page.PageId())
	}

	// Test fetching page that is already in bpm
	page1, _ := bpm.FetchPage(fid)
	if (page1.PageId() != fid) {
		t.Errorf("Expected page id to be %d, but got %d", fid, page.PageId())
	}

	// Test fetching when bpm is full
	bpm.FetchPage(storage.PageId(3))
	fid2 := storage.PageId(1)
	page2, _ := bpm.FetchPage(fid2)
	if (page2.PageId() != storage.PageId(1)) {
		t.Errorf("Expected page id to be %d, but got %d", fid2, page2.PageId())
	}
}

func TestUnpinPage(t *testing.T) {
	replacer := NewClockReplacer(3)
	bpm := NewBufferPoolManager(replacer)

	np1 := storage.NewPage(1)
	np3 := storage.NewPage(3)
	bpm.diskManager.WritePage(np1)
	bpm.diskManager.WritePage(np3)
	bpm.diskManager.WritePage(storage.NewPage(2))
	
	// Test unpinning with empty bpm
	ok := bpm.UnpinPage(2)
	if (ok) {
		t.Errorf("Expected to return %t, but got %t", false, ok)
	}
	
	// Test unpinning already unpinned page
	bpm.FetchPage(1)
	np1.SetPinCount(0)
	ok1 := bpm.UnpinPage(1)
	if (ok1) {
		t.Errorf("Expected to return %t, but got %t", false, ok)
	}

	// Test unpinning pinned page (FetchPage() also pins it)
	bpm.FetchPage(3)
	ok2 := bpm.UnpinPage(3)
	if (!ok2 || np3.PinCount() != 0) {
		t.Errorf("Expected to return %t and page pin count to be 0 when unpinning pinned page, but got %t and %d", true, ok2, np3.PinCount())
	}
}
