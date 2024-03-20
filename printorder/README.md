## Go Routine Execution Order with Channels

This code demonstrates a controlled execution order for three goroutines (`first`, `second`, `third`) using channels for synchronization. The output will always be "firstsecondthird" regardless of the order in which the goroutines are launched.

## Functionality

1. **Channels:** Two channels, `f` and `s` of type `bool`, are created to control the sequence of execution.
2. **Wait Group:** A `sync.WaitGroup` (`wg`) is used to track the completion of all goroutines.
3. **`first` Routine:** This goroutine prints "first" and then sends `true` to the `f` channel, signaling the `second` routine to proceed.
4. **`second` Routine:** Waits to receive `true` from the `f` channel before printing "second".  It then sends `true` to the `s` channel, allowing the `third` routine to execute.
5. **`third` Routine:** Waits to receive `true` from the `s` channel before printing "third".
6. **`main` Function:**
    - Initializes `wg`, `f`, and `s`.
    - Increments `wg` by 3 for the three goroutines.
    - Launches `first`, `third`, and `second` as Go routines, but in any order. This demonstrates the controlled execution despite launch order.
    - Waits for all goroutines to finish using `wg.Wait()`.

## Key Points

- Channels (`f` and `s`) ensure that the goroutines execute in a specific order ("first", "second", "third").
- The order of launching the goroutines in `main` doesn't affect the final output due to the synchronization with channels.

## Considerations

- This is a simple example demonstrating channel-based synchronization.
- Error handling is not explicitly implemented here. You can extend the code to catch and handle potential errors.
