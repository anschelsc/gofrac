all: test

frac.6: frac.go
	6g frac.go

test.6: frac.6 test.go
	6g test.go

test: test.6
	6l -o test test.6

clean:
	rm -f test test.6 frac.6
