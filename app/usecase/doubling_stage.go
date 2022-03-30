package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

type DoublingStageService interface {
	Doubling(in <-chan int) <-chan int
}

type doublingStageService struct{}

func NewDoublingStageService() DoublingStageService {
	return &doublingStageService{}
}

func (s *doublingStageService) Doubling(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			fmt.Printf("\tDoubling: %4d", v*v)
			out <- v
		}
		close(out)
	}()
	return out
}
