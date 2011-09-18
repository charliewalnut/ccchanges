include $(GOROOT)/src/Make.inc

TARG=ccchanges
GOFILES=\
	ccchanges.go\
	parse_entry.go\

include $(GOROOT)/src/Make.pkg
