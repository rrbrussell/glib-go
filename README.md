# glib-go
A pure go implementation of glib.

This is designed around glib-2.62.4.

## gversion.{h,c} & gversionmacros.h
Since go is statically compiled new import paths will be created when breaking API changes happen.
As such the contents of these files are redundant.

## gtypes.h
Go already provides stable variants for most of these types or the types are not required with how go is structured.
So this fule considered redundant for now.

## gmacros.h
Provided by the os package
- G_IS_DIR_SEPARATOR: os.IsPathSeparator
- G_OS_WIN32: test os.GOOS for runtime testing for buildtime use Go's Build Constraints
- G_OS_UNIX: see G_OS_WIN32
- G_DIR_SEPARATOR: os.PathSeparator
- G_DIR_SEPARATOR_S: see G_DIR_SEPARATOR
- G_SEARCHPATH_SEPARATOR: os.PathListSeparator
- G_SEARCHPATH_SEPARATOR_S: see G_SEARCHPATH_SEPARATOR

Provided by the math package
- ABS: math.Abs()
- MAX: math.Max()
- MIN: math.Min()

Provided by laguage
- G_N_ELEMENTS: len()
- TRUE: true
- FALSE: false
- NULL: nil

Probably unneeded due to language differences
- G_STRUCT_MEMBER
- G_STRUCT_MEMBER
- G_STRUCT_OFFSET
- G_ALIGNOF
- G_MEM_ALIGN

Depreciated
- G_CONST_RETURN

## glibconfig.h.in
Probably unneeded due to language differences
- GINT_TO_POINTER
- GPOINTER_TO_INT
- GUINT_TO_POINTER
- GPOINTER_TO_UINT
- GSIZE_TO_POINTER
- GPOINTER_TO_SIZE

## gtypes.h
The byte order macros are not implemented currently.
The "encoding.binary" package likely provides everything you need.

implemented
- 
