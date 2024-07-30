package Collections

import "reflect"

// Represents a collection of keys and values.
type Dictionary[TKey comparable, TValue any] map[TKey]TValue

// Gets a collection containing the keys in the Dictionary<TKey,TValue>.
//
// Returns:
//	List[TKey] - The collection containing the keys in the Dictionary<TKey,TValue>.
func (dictionary Dictionary[TKey, TValue]) GetKeys() List[TKey] {
	keys := make([]TKey, 0)
	for key := range dictionary {
		keys = append(keys, key)
	}
	return keys
}

// Gets a collection containing the values in the Dictionary<TKey,TValue>.
//
// Returns:
//	List[TValue] - The collection containing the values in the Dictionary<TKey,TValue>.
func (dictionary Dictionary[TKey, TValue]) GetValues() List[TValue] {
	values := make([]TValue, 0)
	for _, value := range dictionary {
		values = append(values, value)
	}
	return values
}

// Adds the specified key and value to the dictionary.
//
// Parameters:
//	key TKey - The key of the element to add.
//	value TValue - The value of the element to add.
func (dictionary Dictionary[TKey, TValue]) Add(key TKey, value TValue) {
	dictionary[key] = value
}

// Removes all keys and values from the Dictionary<TKey,TValue>.
func (dictionary Dictionary[TKey, TValue]) Clear() {
	for key := range dictionary {
		delete(dictionary, key)
	}
}

// Determines whether the Dictionary<TKey,TValue> contains the specified key.
//
// Parameters:
//	key TKey - The key to locate in the Dictionary<TKey,TValue>.
//
// Returns:
//	bool - true if the Dictionary<TKey,TValue> contains an element with the specified key; otherwise, false.
func (dictionary Dictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	_, ok := dictionary[key]
	return ok
}

// Determines whether the Dictionary<TKey,TValue> contains the specified key.
//
// Parameters:
//	value TValue - The value to locate in the Dictionary<TKey,TValue>.
//
// Returns:
//	bool - true if the Dictionary<TKey,TValue> contains an element with the specified value; otherwise, false.
func (dictionary Dictionary[TKey, TValue]) ContainsValue(value TValue) bool {
	for _, v := range dictionary {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (dictionary Dictionary[TKey, TValue]) GetEnumerator() IEnumerator[KeyValue[TKey, TValue]] {
	return &keyValueEnumerator[TKey, TValue]{
		Index:  0,
		Object: (map[TKey]TValue)(dictionary),
		Keys:   dictionary.GetKeys(),
	}
}

// Removes the value with the specified key from the Dictionary<TKey,TValue>, and returns copied removed element.
//
// Parameters:
//	key TKey - The key of the element to remove.
//
// Returns:
//	value TValue - The removed element.
//	ok bool - true if the element is successfully found and removed; otherwise, false.
func (dictionary Dictionary[TKey, TValue]) Remove(key TKey) (value TValue, ok bool) {
	value, ok = dictionary[key]
	delete(dictionary, key)
	return value, ok
}

// Attempts to add the specified key and value to the dictionary.
//
// Parameters:
//	key TKey - The key of the element to add.
//	value TValue - The value of the element to add.
//
// Returns:
//	bool - true if the key/value pair was added to the dictionary successfully; otherwise, false.
func (dictionary Dictionary[TKey, TValue]) TryAdd(key TKey, value TValue) bool {
	_, ok := dictionary[key]
	if !ok {
		dictionary.Add(key, value)
	}
	return ok
}

func (dictionary Dictionary[TKey, TValue]) TryGetValue(key TKey) (value TValue, ok bool) {
	value, ok = dictionary[key]
	return value, ok
}

func (dictionary Dictionary[TKey, TValue]) Item(key TKey) (value TValue) {
	return dictionary[key]
}
