package localcache

import "fmt"

func (c *cache) Get() {
	fmt.Println("this is get method")
}

func (c *cache) Set() {
	fmt.Println("this is set method")
}

func New() (c Cache) {
	c = &cache{}
	return c
}
