#define _GNU_SOURCE

#include <stdio.h>
#include <sched.h>


// cpu affinity init
void ca_init();

// set cpu affinity mask
void ca_set_mask(int);

int ca_call_affinity(int pid);
