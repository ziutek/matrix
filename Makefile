include $(GOROOT)/src/Make.inc

GC = $Og -N
#GC = $Og -B
TARG=matrix

#OFILES_amd64=\

OFILES=\
	$(OFILES_$(GOARCH))

ALLGOFILES=\
	   matrix.go\
	   utils.go\
	   add.go\

NOGOFILES=\
	$(subst _$(GOARCH).$O,.go,$(OFILES_$(GOARCH)))

GOFILES=\
	$(filter-out $(NOGOFILES),$(ALLGOFILES))\
	$(subst .go,_decl.go,$(NOGOFILES))\

include $(GOROOT)/src/Make.pkg
