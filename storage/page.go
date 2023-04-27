package storage

const maxPageSize = 5

type PageId int

type Page struct {
	Id       PageId
	data     [maxPageSize]byte
	IsDirty  bool // indicates if the page has been modified since fetched
	PinCount int  // number of current page users
}

func NewPage(id PageId) *Page {
	var data [maxPageSize]byte
	return &Page{
		Id: id,
		data: data,
		IsDirty: false,
		PinCount: 0,
	}
}
