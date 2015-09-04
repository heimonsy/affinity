#include "affinity_linux.h"

cpu_set_t ca_set;

void ca_init() {
  CPU_ZERO(&ca_set);
}

void ca_set_mask(int m) {
   CPU_SET(m, &ca_set);
}

int ca_call_affinity(int pid) {
   return sched_setaffinity(pid, sizeof(ca_set), &ca_set);
}
