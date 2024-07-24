package List

import "github.com/golangframework/system/Collections"

func (list *List[T]) InsertRange(index int, collection Collections.IEnumerable[T]) (err error) {
	if index < 0 {
		return Collections.ErrIndexIsLessThanZero
	}
	if index > len(*list) {
		return Collections.ErrCollectionContainsNoElements
	}
	enumerator := collection.GetEnumerator()
	item, ok := enumerator.GetNext()
	if !ok {
		return nil
	}
	temp := append((*list)[:index], item)
	for ok {
		temp.Add(item)
		item, ok = enumerator.GetNext()
	}
	(*list) = append(temp, (*list)[index:]...)
	return nil
}
