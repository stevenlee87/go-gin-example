package cache_service

import (
	"strconv"
	"strings"

	"github.com/stevenlee87/go-gin-example/pkg/e"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
		//fmt.Printf("a.ID is %d, keys is %s\n", a.ID, keys)
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
		//fmt.Printf("a.TagID is %d, keys is %s\n", a.TagID, keys)
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
		//fmt.Printf("a.State is %d, keys is %s\n", a.State, keys)
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
		//fmt.Printf("a.PageNum is %d, keys is %s\n", a.PageNum, keys)
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
		//fmt.Printf("a.PageSize is %d, keys is %s\n", a.PageSize, keys)
	}
	//fmt.Print("keys is:", keys)
	return strings.Join(keys, "_")
}
