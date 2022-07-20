# Golang ML Mutants
![Coverage](https://img.shields.io/badge/Coverage-88.7%25-brightgreen)


## Run locally

- Build and run:

```bash
$ go build .
$ ./ml-mutants

Starting server on http://localhost:8080
```



## Testing

```bash
$ go test -v -cover    
=== RUN   TestIsMutantDiagonal
Diagonal
m[0][0]: A A A A 
Time elapsed: 94.088µs
Is Mutant:  true
--- PASS: TestIsMutantDiagonal (0.00s)
=== RUN   TestIsMutantHorizontal
Horizontal
m[4][0]: C C C C 
Time elapsed: 10.948µs
Is Mutant:  true
--- PASS: TestIsMutantHorizontal (0.00s)
=== RUN   TestIsMutantVerticual
Vertical
m[0][4]: G G G G 
Time elapsed: 7.712µs
Is Mutant:  true
--- PASS: TestIsMutantVerticual (0.00s)
=== RUN   TestIsNotMutant
Time elapsed: 4.516µs
Is Mutant:  false
--- PASS: TestIsNotMutant (0.00s)
=== RUN   TestIsNotMutantInvaildDNA
Time elapsed: 2.951µs
Is Mutant:  false
--- PASS: TestIsNotMutantInvaildDNA (0.00s)
PASS
coverage: 88.9% of statements
ok      github.com/InmortalRegis/ml-mutants     0.109s
```
