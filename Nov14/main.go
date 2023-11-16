package main

import "fmt"

type Entity struct {
	ID string
	Name string
	Queries map[string]string
	Status bool
}
var cache = NewCache()

type Cache struct {
	keyValue map[string]Entity
}

func NewCache() *Cache {
	return &Cache{
		keyValue: make(map[string]Entity),
	}
}

type Store interface {
	Set(id string, entity Entity) error
	Get(id string) (Entity, error)
}

func Set(id string, entity Entity) error {
	err:=cache.Set(id,entity)
	if err!=nil{
		return err
	}
	return nil
}
func Get(id string) (Entity, error){
	user:=cache.GetFromCache(id)
	fmt.Println(user.(Entity))
	if len(user.(Entity).ID) <=0 {
		return user.(Entity),fmt.Errorf("unable to find the record %s",id)
	}
	return user.(Entity),nil
}
func (c *Cache) Set(key string, value interface{}) error{
	_,k:=c.keyValue[key]
	if !k{
	c.keyValue[key] = value.(Entity)
	}
	return fmt.Errorf("unable to set the key as already exists")
}

func (c *Cache) GetFromCache(key string) interface{} {
	 return c.keyValue[key]
}

func (c *Cache) Delete(key string) {
	delete(c.keyValue, key)
}