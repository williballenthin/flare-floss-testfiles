NAME=test-go-garble

BUILDDIR=bin
LINUX=$(NAME)
WINDOWS=$(NAME).exe
WINDOWS64=$(NAME)64.exe

# garble v0.18.0 (go install mvdan.cc/garble@v0.18.0), with Go 1.27.4 on the PATH.
GARBLE?=garble
FLAGS=-literals -seed=HAAAAAAAAAA build -trimpath -ldflags="-s -w -buildid="
dir=@mkdir -p $(@D)

$(BUILDDIR)/$(LINUX): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GARBLE) $(FLAGS) -o $@ .


$(BUILDDIR)/$(WINDOWS): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GARBLE) $(FLAGS) -o $@ .


$(BUILDDIR)/$(WINDOWS64): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GARBLE) $(FLAGS) -o $@ .


clean:
	rm -f $(BUILDDIR)/$(NAME) $(BUILDDIR)/*.exe


all: $(BUILDDIR)/$(LINUX) $(BUILDDIR)/$(WINDOWS) $(BUILDDIR)/$(WINDOWS64)
