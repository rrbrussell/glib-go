# glib-go: A pure go implementation of glib.

This is designed around glib-2.62.4. I am not planning this library as a library that would allow
an automated conversion of glib client C code to Golang.
All of the implemented names have been changed to match Golang's requirements and recommendations.
Several of the functions that accepted a pointer argument so they could return an error have been
converted to use Golang's error types and return multiple items.

