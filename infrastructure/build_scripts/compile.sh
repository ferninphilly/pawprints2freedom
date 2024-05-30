#!/bin/bash
# loop through each lambda directory in each stack.  In each lambda sub directory compile and zip payload.
if [ -d $1 ]; then
	ARTIFACTS_DIR=../../bin/
	for dir in $1/*/;
		do		
			FILE_NAME=$ARTIFACTS_DIR$(basename "$dir")/bootstrap;
			echo "About to compile $FILE_NAME"
		 	cd $dir && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags=-buildid="" -o $FILE_NAME main.go && chmod 755 $FILE_NAME
 		done
fi
