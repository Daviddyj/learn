package _4_channel_with_direction

import "testing"

func TestWithDirection(t *testing.T) {
	ch := make(chan int, 10)
	inOnly(ch)
}

func inOnly(ch chan<- int) {
	ch <- 5

}
func outOnly(ch <-chan int) {
	<-ch
}
