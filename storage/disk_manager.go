package storage

type DiskManager interface {
	ReadPage(PageId) (*Page, error) // read the contents of the specified page
	WritePage(*Page) bool   // write the contents of the specified page into disk file
}