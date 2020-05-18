# gomasscan
a golang wrapper around masscan, based on https://github.com/dean2021/go-masscan/

## example:
```golang

m := gomasscan.New()

ctx, cancel := context.WithCancel(context.Background())
defer cancel()
m.SetPorts("0-65535")
m.SetContext(ctx)
m.SetMasscanOutfile("masscan.out")
m.SetParsedOutfile("parsed.out")
m.SetRanges("10.0.0.1")
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
```