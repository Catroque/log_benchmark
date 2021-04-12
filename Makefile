

all: logrus zerolog


logrus:
	cd internal/logrus/; go test -gcflags "-m" -run "" -bench . -benchmem -memprofile mem.out -cpuprofile cpu.out 2>&1 | tee report.txt

zerolog:
	cd internal/zerolog/; go test -gcflags "-m" -run "" -bench . -benchmem -memprofile mem.out -cpuprofile cpu.out 2>&1 | tee report.txt

gologwrp:
	cd internal/gologwrp/; go test -gcflags "-m" -run "" -bench . -benchmem -memprofile mem.out -cpuprofile cpu.out 2>&1 | tee report.txt
	#env GIT_TERMINAL_PROMPT=1 go get github.com/examplesite/myprivaterepo
	#go env -w GOPRIVATE=github.com/ademirfteo/gologwrp

loggo:
	cd internal/loggo/; go test -gcflags "-m" -run "" -bench . -benchmem -memprofile mem.out -cpuprofile cpu.out 2>&1 | tee report.txt
