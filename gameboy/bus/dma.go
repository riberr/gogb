package bus

type dma struct {
	transferInProgress bool
	transferRestarted  bool
	from               int
	ticks              int
	dmaRegister        uint8
}

func newDMA() *dma {
	return &dma{dmaRegister: 0xFF}
}

func (dma *dma) getDmaRegister() uint8 {
	return dma.dmaRegister
}

func (dma *dma) isOAMBlocked() bool {
	return dma.transferRestarted || dma.transferInProgress && dma.ticks >= 5
}

func (dma *dma) setDmaRegister(value uint8) {
	dma.from = int(value) * 0x100
	dma.transferRestarted = dma.isOAMBlocked()
	dma.ticks = 0
	dma.transferInProgress = true
	dma.dmaRegister = value
}

func (dma *dma) Tick(cycles int, bus *Bus) {
	for i := 0; i < cycles; i++ {
		dma.tick(bus)
	}
}

func (dma *dma) tick(bus *Bus) {
	if !dma.transferInProgress {
		return
	}

	// 162 * 4
	dma.ticks++
	if dma.ticks < 648 {
		return
	}

	dma.transferInProgress = false
	dma.transferRestarted = false
	dma.ticks = 0

	for i := uint16(0); i < 0xA0; i++ {
		bus.Write(0xFE00+i, bus.Read(i+uint16(dma.from)))
	}
}
