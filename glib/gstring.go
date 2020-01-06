package glib

// GString - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gstring.h:39
type GString struct {
	str          []gchar
	len          gsize
	allocatedLen gsize
}

// GStringAppendCInline - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gstring.h:160
func GStringAppendCInline(gstring []GString, c gchar) []GString {
	if gsize(gstring[0].len)+gsize((1)) < gsize(gstring[0].allocatedLen) {
		// GLIB - Library of useful routines for C programming
		// * Copyright (C) 1995-1997  Peter Mattis, Spencer Kimball and Josh MacDonald
		// *
		// * This library is free software; you can redistribute it and/or
		// * modify it under the terms of the GNU Lesser General Public
		// * License as published by the Free Software Foundation; either
		// * version 2.1 of the License, or (at your option) any later version.
		// *
		// * This library is distributed in the hope that it will be useful,
		// * but WITHOUT ANY WARRANTY; without even the implied warranty of
		// * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
		// * Lesser General Public License for more details.
		// *
		// * You should have received a copy of the GNU Lesser General Public
		// * License along with this library; if not, see <http://www.gnu.org/licenses/>.
		//
		//
		// * Modified by the GLib Team and others 1997-2000.  See the AUTHORS
		// * file for a list of people on the GLib Team.  See the ChangeLog
		// * files for a list of changes.  These files are distributed with
		// * GLib at ftp://ftp.gtk.org/pub/gtk/.
		//
		// -- optimize g_strig_append_c ---
		gstring[0].str[func() gsize {
			tempVar := &gstring[0].len
			defer func() {
				*tempVar++
			}()
			return *tempVar
		}()] = c
		gstring[0].str[gsize(gstring[0].len)] = gchar(0)
	} else {
		//g_string_insert_c(gstring, gssize((-1)), c)
	}
	return gstring
}
