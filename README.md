# Golang Pulse

This module implements the same functionality as `time.Ticker` but sends an initial pulse down the `C` channel after calling `New()`.
Useful for functions that run on an interval and need to start before the first tick is sent.
