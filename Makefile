all:
	go build -o a.exe
	a.exe
mac:
	GOOS=windows go build -o a.exe
	wine64 a.exe
res:
	.\tmp\DXArchive\DxaEncode.exe .\data\images
	.\tmp\DXArchive\DxaEncode.exe .\data\sounds
	git add .\data\images.dxa .\data\sounds.dxa
