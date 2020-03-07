#!/bin/sh

fmt_cmd="go fmt ./..."

echo "$fmt_cmd"
fmt_res=`$fmt_cmd`

if [ "$fmt_res" == "" ]; then
  echo "All right"
  exit 0
else
  echo "Find not formatted files:"
  echo $fmt_res
  exit 1
fi