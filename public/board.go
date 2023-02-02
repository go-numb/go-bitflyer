package public

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/emirpasic/gods/maps/treebidimap"
	"github.com/emirpasic/gods/utils"
	"github.com/google/go-querystring/query"
)

type Board struct {
	ProductCode string `url:"product_code,omitempty"`
}

type ResponseForBoard struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Book  `json:"bids"`
	Asks     []Book  `json:"asks"`
}

type Book struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

func (req *Board) IsPrivate() bool {
	return false
}

func (req *Board) Path() string {
	return "board"
}

func (req *Board) Method() string {
	return http.MethodGet
}

func (req *Board) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Board) Payload() []byte {
	return nil
}

// Optional

type SortedBooks struct {
	MidPrice float64
	Asks     SortedBook
	Bids     SortedBook
}

type SortedBook struct {
	sync.RWMutex

	*treebidimap.Map
}

func NewSortedBooks() *SortedBooks {
	return &SortedBooks{
		Asks: SortedBook{
			sync.RWMutex{},
			treebidimap.NewWith(utils.StringComparator, utils.StringComparator),
		},
		Bids: SortedBook{
			sync.RWMutex{},
			treebidimap.NewWith(utils.StringComparator, utils.StringComparator),
		},
	}
}

func (p *SortedBook) Set(books []Book) {
	p.Lock()
	defer p.Unlock()

	for i := 0; i < len(books); i++ {
		key := fmt.Sprintf("%f", books[i].Price)
		if books[i].Size == 0 {
			p.Remove(key)
			continue
		}
		p.Put(key, fmt.Sprintf("%f", books[i].Size))
	}
}

func (p *SortedBook) GetVolume(price float64) (size float64) {
	p.RLock()
	defer p.RUnlock()

	v, ok := p.Get(fmt.Sprintf("%f", price))
	if v == nil || !ok {
		return 0
	}

	size, _ = strconv.ParseFloat(v.(string), 64)
	return size
}

// Sort : limitdepth is book length
func (p *SortedBook) Sort(ascending bool) []Book {
	p.RLock()
	defer p.RUnlock()

	books := make([]Book, p.Size())
	it := p.Iterator()
	if ascending {
		i := 0
		for it.End(); it.Prev(); {
			key, size := it.Key(), it.Value()
			f, _ := strconv.ParseFloat(key.(string), 64)
			v, _ := strconv.ParseFloat(size.(string), 64)
			books[i] = Book{
				Price: f,
				Size:  v,
			}
			i++
		}
	} else {
		i := 0
		for it.Begin(); it.Next(); {
			key, size := it.Key(), it.Value()
			f, _ := strconv.ParseFloat(key.(string), 64)
			v, _ := strconv.ParseFloat(size.(string), 64)
			books[i] = Book{
				Price: f,
				Size:  v,
			}
			i++
		}
	}

	return books
}
