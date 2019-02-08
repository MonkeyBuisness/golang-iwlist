all: rpi3

deployDir  = ./deploy
deployFile = wlist

rpi3:
	rm -rf ${deployDir}/*
	GOOS=linux GOARCH=arm GOARM=5 go build -o ${deployDir}/${deployFile} ./util/wlist_util.go
	chmod +x ${deployDir}/${deployFile}