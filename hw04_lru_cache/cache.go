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

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity, NewList(), make(map[Key]*listItem, capacity)}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	cache := cacheItem{Key: key, Value: value}
	if item, ok := l.items[key]; ok {
		item.Value = cache
		l.queue.MoveToFront(item)
		return true
	}

	item := l.queue.PushFront(cache)
	l.items[key] = item

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := l.items[key]
	if ok == false {
		return nil, ok
	}

	l.queue.MoveToFront(item)
	cache := item.Value.(cacheItem)

	return cache.Value, ok
}

func (l *lruCache) Clear() {

}
