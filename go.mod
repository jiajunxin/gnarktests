module github.com/gnarktests

go 1.19

require (
	github.com/consensys/gnark v0.7.1
	github.com/consensys/gnark-crypto v0.10.0
)

require (
	github.com/bits-and-blooms/bitset v1.7.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

replace (
	github.com/consensys/gnark => github.com/bnb-chain/gnark v0.7.1-0.20230203031713-0d81c67d080a
	github.com/consensys/gnark-crypto => github.com/bnb-chain/gnark-crypto v0.7.1-0.20230203031630-7c643ad11891
)