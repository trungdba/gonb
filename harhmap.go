package coll

import "reflect"

type Harshmap[V any] map[interface{}]V

// constructor creates a new harshmap
func NewHarshMap[V any]() Harshmap[V] {
	return make(Harshmap[V])
}

// method to add a new element to the harshmap
func (hm Harshmap[V]) Add(key interface{}, value V) {
	hm[key] = value
}

// method to add an element to the harshmap but only if an element with the same key does not already exist
func (hm Harshmap[V]) AddIfAbsent(key interface{}, value V) {
	if !hm.ContainsKey(key) {
		hm[key] = value
	}
}

// method to all all of the elements from another harshmap into this one
func (hm Harshmap[V]) AddAll(m Harshmap[V]) {
	if m == nil {
		return
	}
	for k, v := range m {
		hm[k] = v
	}
}

// method to remove the element from the harshmap
func (hm Harshmap[V]) Remove(key interface{}) {
	delete(hm, key)
}

// method to remove all elements from the harshmap
func (hm *Harshmap[V]) RemoveAll() {
	*hm = make(Harshmap[V])
}

// method to get the element from the harshmap
func (hm Harshmap[V]) Get(key interface{}) V {
	return hm[key]
}

// method to check if a key belongs to the harshmap
func (hm Harshmap[V]) ContainsKey(key interface{}) bool {
	_, found := hm[key]
	return found
}

// method to check if a value belongs to the harshmap
func (hm Harshmap[V]) ContainsValue(value V) bool {
	found := false
	for _, v := range hm {
		if reflect.ValueOf(v).Interface() == reflect.ValueOf(value).Interface() {
			found = true
		}
	}
	return found
}

// method to clear the harshmap
func (hm *Harshmap[V]) Clear() {
	*hm = nil
}

// method to count the total elements of the harshmap
func (hm Harshmap[V]) Length() int {
	return len(hm)
}

// method to convert the harshmap to a array of keys
func (hm Harshmap[V]) Keys() []interface{} {
	var keys []interface{}
	for k := range hm {
		keys = append(keys, k)
	}
	return keys
}

// method to convert the harshmap to a array of values
func (hm Harshmap[V]) Values() []V {
	var values []V
	for _, v := range hm {
		values = append(values, v)
	}
	return values
}
