NAME=test-go-gobfuscate

BUILDDIR=bin
LINUX=$(NAME)
WINDOWS=$(NAME).exe
WINDOWS64=$(NAME)64.exe

# GO: the Go toolchain that builds the obfuscated program (here, 1.27.4).
# GOBFUSCATE: gobfuscate, built with Go 1.19 or older (see build.sh):
#   GO111MODULE=on go1.17.13 install github.com/unixpickle/gobfuscate@13796335c727f53349ac9c3aa3d2fc812fa8c2d4
GO?=go
GOBFUSCATE?=gobfuscate
export GOBFUSCATE
dir=@mkdir -p $(@D)

$(BUILDDIR)/$(LINUX): main.go build.sh
	$(dir)
	./build.sh $(GO) $@ linux amd64


$(BUILDDIR)/$(WINDOWS): main.go build.sh
	$(dir)
	./build.sh $(GO) $@ windows 386


$(BUILDDIR)/$(WINDOWS64): main.go build.sh
	$(dir)
	./build.sh $(GO) $@ windows amd64


clean:
	rm -f $(BUILDDIR)/$(NAME) $(BUILDDIR)/*.exe


all: $(BUILDDIR)/$(LINUX) $(BUILDDIR)/$(WINDOWS) $(BUILDDIR)/$(WINDOWS64)
