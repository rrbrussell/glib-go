package glib

// GLogLevelFlags says how important a log message is
type GLogLevelFlags int

// G_LOG_FLAG_RECURSION - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:52
// Warning (*ast.EnumDecl):  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:52 :add support of continues counter for Go type : &ast.ParenExpr{Lparen:0, X:(*ast.CallExpr)(0xc0009fa500), Rparen:0}
const (
	GLogFlagRecursion GLogLevelFlags = 1 << 0
	GLogFlagFatal                    = 1 << 1
	GLogLevelError                   = 1 << 2
	GLogLevelCritical                = 1 << 3
	GLogLevelWarning                 = 1 << 4
	GLogLevelMessage                 = 1 << 5
	GLogLevelInfo                    = 1 << 6
	GLogLevelDebug                   = 1 << 7
	GLogLevelMask                    = GLogFlagRecursion | GLogFlagFatal
)

// GLogFatalMask defines GLib log levels that are considered fatal by default
const GLogFatalMask = GLogFlagRecursion | GLogLevelError

// GLogFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:72
type GLogFunc = func([]gchar, int32, []gchar, gpointer)

// GLogWriterOutput - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:133
type GLogWriterOutput = int32

// G_LOG_WRITER_HANDLED - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:133
const (
	GLogWriteHandled   GLogWriterOutput = 1
	GLogWriteUnHandled                  = 0
)

// GLogField - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:155
type GLogField = _GLogField

// _GLogField - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:156
type _GLogField struct {
	key    []gchar
	value  gconstpointer
	length gssize
}

// GLogWriterFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:193
type GLogWriterFunc = func(int32, []GLogField, gsize, gpointer) int32

// GPrintFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmessages.h:463
type GPrintFunc = func([]gchar)
