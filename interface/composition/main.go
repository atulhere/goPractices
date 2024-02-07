package main

type Mover interface {
	move()
}

type Locker interface {
	locker()
}
type MoverLocker interface {
	Mover
	Locker
}
