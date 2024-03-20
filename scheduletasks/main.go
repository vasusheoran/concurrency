package scheduletasks

import "time"

type Command interface {
	Run(cmd string)
}

type taskScheduler struct {
	Command Command
}

// Schedule creates and executes a one-shot action that becomes enabled after the given delay.
func (ts taskScheduler) Schedule(cmd string, inDelay int64, unit time.Duration) {
	//inDelay = 10 time
	t := time.NewTimer(time.Duration(inDelay) * unit)
	<-t.C
	ts.Command.Run(cmd)
}

// ScheduleWithFixedDelay creates and executes a periodic action that becomes enabled first after the given initial delay, and
// subsequently with the given period; that is executions will commence after initialDelay then
// initialDelay+period, then initialDelay + 2 * period, and so on.
func (ts taskScheduler) ScheduleWithFixedDelay(cmd string, inDelay int64, unit time.Duration, delay int) {
	ts.Schedule(cmd, inDelay, unit)
	for {
		// call next after this delay
		time.Sleep(time.Duration(delay) * unit)
		ts.Schedule(cmd, 1, unit)
	}
}

// ScheduleAtFixedRate creates and executes a periodic action that becomes enabled first after the given initial delay, and
// subsequently with the given delay between the termination of one execution and the commencement of the next.
func (ts taskScheduler) ScheduleAtFixedRate(cmd string, inDelay int64, unit time.Duration, period int) {
	timer := time.NewTimer(time.Duration(inDelay) * unit)

	<-timer.C

	ticker := time.NewTicker(time.Duration(period) * unit)
	defer ticker.Stop()

	go func() {

		for {
			select {
			case <-ticker.C:
				go ts.Command.Run(cmd)

			}
		}

	}()
}
