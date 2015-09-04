package affinity

/*
#include "affinity_linux.h"       // for free()
*/
import "C"
import (
	"os"
	"strconv"
	"time"
)

func init() {
	C.ca_init()
}

func SetMask(m int) {
	C.ca_set_mask(C.int(m))
}

func CallAffinity(pid int) bool {
	a := C.ca_call_affinity(C.int(pid))
	return a != -1
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
