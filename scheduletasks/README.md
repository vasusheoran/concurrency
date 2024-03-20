## Go Package for Scheduling Tasks

This package, `scheduletasks`, provides functionalities for scheduling tasks to run at specific times or with fixed intervals in Go applications.

## Functionality

1. **Command Interface:**
    - Defines a single method `Run(cmd string)` that needs to be implemented by the actual command object you want to schedule. This method takes the command string as an argument and is responsible for executing the desired action.

2. **`taskScheduler` Struct:**
    - Holds a `Command` interface object, which represents the command to be executed when a task is scheduled.

3. **`Schedule` Method:**
    - Takes three arguments:
        - `cmd` (string): The command string to be passed to the `Run` method of the `Command` interface.
        - `inDelay` (int64): The initial delay before executing the task.
        - `unit` (time.Duration): The unit of time for the `inDelay` argument (e.g., `time.Second`, `time.Nanosecond`).
    - Creates a new `time.Timer` with the specified delay and unit.
    - Waits for the timer to expire using a channel receive operation (`<-t.C`).
    - Once the timer expires, calls the `Run` method of the `Command` interface with the provided `cmd` string.

4. **`ScheduleWithFixedDelay` Method:**
    - Takes four arguments:
        - `cmd` (string): The command string to be passed to the `Run` method of the `Command` interface.
        - `inDelay` (int64): The initial delay before the first execution.
        - `unit` (time.Duration): The unit of time for the `inDelay` and `delay` arguments (e.g., `time.Second`, `time.Nanosecond`).
        - `delay` (int): The delay between subsequent executions of the task.
    - Calls the `Schedule` method with the provided arguments to execute the task after the initial delay.
    - Enters an infinite loop:
        - Sleeps for the specified `delay` using the provided `unit`.
        - Calls the `Schedule` method again with a minimal delay (1 unit) to trigger the next execution of the command.

5. **`ScheduleAtFixedRate` Method:**
    - Takes four arguments:
        - `cmd` (string): The command string to be passed to the `Run` method of the `Command` interface.
        - `inDelay` (int64): The initial delay before the first execution.
        - `unit` (time.Duration): The unit of time for the `inDelay` and `period` arguments (e.g., `time.Second`, `time.Nanosecond`).
        - `period` (int): The interval between task executions.
    - Creates a new `time.Timer` with the specified initial delay and unit.
    - Waits for the timer to expire using a channel receive operation (`<-timer.C`).
    - Creates a new `time.Ticker` with the specified `period` and unit.
    - Starts a goroutine:
        - The goroutine continuously waits on the `time.Ticker` channel (`<-ticker.C`).
        - Upon receiving a tick from the ticker, it calls the `Run` method of the `Command` interface in a separate goroutine to avoid blocking the main loop.

## Usage

1. Import the `scheduletasks` package in your Go program.
2. Implement the `Command` interface by creating a concrete struct that defines the actual functionality you want to execute when a task is scheduled. The `Run` method of this struct should take the command string as input and perform the desired action.
3. Create a new `taskScheduler` instance.
4. Use the appropriate scheduling method (`Schedule`, `ScheduleWithFixedDelay`, or `ScheduleAtFixedRate`) of the `taskScheduler` object to schedule your task execution based on your desired timing needs. Remember to provide the command string, initial delay, time unit, and additional parameters (delay for `ScheduleWithFixedDelay` and period for `ScheduleAtFixedRate`) as arguments.

This package provides a simple and flexible way to schedule tasks in your Go applications using timers and goroutines. 
