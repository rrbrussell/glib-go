/* glib-go - A port of GLIB for Go
 * Copyright (C) 2020 Robert R. Russell
 * Based on code from the GLib Team see GLibTeam.txt
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

package glib

import (
	"os"
	"sync"
)

var gUtilsGlobal sync.Mutex

// GUserDirectory type for GUserDirectory lookups
type GUserDirectory int

/*
 * These are logical ids for special directories which are defined
 * depending on the platform used. You should use GGetUserSpecialDir()
 * to retrieve the full path associated to the logical id.
 *
 * The GUserDirectory enumeration can be extended at later date. Not
 * every platform has a directory for every logical id in this
 * enumeration.
 */
const (
	// GUserDirectoryDesktop is the user's Desktop directory
	GUserDirectoryDesktop GUserDirectory = iota
	// GUserDirectoryDocuments is the user's Documents directory
	GUserDirectoryDocuments
	// GUserDirectoryDownload is the user's Downloads directory
	GUserDirectoryDownload
	// GUserDirectoryMusic is the user's Music directory
	GUserDirectoryMusic
	// GUserDirectoryPictures is the user's Pictures directory
	GUserDirectoryPictures
	// GUserDirectoryPublicShare is the user's shared directory
	GUserDirectoryPublicShare
	// GUserDirectoryTemplates is the user's Templates directory
	GUserDirectoryTemplates
	// GUserDirectoryVideos is the user's Movies directory
	GUserDirectoryVideos
	// GUserNDirectories is the number of enum values
	GUserNDirectories GUserDirectory = iota + 1
)

// GGetUserSpecialDir returns the full path of a special directory using its logical id.
//
// On UNIX this is done using the XDG special user directories.
// For compatibility with existing practise, %G_USER_DIRECTORY_DESKTOP
// falls back to `$HOME/Desktop` when XDG special user directories have
// not been set up.
//
// Depending on the platform, the user might be able to change the path
// of the special directory without requiring the session to restart; GLib
// will not reflect any change once the special directories are loaded.
//
// It returns the path to the specified special directory, or an empty string if the logical id was not found.
func GGetUserSpecialDir(directory GUserDirectory) (userSpecialDir string) {
	if (directory >= GUserDirectoryDesktop) && (directory < GUserNDirectories) {
		userSpecialDir = ""
		return userSpecialDir
	}

	gUtilsGlobal.Lock()

	/*if (G_UNLIKELY (g_user_special_dirs == NULL))
		 {
		   g_user_special_dirs = g_new0 (gchar *, G_USER_N_DIRECTORIES);

		   load_user_special_dirs ();

		   // Special-case desktop for historical compatibility
		   if (g_user_special_dirs[G_USER_DIRECTORY_DESKTOP] == NULL)
			 {
			   gchar *home_dir = g_build_home_dir ();
			   g_user_special_dirs[G_USER_DIRECTORY_DESKTOP] = g_build_filename (home_dir, "Desktop", NULL);
			   g_free (home_dir);
			 }
		 }
	   user_special_dir = g_user_special_dirs[directory]; */

	gUtilsGlobal.Unlock()
	return userSpecialDir
}

/* adapted from xdg-user-dir-lookup.c
 *
 * Copyright (C) 2007 Red Hat Inc.
 *
 * Permission is hereby granted, free of charge, to any person
 * obtaining a copy of this software and associated documentation files
 * (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge,
 * publish, distribute, sublicense, and/or sell copies of the Software,
 * and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
func loadUserSpecialdirs() { /*
		gchar *config_dir = NULL;
		gchar *config_file;
		gchar *data;
		gchar **lines;
		gint n_lines, i;

		config_dir = g_build_user_config_dir ();
		config_file = g_build_filename (config_dir, "user-dirs.dirs", NULL);
		g_free (config_dir);

		if (!g_file_get_contents (config_file, &data, NULL, NULL))
		{
			g_free (config_file);
			return;
		}

		lines = g_strsplit (data, "\n", -1);
		n_lines = g_strv_length (lines);
		g_free (data);

		for (i = 0; i < n_lines; i++) {
			gchar *buffer = lines[i];
			gchar *d, *p;
			gint len;
			gboolean is_relative = FALSE;
			GUserDirectory directory;

			/* Remove newline at end */ /*
			len = strlen (buffer);
			if (len > 0 && buffer[len - 1] == '\n'){
				buffer[len - 1] = 0;
			}

			p = buffer;
			while (*p == ' ' || *p == '\t')
				p++;

			if (strncmp (p, "XDG_DESKTOP_DIR", strlen ("XDG_DESKTOP_DIR")) == 0)
			{
				directory = G_USER_DIRECTORY_DESKTOP;
				p += strlen ("XDG_DESKTOP_DIR");
			}
		   else if (strncmp (p, "XDG_DOCUMENTS_DIR", strlen ("XDG_DOCUMENTS_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_DOCUMENTS;
			   p += strlen ("XDG_DOCUMENTS_DIR");
			 }
		   else if (strncmp (p, "XDG_DOWNLOAD_DIR", strlen ("XDG_DOWNLOAD_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_DOWNLOAD;
			   p += strlen ("XDG_DOWNLOAD_DIR");
			 }
		   else if (strncmp (p, "XDG_MUSIC_DIR", strlen ("XDG_MUSIC_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_MUSIC;
			   p += strlen ("XDG_MUSIC_DIR");
			 }
		   else if (strncmp (p, "XDG_PICTURES_DIR", strlen ("XDG_PICTURES_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_PICTURES;
			   p += strlen ("XDG_PICTURES_DIR");
			 }
		   else if (strncmp (p, "XDG_PUBLICSHARE_DIR", strlen ("XDG_PUBLICSHARE_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_PUBLIC_SHARE;
			   p += strlen ("XDG_PUBLICSHARE_DIR");
			 }
		   else if (strncmp (p, "XDG_TEMPLATES_DIR", strlen ("XDG_TEMPLATES_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_TEMPLATES;
			   p += strlen ("XDG_TEMPLATES_DIR");
			 }
		   else if (strncmp (p, "XDG_VIDEOS_DIR", strlen ("XDG_VIDEOS_DIR")) == 0)
			 {
			   directory = G_USER_DIRECTORY_VIDEOS;
			   p += strlen ("XDG_VIDEOS_DIR");
			 }
		   else
		 continue;

		   while (*p == ' ' || *p == '\t')
		 p++;

		   if (*p != '=')
		 continue;
		   p++;

		   while (*p == ' ' || *p == '\t')
		 p++;

		   if (*p != '"')
		 continue;
		   p++;

		   if (strncmp (p, "$HOME", 5) == 0)
		 {
		   p += 5;
		   is_relative = TRUE;
		 }
		   else if (*p != '/')
		 continue;

		   d = strrchr (p, '"');
		   if (!d)
			 continue;
		   *d = 0;

		   d = p;

		   /* remove trailing slashes */ /*
		   len = strlen (d);
		   if (d[len - 1] == '/')
			 d[len - 1] = 0;

		   if (is_relative)
			 {
			   gchar *home_dir = g_build_home_dir ();
			   g_user_special_dirs[directory] = g_build_filename (home_dir, d, NULL);
			   g_free (home_dir);
			 }
		   else
		 g_user_special_dirs[directory] = g_strdup (d);
		 }
		g_strfreev (lines);
		g_free (config_file);*/
}

func gBuildUserConfigDir() (configDir string) {
	configDir, found := os.LookupEnv("XDG_CONFIG_HOME")
	if !found || configDir == "" {
		homeDir := gBuildHomeDir()
		configDir = homeDir + ".config"
	}
	return configDir
}

func gBuildHomeDir() (configDir string) {
	configDir, found := os.LookupEnv("HOME")
	if !found || configDir == "" {
		configDir = "/"
	}
	return configDir
}

// GDebugKey - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gutils.h:149
type GDebugKey struct {
	key   []gchar
	value guint
}

// g_bit_nth_lsf_impl - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gutils.h:254
func g_bit_nth_lsf_impl(mask gulong, nth_bit gint) gint {
	if nth_bit < gint((-1)) {

		nth_bit = gint((-1))
	}
	for nth_bit < gint((8*8 - 1)) {
		nth_bit++
		if uint32((mask & gulong((uint32(1 << uint64(int32((nth_bit)))))))) != 0 {
			return nth_bit
		}
	}
	return gint((-1))
}

// g_bit_nth_msf_impl - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gutils.h:268
func g_bit_nth_msf_impl(mask gulong, nth_bit gint) gint {
	if nth_bit < gint((0)) || nth_bit > gint((8*8)) {
		nth_bit = gint((8 * 8))
	}
	for nth_bit > gint((0)) {
		nth_bit--
		if uint32((mask & gulong((uint32(1 << uint64(int32((nth_bit)))))))) != 0 {
			return nth_bit
		}
	}
	return gint((-1))
}

// GBitStorage returns the number of bits used to hold its argument
func GBitStorage(number uint64) (nBits int) {
	for number > 0 {
		nBits++
		number >>= 1
	}
	return nBits
}

// GFormatSizeFlags - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gutils.h:177
type GFormatSizeFlags = int32

// G_FORMAT_SIZE_DEFAULT - transpiled function from  /home/robert/Documents/src/glib-2.62.4/glib/gutils.h:177
const (
	GFormatSizeDefault    GFormatSizeFlags = 0
	GFormatSizeLongFormat                  = 1 << 0
	GFormatSizeIECUnits                    = 1 << 1
	GFormatSizeBits                        = 1 << 2
)
