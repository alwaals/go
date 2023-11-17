package main

import (
	"fmt"
	"log"
)

func main() {
	var s my_interface1 = &Store{}
	s.SetStore(355, Store{Name: "Indian Bazaar", Location: "LittleEm", Employees: 250})
	ret, err := s.GetStore(355)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(ret)
}

type my_interface1 interface {
	GetStore(StoreId int) (Store, error)
	SetStore(StoreId int, s Store) error
}

type Store struct {
	StoreId   int
	Name      string
	Location  string
	Employees int
}

var stores = []Store{{Name: "Walmart", Location: "Lewisville", Employees: 100},
	{Name: "Costo", Location: "Dallas", Employees: 150},
	{Name: "DollarTree", Location: "Texas", Employees: 180}}

func (s1 *Store) GetStore(StoreId int) (Store, error) {
	for _, v := range stores {
		if v.StoreId == StoreId {
			return v, nil
		}
	}
	return Store{}, fmt.Errorf("unable to find the storeID:%v", StoreId)
}
func (s1 *Store) SetStore(StoreId int, s Store) error {
	initialLen := len(stores)
	fmt.Println(initialLen)
	s.StoreId = StoreId
	stores = append(stores, s)
	afterLen := len(stores)
	fmt.Println(afterLen)
	return nil
}
