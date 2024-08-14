package coll

type Arraylist []interface{}

// constructor creates a new array list
func NewArrayList() Arraylist {
	return *(new(Arraylist))
}

// method to add a new element to the array list
func (al *Arraylist) Add(value interface{}) {
	*al = append(*al, value)
}

// method to insert a specific element at a specific position index in the array list
func (al *Arraylist) Replace(index int, value interface{}) {
	if index < len(*al) {
		(*al)[index] = value
	}
}

// method to remove the element from the array list
func (al *Arraylist) Remove(index int) {
	if index < len(*al) {
		*al = append((*al)[:index], (*al)[index+1:]...)
	}
}

// method to remove the first occurrence of the specified element from the array list
func (al *Arraylist) RemoveValue(value interface{}) {
	index := -1
	for idx := range *al {
		if (*al)[idx] == value {
			index = idx
			break
		}
	}
	if index != -1 {
		*al = append((*al)[:index], (*al)[index+1:]...)
	}
}

// methos to remove from the array list all of the elements whose index is between fromIndex (inclusive) and toIndex (exclusive)
func (al *Arraylist) RemoveRange(fromIndex int, toIndex int) {
	if fromIndex > toIndex {
		return
	}
	if toIndex < len(*al) {
		*al = append((*al)[:fromIndex], (*al)[toIndex:]...)
	}
}

// method to get the element from the array list
func (al *Arraylist) Get(index int) interface{} {
	if index < len(*al) {
		return (*al)[index]
	}
	return nil
}

// method to count the total elements of the array list
func (al *Arraylist) Length() int {
	return len(*al)
}

// method to check if a value belongs to the array list
func (al *Arraylist) Contains(value interface{}) bool {
	for idx := range *al {
		if (*al)[idx] == value {

			return true
		}
	}
	return false
}

// method to clear the array list
func (al *Arraylist) Clear() {
	*al = nil
}
