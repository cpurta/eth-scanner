package block

import "sync"

type BlockManager struct {
	workers    []*BlockWorker
	blockRange *BlockRange
	waitGroup  *sync.WaitGroup
}

func NewBlockManager(workers []*BlockWorker, blockRange *BlockRange, wg *sync.WaitGroup) *BlockManager {
	manager := &BlockManager{
		workers:    workers,
		blockRange: blockRange,
		waitGroup:  wg,
	}

	manager.setWorkersBlockRange()

	return manager
}

func (manager *BlockManager) StartWorkers() {
	defer manager.waitGroup.Done()
	for _, worker := range manager.workers {
		go worker.Start()
	}
}

func (manager *BlockManager) Stop() {
	for _, worker := range manager.workers {
		worker.Stop()
	}
}

func (manager *BlockManager) BlocksRetreived() (int64, int64) {
	var (
		sum   int64
		total int64
	)
	for _, worker := range manager.workers {
		completed, rangelength := worker.GetProgress()
		sum += completed
		total += rangelength
	}

	return sum, total
}

func (manager *BlockManager) setWorkersBlockRange() {
	partitionRange := manager.blockRange.Len() / int64(len(manager.workers))

	for i, worker := range manager.workers {
		blockRangeMin := manager.blockRange.Min() + partitionRange*int64(i)
		blockRangeMax := blockRangeMin + partitionRange

		var workerBlockRange *BlockRange
		if i == len(manager.workers)-1 {
			workerBlockRange = NewBlockRange(blockRangeMin, manager.blockRange.Max())
		} else {
			workerBlockRange = NewBlockRange(blockRangeMin, blockRangeMax)
		}
		worker.SetBlockRange(workerBlockRange)
	}
}
