package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/go-redis/redis/v8"
)

var ctx=context.Background()
var client = SetUpRedisCluster()
func main() {
	var s Store = Store{StoreId: "101",Name: "Walmart",Location: "Texas",Employees: 1000}
	var store StoreInterface = s
	err:=store.SetStoreDetails("101",s)
	if err!=nil{
		log.Fatal(err.Error())
	}
	retStore,errG:=store.GetStoreDetails("101")
	if errG!=nil{
		log.Fatal(errG.Error())
	}
	fmt.Printf("Store details:%+v\n",retStore)
	
}
type StoreInterface interface {
	GetStoreDetails(storeId string) (Store, error)
	SetStoreDetails(storeId string, s Store) error
}
type Store struct {
	StoreId   string `json:"StoreID"`
	Name      string `json:"Name"`
	Location  string `json:"Location"`
	Employees int    `json:"Employees"`
}
func SetUpRedisCluster() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		//DB:       0,
	})
}
func (s Store) GetStoreDetails(storeId string) (Store, error){
	value,err:=client.Get(ctx,storeId).Result()
	if err!=nil{
		return Store{},fmt.Errorf("unable to get the key %s with error:%s",storeId,err.Error())
	}
	fmt.Printf("%+v\n",value)
	var store Store
	json.Unmarshal([]byte(value),&store)
	return store,nil
}
func (s Store) SetStoreDetails(storeId string, s1 Store) error{
	bytes,err:=json.Marshal(s)
	if err!=nil{
		log.Fatal(err.Error())
	}
	return client.Set(ctx,storeId,bytes,0).Err()
}