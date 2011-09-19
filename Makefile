include $(GOROOT)/src/Make.inc

TARG=bin/ccchanges
GOFILES=\
	parse_entry.go\
	parse_log.go\
	main.go\

include $(GOROOT)/src/Make.cmd
