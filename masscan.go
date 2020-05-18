package gomasscan

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Masscan struct {
	Ctx            context.Context
	BinaryPath     string
	Args           []string
	Exclude        string
	ExcludedFile   string
	Ports          string
	Ranges         string
	Rate           string
	MasscanOutfile string
	ParsedOutfile  string
	InputFile      string
	Result         []byte
}

func (m *Masscan) SetBinaryPath(Path string) {
	m.BinaryPath = Path
}
func (m *Masscan) SetContext(Ctx context.Context) {
	m.Ctx = Ctx
}
func (m *Masscan) SetInputFile(File string) {
	m.InputFile = File
}
func (m *Masscan) SetMasscanOutfile(File string) {
	m.MasscanOutfile = File
}
func (m *Masscan) SetParsedOutfile(File string) {
	m.ParsedOutfile = File
}
func (m *Masscan) SetArgs(arg ...string) {
	m.Args = arg
}
func (m *Masscan) SetPorts(ports string) {
	m.Ports = ports
}
func (m *Masscan) SetRanges(ranges string) {
	m.Ranges = ranges
}

func (m *Masscan) SetRate(rate string) {
	m.Rate = rate
}
func (m *Masscan) SetExclude(exclude string) {
	m.Exclude = exclude
}
func (m *Masscan) SetExcludedFile(excluded string) {
	m.ExcludedFile = excluded
}
func (m *Masscan) Run() error {
	var cmd *exec.Cmd
	var outb, errs bytes.Buffer
	_, err := os.Stat(m.BinaryPath)
	if err != nil {
		return fmt.Errorf("masscan could not be run: %v", err)
	}
	if m.Rate != "" {
		m.Args = append(m.Args, "--rate")
		m.Args = append(m.Args, m.Rate)
	}
	if m.ExcludedFile != "" {
		m.Args = append(m.Args, "--excludefile")
		m.Args = append(m.Args, m.ExcludedFile)
	}
	if m.Ranges != "" {
		m.Args = append(m.Args, "--range")
		m.Args = append(m.Args, m.Ranges)
	}
	if m.InputFile != "" {
		m.Args = append(m.Args, "-iL")
		m.Args = append(m.Args, m.InputFile)
	}
	if m.Ports != "" {
		m.Args = append(m.Args, "-p")
		m.Args = append(m.Args, m.Ports)
	}
	if m.Exclude != "" {
		m.Args = append(m.Args, "--exclude")
		m.Args = append(m.Args, m.Exclude)
	}
	m.Args = append(m.Args, "-oL")
	m.Args = append(m.Args, m.MasscanOutfile)

	if m.Ctx == nil {
		m.Ctx = context.TODO()
	}

	cmd = exec.CommandContext(m.Ctx, m.BinaryPath, m.Args...)
	cmd.Stdout = &outb
	cmd.Stderr = &errs
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		if errs.Len() > 0 {
			return errors.New(errs.String())
		}
		return err
	}
	m.Result = outb.Bytes()
	return nil
}
func (m *Masscan) Parse() error {
	_, err := os.Stat(m.MasscanOutfile)
	if err != nil {
		return err
	}
	f, err := os.Create(m.ParsedOutfile)
	if err != nil {
		return err
	}
	defer f.Close()
	ifp, err := os.Open(m.MasscanOutfile)
	if err != nil {
		return err
	}
	defer ifp.Close()
	scanner := bufio.NewScanner(ifp)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		if strings.HasPrefix(scanner.Text(), "open") {
			dat := strings.Split(scanner.Text(), " ")
			if len(dat) > 3 {
				f.WriteString(fmt.Sprintf("%s:%s\n", dat[3], dat[2]))
			}
		}
	}

	return nil
}
func (m *Masscan) Clean() error {
	_, err := os.Stat(m.MasscanOutfile)
	if err != nil {
		return err
	}
	err = os.Remove(m.MasscanOutfile)
	if err != nil {
		return err
	}
	return nil
}
func New() *Masscan {
	return &Masscan{
		BinaryPath: "/usr/local/bin/masscan",
	}
}
