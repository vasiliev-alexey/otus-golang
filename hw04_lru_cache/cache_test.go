package hw04_lru_cache //nolint:golint,stylecheck
import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {

	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("evict_uno", func(t *testing.T) {
		//  t.Skip("implemented in separates  suit")
		//  на логику выталкивания элементов из-за размера очереди (например: n = 3, добавили 4 элемента - 1й из кэша вытолкнулся);
		cCacheSize := rand.Intn(5)
		c := NewCache(cCacheSize)

		sampleData := []string{}

		for i := 0; i < rand.Intn(100); i++ {
			sampleData = append(sampleData, RandStringRunes(i))
		}

		// прогреваем кеш

		for _, datum := range sampleData {
			c.Set(Key(datum), datum)
		}

		// Считаем промахи и попадания
		cashHit, cacheMiss := 0, 0
		for _, datum := range sampleData {
			_, ok := c.Get(Key(datum))
			if ok {
				cashHit++
			} else {
				cacheMiss++
			}
		}

		require.True(t, cashHit == cCacheSize, "Вытеснение кеша  незафиксированное")
		require.True(t, cacheMiss > 0, "Вытеснение кеша  зафиксированное")

	})

	t.Run("evict_second", func(t *testing.T) {
		//  на логику выталкивания редкоиспользуемых элементов (например: n = 3, добавили 3 элемента,
		// обратились много раз к разным элементам: изменили значение,
		//получили значение и пр. - добавили 4й элемент, из первой тройки вытолкнулся наименее используемый).
		cCacheSize := rand.Intn(10) + 3
		c := NewCache(cCacheSize)

		sampleData := []string{}

		for i := 0; i < rand.Intn(100); i++ {
			sampleData = append(sampleData, RandStringRunes(i))
		}

		// прогреваем кеш
		for j := cCacheSize; 0 < j; j-- {

			for i, datum := range sampleData {
				if i < j {
					c.Set(Key(datum), datum)
				}
			}

			for i, datum := range sampleData {
				if i < j {
					_, _ = c.Get(Key(datum))
				}
			}

		}

		// Считаем промахи и попадания

		cacheItems := make(map[string]interface{})

		for _, datum := range sampleData {
			item, ok := c.Get(Key(datum))
			if ok {
				cacheItems[item.(string)] = nil
			}
		}

		c.Set("datum", "datum")
		for _, datum := range sampleData {
			item, ok := c.Get(Key(datum))
			if ok {
				delete(cacheItems, item.(string))
			}
		}
		require.True(t, len(cacheItems) == 1, "В живых остался только один")

	})

}

func TestCacheClear(t *testing.T) {

	t.Run("method Clear cache test suit ", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("ttt", "empty")
		require.False(t, ok, "not in cache")

		c.Clear()

		rez, ok := c.Get("ttt")

		require.False(t, ok)
		require.Nil(t, rez)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

}

func TestCacheSetters(t *testing.T) {

	t.Run("method Set cache test 1 ", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("ttt", "empty")
		require.False(t, ok, "value set first chance")

		rez, ok := c.Get("ttt")

		require.True(t, ok)
		require.Equal(t, "empty", rez)
	})

	t.Run("method Set cache test 2 ", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("ttt", "empty")

		require.False(t, ok)
		require.True(t, c.Set("ttt", "empty"))

	})

}

func TestCacheGetters(t *testing.T) {

	t.Run("method Set cache test 1 ", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("ttt", "empty")
		require.False(t, ok, "value not set earle")

		rez, ok := c.Get("ttt")

		require.True(t, ok)
		require.Equal(t, "empty", rez)
	})

	t.Run("method Set cache test multy gets  ", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("ttt", "empty")
		require.False(t, ok, "value not set earle")

		_, _ = c.Get("ttt")
		_, _ = c.Get("ttt")
		_, _ = c.Get("ttt")
		rez, ok := c.Get("ttt")

		require.True(t, ok)
		require.Equal(t, "empty", rez)
	})

}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // NeedRemove if task with asterisk completed

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
