package storage

const maxPageSize = 5

type PageId int

type Page struct {
	id       PageId
	data     [maxPageSize]byte
	isDirty  bool // indicates if the page has been modified since fetched
	pinCount int  // number of current page users
}

func NewPage(id PageId) *Page {
	var data [maxPageSize]byte
	return &Page{
		id: id,
		data: data,
		isDirty: false,
		pinCount: 0,
	}
}

func (p Page) PinCount() int {
	return p.pinCount
}

func (p *Page) SetPinCount(pinCount int) {
	p.pinCount = pinCount
}

func (p Page) PageId() PageId {
	return p.id
}
