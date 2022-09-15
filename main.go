package main

import (
	"fmt"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
)

func main() {
	sshConfig := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Start connection to network device
	s, err := netconf.DialSSH("192.168.0.1", sshConfig)

	if err != nil {
		fmt.Printf("SSH response indicate problem: %v\n", err)

		if s != nil {
			defer s.Close()
		}
	} else {
		// Close connetion to device when this function is done executing
		defer s.Close()

		methodToExec := netconf.RawMethod("<get/>")

		startTimeGetConf := time.Now().UnixNano()

		// Request data from switch
		reply, err := s.Exec(methodToExec)

		fmt.Printf("Time to get response without session creation is: %v\n", time.Now().UnixNano()-startTimeGetConf)

		if err != nil {
			fmt.Printf("Failed response: %v\n", err)
		} else {
			fmt.Printf("Response from switch: %v\n", reply.Data)
		}
	}

	// Keep service alive, to not have it restart constantly
	for {
		time.Sleep(time.Second * 10)
	}
}
