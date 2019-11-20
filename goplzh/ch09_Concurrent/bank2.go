package main

var(
	sema = make(chan struct{}, 1)
	balacne int
)

func Deposit2(amount int)  {
	sema <- struct{}{}
	balacne = balacne + amount
	<- sema
}

func Balace2() int {
	sema <- struct{}{}
	b := balacne
	<-sema // release token
	return b
}

