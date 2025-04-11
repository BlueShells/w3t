package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getSignalName(sig os.Signal) string {
	switch sig {
	case syscall.SIGHUP:
		return "SIGHUP"
	case syscall.SIGINT:
		return "SIGINT"
	case syscall.SIGQUIT:
		return "SIGQUIT"
	case syscall.SIGILL:
		return "SIGILL"
	case syscall.SIGTRAP:
		return "SIGTRAP"
	case syscall.SIGABRT:
		return "SIGABRT"
	case syscall.SIGBUS:
		return "SIGBUS"
	case syscall.SIGFPE:
		return "SIGFPE"
	case syscall.SIGUSR1:
		return "SIGUSR1"
	case syscall.SIGSEGV:
		return "SIGSEGV"
	case syscall.SIGUSR2:
		return "SIGUSR2"
	case syscall.SIGPIPE:
		return "SIGPIPE"
	case syscall.SIGALRM:
		return "SIGALRM"
	case syscall.SIGTERM:
		return "SIGTERM"
	case syscall.SIGCHLD:
		return "SIGCHLD"
	case syscall.SIGCONT:
		return "SIGCONT"
	case syscall.SIGSTOP:
		return "SIGSTOP"
	case syscall.SIGTSTP:
		return "SIGTSTP"
	case syscall.SIGTTIN:
		return "SIGTTIN"
	case syscall.SIGTTOU:
		return "SIGTTOU"
	case syscall.SIGURG:
		return "SIGURG"
	case syscall.SIGXCPU:
		return "SIGXCPU"
	case syscall.SIGXFSZ:
		return "SIGXFSZ"
	case syscall.SIGVTALRM:
		return "SIGVTALRM"
	case syscall.SIGPROF:
		return "SIGPROF"
	case syscall.SIGWINCH:
		return "SIGWINCH"
	case syscall.SIGIO:
		return "SIGIO"
	case syscall.SIGSYS:
		return "SIGSYS"
	default:
		return fmt.Sprintf("UNKNOWN(%v)", sig)
	}
}

func checkProcess() bool {
	// 检查进程1是否还在运行
	process, err := os.FindProcess(1)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func main() {
	// 创建一个 channel 用于接收信号
	sigs := make(chan os.Signal, 1)

	// 定义要监听的信号
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

	fmt.Printf("Signal logger started at %s. Waiting for signals...\n", time.Now().Format("2006-01-02 15:04:05.000000000"))

	// 无限循环接收信号
	for {
		sig := <-sigs
		signalName := getSignalName(sig)
		fmt.Printf("Received %s at %s\n", signalName, time.Now().Format("2006-01-02 15:04:05.000000000"))

		// 如果是终止信号，开始监控进程状态
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			for {
				if !checkProcess() {
					fmt.Printf("Container was forcefully terminated at %s\n", time.Now().Format("2006-01-02 15:04:05.000000000"))
					os.Exit(1)
				}
				time.Sleep(time.Second)
			}
		}
	}
}
