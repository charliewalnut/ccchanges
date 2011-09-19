include $(GOROOT)/src/Make.inc

TARG=bin/ccchanges
GOFILES=\
	parse_entry.go\
	parse_log.go\
	ccchanges.go\

include $(GOROOT)/src/Make.cmd
