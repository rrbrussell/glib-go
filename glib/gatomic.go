package glib

import "sync/atomic"

// GAtomicIntGet wraps atomic.LoadInt32
func GAtomicIntGet(address *int32) (value int32) {
	value = atomic.LoadInt32(address)
	return value
}

// GAtomicIntSet wraps atomic.StoreInt32
/*
GLIB_AVAILABLE_IN_ALL
void                    g_atomic_int_set                      (volatile gint  *atomic,
															   gint            newval);
*/
func GAtomicIntSet(address *int32, newValue int32) {
	atomic.StoreInt32(address, newValue)
}

// GAtomicIntInc wraps atomic.AddInt32
/*
GLIB_AVAILABLE_IN_ALL
void                    g_atomic_int_inc                      (volatile gint  *atomic);
*/
func GAtomicIntInc(address *int32) {
	atomic.AddInt32(address, 1)
}

// GAtomicIntCompareAndExchange wraps atomic.CompareAndSwapInt32
/*
GLIB_AVAILABLE_IN_ALL
gboolean                g_atomic_int_compare_and_exchange     (volatile gint  *atomic,
                                                               gint            oldval,
															   gint            newval);
*/
func GAtomicIntCompareAndExchange(address *int32, old, new int32) (swapped bool) {
	swapped = atomic.CompareAndSwapInt32(address, old, new)
	return swapped
}

// GAtomicIntAdd wraps atomic.AddInt32
/*
GLIB_AVAILABLE_IN_ALL
gint                    g_atomic_int_add                      (volatile gint  *atomic,
                                                               gint            val);
*/
func GAtomicIntAdd(address *int32, value int32) {
	atomic.AddInt32(address, value)
}
