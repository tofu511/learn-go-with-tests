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

type DefaultSleeper struct {

}

func (d *DefaultSleeper) Sleep()  {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}