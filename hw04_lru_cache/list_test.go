package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
func TestListPushFront(t *testing.T) {

	t.Run("PushFront", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Front().Value, 10)

		require.Equal(t, l.Front(), l.Back(), "если 1 элемент то и хвост и голова равны")

		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)

		l.PushFront(20) // [10]

		require.Equal(t, 2, l.Len())
		require.Equal(t, l.Front().Value, 20)
		require.Equal(t, l.Front().Next.Value, 10)
		require.Nil(t, l.Front().Prev)
		require.NotNil(t, l.Front().Next)
	})
}

func TestListPushBack(t *testing.T) {

	t.Run("PushBack", func(t *testing.T) {
		l := NewList()

		l.PushBack(90) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Front().Value, 90)
		require.Equal(t, l.Back(), l.Front(), "если 1 элемент то и хвост и голова равны")

		l.PushBack(99) // [10]

		require.Equal(t, 2, l.Len())
		require.Equal(t, l.Back().Value, 99)
		require.Equal(t, l.Back().Prev.Value, 90)
		require.Nil(t, l.Back().Next)
		require.NotNil(t, l.Back().Prev)
	})

}

func TestListMoveToFront(t *testing.T) {

	t.Run("MoveFrontEmpty", func(t *testing.T) {
		l := NewList()
		l.MoveToFront(l.Back())
		require.Nil(t, l.Front(), "Для пустого списка дб пусто")
	})

	t.Run("MoveFrontSingle", func(t *testing.T) {
		l := NewList()
		tmp := l.PushBack(10)
		l.MoveToFront(l.Back())
		require.Equal(t, tmp, l.Front())
		require.NotNil(t, tmp)

	})

	t.Run("MoveFrontMany", func(t *testing.T) {
		l := NewList()
		l.PushBack(999)
		l.PushBack(777)
		l.PushBack(10) // 999 777 10
		l.MoveToFront(l.Back())

		require.Equal(t, 10, l.Front().Value) //10

		require.Nil(t, l.Front().Prev)
		require.NotNil(t, l.Front().Next)

		require.Equal(t, 999, l.Front().Next.Value) //10
		require.NotNil(t, l.Front(), &l.Front().Prev)
	})

}

func TestListRemove(t *testing.T) {
	t.Run("RemoveEmpty", func(t *testing.T) {
		l := NewList()
		l.Remove(l.Back())
		require.Nil(t, l.Front(), "Для пустого списка дб пусто")
		require.Nil(t, l.Back(), "Для пустого списка дб пусто")
		require.Equal(t, 0, l.Len(), "Для пустого списка дб 0")
	})

	t.Run("RemoveSingle", func(t *testing.T) {
		l := NewList()
		tmp := l.PushBack(100)
		l.Remove(tmp)
		require.Nil(t, l.Front(), "Для пустого списка дб пусто")
		require.Nil(t, l.Back(), "Для пустого списка дб пусто")
		require.Equal(t, 0, l.Len(), "Для пустого списка дб 0")
	})

	t.Run("RemoveFirst", func(t *testing.T) {
		l := NewList()
		tmp := l.PushBack(100)
		l.PushBack(200)
		l.PushBack(300)
		//100 200 300
		l.Remove(tmp)

		require.Equal(t, 200, l.Front().Value, "ДБ 200")
		require.Equal(t, 300, l.Back().Value, "ДБ 200")
		require.Equal(t, 2, l.Len(), "Для пустого списка дб 0")

		require.Nil(t, l.Front().Prev, "ДБ пусто раз 1")
		require.NotNil(t, l.Front().Next, "ДБ не пусто  ")

	})

	t.Run("RemoveLast", func(t *testing.T) {
		l := NewList()
		l.PushBack(100)
		l.PushBack(200)
		tmp := l.PushBack(300)
		l.Remove(tmp)

		require.Equal(t, 100, l.Front().Value, "ДБ 200")
		require.Equal(t, 200, l.Back().Value, "ДБ 200")
		require.Equal(t, 2, l.Len(), "Для пустого списка дб 0")

		require.Nil(t, l.Back().Next, "ДБ пусто раз 1")
		require.NotNil(t, l.Back().Prev, "ДБ не пусто  ")

	})

	t.Run("RemoveMiddle", func(t *testing.T) {
		l := NewList()
		_ = l.PushBack(100)
		tmp := l.PushBack(200)
		_ = l.PushBack(300)

		l.Remove(tmp)

		require.NotNil(t, l.Back().Prev, "ДБ  не пусто")
		require.Nil(t, l.Back().Next, "ДБ пусто раз 1")

		require.NotNil(t, l.Front().Next, "ДБ  не пусто")
		require.Nil(t, l.Front().Prev, "ДБ пусто раз 1")

	})

}

func TestListFront(t *testing.T) {
	t.Run("Front", func(t *testing.T) {
		l := NewList()
		empty := l.Front()
		require.Nil(t, empty, "Для пустого списка дб пусто")
	})
}
func TestListBack(t *testing.T) {

	t.Run("Back", func(t *testing.T) {
		l := NewList()
		empty := l.Back()
		require.Nil(t, empty, "Для пустого списка дб пусто")
	})
}
func TestListLen(t *testing.T) {

	t.Run("Back", func(t *testing.T) {
		l := NewList()
		require.Equal(t, l.Len(), 0, "Для пустого списка длина 0")
	})
}
