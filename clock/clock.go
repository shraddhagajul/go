package clock
import("fmt")
type Clock int
func New(hour, minute int) Clock{
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock(time)
}

/*String returns a clock in digital form hh:mm.*/
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

/*Add move the time by a number of minutes.*/
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

func (c Clock) Subtract(minutes int) Clock{
	return New(0, int(c)-minutes)
}