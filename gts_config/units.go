package gts_config

import (
	"strconv"
	"strings"
	"regexp"
	"time"
	"github.com/pivotal-golang/bytefmt"
)


/***********************************************************************************************************************
 * Units
 */

type Size struct {
	Bytes uint64
}

func (self *Size) UnmarshalText(text []byte) error {
	var err error
	self.Bytes, err = bytefmt.ToMegabytes(string(text))
	self.Bytes *= bytefmt.MEGABYTE
	return err
}

type Duration struct {
	time.Duration
}

func (self *Duration) UnmarshalText(text []byte) error {
	var err error
	self.Duration, err = time.ParseDuration(string(text))
	return err
}

type Regex struct {
	Rex *regexp.Regexp
}

func (self *Regex) UnmarshalText(text []byte) error {
	t := string(text)
	var err error
	self.Rex, err = regexp.Compile(t)
	return err
}

type Levels struct {
	Levels []int
}

func (self *Levels) UnmarshalText(text []byte) error {
	arr := strings.Split(string(text), ":")
	var err error
	for _, t := range (arr) {
		i, err := strconv.Atoi(t)
		if nil != err {
			return err
		}
		self.Levels = append(self.Levels, i)
	}
	return err
}
