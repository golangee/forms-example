package livebuilder

import (
	"encoding/hex"
	"fmt"
	"github.com/golangee/forms-example/server/internal/builder"
	"github.com/golangee/forms-example/server/internal/fsnotify"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"sync"
	"time"
)

// Builder provides an automatic live builder which rebuilds an idiomatic golangee wasm project any time it
// recognizes a change.
type Builder struct {
	schedulerStopChan chan bool
	logger            log.Logger
	lastBuildHash     []byte
	srcDir, dstDir    string
	buildLock         sync.Mutex
	watcher           *fsnotify.Watcher
}

func NewBuilder(dstDir, srcDir string) (*Builder, error) {
	b := &Builder{
		srcDir: srcDir,
		dstDir: dstDir,
	}

	b.logger = log.NewLogger(ecs.Log("livebuilder"))

	// TODO polling is a bad thing, implement inotify et al
	/*	b.schedulerStopChan = scheduleDelayed(10*time.Second, func() {
			hash, err := builder.HashFileTree(srcDir)
			if err != nil {
				b.logger.Print(ecs.Msg("failed to calculate file tree hash"), ecs.ErrMsg(err))
				return
			}

			b.buildLock.Lock()
			hashCopy := append([]byte{}, b.lastBuildHash...)
			b.buildLock.Unlock()

			if bytes.Compare(hashCopy, hash) != 0 {
				if err := b.Build(); err != nil {
					b.logger.Print(ecs.Msg("failed to build project"), ecs.ErrMsg(err))
					return
				}
			}

		})

	*/
	w, err := fsnotify.NewWatcher(srcDir, func() {
		b.logger.Print(ecs.Msg("rebuilding..."))
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init fsnotify watcher: %w", err)
	}

	b.watcher = w
	b.logger.Print(ecs.Msg("start watching " + srcDir))

	return b, nil
}

// Build triggers a build now
func (b *Builder) Build() error {
	b.buildLock.Lock()
	defer b.buildLock.Unlock()

	start := time.Now()
	defer func() {
		b.logger.Print(ecs.Msg("build duration " + time.Now().Sub(start).String()))
	}()
	hash, err := builder.HashFileTree(b.srcDir)
	if err != nil {
		return err
	}
	b.logger.Print(ecs.Msg("building " + hex.EncodeToString(hash)))

	if err := builder.BuildProject(b.srcDir, b.dstDir); err != nil {
		return fmt.Errorf("unable to build wasm project: %w", err)
	}

	b.lastBuildHash = hash
	return nil
}

func (b *Builder) Close() error {
	b.schedulerStopChan <- true
	close(b.schedulerStopChan)
	return nil
}
