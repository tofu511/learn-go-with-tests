package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type CountdownOperationsSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s * CountdownOperationsSpy) Write(p []byte) (n int, err error)  {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

const write  = "write"
const sleep  = "sleep"

const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer, sleeper Sleeper)  {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main()  {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}