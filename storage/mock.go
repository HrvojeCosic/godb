package storage

import "errors"

// TODO: ACTUAL DISK MANAGER (THIS IS ONLY TO IMPLEMENT BUFFER POOL)
type DiskManagerMock struct {
	diskFile map[PageId]*Page
}

func NewDiskManagerMock() *DiskManagerMock {
	return &DiskManagerMock{
		diskFile: make(map[PageId]*Page, 0),
	}
}

func (dm *DiskManagerMock) ReadPage(pageId PageId) (*Page, error) {
	page, ok := dm.diskFile[pageId]
	if (ok) {
		return page, nil
	}
	return nil, errors.New("requested page does not exist")
}

func (dm *DiskManagerMock) WritePage(pageId PageId, content *Page) bool {
	dm.diskFile[pageId] = content
	return true	
}