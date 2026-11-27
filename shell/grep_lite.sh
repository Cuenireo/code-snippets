#!/bin/bash
# 简易 grep：在文件里查找包含关键字的行，输出行号和内容。
#
# 用法: bash grep_lite.sh <关键字> <文件路径> [-i 忽略大小写]
ignore=""
if [ "$3" = "-i" ]; then
  ignore="-i"
fi
if [ $# -lt 2 ]; then
  echo "用法: $0 <关键字> <文件路径> [-i]"
  exit 1
fi
grep -n $ignore -- "$1" "$2"
