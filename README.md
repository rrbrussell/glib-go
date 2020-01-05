# glib-go: A pure go implementation of glib

This is designed around glib-2.62.4. Currently the goal is minimal conversion of glib into Golang.
Language specific optimizations should be internally only right now.

The long range plan is to get enough of this and the GTK2 stack ported to Golang that I can write
Golang apps with a fully supported GTK2 based gui. The reason for targeting GTK2 instead of GTK3
is a combination of design decisions made in GTK3 and stability of the API. The really longterm plan
is to get this code polished well enough that the GTK2 versions of Mate, XFCE, or maybe GNOME 2 itself
can be converted into Golang and work on Linux and maybe Windows.

# Files in glib to port

## galloca.h

- [ ] glib/gtypes.h

## garcbox.c

- [ ] config.h
- [ ] glib_trace.h
- [ ] gmessages.h
- [ ] grcboxprivate.h
- [ ] grefcount.h
- [ ] string.h
- [ ] valgrind.h

## garray.c

- [ ] config.h
- [ ] garray.h
- [ ] gbytes.h
- [ ] ghash.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gqsort.h
- [ ] grefcount.h
- [ ] gslice.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] stdlib.h
- [ ] string.h

## garray.h

- [ ] glib/gtypes.h

## gasyncqueue.c

- [ ] config.h
- [ ] deprecated/gthread.h
- [ ] gasyncqueue.h
- [ ] gasyncqueueprivate.h
- [ ] gmain.h
- [ ] gmem.h
- [ ] gqueue.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gtimer.h

## gasyncqueue.h

- [ ] glib/gthread.h

## gasyncqueueprivate.h

- [ ] gasyncqueue.h

## gatomic.c - Partial

- [X] gatomic.h

## gatomic.h - Partial

- [X] glib/gtypes.h

## gbacktrace.c

- [ ] config.h
- [ ] gbacktrace.h
- [ ] glibconfig.h
- [ ] gmain.h
- [ ] gprintfint.h
- [ ] gtypes.h
- [ ] gutils.h
- [ ] signal.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/select.h
- [ ] sys/time.h
- [ ] sys/types.h
- [ ] sys/wait.h
- [ ] time.h
- [ ] unistd.h

## gbacktrace.h

- [ ] glib/gtypes.h
- [ ] signal.h
- [ ] sys/select.h

## gbase64.c

- [ ] config.h
- [ ] gbase64.h
- [ ] glibintl.h
- [ ] gtestutils.h
- [ ] string.h

## gbase64.h

- [ ] glib/gtypes.h

## gbitlock.c

- [ ] config.h
- [ ] gbitlock.h
- [ ] glib/gatomic.h
- [ ] glib/gmessages.h
- [ ] glib/gslice.h
- [ ] glib/gslist.h
- [ ] glib/gthread.h
- [ ] gthreadprivate.h
- [ ] linux/futex.h
- [ ] sys/syscall.h
- [ ] unistd.h

## gbitlock.h

- [ ] glib/gtypes.h

## gbookmarkfile.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gbookmarkfile.h
- [ ] gconvert.h
- [ ] gdataset.h
- [ ] gerror.h
- [ ] gfileutils.h
- [ ] ghash.h
- [ ] glibintl.h
- [ ] glist.h
- [ ] gmain.h
- [ ] gmarkup.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gshell.h
- [ ] gslice.h
- [ ] gstdio.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtimer.h
- [ ] gutils.h
- [ ] locale.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] time.h

## gbookmarkfile.h

- [ ] glib/gerror.h
- [ ] time.h

## gbsearcharray.h

- [ ] glib.h
- [ ] string.h

## gbytes.c

- [ ] config.h
- [ ] gbytes.h
- [ ] glib/garray.h
- [ ] glib/gatomic.h
- [ ] glib/gmem.h
- [ ] glib/gmessages.h
- [ ] glib/grefcount.h
- [ ] glib/gslice.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gtestutils.h
- [ ] string.h

## gbytes.h

- [ ] glib/garray.h
- [ ] glib/gtypes.h

## gcharset.c

- [ ] config.h
- [ ] garray.h
- [ ] gcharset.h
- [ ] gcharsetprivate.h
- [ ] genviron.h
- [ ] ghash.h
- [ ] gmessages.h
- [ ] gstrfuncs.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] gwin32.h
- [ ] libcharset/libcharset.h
- [ ] stdio.h
- [ ] string.h
- [ ] windows.h

## gcharset.h

- [ ] glib/gtypes.h

## gcharsetprivate.h

- [ ] gcharset.h

## gchecksum.c

- [ ] config.h
- [ ] gchecksum.h
- [ ] glibintl.h
- [ ] gmem.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gtypes.h
- [ ] string.h

## gchecksum.h

- [ ] glib/gbytes.h
- [ ] glib/gtypes.h

## gconstructor.h

- [ ] stdlib.h

## gconvert.c

- [ ] config.h
- [ ] errno.h
- [ ] gcharsetprivate.h
- [ ] gconvert.h
- [ ] gfileutils.h
- [ ] glibconfig.h
- [ ] glibintl.h
- [ ] gslist.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] gunicode.h
- [ ] iconv.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] win_iconv.c
- [ ] windows.h

## gconvert.h

- [ ] glib/gerror.h

## gdataset.c

- [ ] config.h
- [ ] gbitlock.h
- [ ] gdataset.h
- [ ] gdatasetprivate.h
- [ ] ghash.h
- [ ] glib_trace.h
- [ ] gquark.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] string.h

## gdataset.h

- [ ] glib/gquark.h

## gdatasetprivate.h

- [ ] gatomic.h

## gdate.c

- [ ] config.h
- [ ] garray.h
- [ ] gconvert.h
- [ ] gdate.h
- [ ] glibconfig.h
- [ ] gmem.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gunicode.h
- [ ] locale.h
- [ ] stdlib.h
- [ ] string.h
- [ ] time.h
- [ ] windows.h

## gdate.h

- [ ] glib/gquark.h
- [ ] glib/gtypes.h
- [ ] time.h

## gdatetime.c

- [ ] config.h
- [ ] gatomic.h
- [ ] gcharset.h
- [ ] gconvert.h
- [ ] gdatetime.h
- [ ] gfileutils.h
- [ ] ghash.h
- [ ] glibintl.h
- [ ] gmain.h
- [ ] gmappedfile.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gtimezone.h
- [ ] langinfo.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/time.h
- [ ] time.h

## gdatetime.h

- [ ] glib/gtimezone.h

## gdir.c

- [ ] config.h
- [ ] dirent.h
- [ ] dirent/dirent.h
- [ ] errno.h
- [ ] gconvert.h
- [ ] gdir.h
- [ ] gfileutils.h
- [ ] glib-private.h
- [ ] glibintl.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] stdio.h
- [ ] string.h
- [ ] sys/stat.h
- [ ] sys/types.h

## gdir.h

- [ ] dirent.h
- [ ] glib/gerror.h

## gen-unicode-tables.pl

- [ ] print
- [ ] print

## genviron.c

- [ ] config.h
- [ ] crt_externs.h
- [ ] gconvert.h
- [ ] genviron.h
- [ ] glib-private.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gquark.h
- [ ] gstrfuncs.h
- [ ] gunicode.h
- [ ] stdlib.h
- [ ] string.h
- [ ] windows.h

## genviron.h

- [ ] glib/gtypes.h

## gerror.c

- [ ] config.h
- [ ] gerror.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h

## gerror.h

- [ ] glib/gquark.h
- [ ] stdarg.h

## gfileutils.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gfileutils.h
- [ ] glibconfig.h
- [ ] glibintl.h
- [ ] gstdio.h
- [ ] gstdioprivate.h
- [ ] io.h
- [ ] linux/magic.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/stat.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] sys/vfs.h
- [ ] unistd.h
- [ ] windows.h

## gfileutils.h

- [ ] glib/gerror.h
- [ ] glibconfig.h

## ggettext.c

- [ ] *
- [ ] *
- [ ] *
- [ ] config.h
- [ ] galloca.h
- [ ] gfileutils.h
- [ ] ggettext.h
- [ ] glib-init.h
- [ ] glib-private.h
- [ ] glibintl.h
- [ ] gmem.h
- [ ] gstrfuncs.h
- [ ] gthread.h
- [ ] gwin32.h
- [ ] libintl.h
- [ ] locale.h
- [ ] string.h

## ggettext.h

- [ ] glib/gtypes.h

## ghash.c

- [ ] config.h
- [ ] gatomic.h
- [ ] ghash.h
- [ ] glib-private.h
- [ ] gmacros.h
- [ ] grefcount.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gvalgrind.h
- [ ] string.h

## ghash.h

- [ ] glib/glist.h
- [ ] glib/gtypes.h

## ghmac.c

- [ ] config.h
- [ ] gatomic.h
- [ ] ghmac.h
- [ ] glib/galloca.h
- [ ] glibintl.h
- [ ] gmem.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gtypes.h
- [ ] string.h

## ghmac.h

- [ ] gchecksum.h
- [ ] glib/gtypes.h

## ghook.c

- [ ] config.h
- [ ] ghook.h
- [ ] gslice.h
- [ ] gtestutils.h

## ghook.h

- [ ] glib/gmem.h

## ghostutils.c

- [ ] config.h
- [ ] garray.h
- [ ] ghostutils.h
- [ ] glibintl.h
- [ ] gmem.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] string.h

## ghostutils.h

- [ ] glib/gtypes.h

## gi18n-lib.h

- [ ] glib.h
- [ ] libintl.h
- [ ] string.h

## gi18n.h

- [ ] glib.h
- [ ] libintl.h
- [ ] string.h

## giochannel.c

- [ ] config.h
- [ ] errno.h
- [ ] giochannel.h
- [ ] glibintl.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] string.h

## giochannel.h

- [ ] glib/gconvert.h
- [ ] glib/gmain.h
- [ ] glib/gstring.h

## giounix.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gerror.h
- [ ] gfileutils.h
- [ ] giochannel.h
- [ ] glib/gstdio.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] stdio.h
- [ ] string.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] unistd.h

## giowin32.c

- [ ] config.h
- [ ] conio.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] glib.h
- [ ] glibintl.h
- [ ] gstdio.h
- [ ] io.h
- [ ] process.h
- [ ] stdlib.h
- [ ] sys/stat.h
- [ ] windows.h
- [ ] winsock2.h

## gkeyfile.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gconvert.h
- [ ] gdataset.h
- [ ] gerror.h
- [ ] gfileutils.h
- [ ] ghash.h
- [ ] gkeyfile.h
- [ ] glibintl.h
- [ ] glist.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gslist.h
- [ ] gstdio.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gutils.h
- [ ] gutils.h
- [ ] io.h
- [ ] locale.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] unistd.h

## gkeyfile.h

- [ ] glib/gbytes.h
- [ ] glib/gerror.h

## glib-init.c

- [ ] config.h
- [ ] ctype.h
- [ ] gconstructor.h
- [ ] glib-init.h
- [ ] gmacros.h
- [ ] gmem.h
- [ ] gtypes.h
- [ ] gutils.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h

## glib-init.h

- [ ] gmessages.h
- [ ] windows.h

## glib-object.h

- [ ] gobject/gbinding.h
- [ ] gobject/gboxed.h
- [ ] gobject/genums.h
- [ ] gobject/glib-enumtypes.h
- [ ] gobject/gobject-autocleanups.h
- [ ] gobject/gobject.h
- [ ] gobject/gparam.h
- [ ] gobject/gparamspecs.h
- [ ] gobject/gsignal.h
- [ ] gobject/gsourceclosure.h
- [ ] gobject/gtype.h
- [ ] gobject/gtypemodule.h
- [ ] gobject/gtypeplugin.h
- [ ] gobject/gvalue.h
- [ ] gobject/gvaluearray.h
- [ ] gobject/gvaluetypes.h

## glib-private.c

- [ ] config.h
- [ ] glib-init.h
- [ ] glib-private.h

## glib-private.h

- [ ] glib.h
- [ ] gstdioprivate.h
- [ ] gwakeup.h

## glib-unix.c

- [ ] config.h
- [ ] glib-unix.h
- [ ] gmain-internal.h
- [ ] string.h

## glib-unix.h

- [ ] errno.h
- [ ] fcntl.h
- [ ] glib.h
- [ ] stdlib.h
- [ ] sys/wait.h
- [ ] unistd.h

## glib.h

- [ ] glib/deprecated/gallocator.h
- [ ] glib/deprecated/gcache.h
- [ ] glib/deprecated/gcompletion.h
- [ ] glib/deprecated/gmain.h
- [ ] glib/deprecated/grel.h
- [ ] glib/deprecated/gthread.h
- [ ] glib/galloca.h
- [ ] glib/garray.h
- [ ] glib/gasyncqueue.h
- [ ] glib/gatomic.h
- [ ] glib/gbacktrace.h
- [ ] glib/gbase64.h
- [ ] glib/gbitlock.h
- [ ] glib/gbookmarkfile.h
- [ ] glib/gbytes.h
- [ ] glib/gcharset.h
- [ ] glib/gchecksum.h
- [ ] glib/gconvert.h
- [ ] glib/gdataset.h
- [ ] glib/gdate.h
- [ ] glib/gdatetime.h
- [ ] glib/gdir.h
- [ ] glib/genviron.h
- [ ] glib/gerror.h
- [ ] glib/gfileutils.h
- [ ] glib/ggettext.h
- [ ] glib/ghash.h
- [ ] glib/ghmac.h
- [ ] glib/ghook.h
- [ ] glib/ghostutils.h
- [ ] glib/giochannel.h
- [ ] glib/gkeyfile.h
- [ ] glib/glib-autocleanups.h
- [ ] glib/glist.h
- [ ] glib/gmacros.h
- [ ] glib/gmain.h
- [ ] glib/gmappedfile.h
- [ ] glib/gmarkup.h
- [ ] glib/gmem.h
- [ ] glib/gmessages.h
- [ ] glib/gnode.h
- [ ] glib/goption.h
- [ ] glib/gpattern.h
- [ ] glib/gpoll.h
- [ ] glib/gprimes.h
- [ ] glib/gqsort.h
- [ ] glib/gquark.h
- [ ] glib/gqueue.h
- [ ] glib/grand.h
- [ ] glib/grcbox.h
- [ ] glib/grefcount.h
- [ ] glib/grefstring.h
- [ ] glib/gregex.h
- [ ] glib/gscanner.h
- [ ] glib/gsequence.h
- [ ] glib/gshell.h
- [ ] glib/gslice.h
- [ ] glib/gslist.h
- [ ] glib/gspawn.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gstring.h
- [ ] glib/gstringchunk.h
- [ ] glib/gtestutils.h
- [ ] glib/gthread.h
- [ ] glib/gthreadpool.h
- [ ] glib/gtimer.h
- [ ] glib/gtimezone.h
- [ ] glib/gtrashstack.h
- [ ] glib/gtree.h
- [ ] glib/gtypes.h
- [ ] glib/gunicode.h
- [ ] glib/gurifuncs.h
- [ ] glib/gutils.h
- [ ] glib/guuid.h
- [ ] glib/gvariant.h
- [ ] glib/gvarianttype.h
- [ ] glib/gversion.h
- [ ] glib/gversionmacros.h
- [ ] glib/gwin32.h

## glib.rc.in

- [ ] winver.h

## glib_trace.h

- [ ] glib_probes.h

## glibconfig.h.in

- [ ] float.h
- [ ] glib/gmacros.h
- [ ] limits.h

## glibintl.h

- [ ] libintl.h

## glist.c

- [ ] config.h
- [ ] glist.h
- [ ] gmessages.h
- [ ] gslice.h
- [ ] gtestutils.h

## glist.h

- [ ] glib/gmem.h
- [ ] glib/gnode.h

## gmacros.h

- [ ] stddef.h

## gmain-internal.h

- [ ] gmain.h

## gmain.c

- [ ] config.h
- [ ] errno.h
- [ ] garray.h
- [ ] ghash.h
- [ ] ghook.h
- [ ] giochannel.h
- [ ] glib-init.h
- [ ] glib-private.h
- [ ] glib-unix.h
- [ ] glib_trace.h
- [ ] glib_trace.h
- [ ] glibconfig.h
- [ ] gmain-internal.h
- [ ] gmain.h
- [ ] gqueue.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthreadprivate.h
- [ ] gtimer.h
- [ ] gwakeup.h
- [ ] gwin32.h
- [ ] mach/mach_time.h
- [ ] pthread.h
- [ ] signal.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/eventfd.h
- [ ] sys/time.h
- [ ] sys/types.h
- [ ] time.h
- [ ] unistd.h
- [ ] windows.h

## gmain.h

- [ ] glib/gpoll.h
- [ ] glib/gslist.h
- [ ] glib/gthread.h

## gmappedfile.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gatomic.h
- [ ] gconvert.h
- [ ] gerror.h
- [ ] gfileutils.h
- [ ] glibconfig.h
- [ ] glibintl.h
- [ ] gmappedfile.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gstdio.h
- [ ] gstrfuncs.h
- [ ] io.h
- [ ] sys/mman.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] unistd.h
- [ ] windows.h

## gmappedfile.h

- [ ] glib/gbytes.h
- [ ] glib/gerror.h

## gmarkup.c

- [ ] config.h
- [ ] errno.h
- [ ] galloca.h
- [ ] gatomic.h
- [ ] glibintl.h
- [ ] gmarkup.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h

## gmarkup.h

- [ ] glib/gerror.h
- [ ] glib/gslist.h
- [ ] stdarg.h

## gmem.c

- [ ] config.h
- [ ] gbacktrace.h
- [ ] glib_trace.h
- [ ] gmem.h
- [ ] gslice.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] signal.h
- [ ] stdlib.h
- [ ] string.h

## gmem.h

- [ ] glib/gutils.h

## gmessages.c

- [ ] config.h
- [ ] crtdbg.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] galloca.h
- [ ] gbacktrace.h
- [ ] gcharset.h
- [ ] gconvert.h
- [ ] genviron.h
- [ ] glib-init.h
- [ ] gmain.h
- [ ] gmem.h
- [ ] gpattern.h
- [ ] gprintfint.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gwin32.h
- [ ] io.h
- [ ] locale.h
- [ ] process.h
- [ ] signal.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/socket.h
- [ ] sys/types.h
- [ ] sys/uio.h
- [ ] sys/un.h
- [ ] unistd.h

## gmessages.h

- [ ] glib/gmacros.h
- [ ] glib/gtypes.h
- [ ] glib/gvariant.h
- [ ] stdarg.h

## gnode.c

- [ ] config.h
- [ ] gnode.h
- [ ] gslice.h
- [ ] gtestutils.h

## gnode.h

- [ ] glib/gmem.h

## goption.c

- [ ] config.h
- [ ] errno.h
- [ ] glibintl.h
- [ ] goption.h
- [ ] gprintf.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/sysctl.h
- [ ] unistd.h
- [ ] windows.h

## goption.h

- [ ] glib/gerror.h
- [ ] glib/gquark.h

## gosxutils.m

- [ ] Cocoa/Cocoa.h
- [ ] config.h
- [ ] gstrfuncs.h
- [ ] gutils.h

## gpattern.c

- [ ] config.h
- [ ] gmacros.h
- [ ] gmem.h
- [ ] gmessages.h
- [ ] gpattern.h
- [ ] gunicode.h
- [ ] gutils.h
- [ ] string.h

## gpattern.h

- [ ] glib/gtypes.h

## gpoll.c

- [ ] config.h
- [ ] errno.h
- [ ] giochannel.h
- [ ] glibconfig.h
- [ ] gpoll.h
- [ ] gprintf.h
- [ ] process.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/select.h
- [ ] sys/time.h
- [ ] sys/types.h
- [ ] time.h
- [ ] unistd.h
- [ ] windows.h

## gpoll.h

- [ ] glib/gtypes.h
- [ ] glibconfig.h

## gprimes.c

- [ ] config.h
- [ ] gprimes.h

## gprimes.h

- [ ] glib/gtypes.h

## gprintf.c

- [ ] config.h
- [ ] gprintf.h
- [ ] gprintfint.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h

## gprintf.h

- [ ] glib.h
- [ ] stdarg.h
- [ ] stdio.h

## gprintfint.h

- [ ] gnulib/printf.h

## gqsort.c

- [ ] config.h
- [ ] galloca.h
- [ ] gmem.h
- [ ] gqsort.h
- [ ] gtestutils.h
- [ ] limits.h
- [ ] stdlib.h
- [ ] string.h

## gqsort.h

- [ ] glib/gtypes.h

## gquark.c

- [ ] config.h
- [ ] ghash.h
- [ ] glib-init.h
- [ ] glib_trace.h
- [ ] gquark.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] string.h

## gquark.h

- [ ] glib/gtypes.h

## gqueue.c

- [ ] config.h
- [ ] gqueue.h
- [ ] gslice.h
- [ ] gtestutils.h

## gqueue.h

- [ ] glib/glist.h

## grand.c

- [ ] config.h
- [ ] errno.h
- [ ] genviron.h
- [ ] gmain.h
- [ ] gmem.h
- [ ] grand.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gtimer.h
- [ ] math.h
- [ ] process.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/types.h
- [ ] unistd.h

## grand.h

- [ ] glib/gtypes.h

## grcbox.c

- [ ] config.h
- [ ] glib_trace.h
- [ ] gmessages.h
- [ ] grcboxprivate.h
- [ ] grefcount.h
- [ ] gtestutils.h
- [ ] string.h
- [ ] valgrind.h

## grcbox.h

- [ ] glib/gmem.h

## grcboxprivate.h

- [ ] grcbox.h
- [ ] gtypes.h

## grefcount.c

- [ ] config.h
- [ ] gatomic.h
- [ ] gmessages.h
- [ ] grefcount.h

## grefcount.h

- [ ] glib/gatomic.h
- [ ] glib/gtypes.h

## grefstring.c

- [ ] config.h
- [ ] ghash.h
- [ ] gmessages.h
- [ ] grcbox.h
- [ ] grefstring.h
- [ ] gthread.h
- [ ] string.h

## grefstring.h

- [ ] gmacros.h
- [ ] gmem.h

## gregex.c

- [ ] config.h
- [ ] gatomic.h
- [ ] glibintl.h
- [ ] glist.h
- [ ] gmessages.h
- [ ] gregex.h
- [ ] gstrfuncs.h
- [ ] gthread.h
- [ ] gtypes.h
- [ ] pcre.h
- [ ] pcre/pcre.h
- [ ] string.h

## gregex.h

- [ ] glib/gerror.h
- [ ] glib/gstring.h

## gscanner.c

- [ ] config.h
- [ ] errno.h
- [ ] gprintfint.h
- [ ] gscanner.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] io.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] unistd.h

## gscanner.h

- [ ] glib/gdataset.h
- [ ] glib/ghash.h

## gsequence.c

- [ ] config.h
- [ ] gmem.h
- [ ] gsequence.h
- [ ] gslice.h
- [ ] gtestutils.h

## gsequence.h

- [ ] glib/gtypes.h

## gshell.c

- [ ] config.h
- [ ] glibintl.h
- [ ] gshell.h
- [ ] gslist.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] string.h

## gshell.h

- [ ] glib/gerror.h

## gslice.c

- [ ] config.h
- [ ] errno.h
- [ ] glib_trace.h
- [ ] glibconfig.h
- [ ] gmain.h
- [ ] gmem.h
- [ ] gprintf.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] gtrashstack.h
- [ ] gutils.h
- [ ] gvalgrind.h
- [ ] malloc.h
- [ ] process.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] unistd.h
- [ ] windows.h

## gslice.h

- [ ] glib/gtypes.h

## gslist.c

- [ ] config.h
- [ ] gslice.h
- [ ] gslist.h
- [ ] gtestutils.h

## gslist.h

- [ ] glib/gmem.h
- [ ] glib/gnode.h

## gspawn-private.h

- [ ] config.h
- [ ] errno.h
- [ ] gspawn.h

## gspawn-win32-helper.c

- [ ] config.h
- [ ] crtdbg.h
- [ ] fcntl.h
- [ ] glib.h
- [ ] gspawn-win32.c

## gspawn-win32.c

- [ ] config.h
- [ ] direct.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] glib-private.h
- [ ] glib.h
- [ ] glibintl.h
- [ ] gprintfint.h
- [ ] gspawn-private.h
- [ ] gthread.h
- [ ] io.h
- [ ] process.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] wchar.h
- [ ] windows.h

## gspawn.c

- [ ] config.h
- [ ] crt_externs.h
- [ ] dirent.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] genviron.h
- [ ] glib-unix.h
- [ ] glib/gstdio.h
- [ ] glibintl.h
- [ ] gmem.h
- [ ] gshell.h
- [ ] gspawn-private.h
- [ ] gspawn.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gutils.h
- [ ] signal.h
- [ ] spawn.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/resource.h
- [ ] sys/select.h
- [ ] sys/syscall.h
- [ ] sys/time.h
- [ ] sys/types.h
- [ ] sys/wait.h
- [ ] unistd.h

## gspawn.h

- [ ] glib/gerror.h

## gstdio.c

- [ ] config.h
- [ ] direct.h
- [ ] errno.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] glibconfig.h
- [ ] gstdio-private.c
- [ ] gstdio.h
- [ ] gstdioprivate.h
- [ ] io.h
- [ ] stdlib.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] sys/utime.h
- [ ] unistd.h
- [ ] utime.h
- [ ] wchar.h
- [ ] windows.h

## gstdio.h

- [ ] glib/gprintf.h
- [ ] sys/stat.h

## gstrfuncs.c

- [ ] config.h
- [ ] ctype.h
- [ ] errno.h
- [ ] glibintl.h
- [ ] gprintf.h
- [ ] gprintfint.h
- [ ] gstrfuncs.h
- [ ] locale.h
- [ ] locale.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] windows.h
- [ ] xlocale.h

## gstrfuncs.h

- [ ] glib/gerror.h
- [ ] glib/gmacros.h
- [ ] glib/gtypes.h
- [ ] stdarg.h

## gstring.c

- [ ] config.h
- [ ] ctype.h
- [ ] gprintf.h
- [ ] gstring.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h

## gstring.h

- [ ] glib/gbytes.h
- [ ] glib/gtypes.h
- [ ] glib/gunicode.h
- [ ] glib/gutils.h

## gstringchunk.c

- [ ] config.h
- [ ] ghash.h
- [ ] gmessages.h
- [ ] gslist.h
- [ ] gstringchunk.h
- [ ] gutils.h
- [ ] string.h

## gstringchunk.h

- [ ] glib/gtypes.h

## gtester.c

- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] glib-unix.h
- [ ] glib.h
- [ ] gstdio.h
- [ ] signal.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/wait.h
- [ ] unistd.h

## gtestutils.c

- [ ] *
- [ ] *
- [ ] config.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] gfileutils.h
- [ ] glib-private.h
- [ ] glib/gstdio.h
- [ ] gmain.h
- [ ] gpattern.h
- [ ] grand.h
- [ ] gslice.h
- [ ] gspawn.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gtimer.h
- [ ] gutilsprivate.h
- [ ] io.h
- [ ] signal.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/resource.h
- [ ] sys/select.h
- [ ] sys/time.h
- [ ] sys/types.h
- [ ] sys/wait.h
- [ ] unistd.h
- [ ] windows.h

## gtestutils.h

- [ ] glib/gerror.h
- [ ] glib/gmessages.h
- [ ] glib/gslist.h
- [ ] glib/gstring.h
- [ ] string.h

## gthread-posix.c

- [ ] config.h
- [ ] errno.h
- [ ] gmain.h
- [ ] gmessages.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] gutils.h
- [ ] linux/futex.h
- [ ] pthread.h
- [ ] pthread_np.h
- [ ] sched.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/syscall.h
- [ ] sys/time.h
- [ ] unistd.h
- [ ] windows.h

## gthread-win32.c

- [ ] config.h
- [ ] glib-init.h
- [ ] glib.h
- [ ] gslice.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] process.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] windows.h

## gthread.c

- [ ] config.h
- [ ] glib_trace.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gthreadprivate.h
- [ ] string.h
- [ ] sys/time.h
- [ ] time.h
- [ ] unistd.h
- [ ] windows.h

## gthread.h

- [ ] glib/gatomic.h
- [ ] glib/gerror.h
- [ ] glib/gutils.h

## gthreadpool.c

- [ ] config.h
- [ ] gasyncqueue.h
- [ ] gasyncqueueprivate.h
- [ ] gmain.h
- [ ] gtestutils.h
- [ ] gthreadpool.h
- [ ] gtimer.h
- [ ] gutils.h

## gthreadpool.h

- [ ] glib/gthread.h

## gthreadprivate.h

- [ ] deprecated/gthread.h

## gtimer.c

- [ ] config.h
- [ ] errno.h
- [ ] glibconfig.h
- [ ] gmain.h
- [ ] gmem.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gtimer.h
- [ ] stdlib.h
- [ ] sys/time.h
- [ ] time.h
- [ ] unistd.h
- [ ] windows.h

## gtimer.h

- [ ] glib/gtypes.h

## gtimezone.c

- [ ] config.h
- [ ] gbytes.h
- [ ] gdate.h
- [ ] gdatetime.h
- [ ] gfileutils.h
- [ ] ghash.h
- [ ] gmappedfile.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gtimezone.h
- [ ] signal.h
- [ ] stdlib.h
- [ ] string.h
- [ ] wchar.h
- [ ] windows.h

## gtimezone.h

- [ ] glib/gtypes.h

## gtranslit.c

- [ ] config.h
- [ ] glib.h
- [ ] gstrfuncs.h
- [ ] gtranslit-data.h
- [ ] locale.h
- [ ] stdlib.h
- [ ] string.h

## gtrashstack.c

- [ ] config.h
- [ ] gtrashstack.h

## gtrashstack.h

- [ ] glib/gutils.h

## gtree.c

- [ ] config.h
- [ ] gatomic.h
- [ ] gslice.h
- [ ] gtestutils.h
- [ ] gtree.h

## gtree.h

- [ ] glib/gnode.h

## gtypes.h

- [ ] glib/gmacros.h
- [ ] glib/gversionmacros.h
- [ ] glibconfig.h
- [ ] time.h

## gunibreak.c

- [ ] config.h
- [ ] gunibreak.h
- [ ] stdlib.h

## gunibreak.h

- [ ] glib/gtypes.h
- [ ] glib/gunicode.h

## gunicode.h

- [ ] glib/gerror.h
- [ ] glib/gtypes.h

## gunicodeprivate.h

- [ ] gtypes.h

## gunicollate.c

- [ ] CoreServices/CoreServices.h
- [ ] config.h
- [ ] gcharset.h
- [ ] gconvert.h
- [ ] gmem.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gunicode.h
- [ ] gunicodeprivate.h
- [ ] locale.h
- [ ] string.h
- [ ] wchar.h

## gunidecomp.c

- [ ] config.h
- [ ] gmem.h
- [ ] gunicode.h
- [ ] gunicodeprivate.h
- [ ] gunicomp.h
- [ ] gunidecomp.h
- [ ] stdlib.h

## guniprop.c

- [ ] config.h
- [ ] gmem.h
- [ ] gmirroringtable.h
- [ ] gscripttable.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gtypes.h
- [ ] gunichartables.h
- [ ] gunicode.h
- [ ] gunicodeprivate.h
- [ ] gwin32.h
- [ ] locale.h
- [ ] stddef.h
- [ ] stdlib.h
- [ ] string.h

## gurifuncs.c

- [ ] config.h
- [ ] config.h
- [ ] glib/gmem.h
- [ ] glib/gmessages.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gstring.h
- [ ] gurifuncs.h
- [ ] string.h

## gurifuncs.h

- [ ] glib/gtypes.h

## gutf8.c

- [ ] config.h
- [ ] gconvert.h
- [ ] ghash.h
- [ ] glibintl.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gtypes.h
- [ ] langinfo.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] windows.h

## gutils.c

- [ ] config.h
- [ ] crt_externs.h
- [ ] ctype.h
- [ ] errno.h
- [ ] garray.h
- [ ] gconvert.h
- [ ] genviron.h
- [ ] gfileutils.h
- [ ] ggettext.h
- [ ] ghash.h
- [ ] glib-init.h
- [ ] glib-private.h
- [ ] glibintl.h
- [ ] gstdio.h
- [ ] gstrfuncs.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gunicode.h
- [ ] gutils.h
- [ ] gutilsprivate.h
- [ ] gwin32.h
- [ ] langinfo.h
- [ ] locale.h
- [ ] pwd.h
- [ ] stdarg.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/auxv.h
- [ ] sys/param.h
- [ ] sys/stat.h
- [ ] sys/types.h
- [ ] sys/types.h
- [ ] unistd.h

## gutils.h

- [ ] glib/gtypes.h
- [ ] stdarg.h

## gutilsprivate.h

- [ ] gtypes.h

## guuid.c

- [ ] config.h
- [ ] gchecksum.h
- [ ] gi18n.h
- [ ] gmessages.h
- [ ] grand.h
- [ ] gstrfuncs.h
- [ ] guuid.h
- [ ] string.h

## guuid.h

- [ ] glib/gtypes.h

## gvalgrind.h

- [ ] stdint.h
- [ ] valgrind.h

## gvariant-core.c

- [ ] config.h
- [ ] glib/gatomic.h
- [ ] glib/gbitlock.h
- [ ] glib/gbytes.h
- [ ] glib/gmem.h
- [ ] glib/grefcount.h
- [ ] glib/gslice.h
- [ ] glib/gtestutils.h
- [ ] glib/gvariant-core.h
- [ ] glib/gvariant-internal.h
- [ ] glib/gvariant-serialiser.h
- [ ] string.h

## gvariant-core.h

- [ ] glib/gbytes.h
- [ ] glib/gvariant.h
- [ ] glib/gvarianttypeinfo.h

## gvariant-internal.h

- [ ] glib/gtypes.h
- [ ] glib/gvarianttype.h
- [ ] gvariant-serialiser.h
- [ ] gvarianttypeinfo.h

## gvariant-parser.c

- [ ] config.h
- [ ] errno.h
- [ ] gerror.h
- [ ] gquark.h
- [ ] gslice.h
- [ ] gstrfuncs.h
- [ ] gstring.h
- [ ] gtestutils.h
- [ ] gthread.h
- [ ] gvariant-internal.h
- [ ] gvariant.h
- [ ] gvarianttype.h
- [ ] stdlib.h
- [ ] string.h

## gvariant-serialiser.c

- [ ] config.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gtestutils.h
- [ ] glib/gtypes.h
- [ ] glib/gvariant-internal.h
- [ ] gvariant-serialiser.h
- [ ] string.h

## gvariant-serialiser.h

- [ ] gvarianttypeinfo.h

## gvariant.c

- [ ] config.h
- [ ] glib/ghash.h
- [ ] glib/gmem.h
- [ ] glib/gslice.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gtestutils.h
- [ ] glib/gvariant-core.h
- [ ] glib/gvariant-serialiser.h
- [ ] gvariant-internal.h
- [ ] string.h

## gvariant.h

- [ ] glib/gbytes.h
- [ ] glib/gstring.h
- [ ] glib/gvarianttype.h

## gvarianttype.c

- [ ] config.h
- [ ] glib/gstrfuncs.h
- [ ] glib/gtestutils.h
- [ ] glib/gvariant-internal.h
- [ ] gvarianttype.h
- [ ] string.h

## gvarianttype.h

- [ ] glib/gtypes.h

## gvarianttypeinfo.c

- [ ] config.h
- [ ] glib/ghash.h
- [ ] glib/grefcount.h
- [ ] glib/gslice.h
- [ ] glib/gtestutils.h
- [ ] glib/gthread.h
- [ ] gvarianttypeinfo.h

## gvarianttypeinfo.h

- [ ] glib/gvarianttype.h

## gversion.c

- [ ] config.h
- [ ] gversion.h

## gversion.h

- [ ] glib/gtypes.h

## gwakeup.c

- [ ] config.h
- [ ] fcntl.h
- [ ] giochannel.h
- [ ] glib-unix.h
- [ ] glib.h
- [ ] gmessages.h
- [ ] gpoll.h
- [ ] gtypes.h
- [ ] gwakeup.h
- [ ] gwin32.h
- [ ] sys/eventfd.h
- [ ] windows.h

## gwakeup.h

- [ ] glib/gpoll.h

## gwin32.c

- [ ] config.h
- [ ] ctype.h
- [ ] direct.h
- [ ] errno.h
- [ ] errno.h
- [ ] fcntl.h
- [ ] glib-init.h
- [ ] glib.h
- [ ] glibconfig.h
- [ ] gthreadprivate.h
- [ ] gwin32-private.c
- [ ] ntddk.h
- [ ] ntdef.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] sys/cygwin.h
- [ ] wchar.h
- [ ] windows.h
- [ ] winternl.h

## gwin32.h

- [ ] glib/gtypes.h

## valgrind.h

- [ ] stdarg.h

## win_iconv.c

- [ ] errno.h
- [ ] fcntl.h
- [ ] io.h
- [ ] stdio.h
- [ ] stdlib.h
- [ ] string.h
- [ ] windows.h

