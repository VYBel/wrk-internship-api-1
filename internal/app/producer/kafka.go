package producer

import (
	"github.com/ozonmp/wrk-internship-api/internal/app/repo"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/wrk-internship-api/internal/app/sender"
	"github.com/ozonmp/wrk-internship-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n          uint64
	timeout    time.Duration
	repo       repo.EventRepo
	sender     sender.EventSender
	events     <-chan model.InternshipEvent
	workerPool *workerpool.WorkerPool
	wg         *sync.WaitGroup
	done       chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.InternshipEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					err := p.sender.Send(&event)
					if err != nil {
						p.processUpdate([]uint64{event.Id})
						continue
					}
					p.processClean([]uint64{event.Id})
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) processUpdate(eventIDs []uint64) {
	p.workerPool.Submit(func() {
		err := p.repo.Unlock(eventIDs)
		if err != nil {
			log.Printf("update error: %s\n", err)
		}
	})
}

func (p *producer) processClean(eventIDs []uint64) {
	err := p.repo.Remove(eventIDs)
	if err != nil {
		log.Printf("remove error: %s\n", err)
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
