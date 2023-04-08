package buffer

import (
	"sync"

	"github.com/HrvojeCosic/godb/src/utils"
)

type ClockReplacer struct {
	numPages uint                                // max number of pages ClockReplacer is able to store
	frames     *utils.CircularList[FrameId]      // frames tracked by clock replacer as a circular list
	frameTable map[FrameId]bool                  // mapping each frame in ClockReplacer to its reference bit
	hand       *utils.CircularListNode[FrameId]  // current position of the clock hand
	latch      sync.RWMutex
}

func NewClockReplacer(capacity uint) *ClockReplacer {
	return &ClockReplacer{
		numPages: capacity,
		frames: utils.NewCircularList[FrameId](capacity),
		hand: nil,
		frameTable: make(map[FrameId]bool),
		latch: sync.RWMutex{},
	}
}

func (cr *ClockReplacer) Evict() FrameId {
	cr.latch.Lock()
	defer cr.latch.Unlock()
	if (cr.frames.Size() == 0) {
		return -1
	}
	if (cr.hand == nil) {
		cr.hand = cr.frames.Head()
	}

	for {
		currFrameId := cr.hand.Value()
		if (!cr.frameTable[currFrameId]) {
			cr.frames.Remove(currFrameId)
			delete(cr.frameTable, currFrameId)
			cr.hand = cr.frames.Next(cr.hand)
			return currFrameId
		} else {
			cr.frameTable[currFrameId] = false
			cr.hand = cr.frames.Next(cr.hand)
		}
	}
}

func (cr *ClockReplacer) Pin(frameId FrameId) {
	cr.latch.Lock()
	defer cr.latch.Unlock()
	cr.frames.Remove(frameId)
	delete(cr.frameTable, frameId)
}

func (cr *ClockReplacer) Unpin(frameId FrameId) {
	cr.latch.Lock()
	defer cr.latch.Unlock()
	_, err := cr.frames.Insert(frameId)
	if (err != nil) {
		return
	}
	cr.frameTable[frameId] = false
}

func (cr *ClockReplacer) Size() uint {
	cr.latch.RLock()
	defer cr.latch.RUnlock()
	return cr.frames.Size()
}