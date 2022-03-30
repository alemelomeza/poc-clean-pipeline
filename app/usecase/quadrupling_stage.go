package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

type QuadruplingStageService interface {
	Quadrupling(in <-chan int) <-chan int
}

type quadruplingStageService struct{}

func NewQuadrupoublingStageService() QuadruplingStageService {
	return &quadruplingStageService{}
}

func (s *quadruplingStageService) Quadrupling(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			fmt.Printf("\tQuadrupling: %4d", v*v*v*v)
			out <- v
		}
		close(out)
	}()
	return out
}
