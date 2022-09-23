#!/bin/bash

make

Q="TEST1"
echo "#### ${Q} 1 ####"
./ral2jqlog-f
if [[ $? -eq 1 ]]; then
	echo "${Q} OK"
else
	echo "${Q} NG"
	exit 1
fi

Q="TEST2"
echo "#### ${Q} 1 ####"
./ral2jqlog-f -s aaa
if [[ $? -eq 1 ]]; then
	echo "${Q} OK"
else
	echo "${Q} NG"
	exit 1
fi

Q="TEST3"
echo "#### ${Q} 1 ####"
./ral2jqlog-f -s aaa -d bbb
if [[ $? -eq 1 ]]; then
	echo "${Q} OK"
else
	echo "${Q} NG"
	exit 1
fi

Q="TEST4"
echo "#### ${Q} 1 ####"
./ral2jqlog-f -s aaa -d bbb -r ccc
if [[ $? -eq 1 ]]; then
	echo "${Q} OK"
else
	echo "${Q} NG"
	exit 1
fi

Q="TEST5"
echo "#### ${Q} 1 ####"
./ral2jqlog-f -s aaa -d bbb -r ccc --date 2022
if [[ $? -eq 1 ]]; then
	echo "${Q} OK"
else
	echo "${Q} NG"
	exit 1
fi


make clean
