# glib-go: A pure go implementation of glib.

This is designed around glib-2.62.4. I am not planning this library as a library that would allow
an automated conversion of glib client C code to Golang.
All of the implemented names have been changed to match Golang's requirements and recommendations.
Several of the functions that accepted a pointer argument so they could return an error have been
converted to use Golang's error types and return multiple items.

## gversion.{h,c} & gversionmacros.h
Since go is statically compiled new import paths will be created when breaking API changes happen.
As such the contents of these files are redundant.

## gtypes.h
### Type Macros and definitions
Go already provides stable variants for most of these types or the types are not required with how go is structured.

### Byte Order Macros
The byte order macros are not implemented currently.
The "encoding.binary" package likely provides everything you need.

### Bounds-Checking Integer Arithmetic
- g_uint_checked_add(): implemented with a signature change func UintCheckAdd(a, b) z, error
- g_uint_checked_add(): implemented with a signature change

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

