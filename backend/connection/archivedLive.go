package connection

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"

	"github.com/f1gopher/f1gopherlib/f1log"
)

type archivedLive struct {
	log         *f1log.F1GopherLibLog
	path        string
	archiveFile *bufio.Scanner

	dataFeed chan Payload

	ctx context.Context
	wg  *sync.WaitGroup
}

func CreateArchivedLive(
	ctx context.Context,
	wg *sync.WaitGroup,
	log *f1log.F1GopherLibLog,
	file string) Connection {

	return &archivedLive{
		ctx:      ctx,
		wg:       wg,
		log:      log,
		path:     file,
		dataFeed: make(chan Payload, 100),
	}
}

func (a *archivedLive) Connect() (error, <-chan Payload) {
	f, err := os.Open(a.path)
	if err != nil {
		a.log.Errorf("Archived Live can't open file '%s': %s", a.path, err)
		return err, nil
	}
	a.archiveFile = bufio.NewScanner(f)

	go a.readEntries()

	return nil, a.dataFeed
}

func (a *archivedLive) readEntries() {
	a.wg.Add(1)
	defer a.wg.Done()

	// Will read entries as fast as possible until the channel is full
	// and then wait. Flow control for message timing is handled elsewhere
	for a.archiveFile.Scan() {

		select {
		case <-a.ctx.Done():
			return
		default:
		}

		line1 := a.archiveFile.Text()

		if !a.archiveFile.Scan() {
			a.log.Error("Archived Live unexpected EOF, missing second line")
			return
		}
		line2 := []byte(a.archiveFile.Text())

		if !a.archiveFile.Scan() {
			a.log.Error("Archived Live unexpected EOF, missing third line")
			return
		}
		line3 := a.archiveFile.Text()

		a.dataFeed <- Payload{
			Name:      line1,
			Data:      line2,
			Timestamp: line3,
		}
	}
}

func (a *archivedLive) IncrementTime(amount time.Duration) {}

func (a *archivedLive) JumpToStart() time.Time { return time.Time{} }
