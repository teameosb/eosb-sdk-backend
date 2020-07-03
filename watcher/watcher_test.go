package watcher

import (
	"context"
	"github.com/teameosb/eosb-sdk-backend/common"
	"github.com/teameosb/eosb-sdk-backend/sdk"
	"github.com/stretchr/testify/suite"
	"testing"
)

type watcherTestSuit struct {
	suite.Suite
}

func (s *watcherTestSuit) InitWatcher() *Watcher {
	return &Watcher{
		Ctx:         context.Background(),
		Eosb:       sdk.NewMockEosb(),
		KVClient:    &common.MockKVStore{},
		QueueClient: &common.MockQueue{},
	}
}

func (s *watcherTestSuit) SetupSuite() {
}

func (s *watcherTestSuit) TearDownSuite() {
}

func (s *watcherTestSuit) TearDownTest() {
}

func (s *watcherTestSuit) TestInitLastBlockNumberWithCache() {
	watcher := s.InitWatcher()

	watcher.KVClient.(*common.MockKVStore).On("Get", common.EOSB_WATCHER_BLOCK_NUMBER_CACHE_KEY).Return("10086", nil)

	watcher.initBlockNumber()
	s.Equal(uint64(10086), watcher.lastSyncedBlockNumber)
}

func (s *watcherTestSuit) TestInitLastBlockNumberWithoutCache() {
	watcher := s.InitWatcher()

	watcher.KVClient.(*common.MockKVStore).On("Get", common.EOSB_WATCHER_BLOCK_NUMBER_CACHE_KEY).Return("", common.KVStoreEmpty)
	watcher.Eosb.(*sdk.MockEosb).BlockChain.(*sdk.MockBlockchain).On("GetBlockNumber").Return(uint64(10086), nil)

	watcher.initBlockNumber()
	s.Equal(uint64(10086), watcher.lastSyncedBlockNumber)
}

func TestWatcherSuite(t *testing.T) {
	suite.Run(t, new(watcherTestSuit))
}
