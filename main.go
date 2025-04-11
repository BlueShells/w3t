package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个 channel 用于接收信号
	sigs := make(chan os.Signal, 1)

	// 定义要监听的信号（几乎所有可捕捉的信号）
	signal.Notify(sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGILL,
		syscall.SIGTRAP,
		syscall.SIGABRT,
		syscall.SIGBUS,
		syscall.SIGFPE,
		syscall.SIGUSR1,
		syscall.SIGSEGV,
		syscall.SIGUSR2,
		syscall.SIGPIPE,
		syscall.SIGALRM,
		syscall.SIGTERM,
		syscall.SIGCHLD,
		syscall.SIGCONT,
		syscall.SIGSTOP,
		syscall.SIGTSTP,
		syscall.SIGTTIN,
		syscall.SIGTTOU,
		syscall.SIGURG,
		syscall.SIGXCPU,
		syscall.SIGXFSZ,
		syscall.SIGVTALRM,
		syscall.SIGPROF,
		syscall.SIGWINCH,
		syscall.SIGIO,
		syscall.SIGSYS,
	)

	fmt.Println("Signal logger started. Waiting for signals...")

	// 无限循环接收信号
	for {
		sig := <-sigs
		fmt.Printf("Received signal: %v\n", sig)
	}
}
