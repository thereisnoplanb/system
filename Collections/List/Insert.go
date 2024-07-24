package List

import "github.com/golangframework/system/Collections"

func (list *List[T]) Insert(index int, item T) (err error) {
	if index < 0 {
		return Collections.ErrIndexIsLessThanZero
	}
	if index > len(*list) {
		return Collections.ErrCollectionContainsNoElements
	}
	temp := append((*list)[:index], item)
	(*list) = append(temp, (*list)[index:]...)
	return nil
}
