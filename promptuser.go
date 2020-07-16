package promptuser

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Echo will prompt the user for input and echo their input
func Echo(prompt string) string {
	return base(prompt, true)
}

// NoEcho will prompt the user and hide their input
func NoEcho(prompt string) string {
	return base(prompt, false)
}

func base(prompt string, echo bool) string {
	initialTermState, e1 := terminal.GetState(syscall.Stdin)
	if e1 != nil {
		panic(e1)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	defer signal.Stop(c)
	go func() {
		<-c
		terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	fmt.Print(prompt)

	if echo {
		reader := bufio.NewReader(os.Stdin)
		u, _, _ := reader.ReadLine()
		return string(u)
	} else {
		p, err := terminal.ReadPassword(syscall.Stdin)
		fmt.Println("")
		if err != nil {
			panic(err)
		}
		return string(p)
	}
}
