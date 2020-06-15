package cmd

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Exec Execute a command and collect the output
func Exec(args ...string) (string, error) {

	baseCmd := args[0]
	cmdArgs := args[1:]

	log.Tracef("Exec: %s %s", baseCmd, cmdArgs)

	cmd := exec.Command(baseCmd, cmdArgs...)
	res, err := cmd.CombinedOutput()

	return string(res), err
}

// ExecInteract can execute semi-interactive commands like the new btmgmt executable of BlueZ >= 5.50,
// that in therory should be executable by Exec when command line arguments are provided, but in fact
// seem to misbehave on their cli...
func ExecInteract(args ...string) (string, error) {
	baseCmd := args[0]
	cmdArgs := args[1:]

	log.Tracef("ExecInteract: %s %s", baseCmd, cmdArgs)

	subProcess := exec.Command(baseCmd, cmdArgs...)
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close() // close process afterwards

	// create and attach pipes and bufio-readers for stdout and stderr
	pstdoutr, pstdoutw, _ := os.Pipe()
	biostdoutr := bufio.NewReader(pstdoutr)
	pstderrr, pstderrw, _ := os.Pipe()
	biostderrr := bufio.NewReader(pstderrr)
	subProcess.Stdout = pstdoutw
	subProcess.Stderr = pstderrw

	// start command
	if err = subProcess.Start(); err != nil { //Use start, not run...
		return "", err
	}

	res := ""
	err = nil

	// stdout readline routine
	cline := make(chan string, 1)
	cerr := make(chan error, 1)
	// stderr readline routine
	cerrline := make(chan string, 1)

loop:
	for {
		go func() {
			line, err := biostdoutr.ReadString('\n')
			if err != nil {
				if err.Error() != "EOF" {
					cerr <- err
				}
			} else {
				log.Trace("cmd.ExecInteract('" + baseCmd + "',...) STDOUT: " + line)
				cline <- line
			}
		}()

		go func() {
			line, err := biostderrr.ReadString('\n')
			if err != nil {
				if err.Error() != "EOF" {
					cerr <- err
				}
			} else {
				log.Trace("cmd.ExecInteract('" + baseCmd + "',...) STDERR: " + line)
				cerrline <- line
			}
		}()

		select {
		case newline := <-cline:
			res += newline
		case newerrline := <-cerrline:
			err = errors.New(newerrline)
			return "", errors.New(newerrline)
		case newerr := <-cerr:
			err = newerr
			break loop
		case <-time.After(100 * time.Millisecond):
			break loop
		}
	}

	close(cline)
	close(cerrline)
	close(cerr)

	// Sadly, this seems to be the only way to detect errors with new btmgmt exe... :(
	if baseCmd == "btmgmt" && strings.Contains(res, "Invalid command") {
		return "", errors.New(res)
	}

	return res, nil
}
