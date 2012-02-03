include $(GOROOT)/src/Make.inc

GC = $Og -N
#GC = $Og -B
TARG=github.com/ziutek/matrix

#OFILES_amd64=\

OFILES=\
	$(OFILES_$(GOARCH))

ALLGOFILES=\
	   dense.go\
	   dmulby.go\

NOGOFILES=\
	$(subst _$(GOARCH).$O,.go,$(OFILES_$(GOARCH)))

GOFILES=\
	$(filter-out $(NOGOFILES),$(ALLGOFILES))\
	$(subst .go,_decl.go,$(NOGOFILES))\

include $(GOROOT)/src/Make.pkg
