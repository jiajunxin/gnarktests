package main

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type LoopCircuit struct {
	LoopNum frontend.Variable // secret value to determine how many rounds to run
	Secret  frontend.Variable // pre-image of the hash secret known to the prover only
	Hash    frontend.Variable `gnark:",public"` // hash of the secret known to all
}

func (circuit *LoopCircuit) Init() {
	circuit.LoopNum = 1
	circuit.Secret = 0xdeadf00d
	circuit.Hash = 1037254799353855871006189384309576393135431139055333626960622147300727796413
}

func (circuit *LoopCircuit) Define(api frontend.API) error {
	// hash function
	mimcInput, _ := mimc.NewMiMC(api)

	// hash the secret
	mimcInput.Write(circuit.Secret)

	// ensure hashes match
	api.AssertIsEqual(circuit.Hash, mimcInput.Sum())
	return nil
}
