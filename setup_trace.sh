#!/bin/sh
set +x
set +e
cd /sys/kernel/debug/tracing
echo 0 > tracing_on
echo $1 > set_ftrace_pid
echo function_graph > current_tracer
