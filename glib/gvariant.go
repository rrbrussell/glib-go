package glib

import "unsafe"

// GVariant - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:34
type GVariant = rune

// GVariantType - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvarianttype.h:41
type GVariantType = int32

// GVariantClass - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:36
type GVariantClass = uint

// GVariantClassBoolean - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:36
const (
	GVariantClassBoolean    GVariant = 'b'
	GVariantClassByte                = 'y'
	GVariantClassInt16               = 'n'
	GVariantClassUInt16              = 'q'
	GVariantClassInt32               = 'i'
	GVariantClassUInt32              = 'u'
	GVariantClassInt64               = 'x'
	GVariantClassUInt64              = 't'
	GVariantClassHandle              = 'h'
	GVariantClassDouble              = 'd'
	GVariantClassSting               = 's'
	GVariantClassObjectPath          = 'o'
	GVariantClassSignature           = 'g'
	GVariantClassVariant             = 'v'
	GVariantClassMaybe               = 'm'
	GVariantClassArray               = 'a'
	GVariantClassTuple               = '('
	GVariantClassDictEntry           = '{'
)

// GVariantIter - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:268
type GVariantIter struct {
	x [16]uint
}

// GVariantBuilder - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:297
type GVariantBuilder = _GVariantBuilder

// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:298
// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:300
// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:302
type __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_ struct {
	partial_magic gsize
	type_         []GVariantType
	y             [14]gsize
}
type __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_ struct{ memory unsafe.Pointer }

func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_) copy() __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_ {
	var buffer [152]byte
	for i := range buffer {
		buffer[i] = (*((*[152]byte)(unionVar.memory)))[i]
	}
	var newUnion __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_) s() *__struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_ {
	if unionVar.memory == nil {
		var buffer [152]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_302_5_)(unionVar.memory)
}
func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_) x() *[16]gsize {
	if unionVar.memory == nil {
		var buffer [152]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[16]gsize)(unionVar.memory)
}

type _GVariantBuilder struct {
	u __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_300_3_
}

// G_VARIANT_PARSE_ERROR_FAILED - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:311
const (
	G_VARIANT_PARSE_ERROR_FAILED                       int32 = 0
	G_VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED                = 1
	G_VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE                  = 2
	G_VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED             = 3
	G_VARIANT_PARSE_ERROR_INPUT_NOT_AT_END                   = 4
	G_VARIANT_PARSE_ERROR_INVALID_CHARACTER                  = 5
	G_VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING              = 6
	G_VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH                = 7
	G_VARIANT_PARSE_ERROR_INVALID_SIGNATURE                  = 8
	G_VARIANT_PARSE_ERROR_INVALID_TYPE_STRING                = 9
	G_VARIANT_PARSE_ERROR_NO_COMMON_TYPE                     = 10
	G_VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE                = 11
	G_VARIANT_PARSE_ERROR_NUMBER_TOO_BIG                     = 12
	G_VARIANT_PARSE_ERROR_TYPE_ERROR                         = 13
	G_VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN                   = 14
	G_VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD                    = 15
	G_VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT       = 16
	G_VARIANT_PARSE_ERROR_VALUE_EXPECTED                     = 17
)

// GVariantParseError - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:311
type GVariantParseError = int32

// GVariantDict - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:436
type GVariantDict = _GVariantDict

// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:437
// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:439
// __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_ - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gvariant.h:441
type __struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_ struct {
	asv           []GVariant
	partial_magic gsize
	y             [14]gsize
}
type __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_ struct{ memory unsafe.Pointer }

func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_) copy() __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_ {
	var buffer [152]byte
	for i := range buffer {
		buffer[i] = (*((*[152]byte)(unionVar.memory)))[i]
	}
	var newUnion __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_) s() *__struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_ {
	if unionVar.memory == nil {
		var buffer [152]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__struct_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_441_5_)(unionVar.memory)
}
func (unionVar *__union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_) x() *[16]gsize {
	if unionVar.memory == nil {
		var buffer [152]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[16]gsize)(unionVar.memory)
}

type _GVariantDict struct {
	u __union_at__home_robert_Documents_src_glib_2_62_4_glib_gvariant_h_439_3_
}
