package application

import (
	"os"
	"os/signal"
	"poc-clean-pipeline/app/infrastructure/kafka/consumer"
	"poc-clean-pipeline/app/usecase"
	"syscall"
)

func StartApp() {
	consumer := consumer.NewKafkaConsumer(
		usecase.NewFirstStageService(),
		usecase.NewDoublingStageService(),
		usecase.NewTriplicationStageService(),
		usecase.NewQuadrupoublingStageService(),
		usecase.NewLastStageService())

	go consumer.Created()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
