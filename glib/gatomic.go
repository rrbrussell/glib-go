package glib

import "sync/atomic"

// GAtomicIntGet wraps atomic.LoadInt32
func GAtomicIntGet(address *int32) (value int32) {
	value = atomic.LoadInt32(address)
	return value
}

// GAtomicIntSet wraps atomic.StoreInt32
func GAtomicIntSet(address *int32, newValue int32) {
	atomic.StoreInt32(address, newValue)
}

// GAtomicIntInc wraps atomic.AddInt32
func GAtomicIntInc(address *int32) {
	atomic.AddInt32(address, 1)
}

// GAtomicIntCompareAndExchange wraps atomic.CompareAndSwapInt32
func GAtomicIntCompareAndExchange(address *int32, old, new int32) (swapped bool) {
	swapped = atomic.CompareAndSwapInt32(address, old, new)
	return swapped
}

// GAtomicIntAdd wraps atomic.AddInt32
func GAtomicIntAdd(address *int32, value int32) {
	atomic.AddInt32(address, value)
}
