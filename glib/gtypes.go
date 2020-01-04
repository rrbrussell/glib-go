package glib

import (
	"math"
	"math/big"
)

/* glib-go - A port of GLIB for Golang
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.	 See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, see <http://www.gnu.org/licenses/>.
 *
 * These files are distributed with glib-go @ https://github.com/rrbrussell/glib-go.
 */

/* Form my personal Linux system
sizeof(char): 1
sizeof(short): 2
sizeof(int): 4
sizeof(long): 8
sizeof(long long): 8
sizeof(float): 4
sizeof(double): 8
sizeof(long double): 16
sizeof(bool): 1
sizeof(size_t): 8
*/

// Gchar int8
type Gchar rune

// Gint8 int8
type Gint8 int8

// Guint8 unit8
type Guint8 uint8

// Gint16 int16
type Gint16 int16

// Guint16 int16
type Guint16 uint16

// Gint32 int32
type Gint32 int32

// Guint32 int32
type Guint32 uint32

// Gint64 int64
type Gint64 int64

// Guint64 int64
type Guint64 uint64

// Gshort int16
type Gshort int16

// Glong int32
type Glong int32

// Gint int
type Gint int

// Gboolean bool
type Gboolean bool

//Gushort uint16
type Gushort uint16

// Gulong uint32
type Gulong uint32

// Guint uint
type Guint uint

// Gfloat float32
type Gfloat float32

//Gdouble float64
type Gdouble float64

// The min and maximums for integers
const (
	GMinInt8   Gint8   = math.MinInt8
	GMaxInt8   Gint8   = math.MaxInt8
	GMaxUint8  Guint8  = math.MaxUint8
	GMinInt16  Gint16  = math.MinInt16
	GMaxInt16  Gint16  = math.MaxInt16
	GMaxUint16 Guint16 = math.MaxUint16
	GMinInt32  Gint32  = math.MinInt32
	GMaxInt32  Gint32  = math.MaxInt32
	GMaxUint32 Guint32 = math.MaxUint32
	GMinInt64  Gint64  = math.MinInt64
	GMaxInt64  Gint64  = math.MaxInt64
	GMaxUint64 Guint64 = math.MaxUint64
)

// Some mathematical constants
const (
	E       Gdouble = 2.7182818284590452353602874713526624977572470937000
	LN2     Gdouble = 0.69314718055994530941723212145817656807550013436026
	LN10    Gdouble = 2.3025850929940456840179914546843642076011014886288
	PI      Gdouble = 3.1415926535897932384626433832795028841971693993751
	PIOver2 Gdouble = 1.5707963267948966192313216916397514420985846996876
	PIOver4 Gdouble = 0.78539816339744830961566084581987572104929234984378
	SQRT2   Gdouble = 1.4142135623730950488016887242096980785696718753769
)

// Uint64CheckedAdd adds two uint64s together and returns 0, true if the result won't fit in a uint64
func Uint64CheckedAdd(a uint64, b uint64) (val uint64, overflow bool) {
	aBig := big.NewInt(0)
	bBig := big.NewInt(0)
	aBig.SetUint64(a)
	bBig.SetUint64(b)
	aBig.Add(aBig, bBig)
	if !aBig.IsUint64() {
		return 0, true
	}
	return aBig.Uint64(), false
}

// Uint64CheckedMul adds two uint64s together and returns 0, true if the result won't fit in a uint64
func Uint64CheckedMul(a uint64, b uint64) (val uint64, overflow bool) {
	aBig := big.NewInt(0)
	bBig := big.NewInt(0)
	aBig.SetUint64(a)
	bBig.SetUint64(b)
	aBig.Mul(aBig, bBig)
	if !aBig.IsUint64() {
		return 0, true
	}
	return aBig.Uint64(), false
}
