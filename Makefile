build: 
	echo Starting to build API
	go build -o ./dist/app.exe ./*.go
	echo API built 

start: build 
	echo Starting API
	./dist/app.exe 
	echo API started