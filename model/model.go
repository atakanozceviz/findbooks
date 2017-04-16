package model

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Img       string `json:"img"`
	Price     string `json:"price"`
	WebSite   string `json:"website"`
}

type Books []Book

var lock sync.Mutex

func Add(b Book, bs *Books) {
	lock.Lock()
	*bs = append(*bs, Book{
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
		Img:       b.Img,
		Price:     b.Price,
		WebSite:   b.WebSite,
	})
	lock.Unlock()
}

func (bs *Books) ToJson() []byte {
	j, err := json.Marshal(bs)
	if err != nil {
		log.Println(err)
	}
	return j
}

var rep = strings.NewReplacer(",", ".", " ", "", "TL", "")

func (bs Books) GetAvg() (float64, error) {
	avg := 0.0
	i := 0.0
	for _, v := range bs {
		num, err := strconv.ParseFloat(rep.Replace(v.Price), 64)
		if err != nil {
			return 0.0, err
		}
		avg += num
		i++
	}
	return avg / i, nil
}
