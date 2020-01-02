package glib

type gSourceFlags int

// Flags for GSources
const (
	GSourceReady      gSourceFlags = 1<<iota + 1
	GSourceCanRecurse              = 1<<iota + 1
	GSourceBlocked                 = 1<<iota + 1
)

// A GSourceList is the list of currently active GSources
type GSourceList struct {
	head, tail *GSource
	priority   int
}

// A GSource is an event source
type GSource struct {
	//gpointer callback_data
	//GSourceCallbackFuncs *callback_funcs

	//const GSourceFuncs *source_funcs
	//guint ref_count

	//GMainContext *context

	priority        int
	flags, sourceID uint

	//GSList *poll_fds

	prev, next *GSource

	name string

	//GSourcePrivate *priv
}
