package Collections

type Dictionary[TKey comparable, TValue any] map[TKey]TValue

// func (dictionary Dictionary[TKey, TValue]) GetKeys() List.List[TKey] {
// 	keys := make([]TKey, 0)
// 	for key := range dictionary {
// 		keys = append(keys, key)
// 	}
// 	return keys
// }

// func (dictionary Dictionary[TKey, TValue]) GetValues() List.List[TValue] {
// 	values := make([]TValue, 0)
// 	for _, value := range dictionary {
// 		values = append(values, value)
// 	}
// 	return values
// }

func (dictionary *Dictionary[TKey, TValue]) Add(key TKey, value TValue) {
	(*dictionary)[key] = value
}

func (dictionary *Dictionary[TKey, TValue]) Remove(key TKey) {
	delete(*dictionary, key)
}

func (dictionary *Dictionary[TKey, TValue]) Item(key TKey) (value TValue) {
	return (*dictionary)[key]
}

func (dictionary Dictionary[TKey, TValue]) GetEnumerator() IEnumerator[KeyValue[TKey, TValue]] {

	keys := make([]TKey, 0)
	for key := range dictionary {
		keys = append(keys, key)
	}

	return &KeyValueEnumerator[TKey, TValue]{
		Index:  0,
		Object: (map[TKey]TValue)(dictionary),
		Keys:   keys,
	}
}
