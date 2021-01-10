package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый элемент списка
	Back() *listItem                   // последний элемент списка
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Prev  *listItem   // предыдущий элемент
	Next  *listItem   // следующий элемент
}

type list struct {
	// Place your code here

	length    int
	firstItem *listItem
	lastItem  *listItem
}

func NewList() List {
	return &list{length: 0}
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.firstItem
}

func (l *list) Back() *listItem {
	return l.lastItem
}
func (l *list) PushFront(v interface{}) *listItem {
	newElement := listItem{Value: v}

	// ничег нет - голова и хвост равны
	if l.length == 0 {
		l.firstItem, l.lastItem = &newElement, &newElement
	} else {
		l.firstItem.Prev = &newElement
		newElement.Next = l.firstItem
		l.firstItem = &newElement
	}
	l.length++
	return l.firstItem
}

func (l *list) PushBack(v interface{}) *listItem {
	newElement := listItem{Value: v}

	if l.length == 0 {
		l.firstItem, l.lastItem = &newElement, &newElement
	} else {
		l.lastItem.Next = &newElement
		newElement.Prev = l.lastItem
		l.lastItem = &newElement
	}

	l.length++
	return &newElement
}

func (l *list) Remove(i *listItem) {
	if l.length == 0 || i == nil {
		return
	}

	if l.length == 1 {
		l.firstItem, l.lastItem = nil, nil
		l.length--
		return
	}

	next := i.Next
	prev := i.Prev

	if next != nil {
		next.Prev = prev
	}

	if prev != nil {
		prev.Next = next
	}

	if i == l.lastItem {
		l.lastItem = i.Prev
		l.lastItem.Next = nil
	}

	if i == l.firstItem {
		l.firstItem = i.Next
		l.firstItem.Prev = nil
	}

	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	if l.length <= 1 || i == nil {
		return
	}

	//  fmt.Println(l.firstItem)
	//  fmt.Println(l.lastItem)
	//	fmt.Printf("MoveToFront = %d %v  %v\n", l.length, l.Front(), l.Back())
	l.Remove(i)
	//  fmt.Printf("MoveToFront = %d %v  %v\n", l.length, l.Front(), l.Back())
	//  fmt.Println(l.firstItem)
	//  fmt.Println(l.lastItem)

	prev := l.firstItem
	//	fmt.Printf("MoveToFront prev =    %v\n", prev)
	prev.Prev = i
	//	fmt.Printf("MoveToFront 2 prev =    %v\n", prev)
	i.Next = prev
	l.firstItem = i
	i.Prev = nil
	l.length++
}
