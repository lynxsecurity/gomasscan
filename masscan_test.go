package gomasscan

import (
	"context"
	"log"
	"testing"
)

func TestMasscanCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := New()
	m.SetContext(ctx)
	m.SetPorts("443")
	m.SetMasscanOutfile("masscan.out")
	m.SetParsedOutfile("parsed.out")
	m.SetRanges("rangeHere")
	m.SetRate("3000")
	m.SetExclude("127.0.0.1")

	err := m.Run()
	if err != nil {
		log.Println("scanner failed:", err)
		return
	}
	err = m.Parse()
	if err != nil {
		log.Println("parsing failed:", err)
		return
	}
	m.Clean()
}
