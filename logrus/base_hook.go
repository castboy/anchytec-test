/*******************************************************************************
 * // Copyright AnchyTec Corp. All Rights Reserved.
 * // SPDX-License-Identifier: Apache-2.0
 * // Author: shaozhiming
 ******************************************************************************/

package main

import (
	"bytes"
	"fmt"
	"github.com/shopspring/decimal"
	. "github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

type myHook struct {
	level Level
	ch    chan *Entry
	io    io.Writer

	module string
}

func NewMyHook(level Level, module string, io io.Writer) *myHook {
	return &myHook{level: level, ch: make(chan *Entry, 1024), io: io, module: module}
}

func (hook *myHook) Fire(entry *Entry) error {
	entry.Data["module"] = hook.module

	hook.ch <- entry
	return nil
}

func (hook *myHook) Levels() []Level {
	return []Level{
		hook.level,
	}
}

func (hook *myHook) WriteLoop() {
	go func() {
		for {
			entry := <-hook.ch
			//hook.io.WriteString(entry.Data["key"].(string) + "\n")
			b, _ := entry.Logger.Formatter.Format(entry)

			buf := &bytes.Buffer{}
			buf.WriteString(fmt.Sprintf("%+v>>%+v>>%+v||", entry.Caller.File, entry.Caller.Function, entry.Caller.Line))
			buf.Write(b)

			hook.io.Write(buf.Bytes())
		}
	}()
}

type myFormat struct {
}

type MarketRecord struct {
	Market         string          `json:"market"`
	TimestampLp    int64           `json:"timestamp_lp"`
	TimestampBg    int64           `json:"timestamp_bg"`
	Ask            decimal.Decimal `json:"ask"`
	Bid            decimal.Decimal `json:"bid"`
	AskMarketDepth decimal.Decimal `json:"ask_market_depth"`
	BidMarketDepth decimal.Decimal `json:"bid_market_depth"`
}

func (format *myFormat) Format(entry *Entry) ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(`{"level": ` + entry.Level.String() + `,`)

	for i := range entry.Data {
		buf.WriteString(fmt.Sprintf(`"%s": "%+v",`, i, entry.Data[i]))
	}

	buf.WriteString(`"time": ` + entry.Time.UTC().Format("2006-01-02 15:04:05"))
	buf.WriteString("}\n")

	return buf.Bytes(), nil
}

func main() {
	logger := New()

	hk, err := os.OpenFile("/Users/wmg/GolandProject/src/anchytec/logrus/hook", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err != nil {
		log.Fatal(err)
	}

	logger.SetLevel(TraceLevel)
	logger.SetReportCaller(true)

	hook := NewMyHook(ErrorLevel, "MODULE", hk)
	hook.WriteLoop()

	logger.AddHook(hook)

	lg, err := os.OpenFile("/Users/wmg/GolandProject/src/anchytec/logrus/log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err != nil {
		log.Fatal(err)
	}

	logger.SetOutput(lg)
	logger.SetFormatter(&myFormat{}) // `TextFormatter` / `JSONFormatter` lead to panic sometimes; can implement by self;

	//logger.SetFormatter(&JSONFormatter{})
	//logger.SetFormatter(&TextFormatter{})
	go func() {
		value := &MarketRecord{
			Market:         "H33HKD",
			TimestampLp:    1612386031978,
			TimestampBg:    1612386032210,
			Ask:            decimal.NewFromFloat(29112.51),
			Bid:            decimal.NewFromFloat(29112.51),
			AskMarketDepth: decimal.NewFromFloat(29112.51),
			BidMarketDepth: decimal.NewFromFloat(29112.51),
		}

		for {
			logger.WithField("value", fmt.Sprintf("%+v", *value)).Info()
			logger.WithField("value", fmt.Sprintf("%+v", *value)).Error()
		}
	}()

	time.Sleep(10 * time.Second)

}
