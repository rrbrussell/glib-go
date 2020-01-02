# glib-go: A pure go implementation of glib

This is designed around glib-2.62.4. Currently the goal is minimal conversion of glib into Golang.
Language specific optimizations should be internally only right now.

The long range plan is to get enough of this and the GTK2 stack ported to Golang that I can write
Golang apps with a fully supported GTK2 based gui. The reason for targeting GTK2 instead of GTK3
is a combination of design decisions made in GTK3 and stability of the API. The really longterm plan
is to get this code polished well enough that the GTK2 versions of Mate, XFCE, or maybe GNOME 2 itself
can be converted into Golang and work on Linux and maybe Windows.

## glib/gstdio.go

Implements:

- [ ] gstdio.h
- [ ] gstdio.c
- [ ] gstdioprivate.h
- [ ] gstdio-private.c
