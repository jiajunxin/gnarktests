package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	// compiles our circuit into a R1CS
	var circuit LoopCircuit
	circuit.Init()
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)
	if err != nil {
		fmt.Println(err.Error())
	}
	// groth16 zkSNARK: Setup
	err = groth16.SetupLazyWithDump(r1cs, fileName)
	//pk, vk, _ := groth16.Setup(ccs)

	// witness definition
	witness, err := frontend.NewWitness(&circuit, ecc.BN254)
	if err != nil {
		fmt.Println(err.Error())
	}
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println(err.Error())

	}
	return
}
