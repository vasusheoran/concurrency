## Rate Limiter

This package, `ratelimiter`, provides a thread-safe implementation for rate limiting functions. It offers a clean interface (`RateLimiter`) to manage request limits within a defined window.

## Functionality

1. **`RateLimiter` Interface:**
   - Defines a single method `RateLimiter(func(ID int), int)` which takes a function to execute and an identifier (optional).

2. **`limiter` Struct:**
   - `wg`: A `sync.WaitGroup` for tracking goroutine completion (optional, can be provided during creation).
   - `m`: A `sync.Mutex` ensures synchronized access to the internal state of the rate limiter.
   - `attempts`: Tracks the current number of requests within the `duration` window.
   - `maxAttempts`: Defines the maximum allowed requests within the `duration` window.
   - `duration`: Represents the time window (in seconds) for tracking attempts.
   - `nextReset`: Stores the timestamp for the next reset of the `attempts` counter.

3. **`NewRateLimiter` Function:**
   - Takes a `sync.WaitGroup` (optional), `maxAttempts`, and `duration` as arguments.
   - Creates a new `limiter` instance with the provided parameters and returns it as a `RateLimiter` interface.

4. **`RateLimiter` Method (on `limiter`):**
   - Takes a function (`fn`) to be executed and an optional identifier (`ID`) as arguments.
   - Acquires a lock on the `m` mutex.
   - If this is the first attempt within the window:
      - Starts a timer for the `duration` with `watchTimer`.
   - Checks if the current attempts exceed the `maxAttempts`:
      - Prints a message indicating exceeding the limit and the time remaining until reset.
      - Increments `attempts` for tracking.
      - Releases the lock and returns (doesn't execute the function).
   - Executes the provided function (`fn`) with the `ID` argument.
   - Increments `attempts`.
   - Releases the lock on `m`.

5. **`watchTimer` Method (on `limiter`):**
   - Takes a `time.Timer` as an argument.
   - Signals completion within the provided `wg` (if provided during `NewRateLimiter`).
   - Adds 1 to the `wg` (if provided) to track completion.
   - Waits for the timer to expire (signaling the end of the `duration` window).
   - Acquires the lock on the `m` mutex.
   - Resets the `attempts` counter to 0.
   - Releases the lock on `m`.

## Usage

1. Import the `ratelimiter` package in your Go program.
2. Create a `sync.WaitGroup` (optional) for tracking goroutine completion if needed.
3. Use the `NewRateLimiter` function to create a new rate limiter instance with desired `maxAttempts` and `duration`.
4. Call the `RateLimiter` method on the rate limiter instance, passing the function you want to rate limit and an optional identifier.

This package provides a basic building block for rate limiting functionalities in your Go applications.

## Considerations

- This implementation utilizes a single timer for all rate limiters. Consider using a map or a channel for managing multiple timers with different durations.
- Error handling is not explicitly implemented here. You can extend the code to handle potential errors.
