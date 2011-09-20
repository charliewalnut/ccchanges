include $(GOROOT)/src/Make.inc

TARG=ccchanges
GOFILES=\
	parse_entry.go\
	parse_log.go\
	ccchanges.go\

include $(GOROOT)/src/Make.cmd
