package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init_log() {
	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.Out = os.Stdout
}

func logTcping(code int, address string) {
	if code == 0 {
		log.Info("    TCP     OPEN      ", address)
	} else if code == 1 {
		log.Warn("    TCP     CLOSED    ", address)
	} else if code == 2 {
		log.Warn("    TCP     ERROR     ", address)
	}
}

func logPing(dst *net.IPAddr, dur time.Duration, err error) {
	if err != nil {
		match, _ := regexp.MatchString("operation not permitted", err.Error())
		if match {
			log.Warn(fmt.Sprintf("    ICMP    ERROR     Root permission is required."))
		} else {
			log.Warn(fmt.Sprintf("    ICMP    ERROR     %s", dst.String()))
		}
		return
	}
	log.Info(fmt.Sprintf("    ICMP    OPEN      %s    %s ms", dst.String(), strconv.FormatInt(dur.Milliseconds(), 10)))
}

func logMtr(hops []string, address string) {
	for _, h := range hops {
		log.Info("    MTR     ", h)
	}
}
