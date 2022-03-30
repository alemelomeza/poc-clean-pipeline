package consumer

import (
	"math/rand"
	"poc-clean-pipeline/app/usecase"
	"time"
)

type KafkaConsumer interface {
	Created()
}

type kafkaConsumer struct {
	ucfss usecase.FirstStageService
	ucdss usecase.DoublingStageService
	uctss usecase.TriplicationStageService
	ucqss usecase.QuadruplingStageService
	uclss usecase.LastStageService
}

func NewKafkaConsumer(
	ucfss usecase.FirstStageService,
	ucdss usecase.DoublingStageService,
	uctss usecase.TriplicationStageService,
	ucqss usecase.QuadruplingStageService,
	uclss usecase.LastStageService,
) KafkaConsumer {
	return &kafkaConsumer{
		ucfss: ucfss,
		ucdss: ucdss,
		uctss: uctss,
		ucqss: ucqss,
		uclss: uclss,
	}
}

func (c *kafkaConsumer) Created() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			// cfss := c.ucfss.Start(rand.Intn(10))
			// cdss := c.ucdss.Doubling(cfss)
			// ctss := c.uctss.Triplication(cdss)
			// cqss := c.ucqss.Quadrupling(ctss)
			// c.uclss.End(cqss)
			c.uclss.End(
				c.ucqss.Quadrupling(
					c.uctss.Triplication(
						c.ucdss.Doubling(
							c.ucfss.Start(
								rand.Intn(10))))))
		}
	}
}
