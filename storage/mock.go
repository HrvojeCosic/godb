package storage

// TODO: ACTUAL DISK MANAGER (THIS IS ONLY TO IMPLEMENT BUFFER POOL)
type DiskManagerMock struct {
	diskFile map[PageId][maxPageSize]byte
}

func NewDiskManagerMock() *DiskManagerMock {
	return &DiskManagerMock{
		diskFile: make(map[PageId][maxPageSize]byte, 0),
	}
}

func (dm *DiskManagerMock) ReadPage(pageId PageId) [maxPageSize]byte {
	return dm.diskFile[pageId]
}

func (dm *DiskManagerMock) WritePage(pageId PageId, contents [maxPageSize]byte) bool {
	dm.diskFile[pageId] = contents
	return true
}