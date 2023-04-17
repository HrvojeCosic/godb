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