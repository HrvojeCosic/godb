package storage

const maxPageSize = 5

type PageId int

type Page struct {
	id       PageId
	data     [maxPageSize]byte
	isDirty  bool
	pinCount int
}

func NewPage(id PageId, data [maxPageSize]byte, isDirty bool, pinCount int) *Page {
	return &Page{id, data, isDirty, pinCount}
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
