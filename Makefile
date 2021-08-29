

BINDIR = ./bin
LOGGER_BINARY=$(BINDIR)/logger
LOGGER_SRC=./cmd/logger/logger.go


all:$(LOGGER_BINARY)

$(BINDIR):
	echo $(BINDIR)
	- mkdir $(BINDIR)

$(LOGGER_BINARY): $(BINDIR) $(LOGGER_SRC)
	go build -o $(LOGGER_BINARY) $(LOGGER_SRC)


clean:
	- rm -rf $(BINDIR)
