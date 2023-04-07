package buffer

import (
	"github.com/HrvojeCosic/godb/src/utils"
)

type ClockReplacer struct {
	numPages uint                                // max number of pages ClockReplacer is able to store
	frames     *utils.CircularList[FrameId]      // frames tracked by clock replacer as a circular list
	frameTable map[FrameId]bool                  // mapping each frame in ClockReplacer to its reference bit
	hand       *utils.CircularListNode[FrameId]  // current position of the clock hand
}

func NewClockReplacer(capacity uint) *ClockReplacer {
	return &ClockReplacer{
		numPages: capacity,
		frames: utils.NewCircularList[FrameId](capacity),
		hand: nil,
		frameTable: make(map[FrameId]bool),
	}
}

func (cr *ClockReplacer) Evict() FrameId {
	if (cr.frames.Size() == 0) {
		return -1
	}
	if (cr.hand == nil) {
		cr.hand = cr.frames.Head()
	}

	for {
		currFrameId := cr.hand.Value()
		if (cr.frameTable[currFrameId] == false) {
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
	cr.frames.Remove(frameId)
	delete(cr.frameTable, frameId)
}

func (cr *ClockReplacer) Unpin(frameId FrameId) {
	cr.frames.Insert(frameId)
	cr.frameTable[frameId] = false
}

func (cr ClockReplacer) Size() uint {
	return cr.frames.Size()
}