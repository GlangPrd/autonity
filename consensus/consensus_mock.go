// Code generated by MockGen. DO NOT EDIT.
// Source: consensus/consensus.go

// Package consensus is a generated GoMock package.
package consensus

import (
	context "context"
	common "github.com/autonity/autonity/common"
	state "github.com/autonity/autonity/core/state"
	types "github.com/autonity/autonity/core/types"
	p2p "github.com/autonity/autonity/p2p"
	params "github.com/autonity/autonity/params"
	rpc "github.com/autonity/autonity/rpc"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockChainReader is a mock of ChainReader interface
type MockChainReader struct {
	ctrl     *gomock.Controller
	recorder *MockChainReaderMockRecorder
}

func (m *MockChainReader) GetTd(hash common.Hash, number uint64) *big.Int {
	panic("implement me")
}

func (m *MockChainReader) MinBaseFee(header *types.Header) (*big.Int, error) {
	panic("implement me")
}

// MockChainReaderMockRecorder is the mock recorder for MockChainReader
type MockChainReaderMockRecorder struct {
	mock *MockChainReader
}

// NewMockChainReader creates a new mock instance
func NewMockChainReader(ctrl *gomock.Controller) *MockChainReader {
	mock := &MockChainReader{ctrl: ctrl}
	mock.recorder = &MockChainReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainReader) EXPECT() *MockChainReaderMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockChainReader) Config() *params.ChainConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*params.ChainConfig)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockChainReaderMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockChainReader)(nil).Config))
}

// CurrentHeader mocks base method
func (m *MockChainReader) CurrentHeader() *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentHeader")
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// CurrentHeader indicates an expected call of CurrentHeader
func (mr *MockChainReaderMockRecorder) CurrentHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockChainReader)(nil).CurrentHeader))
}

// GetHeader mocks base method
func (m *MockChainReader) GetHeader(hash common.Hash, number uint64) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", hash, number)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockChainReaderMockRecorder) GetHeader(hash, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockChainReader)(nil).GetHeader), hash, number)
}

// GetHeaderByNumber mocks base method
func (m *MockChainReader) GetHeaderByNumber(number uint64) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByNumber", number)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeaderByNumber indicates an expected call of GetHeaderByNumber
func (mr *MockChainReaderMockRecorder) GetHeaderByNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByNumber", reflect.TypeOf((*MockChainReader)(nil).GetHeaderByNumber), number)
}

// GetHeaderByHash mocks base method
func (m *MockChainReader) GetHeaderByHash(hash common.Hash) *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", hash)
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash
func (mr *MockChainReaderMockRecorder) GetHeaderByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockChainReader)(nil).GetHeaderByHash), hash)
}

// GetBlock mocks base method
func (m *MockChainReader) GetBlock(hash common.Hash, number uint64) *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", hash, number)
	ret0, _ := ret[0].(*types.Block)
	return ret0
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockChainReaderMockRecorder) GetBlock(hash, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockChainReader)(nil).GetBlock), hash, number)
}

// Engine mocks base method
func (m *MockChainReader) Engine() Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Engine")
	ret0, _ := ret[0].(Engine)
	return ret0
}

// Engine indicates an expected call of Engine
func (mr *MockChainReaderMockRecorder) Engine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Engine", reflect.TypeOf((*MockChainReader)(nil).Engine))
}

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Author mocks base method
func (m *MockEngine) Author(header *types.Header) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Author", header)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Author indicates an expected call of Author
func (mr *MockEngineMockRecorder) Author(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Author", reflect.TypeOf((*MockEngine)(nil).Author), header)
}

// VerifyHeader mocks base method
func (m *MockEngine) VerifyHeader(chain ChainReader, header *types.Header, seal bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", chain, header, seal)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader
func (mr *MockEngineMockRecorder) VerifyHeader(chain, header, seal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockEngine)(nil).VerifyHeader), chain, header, seal)
}

// VerifyHeaders mocks base method
func (m *MockEngine) VerifyHeaders(chain ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeaders", chain, headers, seals)
	ret0, _ := ret[0].(chan<- struct{})
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// VerifyHeaders indicates an expected call of VerifyHeaders
func (mr *MockEngineMockRecorder) VerifyHeaders(chain, headers, seals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeaders", reflect.TypeOf((*MockEngine)(nil).VerifyHeaders), chain, headers, seals)
}

// VerifyUncles mocks base method
func (m *MockEngine) VerifyUncles(chain ChainReader, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUncles", chain, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUncles indicates an expected call of VerifyUncles
func (mr *MockEngineMockRecorder) VerifyUncles(chain, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUncles", reflect.TypeOf((*MockEngine)(nil).VerifyUncles), chain, block)
}

// VerifySeal mocks base method
func (m *MockEngine) VerifySeal(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySeal", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySeal indicates an expected call of VerifySeal
func (mr *MockEngineMockRecorder) VerifySeal(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySeal", reflect.TypeOf((*MockEngine)(nil).VerifySeal), chain, header)
}

// Prepare mocks base method
func (m *MockEngine) Prepare(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockEngineMockRecorder) Prepare(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockEngine)(nil).Prepare), chain, header)
}

// Finalize mocks base method
func (m *MockEngine) Finalize(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (types.Committee, *types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(types.Committee)
	ret1, _ := ret[1].(*types.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Finalize indicates an expected call of Finalize
func (mr *MockEngineMockRecorder) Finalize(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockEngine)(nil).Finalize), chain, header, state, txs, uncles, receipts)
}

// FinalizeAndAssemble mocks base method
func (m *MockEngine) FinalizeAndAssemble(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts *[]*types.Receipt) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAndAssemble", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalizeAndAssemble indicates an expected call of FinalizeAndAssemble
func (mr *MockEngineMockRecorder) FinalizeAndAssemble(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAndAssemble", reflect.TypeOf((*MockEngine)(nil).FinalizeAndAssemble), chain, header, state, txs, uncles, receipts)
}

// Seal mocks base method
func (m *MockEngine) Seal(chain ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", chain, block, results, stop)
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockEngineMockRecorder) Seal(chain, block, results, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockEngine)(nil).Seal), chain, block, results, stop)
}

// SealHash mocks base method
func (m *MockEngine) SealHash(header *types.Header) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SealHash", header)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// SealHash indicates an expected call of SealHash
func (mr *MockEngineMockRecorder) SealHash(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SealHash", reflect.TypeOf((*MockEngine)(nil).SealHash), header)
}

// CalcDifficulty mocks base method
func (m *MockEngine) CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcDifficulty", chain, time, parent)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CalcDifficulty indicates an expected call of CalcDifficulty
func (mr *MockEngineMockRecorder) CalcDifficulty(chain, time, parent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcDifficulty", reflect.TypeOf((*MockEngine)(nil).CalcDifficulty), chain, time, parent)
}

// APIs mocks base method
func (m *MockEngine) APIs(chain ChainReader) []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs", chain)
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs
func (mr *MockEngineMockRecorder) APIs(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockEngine)(nil).APIs), chain)
}

// Close mocks base method
func (m *MockEngine) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockEngineMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEngine)(nil).Close))
}

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// NewChainHead mocks base method
func (m *MockHandler) NewChainHead() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChainHead")
	ret0, _ := ret[0].(error)
	return ret0
}

// NewChainHead indicates an expected call of NewChainHead
func (mr *MockHandlerMockRecorder) NewChainHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChainHead", reflect.TypeOf((*MockHandler)(nil).NewChainHead))
}

// HandleMsg mocks base method
func (m *MockHandler) HandleMsg(address common.Address, data p2p.Msg) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMsg", address, data)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleMsg indicates an expected call of HandleMsg
func (mr *MockHandlerMockRecorder) HandleMsg(address, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMsg", reflect.TypeOf((*MockHandler)(nil).HandleMsg), address, data)
}

// SetBroadcaster mocks base method
func (m *MockHandler) SetBroadcaster(arg0 Broadcaster) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBroadcaster", arg0)
}

// SetBroadcaster indicates an expected call of SetBroadcaster
func (mr *MockHandlerMockRecorder) SetBroadcaster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBroadcaster", reflect.TypeOf((*MockHandler)(nil).SetBroadcaster), arg0)
}

// Protocol mocks base method
func (m *MockHandler) Protocol() (string, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// Protocol indicates an expected call of Protocol
func (mr *MockHandlerMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockHandler)(nil).Protocol))
}

// MockPoW is a mock of PoW interface
type MockPoW struct {
	ctrl     *gomock.Controller
	recorder *MockPoWMockRecorder
}

// MockPoWMockRecorder is the mock recorder for MockPoW
type MockPoWMockRecorder struct {
	mock *MockPoW
}

// NewMockPoW creates a new mock instance
func NewMockPoW(ctrl *gomock.Controller) *MockPoW {
	mock := &MockPoW{ctrl: ctrl}
	mock.recorder = &MockPoWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPoW) EXPECT() *MockPoWMockRecorder {
	return m.recorder
}

// Author mocks base method
func (m *MockPoW) Author(header *types.Header) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Author", header)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Author indicates an expected call of Author
func (mr *MockPoWMockRecorder) Author(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Author", reflect.TypeOf((*MockPoW)(nil).Author), header)
}

// VerifyHeader mocks base method
func (m *MockPoW) VerifyHeader(chain ChainReader, header *types.Header, seal bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", chain, header, seal)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader
func (mr *MockPoWMockRecorder) VerifyHeader(chain, header, seal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockPoW)(nil).VerifyHeader), chain, header, seal)
}

// VerifyHeaders mocks base method
func (m *MockPoW) VerifyHeaders(chain ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeaders", chain, headers, seals)
	ret0, _ := ret[0].(chan<- struct{})
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// VerifyHeaders indicates an expected call of VerifyHeaders
func (mr *MockPoWMockRecorder) VerifyHeaders(chain, headers, seals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeaders", reflect.TypeOf((*MockPoW)(nil).VerifyHeaders), chain, headers, seals)
}

// VerifyUncles mocks base method
func (m *MockPoW) VerifyUncles(chain ChainReader, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUncles", chain, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUncles indicates an expected call of VerifyUncles
func (mr *MockPoWMockRecorder) VerifyUncles(chain, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUncles", reflect.TypeOf((*MockPoW)(nil).VerifyUncles), chain, block)
}

// VerifySeal mocks base method
func (m *MockPoW) VerifySeal(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySeal", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySeal indicates an expected call of VerifySeal
func (mr *MockPoWMockRecorder) VerifySeal(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySeal", reflect.TypeOf((*MockPoW)(nil).VerifySeal), chain, header)
}

// Prepare mocks base method
func (m *MockPoW) Prepare(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockPoWMockRecorder) Prepare(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockPoW)(nil).Prepare), chain, header)
}

// Finalize mocks base method
func (m *MockPoW) Finalize(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (types.Committee, *types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(types.Committee)
	ret1, _ := ret[1].(*types.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Finalize indicates an expected call of Finalize
func (mr *MockPoWMockRecorder) Finalize(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockPoW)(nil).Finalize), chain, header, state, txs, uncles, receipts)
}

// FinalizeAndAssemble mocks base method
func (m *MockPoW) FinalizeAndAssemble(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts *[]*types.Receipt) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAndAssemble", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalizeAndAssemble indicates an expected call of FinalizeAndAssemble
func (mr *MockPoWMockRecorder) FinalizeAndAssemble(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAndAssemble", reflect.TypeOf((*MockPoW)(nil).FinalizeAndAssemble), chain, header, state, txs, uncles, receipts)
}

// Seal mocks base method
func (m *MockPoW) Seal(chain ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", chain, block, results, stop)
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockPoWMockRecorder) Seal(chain, block, results, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockPoW)(nil).Seal), chain, block, results, stop)
}

// SealHash mocks base method
func (m *MockPoW) SealHash(header *types.Header) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SealHash", header)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// SealHash indicates an expected call of SealHash
func (mr *MockPoWMockRecorder) SealHash(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SealHash", reflect.TypeOf((*MockPoW)(nil).SealHash), header)
}

// CalcDifficulty mocks base method
func (m *MockPoW) CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcDifficulty", chain, time, parent)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CalcDifficulty indicates an expected call of CalcDifficulty
func (mr *MockPoWMockRecorder) CalcDifficulty(chain, time, parent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcDifficulty", reflect.TypeOf((*MockPoW)(nil).CalcDifficulty), chain, time, parent)
}

// APIs mocks base method
func (m *MockPoW) APIs(chain ChainReader) []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs", chain)
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs
func (mr *MockPoWMockRecorder) APIs(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockPoW)(nil).APIs), chain)
}

// Close mocks base method
func (m *MockPoW) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPoWMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPoW)(nil).Close))
}

// Hashrate mocks base method
func (m *MockPoW) Hashrate() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hashrate")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Hashrate indicates an expected call of Hashrate
func (mr *MockPoWMockRecorder) Hashrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hashrate", reflect.TypeOf((*MockPoW)(nil).Hashrate))
}

// MockBFT is a mock of BFT interface
type MockBFT struct {
	ctrl     *gomock.Controller
	recorder *MockBFTMockRecorder
}

// MockBFTMockRecorder is the mock recorder for MockBFT
type MockBFTMockRecorder struct {
	mock *MockBFT
}

// NewMockBFT creates a new mock instance
func NewMockBFT(ctrl *gomock.Controller) *MockBFT {
	mock := &MockBFT{ctrl: ctrl}
	mock.recorder = &MockBFTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBFT) EXPECT() *MockBFTMockRecorder {
	return m.recorder
}

// Author mocks base method
func (m *MockBFT) Author(header *types.Header) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Author", header)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Author indicates an expected call of Author
func (mr *MockBFTMockRecorder) Author(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Author", reflect.TypeOf((*MockBFT)(nil).Author), header)
}

// VerifyHeader mocks base method
func (m *MockBFT) VerifyHeader(chain ChainReader, header *types.Header, seal bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", chain, header, seal)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader
func (mr *MockBFTMockRecorder) VerifyHeader(chain, header, seal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockBFT)(nil).VerifyHeader), chain, header, seal)
}

// VerifyHeaders mocks base method
func (m *MockBFT) VerifyHeaders(chain ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeaders", chain, headers, seals)
	ret0, _ := ret[0].(chan<- struct{})
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// VerifyHeaders indicates an expected call of VerifyHeaders
func (mr *MockBFTMockRecorder) VerifyHeaders(chain, headers, seals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeaders", reflect.TypeOf((*MockBFT)(nil).VerifyHeaders), chain, headers, seals)
}

// VerifyUncles mocks base method
func (m *MockBFT) VerifyUncles(chain ChainReader, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUncles", chain, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUncles indicates an expected call of VerifyUncles
func (mr *MockBFTMockRecorder) VerifyUncles(chain, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUncles", reflect.TypeOf((*MockBFT)(nil).VerifyUncles), chain, block)
}

// VerifySeal mocks base method
func (m *MockBFT) VerifySeal(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySeal", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySeal indicates an expected call of VerifySeal
func (mr *MockBFTMockRecorder) VerifySeal(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySeal", reflect.TypeOf((*MockBFT)(nil).VerifySeal), chain, header)
}

// Prepare mocks base method
func (m *MockBFT) Prepare(chain ChainReader, header *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", chain, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockBFTMockRecorder) Prepare(chain, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockBFT)(nil).Prepare), chain, header)
}

// Finalize mocks base method
func (m *MockBFT) Finalize(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (types.Committee, *types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(types.Committee)
	ret1, _ := ret[1].(*types.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Finalize indicates an expected call of Finalize
func (mr *MockBFTMockRecorder) Finalize(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockBFT)(nil).Finalize), chain, header, state, txs, uncles, receipts)
}

// FinalizeAndAssemble mocks base method
func (m *MockBFT) FinalizeAndAssemble(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts *[]*types.Receipt) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAndAssemble", chain, header, state, txs, uncles, receipts)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalizeAndAssemble indicates an expected call of FinalizeAndAssemble
func (mr *MockBFTMockRecorder) FinalizeAndAssemble(chain, header, state, txs, uncles, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAndAssemble", reflect.TypeOf((*MockBFT)(nil).FinalizeAndAssemble), chain, header, state, txs, uncles, receipts)
}

// Seal mocks base method
func (m *MockBFT) Seal(chain ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", chain, block, results, stop)
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockBFTMockRecorder) Seal(chain, block, results, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockBFT)(nil).Seal), chain, block, results, stop)
}

// SealHash mocks base method
func (m *MockBFT) SealHash(header *types.Header) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SealHash", header)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// SealHash indicates an expected call of SealHash
func (mr *MockBFTMockRecorder) SealHash(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SealHash", reflect.TypeOf((*MockBFT)(nil).SealHash), header)
}

// CalcDifficulty mocks base method
func (m *MockBFT) CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcDifficulty", chain, time, parent)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CalcDifficulty indicates an expected call of CalcDifficulty
func (mr *MockBFTMockRecorder) CalcDifficulty(chain, time, parent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcDifficulty", reflect.TypeOf((*MockBFT)(nil).CalcDifficulty), chain, time, parent)
}

// APIs mocks base method
func (m *MockBFT) APIs(chain ChainReader) []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs", chain)
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs
func (mr *MockBFTMockRecorder) APIs(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockBFT)(nil).APIs), chain)
}

// Close mocks base method
func (m *MockBFT) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBFTMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBFT)(nil).Close))
}

// Start mocks base method
func (m *MockBFT) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBFTMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBFT)(nil).Start), ctx)
}

// MockSyncer is a mock of Syncer interface
type MockSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncerMockRecorder
}

// MockSyncerMockRecorder is the mock recorder for MockSyncer
type MockSyncerMockRecorder struct {
	mock *MockSyncer
}

// NewMockSyncer creates a new mock instance
func NewMockSyncer(ctrl *gomock.Controller) *MockSyncer {
	mock := &MockSyncer{ctrl: ctrl}
	mock.recorder = &MockSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyncer) EXPECT() *MockSyncerMockRecorder {
	return m.recorder
}

// SyncPeer mocks base method
func (m *MockSyncer) SyncPeer(address common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SyncPeer", address)
}

// SyncPeer indicates an expected call of SyncPeer
func (mr *MockSyncerMockRecorder) SyncPeer(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPeer", reflect.TypeOf((*MockSyncer)(nil).SyncPeer), address)
}

// ResetPeerCache mocks base method
func (m *MockSyncer) ResetPeerCache(address common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetPeerCache", address)
}

// ResetPeerCache indicates an expected call of ResetPeerCache
func (mr *MockSyncerMockRecorder) ResetPeerCache(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPeerCache", reflect.TypeOf((*MockSyncer)(nil).ResetPeerCache), address)
}
