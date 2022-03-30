package usecase

import (
	"fmt"
	"math/rand"
	"time"
)

type LastStageService interface {
	End(in <-chan int)
}

type lastStageService struct{}

func NewLastStageService() LastStageService {
	return &lastStageService{}
}

func (s *lastStageService) End(in <-chan int) {
	for v := range in {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Printf("\tEnd:%4d\n", v)
	}
}
