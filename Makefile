default:
	go run acc.go

watch:
	ls -d * | entr -d make
