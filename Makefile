

BINDIR = ./bin
LOGGER_BINARY=$(BINDIR)/logger
LOGGER_SRC=./cmd/logger/logger.go

LOADER_BINARY=$(BINDIR)/loader
LOADER_SRC=./cmd/loader/loader.go

REAL_LOGGER_BINARY=$(BINDIR)/reallogger
REAL_LOGGER_SRC=./cmd/reallogger/reallogger.go

GO_SRC=./bb/*.go ./db/*.go ./common/*.go

all:$(LOGGER_BINARY)

$(BINDIR):
	echo $(BINDIR)
	- mkdir $(BINDIR)

$(LOGGER_BINARY): $(BINDIR) $(LOGGER_SRC) $(LOADER_SRC) $(REAL_LOGGER_SRC) $(GO_SRC)
	go build -o $(LOGGER_BINARY) $(LOGGER_SRC)
	go build -o $(LOADER_BINARY) $(LOADER_SRC)
	go build -o $(REAL_LOGGER_BINARY) $(REAL_LOGGER_SRC)

clean:
	- rm -rf $(BINDIR)


cloc:
	cloc .

restartdb:
		brew services stop influxdb
	 	brew services start influxdb


deletebucket:
	! influx bucket delete -n btc -o bb

createbucket:
	influx bucket create -n btc -o bb


download:
	wget -O ./TEST_DATA/BTCUSD2021-08-29.csv.gz https://public.bybit.com/trading/BTCUSD/BTCUSD2021-08-29.csv.gz


deletedata:
	influx delete --bucket btc --start 1970-01-01T00:00:00Z --stop 2021-09-10T18:54:09Z


