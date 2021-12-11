package helper

import "time"

var (
	helperName string
	Desc       string
)

// An identifier is exported if it begins with a capital letter.

type StopWatch struct {
	start   time.Time // this is private variable of struct
	stop    time.Time
	total   time.Duration
	running bool
	Desc    string // this variable is public
}

func (s *StopWatch) Start() {
	if !s.running {
		s.start = time.Now()
		s.running = true
	}
}

func (s *StopWatch) Stop() {
	if s.running {
		s.stop = time.Now()
		s.running = false
	}
}

// GetPassingTime get duration (second)
func (s *StopWatch) GetPassingTime() float64 {
	s.total = s.stop.Sub(s.start)

	return s.total.Seconds()
}

func GetHelperName() string {
	setHelperName("Loggo Helper")
	return helperName
}

func setHelperName(name string) {
	helperName = name
}
