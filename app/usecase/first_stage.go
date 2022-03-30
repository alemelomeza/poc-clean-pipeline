package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

type FirstStageService interface {
	Start(values ...int) <-chan int
}

type firstStageService struct{}

func NewFirstStageService() FirstStageService {
	return &firstStageService{}
}

func (s *firstStageService) Start(values ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range values {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			fmt.Printf("Start: %4d", v)
			out <- v
		}
		close(out)
	}()
	return out
}
