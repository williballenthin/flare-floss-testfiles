NAME=test-go-stackstrings

BUILDDIR=bin
LINUX=$(NAME)
WINDOWS=$(NAME).exe
WINDOWS64=$(NAME)64.exe

GO?=go
FLAGS=-trimpath -ldflags="-s -w -buildid="
dir=@mkdir -p $(@D)

$(BUILDDIR)/$(LINUX): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(FLAGS) -o $@ .


$(BUILDDIR)/$(WINDOWS): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GO) build $(FLAGS) -o $@ .


$(BUILDDIR)/$(WINDOWS64): main.go go.mod
	$(dir)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build $(FLAGS) -o $@ .


clean:
	rm -f $(BUILDDIR)/$(NAME) $(BUILDDIR)/*.exe


all: $(BUILDDIR)/$(LINUX) $(BUILDDIR)/$(WINDOWS) $(BUILDDIR)/$(WINDOWS64)
