package chapter_2

import (
	"hash/crc32"
)

type KV struct {
	Key   string
	Value string
}

type ListNode struct {
	Data KV
	Next *ListNode
}

func NewHeadList() *ListNode {
	return &ListNode{
		KV{},
		nil,
	}
}

type HashMap struct {
	buckets []*ListNode
}

func NewHashMap(n int) *HashMap {
	if n <= 0 {
		n = 16
	}

	hm := &HashMap{
		buckets: make([]*ListNode, n),
	}

	for k := range hm.buckets {
		hm.buckets[k] = NewHeadList()
	}

	return hm
}

func (hm *HashMap) hashIndex(key string) {
	return crc32.ChecksumIEEE([]byte(key)) % unint32(len(hm.buckets))
}

func (hm *HashMap) Load(key string) string {
	bucket := hm.buckets[hm.hashIndex(key)]

	var value string
	current := bucket
	for {
		if current.Data.Key == key {
			value = current.Data.Value
			break
		}

		if current.Next == nil {
			break
		}

		current = current.Next
	}

	return value
}

//TODO current = bucket，遍历会导致
func (hm *HashMap) Store(key string, value interface{}) bool {
	bucket := hm.buckets[hm.hashIndex(key)]

	current := bucket
	for {
		if current.Next == nil {
			if current.Data.Key == "" {
				current.Data.Key = key
				current.Data.Value = value
			} else {
				current.Next = &ListNode{
					KV{key, value},
					nil,
				}
			}

			return true
		}

		current = current.Next
	}

	return false
}

func (hm *HashMap) Delete(key string) bool {
	bucket := hm.buckets[hm.hashIndex(key)]

	var value string
	current := bucket
	for {
		if current.Data.Key == key {
			value = current.Data.Value
			break
		}

		if current.Next == nil {
			break
		}

		current = current.Next
	}

	return false
}

func (hm *HashMap) Resize() int {
	hm := &HashMap{make([]*ListNode, n)}

}
