package pulse

import (
	"time"
)

type Pulse struct {
	c      chan time.Time
	C      <-chan time.Time
	ticker *time.Ticker
}

func (p *Pulse) Reset(d time.Duration) {
	p.ticker.Reset(d)
}

func (p *Pulse) Stop() {
	p.ticker.Stop()
	close(p.c)
}

func (p *Pulse) tick() {
	for pulse := range p.ticker.C {
		p.c <- pulse
	}
}

func New(d time.Duration) *Pulse {
	c := make(chan time.Time, 1)
	t := time.NewTicker(d)

	//Send the initial pulse
	c <- time.Now()

	p := &Pulse{
		c:      c,
		C:      c,
		ticker: t,
	}

	//Keep retrieving ticks from ticker
	go p.tick()

	return p
}
