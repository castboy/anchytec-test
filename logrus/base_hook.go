/*******************************************************************************
 * // Copyright AnchyTec Corp. All Rights Reserved.
 * // SPDX-License-Identifier: Apache-2.0
 * // Author: shaozhiming
 ******************************************************************************/

package main

import (
	"fmt"
	. "github.com/sirupsen/logrus"
	"time"
)

type myHook struct {
	level Level
}

func NewMyHook(level Level) Hook {
	return &myHook{level: level}
}

func (hook *myHook) Fire(entry *Entry) error {
	fmt.Printf("%+v\n", *entry)
	time.Sleep(5 * time.Second)
	return nil
}

func (hook *myHook) Levels() []Level {
	return []Level{
		hook.level,
	}
}

func main() {
	fmt.Println("service beginning...")

	hook := NewMyHook(InfoLevel)
	hook2 := NewMyHook(WarnLevel)

	logger := New()

	logger.AddHook(hook)
	logger.AddHook(hook2)

	logger.WithField("wow", "wmg").Info()
	logger.WithField("wow", "wmg").Warn()
}
