package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func sigHandler() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			syscall.SIGINT,  // Ctrl+C
			syscall.SIGTERM, // Termination Request
			syscall.SIGSEGV, // FullDerp
			syscall.SIGABRT, // Abnormal termination
			syscall.SIGILL,  // illegal instruction
			syscall.SIGFPE)  // floating point - this is why we can't have nice things
		sig := <-c
		log.Printf("Signal (%v) Detected, Shutting Down", sig)
	}()
}
