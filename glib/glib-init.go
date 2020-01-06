package glib

// #include </usr/include/strings.h>
// #include </usr/include/string.h>
import "C"

import (
	"bytes"
	"math"
	"unicode"
	"unsafe"

	"github.com/Konstantin8105/c4go/noarch"
)

func init() {
	gMessagesPrefixedInit()
	//g_quark_init()
}

var gLogMsgPrefix GLogLevelFlags = GLogLevelError | GLogLevelWarning | GLogLevelCritical | GLogLevelDebug
var gLogAlwaysFatal GLogLevelFlags = GLogFatalMask

// gint8 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:40
type gint8 = int8

// guint8 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:41
type guint8 = uint8

// gint16 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:43
type gint16 = int16

// guint16 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:44
type guint16 = uint16

// gint32 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:51
type gint32 = int32

// guint32 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:52
type guint32 = uint32

// gint64 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:61
type gint64 = int32

// guint64 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:62
type guint64 = uint32

// gssize - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:77
type gssize = int32

// gsize - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:78
type gsize = uint32

// goffset - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:88
type goffset = gint64

// gintptr - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:104
type gintptr = int32

// guintptr - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:105
type guintptr = uint32

// GPid - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/config.h:198
type GPid = int32

// gchar - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:46
type gchar = byte

// gshort - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:47
type gshort = int16

// glong - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:48
type glong = int32

// gint - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:49
type gint = int32

// gboolean - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:50
type gboolean = gint

// guchar - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:52
type guchar = uint8

// gushort - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:53
type gushort = uint16

// gulong - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:54
type gulong = uint32

// guint - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:55
type guint = uint32

// gfloat - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:57
type gfloat = float32

// gdouble - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:58
type gdouble = float64

// gpointer - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:103
type gpointer = interface{}

// gconstpointer - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:104
type gconstpointer = interface{}

// GCompareFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:106
type GCompareFunc = func(gconstpointer, gconstpointer) gint

// GCompareDataFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:108
type GCompareDataFunc = func(gconstpointer, gconstpointer, gpointer) gint

// GEqualFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:111
type GEqualFunc = func(gconstpointer, gconstpointer) gboolean

// GDestroyNotify - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:113
type GDestroyNotify = func(gpointer)

// GFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:114
type GFunc = func(gpointer, gpointer)

// GHashFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:116
type GHashFunc = func(gconstpointer) guint

// GHFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:117
type GHFunc = func(gpointer, gpointer, gpointer)

// GCopyFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:133
type GCopyFunc = func(gconstpointer, gpointer) gpointer

// GFreeFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:143
type GFreeFunc = func(gpointer)

// GTranslateFunc - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:157
type GTranslateFunc = func([]gchar, gpointer) []gchar

// grefcount - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:557
type grefcount = gint

// gatomicrefcount - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gtypes.h:558
type gatomicrefcount = gint

// GQuark - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gquark.h:36
type GQuark = guint32

// GError - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gerror.h:41
type GError struct {
	domain  GQuark
	code    gint
	message []gchar
}

/*
// gunichar - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:59
type gunichar = guint32

// gunichar2 - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:77
type gunichar2 = guint16

// GUnicodeType - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:116
type GUnicodeType = int32

// G_UNICODE_CONTROL - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:116
const (
	GUnicodeControl            GUnicodeType = 0
	GUnicodeFormat                          = 1
	GUnicodeUnassigned                      = 2
	GUnicodePrivateUse                      = 3
	GUnicodeSurrogate                       = 4
	GUnicodeLowercaseLetter                 = 5
	GUnicodeModifierLetter                  = 6
	GUnicodeOtherLetter                     = 7
	GUnicodeTittleCaseLetter                = 8
	GUnicodeUpperCaseLetter                 = 9
	GUnicodeSpacingMark                     = 10
	GUnicodeEnclosingMark                   = 11
	GUnicodeNonSpacingMark                  = 12
	GUnicodeDecimalNumber                   = 13
	GUnicodeLetterNumber                    = 14
	GUnicodeOtherNumber                     = 15
	GUnicodeConnectPunctuation              = 16
	GUnicodeDashPunctuation                 = 17
	GUnicodeClosePunctuation                = 18
	GUnicodeFinalPunctuation                = 19
	GUnicodeInitialPunctuation              = 20
	GUnicodeOtherPunctuation                = 21
	GUnicodeOpenPunctuation                 = 22
	GUnicodeCurrencySymbol                  = 23
	GUnicodeModifierSymbol                  = 24
	GUnicodeMathSymbol                      = 25
	GUnicodeOtherSymbol                     = 26
	GUnicodeLineSeparator                   = 27
	GUnicodeParagraphSeparator              = 28
	GUnicodeSpaceSeparator                  = 29
)

// GUnicodeBreakType - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:212
type GUnicodeBreakType = int32

// G_UNICODE_BREAK_MANDATORY - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:212
const (
	G_UNICODE_BREAK_MANDATORY                    int32 = 0
	G_UNICODE_BREAK_CARRIAGE_RETURN                    = 1
	G_UNICODE_BREAK_LINE_FEED                          = 2
	G_UNICODE_BREAK_COMBINING_MARK                     = 3
	G_UNICODE_BREAK_SURROGATE                          = 4
	G_UNICODE_BREAK_ZERO_WIDTH_SPACE                   = 5
	G_UNICODE_BREAK_INSEPARABLE                        = 6
	G_UNICODE_BREAK_NON_BREAKING_GLUE                  = 7
	G_UNICODE_BREAK_CONTINGENT                         = 8
	G_UNICODE_BREAK_SPACE                              = 9
	G_UNICODE_BREAK_AFTER                              = 10
	G_UNICODE_BREAK_BEFORE                             = 11
	G_UNICODE_BREAK_BEFORE_AND_AFTER                   = 12
	G_UNICODE_BREAK_HYPHEN                             = 13
	G_UNICODE_BREAK_NON_STARTER                        = 14
	G_UNICODE_BREAK_OPEN_PUNCTUATION                   = 15
	G_UNICODE_BREAK_CLOSE_PUNCTUATION                  = 16
	G_UNICODE_BREAK_QUOTATION                          = 17
	G_UNICODE_BREAK_EXCLAMATION                        = 18
	G_UNICODE_BREAK_IDEOGRAPHIC                        = 19
	G_UNICODE_BREAK_NUMERIC                            = 20
	G_UNICODE_BREAK_INFIX_SEPARATOR                    = 21
	G_UNICODE_BREAK_SYMBOL                             = 22
	G_UNICODE_BREAK_ALPHABETIC                         = 23
	G_UNICODE_BREAK_PREFIX                             = 24
	G_UNICODE_BREAK_POSTFIX                            = 25
	G_UNICODE_BREAK_COMPLEX_CONTEXT                    = 26
	G_UNICODE_BREAK_AMBIGUOUS                          = 27
	G_UNICODE_BREAK_UNKNOWN                            = 28
	G_UNICODE_BREAK_NEXT_LINE                          = 29
	G_UNICODE_BREAK_WORD_JOINER                        = 30
	G_UNICODE_BREAK_HANGUL_L_JAMO                      = 31
	G_UNICODE_BREAK_HANGUL_V_JAMO                      = 32
	G_UNICODE_BREAK_HANGUL_T_JAMO                      = 33
	G_UNICODE_BREAK_HANGUL_LV_SYLLABLE                 = 34
	G_UNICODE_BREAK_HANGUL_LVT_SYLLABLE                = 35
	G_UNICODE_BREAK_CLOSE_PARANTHESIS                  = 36
	G_UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER       = 37
	G_UNICODE_BREAK_HEBREW_LETTER                      = 38
	G_UNICODE_BREAK_REGIONAL_INDICATOR                 = 39
	G_UNICODE_BREAK_EMOJI_BASE                         = 40
	G_UNICODE_BREAK_EMOJI_MODIFIER                     = 41
	G_UNICODE_BREAK_ZERO_WIDTH_JOINER                  = 42
)

// GUnicodeScript - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:437
type GUnicodeScript = int32

// G_UNICODE_SCRIPT_INVALID_CODE - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:437
const (
	G_UNICODE_SCRIPT_INVALID_CODE           int32 = -1
	G_UNICODE_SCRIPT_COMMON                       = 0
	G_UNICODE_SCRIPT_INHERITED                    = 1
	G_UNICODE_SCRIPT_ARABIC                       = 2
	G_UNICODE_SCRIPT_ARMENIAN                     = 3
	G_UNICODE_SCRIPT_BENGALI                      = 4
	G_UNICODE_SCRIPT_BOPOMOFO                     = 5
	G_UNICODE_SCRIPT_CHEROKEE                     = 6
	G_UNICODE_SCRIPT_COPTIC                       = 7
	G_UNICODE_SCRIPT_CYRILLIC                     = 8
	G_UNICODE_SCRIPT_DESERET                      = 9
	G_UNICODE_SCRIPT_DEVANAGARI                   = 10
	G_UNICODE_SCRIPT_ETHIOPIC                     = 11
	G_UNICODE_SCRIPT_GEORGIAN                     = 12
	G_UNICODE_SCRIPT_GOTHIC                       = 13
	G_UNICODE_SCRIPT_GREEK                        = 14
	G_UNICODE_SCRIPT_GUJARATI                     = 15
	G_UNICODE_SCRIPT_GURMUKHI                     = 16
	G_UNICODE_SCRIPT_HAN                          = 17
	G_UNICODE_SCRIPT_HANGUL                       = 18
	G_UNICODE_SCRIPT_HEBREW                       = 19
	G_UNICODE_SCRIPT_HIRAGANA                     = 20
	G_UNICODE_SCRIPT_KANNADA                      = 21
	G_UNICODE_SCRIPT_KATAKANA                     = 22
	G_UNICODE_SCRIPT_KHMER                        = 23
	G_UNICODE_SCRIPT_LAO                          = 24
	G_UNICODE_SCRIPT_LATIN                        = 25
	G_UNICODE_SCRIPT_MALAYALAM                    = 26
	G_UNICODE_SCRIPT_MONGOLIAN                    = 27
	G_UNICODE_SCRIPT_MYANMAR                      = 28
	G_UNICODE_SCRIPT_OGHAM                        = 29
	G_UNICODE_SCRIPT_OLD_ITALIC                   = 30
	G_UNICODE_SCRIPT_ORIYA                        = 31
	G_UNICODE_SCRIPT_RUNIC                        = 32
	G_UNICODE_SCRIPT_SINHALA                      = 33
	G_UNICODE_SCRIPT_SYRIAC                       = 34
	G_UNICODE_SCRIPT_TAMIL                        = 35
	G_UNICODE_SCRIPT_TELUGU                       = 36
	G_UNICODE_SCRIPT_THAANA                       = 37
	G_UNICODE_SCRIPT_THAI                         = 38
	G_UNICODE_SCRIPT_TIBETAN                      = 39
	G_UNICODE_SCRIPT_CANADIAN_ABORIGINAL          = 40
	G_UNICODE_SCRIPT_YI                           = 41
	G_UNICODE_SCRIPT_TAGALOG                      = 42
	G_UNICODE_SCRIPT_HANUNOO                      = 43
	G_UNICODE_SCRIPT_BUHID                        = 44
	G_UNICODE_SCRIPT_TAGBANWA                     = 45
	G_UNICODE_SCRIPT_BRAILLE                      = 46
	G_UNICODE_SCRIPT_CYPRIOT                      = 47
	G_UNICODE_SCRIPT_LIMBU                        = 48
	G_UNICODE_SCRIPT_OSMANYA                      = 49
	G_UNICODE_SCRIPT_SHAVIAN                      = 50
	G_UNICODE_SCRIPT_LINEAR_B                     = 51
	G_UNICODE_SCRIPT_TAI_LE                       = 52
	G_UNICODE_SCRIPT_UGARITIC                     = 53
	G_UNICODE_SCRIPT_NEW_TAI_LUE                  = 54
	G_UNICODE_SCRIPT_BUGINESE                     = 55
	G_UNICODE_SCRIPT_GLAGOLITIC                   = 56
	G_UNICODE_SCRIPT_TIFINAGH                     = 57
	G_UNICODE_SCRIPT_SYLOTI_NAGRI                 = 58
	G_UNICODE_SCRIPT_OLD_PERSIAN                  = 59
	G_UNICODE_SCRIPT_KHAROSHTHI                   = 60
	G_UNICODE_SCRIPT_UNKNOWN                      = 61
	G_UNICODE_SCRIPT_BALINESE                     = 62
	G_UNICODE_SCRIPT_CUNEIFORM                    = 63
	G_UNICODE_SCRIPT_PHOENICIAN                   = 64
	G_UNICODE_SCRIPT_PHAGS_PA                     = 65
	G_UNICODE_SCRIPT_NKO                          = 66
	G_UNICODE_SCRIPT_KAYAH_LI                     = 67
	G_UNICODE_SCRIPT_LEPCHA                       = 68
	G_UNICODE_SCRIPT_REJANG                       = 69
	G_UNICODE_SCRIPT_SUNDANESE                    = 70
	G_UNICODE_SCRIPT_SAURASHTRA                   = 71
	G_UNICODE_SCRIPT_CHAM                         = 72
	G_UNICODE_SCRIPT_OL_CHIKI                     = 73
	G_UNICODE_SCRIPT_VAI                          = 74
	G_UNICODE_SCRIPT_CARIAN                       = 75
	G_UNICODE_SCRIPT_LYCIAN                       = 76
	G_UNICODE_SCRIPT_LYDIAN                       = 77
	G_UNICODE_SCRIPT_AVESTAN                      = 78
	G_UNICODE_SCRIPT_BAMUM                        = 79
	G_UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS         = 80
	G_UNICODE_SCRIPT_IMPERIAL_ARAMAIC             = 81
	G_UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI        = 82
	G_UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN       = 83
	G_UNICODE_SCRIPT_JAVANESE                     = 84
	G_UNICODE_SCRIPT_KAITHI                       = 85
	G_UNICODE_SCRIPT_LISU                         = 86
	G_UNICODE_SCRIPT_MEETEI_MAYEK                 = 87
	G_UNICODE_SCRIPT_OLD_SOUTH_ARABIAN            = 88
	G_UNICODE_SCRIPT_OLD_TURKIC                   = 89
	G_UNICODE_SCRIPT_SAMARITAN                    = 90
	G_UNICODE_SCRIPT_TAI_THAM                     = 91
	G_UNICODE_SCRIPT_TAI_VIET                     = 92
	G_UNICODE_SCRIPT_BATAK                        = 93
	G_UNICODE_SCRIPT_BRAHMI                       = 94
	G_UNICODE_SCRIPT_MANDAIC                      = 95
	G_UNICODE_SCRIPT_CHAKMA                       = 96
	G_UNICODE_SCRIPT_MEROITIC_CURSIVE             = 97
	G_UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS         = 98
	G_UNICODE_SCRIPT_MIAO                         = 99
	G_UNICODE_SCRIPT_SHARADA                      = 100
	G_UNICODE_SCRIPT_SORA_SOMPENG                 = 101
	G_UNICODE_SCRIPT_TAKRI                        = 102
	G_UNICODE_SCRIPT_BASSA_VAH                    = 103
	G_UNICODE_SCRIPT_CAUCASIAN_ALBANIAN           = 104
	G_UNICODE_SCRIPT_DUPLOYAN                     = 105
	G_UNICODE_SCRIPT_ELBASAN                      = 106
	G_UNICODE_SCRIPT_GRANTHA                      = 107
	G_UNICODE_SCRIPT_KHOJKI                       = 108
	G_UNICODE_SCRIPT_KHUDAWADI                    = 109
	G_UNICODE_SCRIPT_LINEAR_A                     = 110
	G_UNICODE_SCRIPT_MAHAJANI                     = 111
	G_UNICODE_SCRIPT_MANICHAEAN                   = 112
	G_UNICODE_SCRIPT_MENDE_KIKAKUI                = 113
	G_UNICODE_SCRIPT_MODI                         = 114
	G_UNICODE_SCRIPT_MRO                          = 115
	G_UNICODE_SCRIPT_NABATAEAN                    = 116
	G_UNICODE_SCRIPT_OLD_NORTH_ARABIAN            = 117
	G_UNICODE_SCRIPT_OLD_PERMIC                   = 118
	G_UNICODE_SCRIPT_PAHAWH_HMONG                 = 119
	G_UNICODE_SCRIPT_PALMYRENE                    = 120
	G_UNICODE_SCRIPT_PAU_CIN_HAU                  = 121
	G_UNICODE_SCRIPT_PSALTER_PAHLAVI              = 122
	G_UNICODE_SCRIPT_SIDDHAM                      = 123
	G_UNICODE_SCRIPT_TIRHUTA                      = 124
	G_UNICODE_SCRIPT_WARANG_CITI                  = 125
	G_UNICODE_SCRIPT_AHOM                         = 126
	G_UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS        = 127
	G_UNICODE_SCRIPT_HATRAN                       = 128
	G_UNICODE_SCRIPT_MULTANI                      = 129
	G_UNICODE_SCRIPT_OLD_HUNGARIAN                = 130
	G_UNICODE_SCRIPT_SIGNWRITING                  = 131
	G_UNICODE_SCRIPT_ADLAM                        = 132
	G_UNICODE_SCRIPT_BHAIKSUKI                    = 133
	G_UNICODE_SCRIPT_MARCHEN                      = 134
	G_UNICODE_SCRIPT_NEWA                         = 135
	G_UNICODE_SCRIPT_OSAGE                        = 136
	G_UNICODE_SCRIPT_TANGUT                       = 137
	G_UNICODE_SCRIPT_MASARAM_GONDI                = 138
	G_UNICODE_SCRIPT_NUSHU                        = 139
	G_UNICODE_SCRIPT_SOYOMBO                      = 140
	G_UNICODE_SCRIPT_ZANABAZAR_SQUARE             = 141
	G_UNICODE_SCRIPT_DOGRA                        = 142
	G_UNICODE_SCRIPT_GUNJALA_GONDI                = 143
	G_UNICODE_SCRIPT_HANIFI_ROHINGYA              = 144
	G_UNICODE_SCRIPT_MAKASAR                      = 145
	G_UNICODE_SCRIPT_MEDEFAIDRIN                  = 146
	G_UNICODE_SCRIPT_OLD_SOGDIAN                  = 147
	G_UNICODE_SCRIPT_SOGDIAN                      = 148
	G_UNICODE_SCRIPT_ELYMAIC                      = 149
	G_UNICODE_SCRIPT_NANDINAGARI                  = 150
	G_UNICODE_SCRIPT_NYIAKENG_PUACHUE_HMONG       = 151
	G_UNICODE_SCRIPT_WANCHO                       = 152
)

// G_NORMALIZE_DEFAULT - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:897
const (
	G_NORMALIZE_DEFAULT         int32 = 0
	G_NORMALIZE_NFD                   = 1
	G_NORMALIZE_DEFAULT_COMPOSE       = 2
	G_NORMALIZE_NFC                   = 3
	G_NORMALIZE_ALL                   = 4
	G_NORMALIZE_NFKD                  = 5
	G_NORMALIZE_ALL_COMPOSE           = 6
	G_NORMALIZE_NFKC                  = 7
)

// GNormalizeMode - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gunicode.h:897
type GNormalizeMode = int32*/

// GBytes - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/garray.h:36
type GBytes = int32

// GArray - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/garray.h:37
type GArray struct {
	data []gchar
	len  guint
}

// GByteArray - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/garray.h:38
type GByteArray struct {
	data []guint8
	len  guint
}

// GPtrArray - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/garray.h:39
type GPtrArray struct {
	pdata []gpointer
	len   guint
}

// GMemVTable - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmem.h:51
type GMemVTable = _GMemVTable

// _GMemVTable - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gmem.h:365
type _GMemVTable struct {
	malloc     func(gsize) gpointer
	realloc    func(gpointer, gsize) gpointer
	free       func(gpointer)
	calloc     func(gsize, gsize) gpointer
	tryMalloc  func(gsize) gpointer
	tryRealloc func(gpointer, gsize) gpointer
}

/*	We assume that "small" enums (those where all values fit in INT32_MIN
	to INT32_MAX) are exactly int-sized. In particular, we assume that if
	an enum has no members that exceed the range of char/short, the
	compiler will make it int-sized anyway, so adding a member later that
	*does* exceed the range of char/short is not an ABI break.
*/

// TestChar Not sure what this does.
type TestChar = int

// TestChar0 Not sure what this does.
const TestChar0 TestChar = 0

// TestShort Not sure what this does.
type TestShort = int

// TestShort0 Not sure what this does.
const (
	TestShort0   TestShort = 0
	TestShort256           = 256
)

// TestInt Not sure what this does.
type TestInt = int32

// TestInt32Min Not sure what this does.
const (
	TestInt32Min TestInt = math.MinInt32
	TestInt32Max         = math.MaxInt32
)

// GLogMsgPrefix - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:78
const GLogMsgPrefix GLogLevelFlags = GLogLevelError | GLogLevelWarning | GLogLevelCritical | GLogLevelDebug

// GLogAlwaysFatal - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:80
const GLogAlwaysFatal GLogLevelFlags = GLogFlagRecursion | GLogLevelError

// debugKeyMatches - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:82
func debugKeyMatches(key []gchar, token []gchar, length guint) gboolean {
	for uint32((length)) != 0 {
		var k byte = byte(func() int32 {
			if int32(byte((key[0]))) == int32('_') {
				return int32('-')
			}
			return tolower(int32(byte((key[0]))))
		}())
		var t byte = byte(func() int32 {
			if int32(byte((token[0]))) == int32('_') {
				return int32('-')
			}
			return tolower(int32(byte((token[0]))))
		}())
		if int32(k) != int32(t) {
			return gboolean((gint((0))))
		}
		length--
		func() []gchar {
			tempVarUnary := key
			defer func() {
				key = key[0+1:]
			}()
			return tempVarUnary
		}()
		func() []gchar {
			tempVarUnary := token
			defer func() {
				token = token[0+1:]
			}()
			return tempVarUnary
		}()
	}
	return gboolean((gint((noarch.BoolToInt(int32(byte((key[0]))) == int32('\x00'))))))
}

// gParseDebugString - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:130

// The GVariant documentation indirectly says that int is at least 32 bits
// * (by saying that b, y, n, q, i, u, h are promoted to int). On any
// * reasonable platform, int is in fact *exactly* 32 bits long, because
// * otherwise, {signed char, short, int} wouldn't be sufficient to provide
// * {int8_t, int16_t, int32_t}.
//
// * g_parse_debug_string:
// * @string: (nullable): a list of debug options separated by colons, spaces, or
// * commas, or %NULL.
// * @keys: (array length=nkeys): pointer to an array of #GDebugKey which associate
// *     strings with bit flags.
// * @nkeys: the number of #GDebugKeys in the array.
// *
// * Parses a string containing debugging options
// * into a %guint containing bit flags. This is used
// * within GDK and GTK+ to parse the debug options passed on the
// * command line or through environment variables.
// *
// * If @string is equal to "all", all flags are set. Any flags
// * specified along with "all" in @string are inverted; thus,
// * "all,foo,bar" or "foo,bar,all" sets all flags except those
// * corresponding to "foo" and "bar".
// *
// * If @string is equal to "help", all the available keys in @keys
// * are printed out to standard error.
// *
// * Returns: the combined set of bit flags.
//

func gParseDebugString(str []gchar, keys []GDebugKey, nkeys guint) guint {
	var i guint
	var result guint
	if len(str) == 0 {

		return 0
	}
	if bytes.Compare(str, []byte("help\x00")) < 0 {
		// this function is used during the initialisation of gmessages, gmem
		//   * and gslice, so it may not do anything that causes memory to be
		//   * allocated or risks messages being emitted.
		//   *
		//   * this means, more or less, that this code may not call anything
		//   * inside GLib.
		//
		// using stdio directly for the reason stated above
		noarch.Fprintf(noarch.Stderr, []byte("Supported debug values:\x00"))
		for i = 0; i < nkeys; i++ {
			noarch.Fprintf(noarch.Stderr, []byte(" %s\x00"), keys[i].key)
		}
		noarch.Fprintf(noarch.Stderr, []byte(" all help\n\x00"))
	} else {
		var p []gchar = str
		var q []gchar
		var invert gboolean = gboolean((gint((0))))
		for byte((p[0])) != 0 {
			q = strpbrk(p, []byte(":;, \t\x00"))
			if q == nil {
				q = p[0+noarch.Strlen(p):]
			}
			if int32((gint((debugKeyMatches([]byte("all\x00"), p, guint((uint32(int32((int64(uintptr(unsafe.Pointer(&q[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&p[0])))/int64(1))))))))))) != 0 {
				invert = gboolean((gint((noarch.BoolToInt(noarch.Not(0))))))
			} else {
				for i = 0; i < nkeys; i++ {
					if int32((gint((debugKeyMatches(keys[i].key, p, guint((uint32(int32((int64(uintptr(unsafe.Pointer(&q[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&p[0])))/int64(1))))))))))) != 0 {
						result |= guint(keys[i].value)
					}
				}
			}
			p = q
			if byte((p[0])) != 0 {
				func() []gchar {
					tempVarUnary := p
					defer func() {
						p = p[0+1:]
					}()
					return tempVarUnary
				}()
			}
		}
		if int32((gint((invert)))) != 0 {
			var allFlags guint
			for i = 0; i < nkeys; i++ {
				allFlags |= guint(keys[i].value)
			}
			result = allFlags & ^guint(result)
		}
	}
	return result
}

// gParseDebugEnvvar - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:199
func gParseDebugEnvvar(envvar []gchar, keys []GDebugKey, nKeys gint, defaultValue guint) guint {
	var value []gchar
	value = noarch.Getenv(envvar)
	if len(value) == 0 {
		return defaultValue
	}
	return gParseDebugString(value, keys, guint((uint32(int32((nKeys))))))
}

// gMessagesPrefixedInit - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/glib-init.c:225
func gMessagesPrefixedInit() {
	var keys []GDebugKey = []GDebugKey{{[]byte("error\x00"), guint(GLogLevelError)}, {[]byte("critical\x00"), guint(GLogLevelCritical)}, {[]byte("warning\x00"), guint(GLogLevelWarning)}, {[]byte("message\x00"), guint(GLogLevelMessage)}, {[]byte("info\x00"), guint(GLogLevelInfo)}, {[]byte("debug\x00"), guint(GLogLevelDebug)}}
	gLogMsgPrefix = GLogLevelFlags(int32(uint32((gParseDebugEnvvar([]byte("G_MESSAGES_PREFIXED\x00"), keys, gint((int32(144 / 24))), guint((uint32(gLogMsgPrefix))))))))
}

// tolower from ctype.h
// c function : int tolower(int)
// dep pkg    : unicode
// dep func   :
func tolower(_c int32) int32 {
	return int32(unicode.ToLower(rune(_c)))
}

// strpbrk - add c-binding for implemention function
func strpbrk(arg0 []byte, arg1 []byte) []byte {
	if arg0 == nil {
		return []byte{}
	}
	if arg1 == nil {
		return []byte{}
	}
	return []byte(C.GoString(C.strpbrk((*C.char)(unsafe.Pointer(&arg0[0])), (*C.char)(unsafe.Pointer(&arg1[0])))))
}
