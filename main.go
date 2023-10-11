package main

import (
	"fmt"
	"math/rand"
)

// all operations need to be O(1)
// class that allows a user to insert
// values can be removed
// a method GetRandom that returns a random number from the class

type DataStore struct {
	items    map[int]int
	itemList []int
}

func NewDataStore() *DataStore {
	return &DataStore{
		items: make(map[int]int),
	}
}

func (d *DataStore) Insert(item int) (int, error) {
	if _, exists := d.items[item]; exists {
		return 0, fmt.Errorf("%d is already in the datastore", item)
	}
	d.itemList = append(d.itemList, item)
	d.items[item] = len(d.itemList) - 1
	return item, nil
}

func (d *DataStore) Remove(item int) (bool, error) {
	if _, exists := d.items[item]; !exists {
		return false, fmt.Errorf("%d doesn't exists in the datastore", item)
	}
	d.itemList[d.items[item]], d.itemList[len(d.itemList)-1] = d.itemList[len(d.itemList)-1], d.itemList[d.items[item]]
	d.itemList = d.itemList[:len(d.itemList)-1]

	d.items[d.itemList[len(d.itemList)-1]] = d.items[item]

	delete(d.items, item)

	return true, nil
}

func (d *DataStore) GetRandom() int {
	randomIndex := rand.Intn(len(d.items))
	return d.itemList[randomIndex]
}

func main() {

	// input
	store := NewDataStore()
	store.Insert(1)
	store.Insert(2)
	store.Insert(3)
	store.Insert(7)
	_, err := store.Insert(1)
	if err != nil {
		fmt.Println(err)
	}
	removeInt := 3
	removed, err := store.Remove(removeInt)
	if err != nil {
		fmt.Printf("couldn't remove %d from store", removeInt)
	}
	fmt.Printf("removed %d from store %t\n", removeInt, removed)

	// 1, 2, 4
	fmt.Println(store.items)

	// get random
	fmt.Println("random", store.GetRandom())

}
