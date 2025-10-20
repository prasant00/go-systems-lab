package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	arr []int       // stores all elements
	pos map[int]int // maps element -> index in arr
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // seed RNG (optional)
	return RandomizedSet{
		arr: make([]int, 0),
		pos: make(map[int]int),
	}
}

// Insert inserts val into the set if not already present.
// Returns true if the value was inserted, false if it already exists.
func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.pos[val]; exists {
		return false
	}
	rs.arr = append(rs.arr, val)
	rs.pos[val] = len(rs.arr) - 1
	return true
}

// Remove removes val from the set if present.
// Returns true if the value was removed, false if it wasn't found.
func (rs *RandomizedSet) Remove(val int) bool {
	idx, exists := rs.pos[val]
	if !exists {
		return false
	}

	lastIdx := len(rs.arr) - 1
	lastVal := rs.arr[lastIdx]

	// Move the last element to the deleted element's index
	rs.arr[idx] = lastVal
	rs.pos[lastVal] = idx

	// Remove the last element
	rs.arr = rs.arr[:lastIdx]
	delete(rs.pos, val)
	return true
}

// GetRandom returns a random element from the set.
// Each element has equal probability of being chosen.
func (rs *RandomizedSet) GetRandom() int {
	return rs.arr[rand.Intn(len(rs.arr))]
}
