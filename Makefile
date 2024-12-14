all:
	go build -o a.exe
	a.exe
mac:
	GOOS=windows go build -o a.exe
	wine64 a.exe
