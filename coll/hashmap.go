package coll

type Hashmap map[interface{}]interface{}

// constructor creates a new hashmap
func NewHashMap() Hashmap {
	return make(Hashmap)
}

// method to add a new element to the hashmap
func (hm Hashmap) Add(key interface{}, value interface{}) {
	hm[key] = value
}

// method to add an element to the hashmap but only if an element with the same key does not already exist
func (hm Hashmap) AddIfAbsent(key interface{}, value interface{}) {
	if !hm.ContainsKey(key) {
		hm[key] = value
	}
}

// method to all all of the elements from another hashmap into this one
func (hm Hashmap) AddAll(m Hashmap) {
	if m == nil {
		return
	}
	for k, v := range m {
		hm[k] = v
	}
}

// method to remove the element from the hashmap
func (hm Hashmap) Remove(key interface{}) {
	delete(hm, key)
}

// method to remove all elements from the hashmap
func (hm *Hashmap) RemoveAll() {
	*hm = make(Hashmap)
}

// method to get the element from the hashmap
func (hm Hashmap) Get(key interface{}) interface{} {
	return hm[key]
}

// method to check if a key belongs to the hashmap
func (hm Hashmap) ContainsKey(key interface{}) bool {
	_, found := hm[key]
	return found
}

// method to check if a value belongs to the hashmap
func (hm Hashmap) ContainsValue(value interface{}) bool {
	found := false
	for _, v := range hm {
		if v == value {
			found = true
		}
	}
	return found
}

// method to clear the hashmap
func (hm *Hashmap) Clear() {
	*hm = nil
}

// method to count the total elements of the hashmap
func (hm Hashmap) Length() int {
	return len(hm)
}

// method to convert the hashmap to a array of keys
func (hm Hashmap) Keys() []interface{} {
	var keys []interface{}
	for k := range hm {
		keys = append(keys, k)
	}
	return keys
}

// method to convert the hashmap to a array of values
func (hm Hashmap) Values() []interface{} {
	var values []interface{}
	for _, v := range hm {
		values = append(values, v)
	}
	return values
}
