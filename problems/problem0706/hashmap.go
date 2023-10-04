package problem0706

/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap.
	If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/

type MyHashMap struct {
	Length  int
	Buckets [][]Pair
}

type Pair [2]int

func Constructor() MyHashMap {
	return MyHashMap{
		Length:  256,
		Buckets: make([][]Pair, 256),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	var inPair = Pair{key, value}
	key = key % m.Length

	// If bucket for key does not exist, create it
	if m.Buckets[key] == nil {
		m.Buckets[key] = []Pair{inPair}
		return
	}

	// Look for key in bucket
	for i := range m.Buckets[key] {
		// Update it if exists
		if m.Buckets[key][i][0] == inPair[0] {
			m.Buckets[key][i][1] = inPair[1]
			return
		}
	}

	// If key does not exists in bucket, add it
	m.Buckets[key] = append(m.Buckets[key], inPair)
}

func (m *MyHashMap) Get(key int) int {
	hashKey := key % m.Length

	// If bucket for key does not exist
	if m.Buckets[hashKey] == nil {
		return -1
	}

	// Look for key in bucket
	for i := range m.Buckets[hashKey] {
		if m.Buckets[hashKey][i][0] == key {
			return m.Buckets[hashKey][i][1]
		}
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	hashKey := key % m.Length

	// If bucket for key does not exist
	if m.Buckets[hashKey] == nil {
		return
	}

	// Look for key in bucket
	for i := range m.Buckets[hashKey] {
		if m.Buckets[hashKey][i][0] == key {
			m.Buckets[hashKey][i][0] = -1
			return
		}
	}
}
