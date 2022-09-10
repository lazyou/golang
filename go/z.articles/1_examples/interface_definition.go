package main

type Car interface {
	NameGet() string
	Run(n int)
	Stop()
}
