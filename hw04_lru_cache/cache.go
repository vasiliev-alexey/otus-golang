package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	storedItem, ok := l.items[key]

	item := cacheItem{key, value}

	if ok {
		l.queue.MoveToFront(storedItem)
		storedItem.Value = item
	} else {
		l.queue.PushFront(item)
		l.items[key] = l.queue.Front()

		if l.queue.Len() > l.capacity {
			delete(l.items, l.queue.Back().Value.(cacheItem).key)
			l.queue.Remove(l.queue.Back())
		}
	}
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	storedItem, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(storedItem)
		return l.queue.Front().Value.(cacheItem).val, ok
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.init()
}

func (l *lruCache) init() *lruCache {
	l.queue = NewList()
	l.items = map[Key]*listItem{}
	return l
}

func NewCache(capacity int) Cache {
	lru := lruCache{capacity: capacity}
	return lru.init()
}

type cacheItem struct {
	key Key
	val interface{}
}
