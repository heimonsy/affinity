package affinity

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var mask [1024 / 64]uint64

func SetMask(cpu uint) {
	mask[cpu/64] |= 1 << (cpu % 64)
}

func CallAffinity(pid int) error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_SCHED_SETAFFINITY, uintptr(pid), uintptr(len(mask)*8), uintptr(unsafe.Pointer(&mask)))
	if errno != 0 {
		fmt.Println(errno)
		return errno
	}
	return nil
}

func AffinityService(pid int, duration time.Duration) {
	go func() {
		time.Sleep(time.Second * 2)

		// for list process tasks id
		f, err := os.Open("/proc/" + strconv.Itoa(pid) + "/task")
		if err != nil {
			return
		}
		defer f.Close()

		for {
			time.Sleep(duration)
			tasks, err := f.Readdirnames(0)
			if err != nil {
				continue
			}
			for _, tid := range tasks {
				nTid, err := strconv.Atoi(tid)
				if err != nil {
					continue
				}
				CallAffinity(nTid)
			}
		}
	}()
}
