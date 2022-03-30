package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

type TriplicationStageService interface {
	Triplication(in <-chan int) <-chan int
}

type triplicationStageService struct{}

func NewTriplicationStageService() TriplicationStageService {
	return &triplicationStageService{}
}

func (s *triplicationStageService) Triplication(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			fmt.Printf("\tTriplication: %4d", v*v*v)
			out <- v
		}
		close(out)
	}()
	return out
}
