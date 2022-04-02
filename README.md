# Golang Pulse

This module implements the same functionality as `time.Ticker` but sends an initial pulse down the `C` channel after calling `NewPulse()`
Useful if you want a function to run on an interval, as well as when starting
