package storage

type DiskManager interface {
	ReadPage(PageId) [maxPageSize]byte        // read the contents of the specified page
	WritePage(PageId, [maxPageSize]byte) bool // write the contents of the specified page into disk file
}