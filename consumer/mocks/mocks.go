package mocks

import (
	context "context"
	"github.com/Shopify/sarama"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConsumerGroupSession is a mock of ConsumerGroupSession interface.
type MockConsumerGroupSession struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerGroupSessionMockRecorder
}

// MockConsumerGroupSessionMockRecorder is the mock recorder for MockConsumerGroupSession.
type MockConsumerGroupSessionMockRecorder struct {
	mock *MockConsumerGroupSession
}

// NewMockConsumerGroupSession creates a new mock instance.
func NewMockConsumerGroupSession(ctrl *gomock.Controller) *MockConsumerGroupSession {
	mock := &MockConsumerGroupSession{ctrl: ctrl}
	mock.recorder = &MockConsumerGroupSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerGroupSession) EXPECT() *MockConsumerGroupSessionMockRecorder {
	return m.recorder
}

// Claims mocks base method.
func (m *MockConsumerGroupSession) Claims() map[string][]int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Claims")
	ret0, _ := ret[0].(map[string][]int32)
	return ret0
}

// Claims indicates an expected call of Claims.
func (mr *MockConsumerGroupSessionMockRecorder) Claims() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Claims", reflect.TypeOf((*MockConsumerGroupSession)(nil).Claims))
}

// Commit mocks base method.
func (m *MockConsumerGroupSession) Commit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Commit")
}

// Commit indicates an expected call of Commit.
func (mr *MockConsumerGroupSessionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockConsumerGroupSession)(nil).Commit))
}

// Context mocks base method.
func (m *MockConsumerGroupSession) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockConsumerGroupSessionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockConsumerGroupSession)(nil).Context))
}

// GenerationID mocks base method.
func (m *MockConsumerGroupSession) GenerationID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerationID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GenerationID indicates an expected call of GenerationID.
func (mr *MockConsumerGroupSessionMockRecorder) GenerationID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerationID", reflect.TypeOf((*MockConsumerGroupSession)(nil).GenerationID))
}

// MarkMessage mocks base method.
func (m *MockConsumerGroupSession) MarkMessage(msg *sarama.ConsumerMessage, metadata string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkMessage", msg, metadata)
}

// MarkMessage indicates an expected call of MarkMessage.
func (mr *MockConsumerGroupSessionMockRecorder) MarkMessage(msg, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkMessage", reflect.TypeOf((*MockConsumerGroupSession)(nil).MarkMessage), msg, metadata)
}

// MarkOffset mocks base method.
func (m *MockConsumerGroupSession) MarkOffset(topic string, partition int32, offset int64, metadata string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkOffset", topic, partition, offset, metadata)
}

// MarkOffset indicates an expected call of MarkOffset.
func (mr *MockConsumerGroupSessionMockRecorder) MarkOffset(topic, partition, offset, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkOffset", reflect.TypeOf((*MockConsumerGroupSession)(nil).MarkOffset), topic, partition, offset, metadata)
}

// MemberID mocks base method.
func (m *MockConsumerGroupSession) MemberID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberID")
	ret0, _ := ret[0].(string)
	return ret0
}

// MemberID indicates an expected call of MemberID.
func (mr *MockConsumerGroupSessionMockRecorder) MemberID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberID", reflect.TypeOf((*MockConsumerGroupSession)(nil).MemberID))
}

// ResetOffset mocks base method.
func (m *MockConsumerGroupSession) ResetOffset(topic string, partition int32, offset int64, metadata string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetOffset", topic, partition, offset, metadata)
}

// ResetOffset indicates an expected call of ResetOffset.
func (mr *MockConsumerGroupSessionMockRecorder) ResetOffset(topic, partition, offset, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetOffset", reflect.TypeOf((*MockConsumerGroupSession)(nil).ResetOffset), topic, partition, offset, metadata)
}

// MockConsumerGroupClaim is a mock of ConsumerGroupClaim interface.
type MockConsumerGroupClaim struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerGroupClaimMockRecorder
}

// MockConsumerGroupClaimMockRecorder is the mock recorder for MockConsumerGroupClaim.
type MockConsumerGroupClaimMockRecorder struct {
	mock *MockConsumerGroupClaim
}

// NewMockConsumerGroupClaim creates a new mock instance.
func NewMockConsumerGroupClaim(ctrl *gomock.Controller) *MockConsumerGroupClaim {
	mock := &MockConsumerGroupClaim{ctrl: ctrl}
	mock.recorder = &MockConsumerGroupClaimMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerGroupClaim) EXPECT() *MockConsumerGroupClaimMockRecorder {
	return m.recorder
}

// HighWaterMarkOffset mocks base method.
func (m *MockConsumerGroupClaim) HighWaterMarkOffset() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HighWaterMarkOffset")
	ret0, _ := ret[0].(int64)
	return ret0
}

// HighWaterMarkOffset indicates an expected call of HighWaterMarkOffset.
func (mr *MockConsumerGroupClaimMockRecorder) HighWaterMarkOffset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HighWaterMarkOffset", reflect.TypeOf((*MockConsumerGroupClaim)(nil).HighWaterMarkOffset))
}

// InitialOffset mocks base method.
func (m *MockConsumerGroupClaim) InitialOffset() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialOffset")
	ret0, _ := ret[0].(int64)
	return ret0
}

// InitialOffset indicates an expected call of InitialOffset.
func (mr *MockConsumerGroupClaimMockRecorder) InitialOffset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialOffset", reflect.TypeOf((*MockConsumerGroupClaim)(nil).InitialOffset))
}

// Messages mocks base method.
func (m *MockConsumerGroupClaim) Messages() <-chan *sarama.ConsumerMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Messages")
	ret0, _ := ret[0].(<-chan *sarama.ConsumerMessage)
	return ret0
}

// Messages indicates an expected call of Messages.
func (mr *MockConsumerGroupClaimMockRecorder) Messages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockConsumerGroupClaim)(nil).Messages))
}

// Partition mocks base method.
func (m *MockConsumerGroupClaim) Partition() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partition")
	ret0, _ := ret[0].(int32)
	return ret0
}

// Partition indicates an expected call of Partition.
func (mr *MockConsumerGroupClaimMockRecorder) Partition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partition", reflect.TypeOf((*MockConsumerGroupClaim)(nil).Partition))
}

// Topic mocks base method.
func (m *MockConsumerGroupClaim) Topic() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockConsumerGroupClaimMockRecorder) Topic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockConsumerGroupClaim)(nil).Topic))
}
