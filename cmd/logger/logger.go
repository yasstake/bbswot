package main

import (
	"bbswot/bb"
	"flag"
	"log"
)

func main() {
	var log_dir = flag.String("log_dir", "/tmp/BB", "log store directory")
	var flag_file = flag.String("flag_file", "", "flag file name, if not specified no flag file used.")
	var exit_wait = flag.Int("exit_wait", 0, "Exit wait minute, when terminated by peer process")
	var enable_compress = flag.Bool("compress", true, "Enable log differential compress mode")

	flag.Parse()

	log.Printf("[LOG DIR]   %s", *log_dir)
	log.Printf("[FLAG FILE] %s", *flag_file)
	log.Printf("[EXIT WAIT] %d", *exit_wait)
	if *enable_compress {
		log.Printf("[Enable Compress]")
		bb.EnableLogCompress()
	}

	writer := bb.CreateWriter(*log_dir)
	bb.Connect(*flag_file, writer, *exit_wait)
}
