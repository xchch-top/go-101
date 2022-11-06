// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/golang/mock/gomock"
)

// MockCmdable is a mock of Cmdable interface.
type MockCmdable struct {
	ctrl     *gomock.Controller
	recorder *MockCmdableMockRecorder
}

// MockCmdableMockRecorder is the mock recorder for MockCmdable.
type MockCmdableMockRecorder struct {
	mock *MockCmdable
}

// NewMockCmdable creates a new mock instance.
func NewMockCmdable(ctrl *gomock.Controller) *MockCmdable {
	mock := &MockCmdable{ctrl: ctrl}
	mock.recorder = &MockCmdableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmdable) EXPECT() *MockCmdableMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockCmdable) Append(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Append indicates an expected call of Append.
func (mr *MockCmdableMockRecorder) Append(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockCmdable)(nil).Append), arg0, arg1, arg2)
}

// BLMove mocks base method.
func (m *MockCmdable) BLMove(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BLMove", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BLMove indicates an expected call of BLMove.
func (mr *MockCmdableMockRecorder) BLMove(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLMove", reflect.TypeOf((*MockCmdable)(nil).BLMove), arg0, arg1, arg2, arg3, arg4, arg5)
}

// BLPop mocks base method.
func (m *MockCmdable) BLPop(arg0 context.Context, arg1 time.Duration, arg2 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BLPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BLPop indicates an expected call of BLPop.
func (mr *MockCmdableMockRecorder) BLPop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLPop", reflect.TypeOf((*MockCmdable)(nil).BLPop), varargs...)
}

// BRPop mocks base method.
func (m *MockCmdable) BRPop(arg0 context.Context, arg1 time.Duration, arg2 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BRPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BRPop indicates an expected call of BRPop.
func (mr *MockCmdableMockRecorder) BRPop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPop", reflect.TypeOf((*MockCmdable)(nil).BRPop), varargs...)
}

// BRPopLPush mocks base method.
func (m *MockCmdable) BRPopLPush(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BRPopLPush", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BRPopLPush indicates an expected call of BRPopLPush.
func (mr *MockCmdableMockRecorder) BRPopLPush(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPopLPush", reflect.TypeOf((*MockCmdable)(nil).BRPopLPush), arg0, arg1, arg2, arg3)
}

// BZPopMax mocks base method.
func (m *MockCmdable) BZPopMax(arg0 context.Context, arg1 time.Duration, arg2 ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMax indicates an expected call of BZPopMax.
func (mr *MockCmdableMockRecorder) BZPopMax(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMax", reflect.TypeOf((*MockCmdable)(nil).BZPopMax), varargs...)
}

// BZPopMin mocks base method.
func (m *MockCmdable) BZPopMin(arg0 context.Context, arg1 time.Duration, arg2 ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMin indicates an expected call of BZPopMin.
func (mr *MockCmdableMockRecorder) BZPopMin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMin", reflect.TypeOf((*MockCmdable)(nil).BZPopMin), varargs...)
}

// BgRewriteAOF mocks base method.
func (m *MockCmdable) BgRewriteAOF(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgRewriteAOF", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgRewriteAOF indicates an expected call of BgRewriteAOF.
func (mr *MockCmdableMockRecorder) BgRewriteAOF(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgRewriteAOF", reflect.TypeOf((*MockCmdable)(nil).BgRewriteAOF), arg0)
}

// BgSave mocks base method.
func (m *MockCmdable) BgSave(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgSave", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgSave indicates an expected call of BgSave.
func (mr *MockCmdableMockRecorder) BgSave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgSave", reflect.TypeOf((*MockCmdable)(nil).BgSave), arg0)
}

// BitCount mocks base method.
func (m *MockCmdable) BitCount(arg0 context.Context, arg1 string, arg2 *redis.BitCount) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitCount indicates an expected call of BitCount.
func (mr *MockCmdableMockRecorder) BitCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitCount", reflect.TypeOf((*MockCmdable)(nil).BitCount), arg0, arg1, arg2)
}

// BitField mocks base method.
func (m *MockCmdable) BitField(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitField", varargs...)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// BitField indicates an expected call of BitField.
func (mr *MockCmdableMockRecorder) BitField(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitField", reflect.TypeOf((*MockCmdable)(nil).BitField), varargs...)
}

// BitOpAnd mocks base method.
func (m *MockCmdable) BitOpAnd(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpAnd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpAnd indicates an expected call of BitOpAnd.
func (mr *MockCmdableMockRecorder) BitOpAnd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpAnd", reflect.TypeOf((*MockCmdable)(nil).BitOpAnd), varargs...)
}

// BitOpNot mocks base method.
func (m *MockCmdable) BitOpNot(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitOpNot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpNot indicates an expected call of BitOpNot.
func (mr *MockCmdableMockRecorder) BitOpNot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpNot", reflect.TypeOf((*MockCmdable)(nil).BitOpNot), arg0, arg1, arg2)
}

// BitOpOr mocks base method.
func (m *MockCmdable) BitOpOr(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpOr", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpOr indicates an expected call of BitOpOr.
func (mr *MockCmdableMockRecorder) BitOpOr(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpOr", reflect.TypeOf((*MockCmdable)(nil).BitOpOr), varargs...)
}

// BitOpXor mocks base method.
func (m *MockCmdable) BitOpXor(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpXor", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpXor indicates an expected call of BitOpXor.
func (mr *MockCmdableMockRecorder) BitOpXor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpXor", reflect.TypeOf((*MockCmdable)(nil).BitOpXor), varargs...)
}

// BitPos mocks base method.
func (m *MockCmdable) BitPos(arg0 context.Context, arg1 string, arg2 int64, arg3 ...int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitPos", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitPos indicates an expected call of BitPos.
func (mr *MockCmdableMockRecorder) BitPos(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitPos", reflect.TypeOf((*MockCmdable)(nil).BitPos), varargs...)
}

// ClientGetName mocks base method.
func (m *MockCmdable) ClientGetName(arg0 context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetName", arg0)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientGetName indicates an expected call of ClientGetName.
func (mr *MockCmdableMockRecorder) ClientGetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetName", reflect.TypeOf((*MockCmdable)(nil).ClientGetName), arg0)
}

// ClientID mocks base method.
func (m *MockCmdable) ClientID(arg0 context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID", arg0)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockCmdableMockRecorder) ClientID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockCmdable)(nil).ClientID), arg0)
}

// ClientKill mocks base method.
func (m *MockCmdable) ClientKill(arg0 context.Context, arg1 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientKill", arg0, arg1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClientKill indicates an expected call of ClientKill.
func (mr *MockCmdableMockRecorder) ClientKill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKill", reflect.TypeOf((*MockCmdable)(nil).ClientKill), arg0, arg1)
}

// ClientKillByFilter mocks base method.
func (m *MockCmdable) ClientKillByFilter(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClientKillByFilter", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientKillByFilter indicates an expected call of ClientKillByFilter.
func (mr *MockCmdableMockRecorder) ClientKillByFilter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKillByFilter", reflect.TypeOf((*MockCmdable)(nil).ClientKillByFilter), varargs...)
}

// ClientList mocks base method.
func (m *MockCmdable) ClientList(arg0 context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientList", arg0)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientList indicates an expected call of ClientList.
func (mr *MockCmdableMockRecorder) ClientList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientList", reflect.TypeOf((*MockCmdable)(nil).ClientList), arg0)
}

// ClientPause mocks base method.
func (m *MockCmdable) ClientPause(arg0 context.Context, arg1 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientPause", arg0, arg1)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ClientPause indicates an expected call of ClientPause.
func (mr *MockCmdableMockRecorder) ClientPause(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientPause", reflect.TypeOf((*MockCmdable)(nil).ClientPause), arg0, arg1)
}

// ClientUnblock mocks base method.
func (m *MockCmdable) ClientUnblock(arg0 context.Context, arg1 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientUnblock", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientUnblock indicates an expected call of ClientUnblock.
func (mr *MockCmdableMockRecorder) ClientUnblock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientUnblock", reflect.TypeOf((*MockCmdable)(nil).ClientUnblock), arg0, arg1)
}

// ClientUnblockWithError mocks base method.
func (m *MockCmdable) ClientUnblockWithError(arg0 context.Context, arg1 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientUnblockWithError", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientUnblockWithError indicates an expected call of ClientUnblockWithError.
func (mr *MockCmdableMockRecorder) ClientUnblockWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientUnblockWithError", reflect.TypeOf((*MockCmdable)(nil).ClientUnblockWithError), arg0, arg1)
}

// ClientUnpause mocks base method.
func (m *MockCmdable) ClientUnpause(arg0 context.Context) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientUnpause", arg0)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ClientUnpause indicates an expected call of ClientUnpause.
func (mr *MockCmdableMockRecorder) ClientUnpause(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientUnpause", reflect.TypeOf((*MockCmdable)(nil).ClientUnpause), arg0)
}

// ClusterAddSlots mocks base method.
func (m *MockCmdable) ClusterAddSlots(arg0 context.Context, arg1 ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterAddSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlots indicates an expected call of ClusterAddSlots.
func (mr *MockCmdableMockRecorder) ClusterAddSlots(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterAddSlots), varargs...)
}

// ClusterAddSlotsRange mocks base method.
func (m *MockCmdable) ClusterAddSlotsRange(arg0 context.Context, arg1, arg2 int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAddSlotsRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlotsRange indicates an expected call of ClusterAddSlotsRange.
func (mr *MockCmdableMockRecorder) ClusterAddSlotsRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlotsRange", reflect.TypeOf((*MockCmdable)(nil).ClusterAddSlotsRange), arg0, arg1, arg2)
}

// ClusterCountFailureReports mocks base method.
func (m *MockCmdable) ClusterCountFailureReports(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountFailureReports", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountFailureReports indicates an expected call of ClusterCountFailureReports.
func (mr *MockCmdableMockRecorder) ClusterCountFailureReports(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountFailureReports", reflect.TypeOf((*MockCmdable)(nil).ClusterCountFailureReports), arg0, arg1)
}

// ClusterCountKeysInSlot mocks base method.
func (m *MockCmdable) ClusterCountKeysInSlot(arg0 context.Context, arg1 int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountKeysInSlot", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountKeysInSlot indicates an expected call of ClusterCountKeysInSlot.
func (mr *MockCmdableMockRecorder) ClusterCountKeysInSlot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountKeysInSlot", reflect.TypeOf((*MockCmdable)(nil).ClusterCountKeysInSlot), arg0, arg1)
}

// ClusterDelSlots mocks base method.
func (m *MockCmdable) ClusterDelSlots(arg0 context.Context, arg1 ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterDelSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlots indicates an expected call of ClusterDelSlots.
func (mr *MockCmdableMockRecorder) ClusterDelSlots(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterDelSlots), varargs...)
}

// ClusterDelSlotsRange mocks base method.
func (m *MockCmdable) ClusterDelSlotsRange(arg0 context.Context, arg1, arg2 int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterDelSlotsRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlotsRange indicates an expected call of ClusterDelSlotsRange.
func (mr *MockCmdableMockRecorder) ClusterDelSlotsRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlotsRange", reflect.TypeOf((*MockCmdable)(nil).ClusterDelSlotsRange), arg0, arg1, arg2)
}

// ClusterFailover mocks base method.
func (m *MockCmdable) ClusterFailover(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterFailover", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterFailover indicates an expected call of ClusterFailover.
func (mr *MockCmdableMockRecorder) ClusterFailover(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterFailover", reflect.TypeOf((*MockCmdable)(nil).ClusterFailover), arg0)
}

// ClusterForget mocks base method.
func (m *MockCmdable) ClusterForget(arg0 context.Context, arg1 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterForget", arg0, arg1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterForget indicates an expected call of ClusterForget.
func (mr *MockCmdableMockRecorder) ClusterForget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterForget", reflect.TypeOf((*MockCmdable)(nil).ClusterForget), arg0, arg1)
}

// ClusterGetKeysInSlot mocks base method.
func (m *MockCmdable) ClusterGetKeysInSlot(arg0 context.Context, arg1, arg2 int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGetKeysInSlot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterGetKeysInSlot indicates an expected call of ClusterGetKeysInSlot.
func (mr *MockCmdableMockRecorder) ClusterGetKeysInSlot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGetKeysInSlot", reflect.TypeOf((*MockCmdable)(nil).ClusterGetKeysInSlot), arg0, arg1, arg2)
}

// ClusterInfo mocks base method.
func (m *MockCmdable) ClusterInfo(arg0 context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterInfo", arg0)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterInfo indicates an expected call of ClusterInfo.
func (mr *MockCmdableMockRecorder) ClusterInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInfo", reflect.TypeOf((*MockCmdable)(nil).ClusterInfo), arg0)
}

// ClusterKeySlot mocks base method.
func (m *MockCmdable) ClusterKeySlot(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterKeySlot", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterKeySlot indicates an expected call of ClusterKeySlot.
func (mr *MockCmdableMockRecorder) ClusterKeySlot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterKeySlot", reflect.TypeOf((*MockCmdable)(nil).ClusterKeySlot), arg0, arg1)
}

// ClusterMeet mocks base method.
func (m *MockCmdable) ClusterMeet(arg0 context.Context, arg1, arg2 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterMeet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterMeet indicates an expected call of ClusterMeet.
func (mr *MockCmdableMockRecorder) ClusterMeet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterMeet", reflect.TypeOf((*MockCmdable)(nil).ClusterMeet), arg0, arg1, arg2)
}

// ClusterNodes mocks base method.
func (m *MockCmdable) ClusterNodes(arg0 context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNodes", arg0)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterNodes indicates an expected call of ClusterNodes.
func (mr *MockCmdableMockRecorder) ClusterNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNodes", reflect.TypeOf((*MockCmdable)(nil).ClusterNodes), arg0)
}

// ClusterReplicate mocks base method.
func (m *MockCmdable) ClusterReplicate(arg0 context.Context, arg1 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterReplicate", arg0, arg1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterReplicate indicates an expected call of ClusterReplicate.
func (mr *MockCmdableMockRecorder) ClusterReplicate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterReplicate", reflect.TypeOf((*MockCmdable)(nil).ClusterReplicate), arg0, arg1)
}

// ClusterResetHard mocks base method.
func (m *MockCmdable) ClusterResetHard(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetHard", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetHard indicates an expected call of ClusterResetHard.
func (mr *MockCmdableMockRecorder) ClusterResetHard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetHard", reflect.TypeOf((*MockCmdable)(nil).ClusterResetHard), arg0)
}

// ClusterResetSoft mocks base method.
func (m *MockCmdable) ClusterResetSoft(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetSoft", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetSoft indicates an expected call of ClusterResetSoft.
func (mr *MockCmdableMockRecorder) ClusterResetSoft(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetSoft", reflect.TypeOf((*MockCmdable)(nil).ClusterResetSoft), arg0)
}

// ClusterSaveConfig mocks base method.
func (m *MockCmdable) ClusterSaveConfig(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSaveConfig", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterSaveConfig indicates an expected call of ClusterSaveConfig.
func (mr *MockCmdableMockRecorder) ClusterSaveConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSaveConfig", reflect.TypeOf((*MockCmdable)(nil).ClusterSaveConfig), arg0)
}

// ClusterSlaves mocks base method.
func (m *MockCmdable) ClusterSlaves(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlaves", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterSlaves indicates an expected call of ClusterSlaves.
func (mr *MockCmdableMockRecorder) ClusterSlaves(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlaves", reflect.TypeOf((*MockCmdable)(nil).ClusterSlaves), arg0, arg1)
}

// ClusterSlots mocks base method.
func (m *MockCmdable) ClusterSlots(arg0 context.Context) *redis.ClusterSlotsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlots", arg0)
	ret0, _ := ret[0].(*redis.ClusterSlotsCmd)
	return ret0
}

// ClusterSlots indicates an expected call of ClusterSlots.
func (mr *MockCmdableMockRecorder) ClusterSlots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterSlots), arg0)
}

// Command mocks base method.
func (m *MockCmdable) Command(arg0 context.Context) *redis.CommandsInfoCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", arg0)
	ret0, _ := ret[0].(*redis.CommandsInfoCmd)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockCmdableMockRecorder) Command(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCmdable)(nil).Command), arg0)
}

// ConfigGet mocks base method.
func (m *MockCmdable) ConfigGet(arg0 context.Context, arg1 string) *redis.MapStringStringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigGet", arg0, arg1)
	ret0, _ := ret[0].(*redis.MapStringStringCmd)
	return ret0
}

// ConfigGet indicates an expected call of ConfigGet.
func (mr *MockCmdableMockRecorder) ConfigGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGet", reflect.TypeOf((*MockCmdable)(nil).ConfigGet), arg0, arg1)
}

// ConfigResetStat mocks base method.
func (m *MockCmdable) ConfigResetStat(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigResetStat", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigResetStat indicates an expected call of ConfigResetStat.
func (mr *MockCmdableMockRecorder) ConfigResetStat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigResetStat", reflect.TypeOf((*MockCmdable)(nil).ConfigResetStat), arg0)
}

// ConfigRewrite mocks base method.
func (m *MockCmdable) ConfigRewrite(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigRewrite", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigRewrite indicates an expected call of ConfigRewrite.
func (mr *MockCmdableMockRecorder) ConfigRewrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigRewrite", reflect.TypeOf((*MockCmdable)(nil).ConfigRewrite), arg0)
}

// ConfigSet mocks base method.
func (m *MockCmdable) ConfigSet(arg0 context.Context, arg1, arg2 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigSet indicates an expected call of ConfigSet.
func (mr *MockCmdableMockRecorder) ConfigSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigSet", reflect.TypeOf((*MockCmdable)(nil).ConfigSet), arg0, arg1, arg2)
}

// Copy mocks base method.
func (m *MockCmdable) Copy(arg0 context.Context, arg1, arg2 string, arg3 int, arg4 bool) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockCmdableMockRecorder) Copy(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockCmdable)(nil).Copy), arg0, arg1, arg2, arg3, arg4)
}

// DBSize mocks base method.
func (m *MockCmdable) DBSize(arg0 context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBSize", arg0)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DBSize indicates an expected call of DBSize.
func (mr *MockCmdableMockRecorder) DBSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBSize", reflect.TypeOf((*MockCmdable)(nil).DBSize), arg0)
}

// DebugObject mocks base method.
func (m *MockCmdable) DebugObject(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugObject", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// DebugObject indicates an expected call of DebugObject.
func (mr *MockCmdableMockRecorder) DebugObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugObject", reflect.TypeOf((*MockCmdable)(nil).DebugObject), arg0, arg1)
}

// Decr mocks base method.
func (m *MockCmdable) Decr(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Decr indicates an expected call of Decr.
func (mr *MockCmdableMockRecorder) Decr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockCmdable)(nil).Decr), arg0, arg1)
}

// DecrBy mocks base method.
func (m *MockCmdable) DecrBy(arg0 context.Context, arg1 string, arg2 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockCmdableMockRecorder) DecrBy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockCmdable)(nil).DecrBy), arg0, arg1, arg2)
}

// Del mocks base method.
func (m *MockCmdable) Del(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockCmdableMockRecorder) Del(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockCmdable)(nil).Del), varargs...)
}

// Dump mocks base method.
func (m *MockCmdable) Dump(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Dump indicates an expected call of Dump.
func (mr *MockCmdableMockRecorder) Dump(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockCmdable)(nil).Dump), arg0, arg1)
}

// Echo mocks base method.
func (m *MockCmdable) Echo(arg0 context.Context, arg1 interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Echo indicates an expected call of Echo.
func (mr *MockCmdableMockRecorder) Echo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockCmdable)(nil).Echo), arg0, arg1)
}

// Eval mocks base method.
func (m *MockCmdable) Eval(arg0 context.Context, arg1 string, arg2 []string, arg3 ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eval", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// Eval indicates an expected call of Eval.
func (mr *MockCmdableMockRecorder) Eval(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockCmdable)(nil).Eval), varargs...)
}

// EvalRO mocks base method.
func (m *MockCmdable) EvalRO(arg0 context.Context, arg1 string, arg2 []string, arg3 ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalRO", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalRO indicates an expected call of EvalRO.
func (mr *MockCmdableMockRecorder) EvalRO(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalRO", reflect.TypeOf((*MockCmdable)(nil).EvalRO), varargs...)
}

// EvalSha mocks base method.
func (m *MockCmdable) EvalSha(arg0 context.Context, arg1 string, arg2 []string, arg3 ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalSha", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalSha indicates an expected call of EvalSha.
func (mr *MockCmdableMockRecorder) EvalSha(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSha", reflect.TypeOf((*MockCmdable)(nil).EvalSha), varargs...)
}

// EvalShaRO mocks base method.
func (m *MockCmdable) EvalShaRO(arg0 context.Context, arg1 string, arg2 []string, arg3 ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalShaRO", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalShaRO indicates an expected call of EvalShaRO.
func (mr *MockCmdableMockRecorder) EvalShaRO(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalShaRO", reflect.TypeOf((*MockCmdable)(nil).EvalShaRO), varargs...)
}

// Exists mocks base method.
func (m *MockCmdable) Exists(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockCmdableMockRecorder) Exists(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockCmdable)(nil).Exists), varargs...)
}

// Expire mocks base method.
func (m *MockCmdable) Expire(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockCmdableMockRecorder) Expire(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockCmdable)(nil).Expire), arg0, arg1, arg2)
}

// ExpireAt mocks base method.
func (m *MockCmdable) ExpireAt(arg0 context.Context, arg1 string, arg2 time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireAt indicates an expected call of ExpireAt.
func (mr *MockCmdableMockRecorder) ExpireAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireAt", reflect.TypeOf((*MockCmdable)(nil).ExpireAt), arg0, arg1, arg2)
}

// ExpireGT mocks base method.
func (m *MockCmdable) ExpireGT(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireGT", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireGT indicates an expected call of ExpireGT.
func (mr *MockCmdableMockRecorder) ExpireGT(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireGT", reflect.TypeOf((*MockCmdable)(nil).ExpireGT), arg0, arg1, arg2)
}

// ExpireLT mocks base method.
func (m *MockCmdable) ExpireLT(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLT", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireLT indicates an expected call of ExpireLT.
func (mr *MockCmdableMockRecorder) ExpireLT(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLT", reflect.TypeOf((*MockCmdable)(nil).ExpireLT), arg0, arg1, arg2)
}

// ExpireNX mocks base method.
func (m *MockCmdable) ExpireNX(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireNX", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireNX indicates an expected call of ExpireNX.
func (mr *MockCmdableMockRecorder) ExpireNX(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireNX", reflect.TypeOf((*MockCmdable)(nil).ExpireNX), arg0, arg1, arg2)
}

// ExpireXX mocks base method.
func (m *MockCmdable) ExpireXX(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireXX", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireXX indicates an expected call of ExpireXX.
func (mr *MockCmdableMockRecorder) ExpireXX(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireXX", reflect.TypeOf((*MockCmdable)(nil).ExpireXX), arg0, arg1, arg2)
}

// FlushAll mocks base method.
func (m *MockCmdable) FlushAll(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockCmdableMockRecorder) FlushAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockCmdable)(nil).FlushAll), arg0)
}

// FlushAllAsync mocks base method.
func (m *MockCmdable) FlushAllAsync(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAllAsync", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAllAsync indicates an expected call of FlushAllAsync.
func (mr *MockCmdableMockRecorder) FlushAllAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAllAsync", reflect.TypeOf((*MockCmdable)(nil).FlushAllAsync), arg0)
}

// FlushDB mocks base method.
func (m *MockCmdable) FlushDB(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDB", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDB indicates an expected call of FlushDB.
func (mr *MockCmdableMockRecorder) FlushDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDB", reflect.TypeOf((*MockCmdable)(nil).FlushDB), arg0)
}

// FlushDBAsync mocks base method.
func (m *MockCmdable) FlushDBAsync(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDBAsync", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDBAsync indicates an expected call of FlushDBAsync.
func (mr *MockCmdableMockRecorder) FlushDBAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDBAsync", reflect.TypeOf((*MockCmdable)(nil).FlushDBAsync), arg0)
}

// GeoAdd mocks base method.
func (m *MockCmdable) GeoAdd(arg0 context.Context, arg1 string, arg2 ...*redis.GeoLocation) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoAdd indicates an expected call of GeoAdd.
func (mr *MockCmdableMockRecorder) GeoAdd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoAdd", reflect.TypeOf((*MockCmdable)(nil).GeoAdd), varargs...)
}

// GeoDist mocks base method.
func (m *MockCmdable) GeoDist(arg0 context.Context, arg1, arg2, arg3, arg4 string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoDist", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// GeoDist indicates an expected call of GeoDist.
func (mr *MockCmdableMockRecorder) GeoDist(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoDist", reflect.TypeOf((*MockCmdable)(nil).GeoDist), arg0, arg1, arg2, arg3, arg4)
}

// GeoHash mocks base method.
func (m *MockCmdable) GeoHash(arg0 context.Context, arg1 string, arg2 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoHash", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoHash indicates an expected call of GeoHash.
func (mr *MockCmdableMockRecorder) GeoHash(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoHash", reflect.TypeOf((*MockCmdable)(nil).GeoHash), varargs...)
}

// GeoPos mocks base method.
func (m *MockCmdable) GeoPos(arg0 context.Context, arg1 string, arg2 ...string) *redis.GeoPosCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoPos", varargs...)
	ret0, _ := ret[0].(*redis.GeoPosCmd)
	return ret0
}

// GeoPos indicates an expected call of GeoPos.
func (mr *MockCmdableMockRecorder) GeoPos(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoPos", reflect.TypeOf((*MockCmdable)(nil).GeoPos), varargs...)
}

// GeoRadius mocks base method.
func (m *MockCmdable) GeoRadius(arg0 context.Context, arg1 string, arg2, arg3 float64, arg4 *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadius", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadius indicates an expected call of GeoRadius.
func (mr *MockCmdableMockRecorder) GeoRadius(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadius", reflect.TypeOf((*MockCmdable)(nil).GeoRadius), arg0, arg1, arg2, arg3, arg4)
}

// GeoRadiusByMember mocks base method.
func (m *MockCmdable) GeoRadiusByMember(arg0 context.Context, arg1, arg2 string, arg3 *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadiusByMember indicates an expected call of GeoRadiusByMember.
func (mr *MockCmdableMockRecorder) GeoRadiusByMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMember", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusByMember), arg0, arg1, arg2, arg3)
}

// GeoRadiusByMemberStore mocks base method.
func (m *MockCmdable) GeoRadiusByMemberStore(arg0 context.Context, arg1, arg2 string, arg3 *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMemberStore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusByMemberStore indicates an expected call of GeoRadiusByMemberStore.
func (mr *MockCmdableMockRecorder) GeoRadiusByMemberStore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMemberStore", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusByMemberStore), arg0, arg1, arg2, arg3)
}

// GeoRadiusStore mocks base method.
func (m *MockCmdable) GeoRadiusStore(arg0 context.Context, arg1 string, arg2, arg3 float64, arg4 *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusStore", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusStore indicates an expected call of GeoRadiusStore.
func (mr *MockCmdableMockRecorder) GeoRadiusStore(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusStore", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusStore), arg0, arg1, arg2, arg3, arg4)
}

// GeoSearch mocks base method.
func (m *MockCmdable) GeoSearch(arg0 context.Context, arg1 string, arg2 *redis.GeoSearchQuery) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoSearch indicates an expected call of GeoSearch.
func (mr *MockCmdableMockRecorder) GeoSearch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearch", reflect.TypeOf((*MockCmdable)(nil).GeoSearch), arg0, arg1, arg2)
}

// GeoSearchLocation mocks base method.
func (m *MockCmdable) GeoSearchLocation(arg0 context.Context, arg1 string, arg2 *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.GeoSearchLocationCmd)
	return ret0
}

// GeoSearchLocation indicates an expected call of GeoSearchLocation.
func (mr *MockCmdableMockRecorder) GeoSearchLocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchLocation", reflect.TypeOf((*MockCmdable)(nil).GeoSearchLocation), arg0, arg1, arg2)
}

// GeoSearchStore mocks base method.
func (m *MockCmdable) GeoSearchStore(arg0 context.Context, arg1, arg2 string, arg3 *redis.GeoSearchStoreQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchStore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoSearchStore indicates an expected call of GeoSearchStore.
func (mr *MockCmdableMockRecorder) GeoSearchStore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchStore", reflect.TypeOf((*MockCmdable)(nil).GeoSearchStore), arg0, arg1, arg2, arg3)
}

// Get mocks base method.
func (m *MockCmdable) Get(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCmdableMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCmdable)(nil).Get), arg0, arg1)
}

// GetBit mocks base method.
func (m *MockCmdable) GetBit(arg0 context.Context, arg1 string, arg2 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBit", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GetBit indicates an expected call of GetBit.
func (mr *MockCmdableMockRecorder) GetBit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBit", reflect.TypeOf((*MockCmdable)(nil).GetBit), arg0, arg1, arg2)
}

// GetDel mocks base method.
func (m *MockCmdable) GetDel(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetDel indicates an expected call of GetDel.
func (mr *MockCmdableMockRecorder) GetDel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockCmdable)(nil).GetDel), arg0, arg1)
}

// GetEx mocks base method.
func (m *MockCmdable) GetEx(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEx", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetEx indicates an expected call of GetEx.
func (mr *MockCmdableMockRecorder) GetEx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEx", reflect.TypeOf((*MockCmdable)(nil).GetEx), arg0, arg1, arg2)
}

// GetRange mocks base method.
func (m *MockCmdable) GetRange(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetRange indicates an expected call of GetRange.
func (mr *MockCmdableMockRecorder) GetRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRange", reflect.TypeOf((*MockCmdable)(nil).GetRange), arg0, arg1, arg2, arg3)
}

// GetSet mocks base method.
func (m *MockCmdable) GetSet(arg0 context.Context, arg1 string, arg2 interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetSet indicates an expected call of GetSet.
func (mr *MockCmdableMockRecorder) GetSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSet", reflect.TypeOf((*MockCmdable)(nil).GetSet), arg0, arg1, arg2)
}

// HDel mocks base method.
func (m *MockCmdable) HDel(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HDel indicates an expected call of HDel.
func (mr *MockCmdableMockRecorder) HDel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockCmdable)(nil).HDel), varargs...)
}

// HExists mocks base method.
func (m *MockCmdable) HExists(arg0 context.Context, arg1, arg2 string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HExists", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HExists indicates an expected call of HExists.
func (mr *MockCmdableMockRecorder) HExists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HExists", reflect.TypeOf((*MockCmdable)(nil).HExists), arg0, arg1, arg2)
}

// HGet mocks base method.
func (m *MockCmdable) HGet(arg0 context.Context, arg1, arg2 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// HGet indicates an expected call of HGet.
func (mr *MockCmdableMockRecorder) HGet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockCmdable)(nil).HGet), arg0, arg1, arg2)
}

// HGetAll mocks base method.
func (m *MockCmdable) HGetAll(arg0 context.Context, arg1 string) *redis.MapStringStringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", arg0, arg1)
	ret0, _ := ret[0].(*redis.MapStringStringCmd)
	return ret0
}

// HGetAll indicates an expected call of HGetAll.
func (mr *MockCmdableMockRecorder) HGetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockCmdable)(nil).HGetAll), arg0, arg1)
}

// HIncrBy mocks base method.
func (m *MockCmdable) HIncrBy(arg0 context.Context, arg1, arg2 string, arg3 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrBy", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HIncrBy indicates an expected call of HIncrBy.
func (mr *MockCmdableMockRecorder) HIncrBy(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrBy", reflect.TypeOf((*MockCmdable)(nil).HIncrBy), arg0, arg1, arg2, arg3)
}

// HIncrByFloat mocks base method.
func (m *MockCmdable) HIncrByFloat(arg0 context.Context, arg1, arg2 string, arg3 float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrByFloat", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// HIncrByFloat indicates an expected call of HIncrByFloat.
func (mr *MockCmdableMockRecorder) HIncrByFloat(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrByFloat", reflect.TypeOf((*MockCmdable)(nil).HIncrByFloat), arg0, arg1, arg2, arg3)
}

// HKeys mocks base method.
func (m *MockCmdable) HKeys(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HKeys", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HKeys indicates an expected call of HKeys.
func (mr *MockCmdableMockRecorder) HKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HKeys", reflect.TypeOf((*MockCmdable)(nil).HKeys), arg0, arg1)
}

// HLen mocks base method.
func (m *MockCmdable) HLen(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HLen", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HLen indicates an expected call of HLen.
func (mr *MockCmdableMockRecorder) HLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HLen", reflect.TypeOf((*MockCmdable)(nil).HLen), arg0, arg1)
}

// HMGet mocks base method.
func (m *MockCmdable) HMGet(arg0 context.Context, arg1 string, arg2 ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// HMGet indicates an expected call of HMGet.
func (mr *MockCmdableMockRecorder) HMGet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockCmdable)(nil).HMGet), varargs...)
}

// HMSet mocks base method.
func (m *MockCmdable) HMSet(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMSet", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HMSet indicates an expected call of HMSet.
func (mr *MockCmdableMockRecorder) HMSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockCmdable)(nil).HMSet), varargs...)
}

// HRandField mocks base method.
func (m *MockCmdable) HRandField(arg0 context.Context, arg1 string, arg2 int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HRandField", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HRandField indicates an expected call of HRandField.
func (mr *MockCmdableMockRecorder) HRandField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HRandField", reflect.TypeOf((*MockCmdable)(nil).HRandField), arg0, arg1, arg2)
}

// HRandFieldWithValues mocks base method.
func (m *MockCmdable) HRandFieldWithValues(arg0 context.Context, arg1 string, arg2 int) *redis.KeyValueSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HRandFieldWithValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.KeyValueSliceCmd)
	return ret0
}

// HRandFieldWithValues indicates an expected call of HRandFieldWithValues.
func (mr *MockCmdableMockRecorder) HRandFieldWithValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HRandFieldWithValues", reflect.TypeOf((*MockCmdable)(nil).HRandFieldWithValues), arg0, arg1, arg2)
}

// HScan mocks base method.
func (m *MockCmdable) HScan(arg0 context.Context, arg1 string, arg2 uint64, arg3 string, arg4 int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HScan", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// HScan indicates an expected call of HScan.
func (mr *MockCmdableMockRecorder) HScan(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HScan", reflect.TypeOf((*MockCmdable)(nil).HScan), arg0, arg1, arg2, arg3, arg4)
}

// HSet mocks base method.
func (m *MockCmdable) HSet(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HSet", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockCmdableMockRecorder) HSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockCmdable)(nil).HSet), varargs...)
}

// HSetNX mocks base method.
func (m *MockCmdable) HSetNX(arg0 context.Context, arg1, arg2 string, arg3 interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSetNX", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HSetNX indicates an expected call of HSetNX.
func (mr *MockCmdableMockRecorder) HSetNX(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSetNX", reflect.TypeOf((*MockCmdable)(nil).HSetNX), arg0, arg1, arg2, arg3)
}

// HVals mocks base method.
func (m *MockCmdable) HVals(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HVals", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HVals indicates an expected call of HVals.
func (mr *MockCmdableMockRecorder) HVals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HVals", reflect.TypeOf((*MockCmdable)(nil).HVals), arg0, arg1)
}

// Incr mocks base method.
func (m *MockCmdable) Incr(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockCmdableMockRecorder) Incr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockCmdable)(nil).Incr), arg0, arg1)
}

// IncrBy mocks base method.
func (m *MockCmdable) IncrBy(arg0 context.Context, arg1 string, arg2 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockCmdableMockRecorder) IncrBy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockCmdable)(nil).IncrBy), arg0, arg1, arg2)
}

// IncrByFloat mocks base method.
func (m *MockCmdable) IncrByFloat(arg0 context.Context, arg1 string, arg2 float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrByFloat", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// IncrByFloat indicates an expected call of IncrByFloat.
func (mr *MockCmdableMockRecorder) IncrByFloat(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrByFloat", reflect.TypeOf((*MockCmdable)(nil).IncrByFloat), arg0, arg1, arg2)
}

// Info mocks base method.
func (m *MockCmdable) Info(arg0 context.Context, arg1 ...string) *redis.StringCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockCmdableMockRecorder) Info(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockCmdable)(nil).Info), varargs...)
}

// Keys mocks base method.
func (m *MockCmdable) Keys(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockCmdableMockRecorder) Keys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCmdable)(nil).Keys), arg0, arg1)
}

// LIndex mocks base method.
func (m *MockCmdable) LIndex(arg0 context.Context, arg1 string, arg2 int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LIndex indicates an expected call of LIndex.
func (mr *MockCmdableMockRecorder) LIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LIndex", reflect.TypeOf((*MockCmdable)(nil).LIndex), arg0, arg1, arg2)
}

// LInsert mocks base method.
func (m *MockCmdable) LInsert(arg0 context.Context, arg1, arg2 string, arg3, arg4 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsert", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsert indicates an expected call of LInsert.
func (mr *MockCmdableMockRecorder) LInsert(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsert", reflect.TypeOf((*MockCmdable)(nil).LInsert), arg0, arg1, arg2, arg3, arg4)
}

// LInsertAfter mocks base method.
func (m *MockCmdable) LInsertAfter(arg0 context.Context, arg1 string, arg2, arg3 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertAfter", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertAfter indicates an expected call of LInsertAfter.
func (mr *MockCmdableMockRecorder) LInsertAfter(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertAfter", reflect.TypeOf((*MockCmdable)(nil).LInsertAfter), arg0, arg1, arg2, arg3)
}

// LInsertBefore mocks base method.
func (m *MockCmdable) LInsertBefore(arg0 context.Context, arg1 string, arg2, arg3 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertBefore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertBefore indicates an expected call of LInsertBefore.
func (mr *MockCmdableMockRecorder) LInsertBefore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertBefore", reflect.TypeOf((*MockCmdable)(nil).LInsertBefore), arg0, arg1, arg2, arg3)
}

// LLen mocks base method.
func (m *MockCmdable) LLen(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LLen", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LLen indicates an expected call of LLen.
func (mr *MockCmdableMockRecorder) LLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LLen", reflect.TypeOf((*MockCmdable)(nil).LLen), arg0, arg1)
}

// LMove mocks base method.
func (m *MockCmdable) LMove(arg0 context.Context, arg1, arg2, arg3, arg4 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LMove", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LMove indicates an expected call of LMove.
func (mr *MockCmdableMockRecorder) LMove(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LMove", reflect.TypeOf((*MockCmdable)(nil).LMove), arg0, arg1, arg2, arg3, arg4)
}

// LPop mocks base method.
func (m *MockCmdable) LPop(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LPop indicates an expected call of LPop.
func (mr *MockCmdableMockRecorder) LPop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockCmdable)(nil).LPop), arg0, arg1)
}

// LPopCount mocks base method.
func (m *MockCmdable) LPopCount(arg0 context.Context, arg1 string, arg2 int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPopCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LPopCount indicates an expected call of LPopCount.
func (mr *MockCmdableMockRecorder) LPopCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPopCount", reflect.TypeOf((*MockCmdable)(nil).LPopCount), arg0, arg1, arg2)
}

// LPos mocks base method.
func (m *MockCmdable) LPos(arg0 context.Context, arg1, arg2 string, arg3 redis.LPosArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPos", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPos indicates an expected call of LPos.
func (mr *MockCmdableMockRecorder) LPos(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPos", reflect.TypeOf((*MockCmdable)(nil).LPos), arg0, arg1, arg2, arg3)
}

// LPosCount mocks base method.
func (m *MockCmdable) LPosCount(arg0 context.Context, arg1, arg2 string, arg3 int64, arg4 redis.LPosArgs) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPosCount", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// LPosCount indicates an expected call of LPosCount.
func (mr *MockCmdableMockRecorder) LPosCount(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPosCount", reflect.TypeOf((*MockCmdable)(nil).LPosCount), arg0, arg1, arg2, arg3, arg4)
}

// LPush mocks base method.
func (m *MockCmdable) LPush(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPush indicates an expected call of LPush.
func (mr *MockCmdableMockRecorder) LPush(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockCmdable)(nil).LPush), varargs...)
}

// LPushX mocks base method.
func (m *MockCmdable) LPushX(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPushX indicates an expected call of LPushX.
func (mr *MockCmdableMockRecorder) LPushX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPushX", reflect.TypeOf((*MockCmdable)(nil).LPushX), varargs...)
}

// LRange mocks base method.
func (m *MockCmdable) LRange(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LRange indicates an expected call of LRange.
func (mr *MockCmdableMockRecorder) LRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRange", reflect.TypeOf((*MockCmdable)(nil).LRange), arg0, arg1, arg2, arg3)
}

// LRem mocks base method.
func (m *MockCmdable) LRem(arg0 context.Context, arg1 string, arg2 int64, arg3 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LRem indicates an expected call of LRem.
func (mr *MockCmdableMockRecorder) LRem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRem", reflect.TypeOf((*MockCmdable)(nil).LRem), arg0, arg1, arg2, arg3)
}

// LSet mocks base method.
func (m *MockCmdable) LSet(arg0 context.Context, arg1 string, arg2 int64, arg3 interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LSet indicates an expected call of LSet.
func (mr *MockCmdableMockRecorder) LSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LSet", reflect.TypeOf((*MockCmdable)(nil).LSet), arg0, arg1, arg2, arg3)
}

// LTrim mocks base method.
func (m *MockCmdable) LTrim(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LTrim", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LTrim indicates an expected call of LTrim.
func (mr *MockCmdableMockRecorder) LTrim(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LTrim", reflect.TypeOf((*MockCmdable)(nil).LTrim), arg0, arg1, arg2, arg3)
}

// LastSave mocks base method.
func (m *MockCmdable) LastSave(arg0 context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSave", arg0)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LastSave indicates an expected call of LastSave.
func (mr *MockCmdableMockRecorder) LastSave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSave", reflect.TypeOf((*MockCmdable)(nil).LastSave), arg0)
}

// MGet mocks base method.
func (m *MockCmdable) MGet(arg0 context.Context, arg1 ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// MGet indicates an expected call of MGet.
func (mr *MockCmdableMockRecorder) MGet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockCmdable)(nil).MGet), varargs...)
}

// MSet mocks base method.
func (m *MockCmdable) MSet(arg0 context.Context, arg1 ...interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSet", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// MSet indicates an expected call of MSet.
func (mr *MockCmdableMockRecorder) MSet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSet", reflect.TypeOf((*MockCmdable)(nil).MSet), varargs...)
}

// MSetNX mocks base method.
func (m *MockCmdable) MSetNX(arg0 context.Context, arg1 ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSetNX", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// MSetNX indicates an expected call of MSetNX.
func (mr *MockCmdableMockRecorder) MSetNX(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSetNX", reflect.TypeOf((*MockCmdable)(nil).MSetNX), varargs...)
}

// MemoryUsage mocks base method.
func (m *MockCmdable) MemoryUsage(arg0 context.Context, arg1 string, arg2 ...int) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MemoryUsage", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// MemoryUsage indicates an expected call of MemoryUsage.
func (mr *MockCmdableMockRecorder) MemoryUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemoryUsage", reflect.TypeOf((*MockCmdable)(nil).MemoryUsage), varargs...)
}

// Migrate mocks base method.
func (m *MockCmdable) Migrate(arg0 context.Context, arg1, arg2, arg3 string, arg4 int, arg5 time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockCmdableMockRecorder) Migrate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockCmdable)(nil).Migrate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Move mocks base method.
func (m *MockCmdable) Move(arg0 context.Context, arg1 string, arg2 int) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockCmdableMockRecorder) Move(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockCmdable)(nil).Move), arg0, arg1, arg2)
}

// ObjectEncoding mocks base method.
func (m *MockCmdable) ObjectEncoding(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectEncoding", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ObjectEncoding indicates an expected call of ObjectEncoding.
func (mr *MockCmdableMockRecorder) ObjectEncoding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectEncoding", reflect.TypeOf((*MockCmdable)(nil).ObjectEncoding), arg0, arg1)
}

// ObjectIdleTime mocks base method.
func (m *MockCmdable) ObjectIdleTime(arg0 context.Context, arg1 string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectIdleTime", arg0, arg1)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// ObjectIdleTime indicates an expected call of ObjectIdleTime.
func (mr *MockCmdableMockRecorder) ObjectIdleTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectIdleTime", reflect.TypeOf((*MockCmdable)(nil).ObjectIdleTime), arg0, arg1)
}

// ObjectRefCount mocks base method.
func (m *MockCmdable) ObjectRefCount(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectRefCount", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ObjectRefCount indicates an expected call of ObjectRefCount.
func (mr *MockCmdableMockRecorder) ObjectRefCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectRefCount", reflect.TypeOf((*MockCmdable)(nil).ObjectRefCount), arg0, arg1)
}

// PExpire mocks base method.
func (m *MockCmdable) PExpire(arg0 context.Context, arg1 string, arg2 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpire", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpire indicates an expected call of PExpire.
func (mr *MockCmdableMockRecorder) PExpire(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpire", reflect.TypeOf((*MockCmdable)(nil).PExpire), arg0, arg1, arg2)
}

// PExpireAt mocks base method.
func (m *MockCmdable) PExpireAt(arg0 context.Context, arg1 string, arg2 time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpireAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpireAt indicates an expected call of PExpireAt.
func (mr *MockCmdableMockRecorder) PExpireAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpireAt", reflect.TypeOf((*MockCmdable)(nil).PExpireAt), arg0, arg1, arg2)
}

// PFAdd mocks base method.
func (m *MockCmdable) PFAdd(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFAdd indicates an expected call of PFAdd.
func (mr *MockCmdableMockRecorder) PFAdd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFAdd", reflect.TypeOf((*MockCmdable)(nil).PFAdd), varargs...)
}

// PFCount mocks base method.
func (m *MockCmdable) PFCount(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFCount", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFCount indicates an expected call of PFCount.
func (mr *MockCmdableMockRecorder) PFCount(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFCount", reflect.TypeOf((*MockCmdable)(nil).PFCount), varargs...)
}

// PFMerge mocks base method.
func (m *MockCmdable) PFMerge(arg0 context.Context, arg1 string, arg2 ...string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFMerge", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// PFMerge indicates an expected call of PFMerge.
func (mr *MockCmdableMockRecorder) PFMerge(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFMerge", reflect.TypeOf((*MockCmdable)(nil).PFMerge), varargs...)
}

// PTTL mocks base method.
func (m *MockCmdable) PTTL(arg0 context.Context, arg1 string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PTTL", arg0, arg1)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// PTTL indicates an expected call of PTTL.
func (mr *MockCmdableMockRecorder) PTTL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PTTL", reflect.TypeOf((*MockCmdable)(nil).PTTL), arg0, arg1)
}

// Persist mocks base method.
func (m *MockCmdable) Persist(arg0 context.Context, arg1 string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", arg0, arg1)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockCmdableMockRecorder) Persist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockCmdable)(nil).Persist), arg0, arg1)
}

// Ping mocks base method.
func (m *MockCmdable) Ping(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockCmdableMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCmdable)(nil).Ping), arg0)
}

// Pipeline mocks base method.
func (m *MockCmdable) Pipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockCmdableMockRecorder) Pipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockCmdable)(nil).Pipeline))
}

// Pipelined mocks base method.
func (m *MockCmdable) Pipelined(arg0 context.Context, arg1 func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelined", arg0, arg1)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelined indicates an expected call of Pipelined.
func (mr *MockCmdableMockRecorder) Pipelined(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelined", reflect.TypeOf((*MockCmdable)(nil).Pipelined), arg0, arg1)
}

// PubSubChannels mocks base method.
func (m *MockCmdable) PubSubChannels(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubChannels", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// PubSubChannels indicates an expected call of PubSubChannels.
func (mr *MockCmdableMockRecorder) PubSubChannels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubChannels", reflect.TypeOf((*MockCmdable)(nil).PubSubChannels), arg0, arg1)
}

// PubSubNumPat mocks base method.
func (m *MockCmdable) PubSubNumPat(arg0 context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubNumPat", arg0)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PubSubNumPat indicates an expected call of PubSubNumPat.
func (mr *MockCmdableMockRecorder) PubSubNumPat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumPat", reflect.TypeOf((*MockCmdable)(nil).PubSubNumPat), arg0)
}

// PubSubNumSub mocks base method.
func (m *MockCmdable) PubSubNumSub(arg0 context.Context, arg1 ...string) *redis.MapStringIntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PubSubNumSub", varargs...)
	ret0, _ := ret[0].(*redis.MapStringIntCmd)
	return ret0
}

// PubSubNumSub indicates an expected call of PubSubNumSub.
func (mr *MockCmdableMockRecorder) PubSubNumSub(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumSub", reflect.TypeOf((*MockCmdable)(nil).PubSubNumSub), varargs...)
}

// PubSubShardChannels mocks base method.
func (m *MockCmdable) PubSubShardChannels(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubShardChannels", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// PubSubShardChannels indicates an expected call of PubSubShardChannels.
func (mr *MockCmdableMockRecorder) PubSubShardChannels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubShardChannels", reflect.TypeOf((*MockCmdable)(nil).PubSubShardChannels), arg0, arg1)
}

// PubSubShardNumSub mocks base method.
func (m *MockCmdable) PubSubShardNumSub(arg0 context.Context, arg1 ...string) *redis.MapStringIntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PubSubShardNumSub", varargs...)
	ret0, _ := ret[0].(*redis.MapStringIntCmd)
	return ret0
}

// PubSubShardNumSub indicates an expected call of PubSubShardNumSub.
func (mr *MockCmdableMockRecorder) PubSubShardNumSub(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubShardNumSub", reflect.TypeOf((*MockCmdable)(nil).PubSubShardNumSub), varargs...)
}

// Publish mocks base method.
func (m *MockCmdable) Publish(arg0 context.Context, arg1 string, arg2 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockCmdableMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockCmdable)(nil).Publish), arg0, arg1, arg2)
}

// Quit mocks base method.
func (m *MockCmdable) Quit(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockCmdableMockRecorder) Quit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockCmdable)(nil).Quit), arg0)
}

// RPop mocks base method.
func (m *MockCmdable) RPop(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPop", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPop indicates an expected call of RPop.
func (mr *MockCmdableMockRecorder) RPop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPop", reflect.TypeOf((*MockCmdable)(nil).RPop), arg0, arg1)
}

// RPopCount mocks base method.
func (m *MockCmdable) RPopCount(arg0 context.Context, arg1 string, arg2 int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// RPopCount indicates an expected call of RPopCount.
func (mr *MockCmdableMockRecorder) RPopCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopCount", reflect.TypeOf((*MockCmdable)(nil).RPopCount), arg0, arg1, arg2)
}

// RPopLPush mocks base method.
func (m *MockCmdable) RPopLPush(arg0 context.Context, arg1, arg2 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopLPush", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPopLPush indicates an expected call of RPopLPush.
func (mr *MockCmdableMockRecorder) RPopLPush(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopLPush", reflect.TypeOf((*MockCmdable)(nil).RPopLPush), arg0, arg1, arg2)
}

// RPush mocks base method.
func (m *MockCmdable) RPush(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPush indicates an expected call of RPush.
func (mr *MockCmdableMockRecorder) RPush(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockCmdable)(nil).RPush), varargs...)
}

// RPushX mocks base method.
func (m *MockCmdable) RPushX(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPushX indicates an expected call of RPushX.
func (mr *MockCmdableMockRecorder) RPushX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPushX", reflect.TypeOf((*MockCmdable)(nil).RPushX), varargs...)
}

// RandomKey mocks base method.
func (m *MockCmdable) RandomKey(arg0 context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomKey", arg0)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RandomKey indicates an expected call of RandomKey.
func (mr *MockCmdableMockRecorder) RandomKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomKey", reflect.TypeOf((*MockCmdable)(nil).RandomKey), arg0)
}

// ReadOnly mocks base method.
func (m *MockCmdable) ReadOnly(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockCmdableMockRecorder) ReadOnly(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockCmdable)(nil).ReadOnly), arg0)
}

// ReadWrite mocks base method.
func (m *MockCmdable) ReadWrite(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWrite", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadWrite indicates an expected call of ReadWrite.
func (mr *MockCmdableMockRecorder) ReadWrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWrite", reflect.TypeOf((*MockCmdable)(nil).ReadWrite), arg0)
}

// Rename mocks base method.
func (m *MockCmdable) Rename(arg0 context.Context, arg1, arg2 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockCmdableMockRecorder) Rename(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockCmdable)(nil).Rename), arg0, arg1, arg2)
}

// RenameNX mocks base method.
func (m *MockCmdable) RenameNX(arg0 context.Context, arg1, arg2 string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameNX", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// RenameNX indicates an expected call of RenameNX.
func (mr *MockCmdableMockRecorder) RenameNX(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameNX", reflect.TypeOf((*MockCmdable)(nil).RenameNX), arg0, arg1, arg2)
}

// Restore mocks base method.
func (m *MockCmdable) Restore(arg0 context.Context, arg1 string, arg2 time.Duration, arg3 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockCmdableMockRecorder) Restore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockCmdable)(nil).Restore), arg0, arg1, arg2, arg3)
}

// RestoreReplace mocks base method.
func (m *MockCmdable) RestoreReplace(arg0 context.Context, arg1 string, arg2 time.Duration, arg3 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreReplace", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// RestoreReplace indicates an expected call of RestoreReplace.
func (mr *MockCmdableMockRecorder) RestoreReplace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreReplace", reflect.TypeOf((*MockCmdable)(nil).RestoreReplace), arg0, arg1, arg2, arg3)
}

// SAdd mocks base method.
func (m *MockCmdable) SAdd(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockCmdableMockRecorder) SAdd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockCmdable)(nil).SAdd), varargs...)
}

// SCard mocks base method.
func (m *MockCmdable) SCard(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SCard", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SCard indicates an expected call of SCard.
func (mr *MockCmdableMockRecorder) SCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SCard", reflect.TypeOf((*MockCmdable)(nil).SCard), arg0, arg1)
}

// SDiff mocks base method.
func (m *MockCmdable) SDiff(arg0 context.Context, arg1 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SDiff indicates an expected call of SDiff.
func (mr *MockCmdableMockRecorder) SDiff(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiff", reflect.TypeOf((*MockCmdable)(nil).SDiff), varargs...)
}

// SDiffStore mocks base method.
func (m *MockCmdable) SDiffStore(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SDiffStore indicates an expected call of SDiffStore.
func (mr *MockCmdableMockRecorder) SDiffStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiffStore", reflect.TypeOf((*MockCmdable)(nil).SDiffStore), varargs...)
}

// SInter mocks base method.
func (m *MockCmdable) SInter(arg0 context.Context, arg1 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInter", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SInter indicates an expected call of SInter.
func (mr *MockCmdableMockRecorder) SInter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInter", reflect.TypeOf((*MockCmdable)(nil).SInter), varargs...)
}

// SInterStore mocks base method.
func (m *MockCmdable) SInterStore(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInterStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SInterStore indicates an expected call of SInterStore.
func (mr *MockCmdableMockRecorder) SInterStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInterStore", reflect.TypeOf((*MockCmdable)(nil).SInterStore), varargs...)
}

// SIsMember mocks base method.
func (m *MockCmdable) SIsMember(arg0 context.Context, arg1 string, arg2 interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIsMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SIsMember indicates an expected call of SIsMember.
func (mr *MockCmdableMockRecorder) SIsMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIsMember", reflect.TypeOf((*MockCmdable)(nil).SIsMember), arg0, arg1, arg2)
}

// SMIsMember mocks base method.
func (m *MockCmdable) SMIsMember(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SMIsMember", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// SMIsMember indicates an expected call of SMIsMember.
func (mr *MockCmdableMockRecorder) SMIsMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMIsMember", reflect.TypeOf((*MockCmdable)(nil).SMIsMember), varargs...)
}

// SMembers mocks base method.
func (m *MockCmdable) SMembers(arg0 context.Context, arg1 string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SMembers indicates an expected call of SMembers.
func (mr *MockCmdableMockRecorder) SMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockCmdable)(nil).SMembers), arg0, arg1)
}

// SMembersMap mocks base method.
func (m *MockCmdable) SMembersMap(arg0 context.Context, arg1 string) *redis.StringStructMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembersMap", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringStructMapCmd)
	return ret0
}

// SMembersMap indicates an expected call of SMembersMap.
func (mr *MockCmdableMockRecorder) SMembersMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembersMap", reflect.TypeOf((*MockCmdable)(nil).SMembersMap), arg0, arg1)
}

// SMove mocks base method.
func (m *MockCmdable) SMove(arg0 context.Context, arg1, arg2 string, arg3 interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMove", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SMove indicates an expected call of SMove.
func (mr *MockCmdableMockRecorder) SMove(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMove", reflect.TypeOf((*MockCmdable)(nil).SMove), arg0, arg1, arg2, arg3)
}

// SPop mocks base method.
func (m *MockCmdable) SPop(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPop", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SPop indicates an expected call of SPop.
func (mr *MockCmdableMockRecorder) SPop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPop", reflect.TypeOf((*MockCmdable)(nil).SPop), arg0, arg1)
}

// SPopN mocks base method.
func (m *MockCmdable) SPopN(arg0 context.Context, arg1 string, arg2 int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPopN", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SPopN indicates an expected call of SPopN.
func (mr *MockCmdableMockRecorder) SPopN(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPopN", reflect.TypeOf((*MockCmdable)(nil).SPopN), arg0, arg1, arg2)
}

// SPublish mocks base method.
func (m *MockCmdable) SPublish(arg0 context.Context, arg1 string, arg2 interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPublish", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SPublish indicates an expected call of SPublish.
func (mr *MockCmdableMockRecorder) SPublish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPublish", reflect.TypeOf((*MockCmdable)(nil).SPublish), arg0, arg1, arg2)
}

// SRandMember mocks base method.
func (m *MockCmdable) SRandMember(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMember", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SRandMember indicates an expected call of SRandMember.
func (mr *MockCmdableMockRecorder) SRandMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMember", reflect.TypeOf((*MockCmdable)(nil).SRandMember), arg0, arg1)
}

// SRandMemberN mocks base method.
func (m *MockCmdable) SRandMemberN(arg0 context.Context, arg1 string, arg2 int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMemberN", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SRandMemberN indicates an expected call of SRandMemberN.
func (mr *MockCmdableMockRecorder) SRandMemberN(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMemberN", reflect.TypeOf((*MockCmdable)(nil).SRandMemberN), arg0, arg1, arg2)
}

// SRem mocks base method.
func (m *MockCmdable) SRem(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SRem indicates an expected call of SRem.
func (mr *MockCmdableMockRecorder) SRem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockCmdable)(nil).SRem), varargs...)
}

// SScan mocks base method.
func (m *MockCmdable) SScan(arg0 context.Context, arg1 string, arg2 uint64, arg3 string, arg4 int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SScan", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// SScan indicates an expected call of SScan.
func (mr *MockCmdableMockRecorder) SScan(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SScan", reflect.TypeOf((*MockCmdable)(nil).SScan), arg0, arg1, arg2, arg3, arg4)
}

// SUnion mocks base method.
func (m *MockCmdable) SUnion(arg0 context.Context, arg1 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnion", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SUnion indicates an expected call of SUnion.
func (mr *MockCmdableMockRecorder) SUnion(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnion", reflect.TypeOf((*MockCmdable)(nil).SUnion), varargs...)
}

// SUnionStore mocks base method.
func (m *MockCmdable) SUnionStore(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnionStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SUnionStore indicates an expected call of SUnionStore.
func (mr *MockCmdableMockRecorder) SUnionStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnionStore", reflect.TypeOf((*MockCmdable)(nil).SUnionStore), varargs...)
}

// Save mocks base method.
func (m *MockCmdable) Save(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockCmdableMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCmdable)(nil).Save), arg0)
}

// Scan mocks base method.
func (m *MockCmdable) Scan(arg0 context.Context, arg1 uint64, arg2 string, arg3 int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockCmdableMockRecorder) Scan(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockCmdable)(nil).Scan), arg0, arg1, arg2, arg3)
}

// ScanType mocks base method.
func (m *MockCmdable) ScanType(arg0 context.Context, arg1 uint64, arg2 string, arg3 int64, arg4 string) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanType", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ScanType indicates an expected call of ScanType.
func (mr *MockCmdableMockRecorder) ScanType(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanType", reflect.TypeOf((*MockCmdable)(nil).ScanType), arg0, arg1, arg2, arg3, arg4)
}

// ScriptExists mocks base method.
func (m *MockCmdable) ScriptExists(arg0 context.Context, arg1 ...string) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScriptExists", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// ScriptExists indicates an expected call of ScriptExists.
func (mr *MockCmdableMockRecorder) ScriptExists(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptExists", reflect.TypeOf((*MockCmdable)(nil).ScriptExists), varargs...)
}

// ScriptFlush mocks base method.
func (m *MockCmdable) ScriptFlush(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptFlush", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptFlush indicates an expected call of ScriptFlush.
func (mr *MockCmdableMockRecorder) ScriptFlush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptFlush", reflect.TypeOf((*MockCmdable)(nil).ScriptFlush), arg0)
}

// ScriptKill mocks base method.
func (m *MockCmdable) ScriptKill(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptKill", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptKill indicates an expected call of ScriptKill.
func (mr *MockCmdableMockRecorder) ScriptKill(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptKill", reflect.TypeOf((*MockCmdable)(nil).ScriptKill), arg0)
}

// ScriptLoad mocks base method.
func (m *MockCmdable) ScriptLoad(arg0 context.Context, arg1 string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptLoad", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ScriptLoad indicates an expected call of ScriptLoad.
func (mr *MockCmdableMockRecorder) ScriptLoad(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptLoad", reflect.TypeOf((*MockCmdable)(nil).ScriptLoad), arg0, arg1)
}

// Set mocks base method.
func (m *MockCmdable) Set(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCmdableMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCmdable)(nil).Set), arg0, arg1, arg2, arg3)
}

// SetArgs mocks base method.
func (m *MockCmdable) SetArgs(arg0 context.Context, arg1 string, arg2 interface{}, arg3 redis.SetArgs) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArgs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetArgs indicates an expected call of SetArgs.
func (mr *MockCmdableMockRecorder) SetArgs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArgs", reflect.TypeOf((*MockCmdable)(nil).SetArgs), arg0, arg1, arg2, arg3)
}

// SetBit mocks base method.
func (m *MockCmdable) SetBit(arg0 context.Context, arg1 string, arg2 int64, arg3 int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetBit indicates an expected call of SetBit.
func (mr *MockCmdableMockRecorder) SetBit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBit", reflect.TypeOf((*MockCmdable)(nil).SetBit), arg0, arg1, arg2, arg3)
}

// SetEx mocks base method.
func (m *MockCmdable) SetEx(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetEx indicates an expected call of SetEx.
func (mr *MockCmdableMockRecorder) SetEx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEx", reflect.TypeOf((*MockCmdable)(nil).SetEx), arg0, arg1, arg2, arg3)
}

// SetNX mocks base method.
func (m *MockCmdable) SetNX(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetNX indicates an expected call of SetNX.
func (mr *MockCmdableMockRecorder) SetNX(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockCmdable)(nil).SetNX), arg0, arg1, arg2, arg3)
}

// SetRange mocks base method.
func (m *MockCmdable) SetRange(arg0 context.Context, arg1 string, arg2 int64, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetRange indicates an expected call of SetRange.
func (mr *MockCmdableMockRecorder) SetRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRange", reflect.TypeOf((*MockCmdable)(nil).SetRange), arg0, arg1, arg2, arg3)
}

// SetXX mocks base method.
func (m *MockCmdable) SetXX(arg0 context.Context, arg1 string, arg2 interface{}, arg3 time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetXX", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetXX indicates an expected call of SetXX.
func (mr *MockCmdableMockRecorder) SetXX(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetXX", reflect.TypeOf((*MockCmdable)(nil).SetXX), arg0, arg1, arg2, arg3)
}

// Shutdown mocks base method.
func (m *MockCmdable) Shutdown(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockCmdableMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCmdable)(nil).Shutdown), arg0)
}

// ShutdownNoSave mocks base method.
func (m *MockCmdable) ShutdownNoSave(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownNoSave", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownNoSave indicates an expected call of ShutdownNoSave.
func (mr *MockCmdableMockRecorder) ShutdownNoSave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownNoSave", reflect.TypeOf((*MockCmdable)(nil).ShutdownNoSave), arg0)
}

// ShutdownSave mocks base method.
func (m *MockCmdable) ShutdownSave(arg0 context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownSave", arg0)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownSave indicates an expected call of ShutdownSave.
func (mr *MockCmdableMockRecorder) ShutdownSave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownSave", reflect.TypeOf((*MockCmdable)(nil).ShutdownSave), arg0)
}

// SlaveOf mocks base method.
func (m *MockCmdable) SlaveOf(arg0 context.Context, arg1, arg2 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlaveOf", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SlaveOf indicates an expected call of SlaveOf.
func (mr *MockCmdableMockRecorder) SlaveOf(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlaveOf", reflect.TypeOf((*MockCmdable)(nil).SlaveOf), arg0, arg1, arg2)
}

// SlowLogGet mocks base method.
func (m *MockCmdable) SlowLogGet(arg0 context.Context, arg1 int64) *redis.SlowLogCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlowLogGet", arg0, arg1)
	ret0, _ := ret[0].(*redis.SlowLogCmd)
	return ret0
}

// SlowLogGet indicates an expected call of SlowLogGet.
func (mr *MockCmdableMockRecorder) SlowLogGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlowLogGet", reflect.TypeOf((*MockCmdable)(nil).SlowLogGet), arg0, arg1)
}

// Sort mocks base method.
func (m *MockCmdable) Sort(arg0 context.Context, arg1 string, arg2 *redis.Sort) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sort", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Sort indicates an expected call of Sort.
func (mr *MockCmdableMockRecorder) Sort(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockCmdable)(nil).Sort), arg0, arg1, arg2)
}

// SortInterfaces mocks base method.
func (m *MockCmdable) SortInterfaces(arg0 context.Context, arg1 string, arg2 *redis.Sort) *redis.SliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortInterfaces", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// SortInterfaces indicates an expected call of SortInterfaces.
func (mr *MockCmdableMockRecorder) SortInterfaces(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortInterfaces", reflect.TypeOf((*MockCmdable)(nil).SortInterfaces), arg0, arg1, arg2)
}

// SortStore mocks base method.
func (m *MockCmdable) SortStore(arg0 context.Context, arg1, arg2 string, arg3 *redis.Sort) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortStore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SortStore indicates an expected call of SortStore.
func (mr *MockCmdableMockRecorder) SortStore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortStore", reflect.TypeOf((*MockCmdable)(nil).SortStore), arg0, arg1, arg2, arg3)
}

// StrLen mocks base method.
func (m *MockCmdable) StrLen(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StrLen", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// StrLen indicates an expected call of StrLen.
func (mr *MockCmdableMockRecorder) StrLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrLen", reflect.TypeOf((*MockCmdable)(nil).StrLen), arg0, arg1)
}

// TTL mocks base method.
func (m *MockCmdable) TTL(arg0 context.Context, arg1 string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", arg0, arg1)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockCmdableMockRecorder) TTL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockCmdable)(nil).TTL), arg0, arg1)
}

// Time mocks base method.
func (m *MockCmdable) Time(arg0 context.Context) *redis.TimeCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time", arg0)
	ret0, _ := ret[0].(*redis.TimeCmd)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockCmdableMockRecorder) Time(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockCmdable)(nil).Time), arg0)
}

// Touch mocks base method.
func (m *MockCmdable) Touch(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Touch", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Touch indicates an expected call of Touch.
func (mr *MockCmdableMockRecorder) Touch(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Touch", reflect.TypeOf((*MockCmdable)(nil).Touch), varargs...)
}

// TxPipeline mocks base method.
func (m *MockCmdable) TxPipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// TxPipeline indicates an expected call of TxPipeline.
func (mr *MockCmdableMockRecorder) TxPipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipeline", reflect.TypeOf((*MockCmdable)(nil).TxPipeline))
}

// TxPipelined mocks base method.
func (m *MockCmdable) TxPipelined(arg0 context.Context, arg1 func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipelined", arg0, arg1)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPipelined indicates an expected call of TxPipelined.
func (mr *MockCmdableMockRecorder) TxPipelined(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipelined", reflect.TypeOf((*MockCmdable)(nil).TxPipelined), arg0, arg1)
}

// Type mocks base method.
func (m *MockCmdable) Type(arg0 context.Context, arg1 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type", arg0, arg1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockCmdableMockRecorder) Type(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockCmdable)(nil).Type), arg0, arg1)
}

// Unlink mocks base method.
func (m *MockCmdable) Unlink(arg0 context.Context, arg1 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unlink", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockCmdableMockRecorder) Unlink(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockCmdable)(nil).Unlink), varargs...)
}

// XAck mocks base method.
func (m *MockCmdable) XAck(arg0 context.Context, arg1, arg2 string, arg3 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XAck", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XAck indicates an expected call of XAck.
func (mr *MockCmdableMockRecorder) XAck(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAck", reflect.TypeOf((*MockCmdable)(nil).XAck), varargs...)
}

// XAdd mocks base method.
func (m *MockCmdable) XAdd(arg0 context.Context, arg1 *redis.XAddArgs) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAdd", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// XAdd indicates an expected call of XAdd.
func (mr *MockCmdableMockRecorder) XAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAdd", reflect.TypeOf((*MockCmdable)(nil).XAdd), arg0, arg1)
}

// XAutoClaim mocks base method.
func (m *MockCmdable) XAutoClaim(arg0 context.Context, arg1 *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaim", arg0, arg1)
	ret0, _ := ret[0].(*redis.XAutoClaimCmd)
	return ret0
}

// XAutoClaim indicates an expected call of XAutoClaim.
func (mr *MockCmdableMockRecorder) XAutoClaim(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaim", reflect.TypeOf((*MockCmdable)(nil).XAutoClaim), arg0, arg1)
}

// XAutoClaimJustID mocks base method.
func (m *MockCmdable) XAutoClaimJustID(arg0 context.Context, arg1 *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaimJustID", arg0, arg1)
	ret0, _ := ret[0].(*redis.XAutoClaimJustIDCmd)
	return ret0
}

// XAutoClaimJustID indicates an expected call of XAutoClaimJustID.
func (mr *MockCmdableMockRecorder) XAutoClaimJustID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaimJustID", reflect.TypeOf((*MockCmdable)(nil).XAutoClaimJustID), arg0, arg1)
}

// XClaim mocks base method.
func (m *MockCmdable) XClaim(arg0 context.Context, arg1 *redis.XClaimArgs) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaim", arg0, arg1)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XClaim indicates an expected call of XClaim.
func (mr *MockCmdableMockRecorder) XClaim(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaim", reflect.TypeOf((*MockCmdable)(nil).XClaim), arg0, arg1)
}

// XClaimJustID mocks base method.
func (m *MockCmdable) XClaimJustID(arg0 context.Context, arg1 *redis.XClaimArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaimJustID", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// XClaimJustID indicates an expected call of XClaimJustID.
func (mr *MockCmdableMockRecorder) XClaimJustID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaimJustID", reflect.TypeOf((*MockCmdable)(nil).XClaimJustID), arg0, arg1)
}

// XDel mocks base method.
func (m *MockCmdable) XDel(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XDel indicates an expected call of XDel.
func (mr *MockCmdableMockRecorder) XDel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XDel", reflect.TypeOf((*MockCmdable)(nil).XDel), varargs...)
}

// XGroupCreate mocks base method.
func (m *MockCmdable) XGroupCreate(arg0 context.Context, arg1, arg2, arg3 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreate indicates an expected call of XGroupCreate.
func (mr *MockCmdableMockRecorder) XGroupCreate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreate", reflect.TypeOf((*MockCmdable)(nil).XGroupCreate), arg0, arg1, arg2, arg3)
}

// XGroupCreateConsumer mocks base method.
func (m *MockCmdable) XGroupCreateConsumer(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateConsumer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupCreateConsumer indicates an expected call of XGroupCreateConsumer.
func (mr *MockCmdableMockRecorder) XGroupCreateConsumer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateConsumer", reflect.TypeOf((*MockCmdable)(nil).XGroupCreateConsumer), arg0, arg1, arg2, arg3)
}

// XGroupCreateMkStream mocks base method.
func (m *MockCmdable) XGroupCreateMkStream(arg0 context.Context, arg1, arg2, arg3 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateMkStream", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreateMkStream indicates an expected call of XGroupCreateMkStream.
func (mr *MockCmdableMockRecorder) XGroupCreateMkStream(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateMkStream", reflect.TypeOf((*MockCmdable)(nil).XGroupCreateMkStream), arg0, arg1, arg2, arg3)
}

// XGroupDelConsumer mocks base method.
func (m *MockCmdable) XGroupDelConsumer(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDelConsumer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDelConsumer indicates an expected call of XGroupDelConsumer.
func (mr *MockCmdableMockRecorder) XGroupDelConsumer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDelConsumer", reflect.TypeOf((*MockCmdable)(nil).XGroupDelConsumer), arg0, arg1, arg2, arg3)
}

// XGroupDestroy mocks base method.
func (m *MockCmdable) XGroupDestroy(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDestroy", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDestroy indicates an expected call of XGroupDestroy.
func (mr *MockCmdableMockRecorder) XGroupDestroy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDestroy", reflect.TypeOf((*MockCmdable)(nil).XGroupDestroy), arg0, arg1, arg2)
}

// XGroupSetID mocks base method.
func (m *MockCmdable) XGroupSetID(arg0 context.Context, arg1, arg2, arg3 string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupSetID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupSetID indicates an expected call of XGroupSetID.
func (mr *MockCmdableMockRecorder) XGroupSetID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupSetID", reflect.TypeOf((*MockCmdable)(nil).XGroupSetID), arg0, arg1, arg2, arg3)
}

// XInfoConsumers mocks base method.
func (m *MockCmdable) XInfoConsumers(arg0 context.Context, arg1, arg2 string) *redis.XInfoConsumersCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoConsumers", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.XInfoConsumersCmd)
	return ret0
}

// XInfoConsumers indicates an expected call of XInfoConsumers.
func (mr *MockCmdableMockRecorder) XInfoConsumers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoConsumers", reflect.TypeOf((*MockCmdable)(nil).XInfoConsumers), arg0, arg1, arg2)
}

// XInfoGroups mocks base method.
func (m *MockCmdable) XInfoGroups(arg0 context.Context, arg1 string) *redis.XInfoGroupsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoGroups", arg0, arg1)
	ret0, _ := ret[0].(*redis.XInfoGroupsCmd)
	return ret0
}

// XInfoGroups indicates an expected call of XInfoGroups.
func (mr *MockCmdableMockRecorder) XInfoGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoGroups", reflect.TypeOf((*MockCmdable)(nil).XInfoGroups), arg0, arg1)
}

// XInfoStream mocks base method.
func (m *MockCmdable) XInfoStream(arg0 context.Context, arg1 string) *redis.XInfoStreamCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStream", arg0, arg1)
	ret0, _ := ret[0].(*redis.XInfoStreamCmd)
	return ret0
}

// XInfoStream indicates an expected call of XInfoStream.
func (mr *MockCmdableMockRecorder) XInfoStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStream", reflect.TypeOf((*MockCmdable)(nil).XInfoStream), arg0, arg1)
}

// XInfoStreamFull mocks base method.
func (m *MockCmdable) XInfoStreamFull(arg0 context.Context, arg1 string, arg2 int) *redis.XInfoStreamFullCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStreamFull", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.XInfoStreamFullCmd)
	return ret0
}

// XInfoStreamFull indicates an expected call of XInfoStreamFull.
func (mr *MockCmdableMockRecorder) XInfoStreamFull(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStreamFull", reflect.TypeOf((*MockCmdable)(nil).XInfoStreamFull), arg0, arg1, arg2)
}

// XLen mocks base method.
func (m *MockCmdable) XLen(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XLen", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XLen indicates an expected call of XLen.
func (mr *MockCmdableMockRecorder) XLen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XLen", reflect.TypeOf((*MockCmdable)(nil).XLen), arg0, arg1)
}

// XPending mocks base method.
func (m *MockCmdable) XPending(arg0 context.Context, arg1, arg2 string) *redis.XPendingCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPending", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.XPendingCmd)
	return ret0
}

// XPending indicates an expected call of XPending.
func (mr *MockCmdableMockRecorder) XPending(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPending", reflect.TypeOf((*MockCmdable)(nil).XPending), arg0, arg1, arg2)
}

// XPendingExt mocks base method.
func (m *MockCmdable) XPendingExt(arg0 context.Context, arg1 *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPendingExt", arg0, arg1)
	ret0, _ := ret[0].(*redis.XPendingExtCmd)
	return ret0
}

// XPendingExt indicates an expected call of XPendingExt.
func (mr *MockCmdableMockRecorder) XPendingExt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPendingExt", reflect.TypeOf((*MockCmdable)(nil).XPendingExt), arg0, arg1)
}

// XRange mocks base method.
func (m *MockCmdable) XRange(arg0 context.Context, arg1, arg2, arg3 string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRange indicates an expected call of XRange.
func (mr *MockCmdableMockRecorder) XRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRange", reflect.TypeOf((*MockCmdable)(nil).XRange), arg0, arg1, arg2, arg3)
}

// XRangeN mocks base method.
func (m *MockCmdable) XRangeN(arg0 context.Context, arg1, arg2, arg3 string, arg4 int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRangeN", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRangeN indicates an expected call of XRangeN.
func (mr *MockCmdableMockRecorder) XRangeN(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRangeN", reflect.TypeOf((*MockCmdable)(nil).XRangeN), arg0, arg1, arg2, arg3, arg4)
}

// XRead mocks base method.
func (m *MockCmdable) XRead(arg0 context.Context, arg1 *redis.XReadArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRead", arg0, arg1)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XRead indicates an expected call of XRead.
func (mr *MockCmdableMockRecorder) XRead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRead", reflect.TypeOf((*MockCmdable)(nil).XRead), arg0, arg1)
}

// XReadGroup mocks base method.
func (m *MockCmdable) XReadGroup(arg0 context.Context, arg1 *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XReadGroup", arg0, arg1)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadGroup indicates an expected call of XReadGroup.
func (mr *MockCmdableMockRecorder) XReadGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadGroup", reflect.TypeOf((*MockCmdable)(nil).XReadGroup), arg0, arg1)
}

// XReadStreams mocks base method.
func (m *MockCmdable) XReadStreams(arg0 context.Context, arg1 ...string) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XReadStreams", varargs...)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadStreams indicates an expected call of XReadStreams.
func (mr *MockCmdableMockRecorder) XReadStreams(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadStreams", reflect.TypeOf((*MockCmdable)(nil).XReadStreams), varargs...)
}

// XRevRange mocks base method.
func (m *MockCmdable) XRevRange(arg0 context.Context, arg1, arg2, arg3 string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRange indicates an expected call of XRevRange.
func (mr *MockCmdableMockRecorder) XRevRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRange", reflect.TypeOf((*MockCmdable)(nil).XRevRange), arg0, arg1, arg2, arg3)
}

// XRevRangeN mocks base method.
func (m *MockCmdable) XRevRangeN(arg0 context.Context, arg1, arg2, arg3 string, arg4 int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRangeN", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRangeN indicates an expected call of XRevRangeN.
func (mr *MockCmdableMockRecorder) XRevRangeN(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRangeN", reflect.TypeOf((*MockCmdable)(nil).XRevRangeN), arg0, arg1, arg2, arg3, arg4)
}

// XTrimMaxLen mocks base method.
func (m *MockCmdable) XTrimMaxLen(arg0 context.Context, arg1 string, arg2 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLen", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLen indicates an expected call of XTrimMaxLen.
func (mr *MockCmdableMockRecorder) XTrimMaxLen(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLen", reflect.TypeOf((*MockCmdable)(nil).XTrimMaxLen), arg0, arg1, arg2)
}

// XTrimMaxLenApprox mocks base method.
func (m *MockCmdable) XTrimMaxLenApprox(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLenApprox", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLenApprox indicates an expected call of XTrimMaxLenApprox.
func (mr *MockCmdableMockRecorder) XTrimMaxLenApprox(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLenApprox", reflect.TypeOf((*MockCmdable)(nil).XTrimMaxLenApprox), arg0, arg1, arg2, arg3)
}

// XTrimMinID mocks base method.
func (m *MockCmdable) XTrimMinID(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinID indicates an expected call of XTrimMinID.
func (mr *MockCmdableMockRecorder) XTrimMinID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinID", reflect.TypeOf((*MockCmdable)(nil).XTrimMinID), arg0, arg1, arg2)
}

// XTrimMinIDApprox mocks base method.
func (m *MockCmdable) XTrimMinIDApprox(arg0 context.Context, arg1, arg2 string, arg3 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinIDApprox", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinIDApprox indicates an expected call of XTrimMinIDApprox.
func (mr *MockCmdableMockRecorder) XTrimMinIDApprox(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinIDApprox", reflect.TypeOf((*MockCmdable)(nil).XTrimMinIDApprox), arg0, arg1, arg2, arg3)
}

// ZAdd mocks base method.
func (m *MockCmdable) ZAdd(arg0 context.Context, arg1 string, arg2 ...redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockCmdableMockRecorder) ZAdd(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockCmdable)(nil).ZAdd), varargs...)
}

// ZAddArgs mocks base method.
func (m *MockCmdable) ZAddArgs(arg0 context.Context, arg1 string, arg2 redis.ZAddArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddArgs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddArgs indicates an expected call of ZAddArgs.
func (mr *MockCmdableMockRecorder) ZAddArgs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddArgs", reflect.TypeOf((*MockCmdable)(nil).ZAddArgs), arg0, arg1, arg2)
}

// ZAddArgsIncr mocks base method.
func (m *MockCmdable) ZAddArgsIncr(arg0 context.Context, arg1 string, arg2 redis.ZAddArgs) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddArgsIncr", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZAddArgsIncr indicates an expected call of ZAddArgsIncr.
func (mr *MockCmdableMockRecorder) ZAddArgsIncr(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddArgsIncr", reflect.TypeOf((*MockCmdable)(nil).ZAddArgsIncr), arg0, arg1, arg2)
}

// ZAddNX mocks base method.
func (m *MockCmdable) ZAddNX(arg0 context.Context, arg1 string, arg2 ...redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddNX indicates an expected call of ZAddNX.
func (mr *MockCmdableMockRecorder) ZAddNX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNX", reflect.TypeOf((*MockCmdable)(nil).ZAddNX), varargs...)
}

// ZAddXX mocks base method.
func (m *MockCmdable) ZAddXX(arg0 context.Context, arg1 string, arg2 ...redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddXX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddXX indicates an expected call of ZAddXX.
func (mr *MockCmdableMockRecorder) ZAddXX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddXX", reflect.TypeOf((*MockCmdable)(nil).ZAddXX), varargs...)
}

// ZCard mocks base method.
func (m *MockCmdable) ZCard(arg0 context.Context, arg1 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", arg0, arg1)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCard indicates an expected call of ZCard.
func (mr *MockCmdableMockRecorder) ZCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockCmdable)(nil).ZCard), arg0, arg1)
}

// ZCount mocks base method.
func (m *MockCmdable) ZCount(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCount indicates an expected call of ZCount.
func (mr *MockCmdableMockRecorder) ZCount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCount", reflect.TypeOf((*MockCmdable)(nil).ZCount), arg0, arg1, arg2, arg3)
}

// ZDiff mocks base method.
func (m *MockCmdable) ZDiff(arg0 context.Context, arg1 ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZDiff indicates an expected call of ZDiff.
func (mr *MockCmdableMockRecorder) ZDiff(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiff", reflect.TypeOf((*MockCmdable)(nil).ZDiff), varargs...)
}

// ZDiffStore mocks base method.
func (m *MockCmdable) ZDiffStore(arg0 context.Context, arg1 string, arg2 ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZDiffStore indicates an expected call of ZDiffStore.
func (mr *MockCmdableMockRecorder) ZDiffStore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffStore", reflect.TypeOf((*MockCmdable)(nil).ZDiffStore), varargs...)
}

// ZDiffWithScores mocks base method.
func (m *MockCmdable) ZDiffWithScores(arg0 context.Context, arg1 ...string) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffWithScores", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZDiffWithScores indicates an expected call of ZDiffWithScores.
func (mr *MockCmdableMockRecorder) ZDiffWithScores(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffWithScores", reflect.TypeOf((*MockCmdable)(nil).ZDiffWithScores), varargs...)
}

// ZIncrBy mocks base method.
func (m *MockCmdable) ZIncrBy(arg0 context.Context, arg1 string, arg2 float64, arg3 string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrBy", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrBy indicates an expected call of ZIncrBy.
func (mr *MockCmdableMockRecorder) ZIncrBy(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrBy", reflect.TypeOf((*MockCmdable)(nil).ZIncrBy), arg0, arg1, arg2, arg3)
}

// ZInter mocks base method.
func (m *MockCmdable) ZInter(arg0 context.Context, arg1 *redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInter", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZInter indicates an expected call of ZInter.
func (mr *MockCmdableMockRecorder) ZInter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInter", reflect.TypeOf((*MockCmdable)(nil).ZInter), arg0, arg1)
}

// ZInterStore mocks base method.
func (m *MockCmdable) ZInterStore(arg0 context.Context, arg1 string, arg2 *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterStore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZInterStore indicates an expected call of ZInterStore.
func (mr *MockCmdableMockRecorder) ZInterStore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterStore", reflect.TypeOf((*MockCmdable)(nil).ZInterStore), arg0, arg1, arg2)
}

// ZInterWithScores mocks base method.
func (m *MockCmdable) ZInterWithScores(arg0 context.Context, arg1 *redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterWithScores", arg0, arg1)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZInterWithScores indicates an expected call of ZInterWithScores.
func (mr *MockCmdableMockRecorder) ZInterWithScores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterWithScores", reflect.TypeOf((*MockCmdable)(nil).ZInterWithScores), arg0, arg1)
}

// ZLexCount mocks base method.
func (m *MockCmdable) ZLexCount(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZLexCount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZLexCount indicates an expected call of ZLexCount.
func (mr *MockCmdableMockRecorder) ZLexCount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZLexCount", reflect.TypeOf((*MockCmdable)(nil).ZLexCount), arg0, arg1, arg2, arg3)
}

// ZMScore mocks base method.
func (m *MockCmdable) ZMScore(arg0 context.Context, arg1 string, arg2 ...string) *redis.FloatSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZMScore", varargs...)
	ret0, _ := ret[0].(*redis.FloatSliceCmd)
	return ret0
}

// ZMScore indicates an expected call of ZMScore.
func (mr *MockCmdableMockRecorder) ZMScore(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZMScore", reflect.TypeOf((*MockCmdable)(nil).ZMScore), varargs...)
}

// ZPopMax mocks base method.
func (m *MockCmdable) ZPopMax(arg0 context.Context, arg1 string, arg2 ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMax indicates an expected call of ZPopMax.
func (mr *MockCmdableMockRecorder) ZPopMax(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMax", reflect.TypeOf((*MockCmdable)(nil).ZPopMax), varargs...)
}

// ZPopMin mocks base method.
func (m *MockCmdable) ZPopMin(arg0 context.Context, arg1 string, arg2 ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMin indicates an expected call of ZPopMin.
func (mr *MockCmdableMockRecorder) ZPopMin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMin", reflect.TypeOf((*MockCmdable)(nil).ZPopMin), varargs...)
}

// ZRandMember mocks base method.
func (m *MockCmdable) ZRandMember(arg0 context.Context, arg1 string, arg2 int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRandMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRandMember indicates an expected call of ZRandMember.
func (mr *MockCmdableMockRecorder) ZRandMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRandMember", reflect.TypeOf((*MockCmdable)(nil).ZRandMember), arg0, arg1, arg2)
}

// ZRandMemberWithScores mocks base method.
func (m *MockCmdable) ZRandMemberWithScores(arg0 context.Context, arg1 string, arg2 int) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRandMemberWithScores", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRandMemberWithScores indicates an expected call of ZRandMemberWithScores.
func (mr *MockCmdableMockRecorder) ZRandMemberWithScores(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRandMemberWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRandMemberWithScores), arg0, arg1, arg2)
}

// ZRange mocks base method.
func (m *MockCmdable) ZRange(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRange indicates an expected call of ZRange.
func (mr *MockCmdableMockRecorder) ZRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRange", reflect.TypeOf((*MockCmdable)(nil).ZRange), arg0, arg1, arg2, arg3)
}

// ZRangeArgs mocks base method.
func (m *MockCmdable) ZRangeArgs(arg0 context.Context, arg1 redis.ZRangeArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgs", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeArgs indicates an expected call of ZRangeArgs.
func (mr *MockCmdableMockRecorder) ZRangeArgs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgs", reflect.TypeOf((*MockCmdable)(nil).ZRangeArgs), arg0, arg1)
}

// ZRangeArgsWithScores mocks base method.
func (m *MockCmdable) ZRangeArgsWithScores(arg0 context.Context, arg1 redis.ZRangeArgs) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgsWithScores", arg0, arg1)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeArgsWithScores indicates an expected call of ZRangeArgsWithScores.
func (mr *MockCmdableMockRecorder) ZRangeArgsWithScores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgsWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeArgsWithScores), arg0, arg1)
}

// ZRangeByLex mocks base method.
func (m *MockCmdable) ZRangeByLex(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByLex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByLex indicates an expected call of ZRangeByLex.
func (mr *MockCmdableMockRecorder) ZRangeByLex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRangeByLex), arg0, arg1, arg2)
}

// ZRangeByScore mocks base method.
func (m *MockCmdable) ZRangeByScore(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockCmdableMockRecorder) ZRangeByScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRangeByScore), arg0, arg1, arg2)
}

// ZRangeByScoreWithScores mocks base method.
func (m *MockCmdable) ZRangeByScoreWithScores(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScoreWithScores", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeByScoreWithScores indicates an expected call of ZRangeByScoreWithScores.
func (mr *MockCmdableMockRecorder) ZRangeByScoreWithScores(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScoreWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeByScoreWithScores), arg0, arg1, arg2)
}

// ZRangeStore mocks base method.
func (m *MockCmdable) ZRangeStore(arg0 context.Context, arg1 string, arg2 redis.ZRangeArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeStore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRangeStore indicates an expected call of ZRangeStore.
func (mr *MockCmdableMockRecorder) ZRangeStore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeStore", reflect.TypeOf((*MockCmdable)(nil).ZRangeStore), arg0, arg1, arg2)
}

// ZRangeWithScores mocks base method.
func (m *MockCmdable) ZRangeWithScores(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeWithScores", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeWithScores indicates an expected call of ZRangeWithScores.
func (mr *MockCmdableMockRecorder) ZRangeWithScores(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeWithScores), arg0, arg1, arg2, arg3)
}

// ZRank mocks base method.
func (m *MockCmdable) ZRank(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRank", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRank indicates an expected call of ZRank.
func (mr *MockCmdableMockRecorder) ZRank(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRank", reflect.TypeOf((*MockCmdable)(nil).ZRank), arg0, arg1, arg2)
}

// ZRem mocks base method.
func (m *MockCmdable) ZRem(arg0 context.Context, arg1 string, arg2 ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockCmdableMockRecorder) ZRem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockCmdable)(nil).ZRem), varargs...)
}

// ZRemRangeByLex mocks base method.
func (m *MockCmdable) ZRemRangeByLex(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByLex", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByLex indicates an expected call of ZRemRangeByLex.
func (mr *MockCmdableMockRecorder) ZRemRangeByLex(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByLex), arg0, arg1, arg2, arg3)
}

// ZRemRangeByRank mocks base method.
func (m *MockCmdable) ZRemRangeByRank(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByRank", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByRank indicates an expected call of ZRemRangeByRank.
func (mr *MockCmdableMockRecorder) ZRemRangeByRank(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByRank", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByRank), arg0, arg1, arg2, arg3)
}

// ZRemRangeByScore mocks base method.
func (m *MockCmdable) ZRemRangeByScore(arg0 context.Context, arg1, arg2, arg3 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByScore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByScore indicates an expected call of ZRemRangeByScore.
func (mr *MockCmdableMockRecorder) ZRemRangeByScore(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByScore), arg0, arg1, arg2, arg3)
}

// ZRevRange mocks base method.
func (m *MockCmdable) ZRevRange(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRange indicates an expected call of ZRevRange.
func (mr *MockCmdableMockRecorder) ZRevRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRange", reflect.TypeOf((*MockCmdable)(nil).ZRevRange), arg0, arg1, arg2, arg3)
}

// ZRevRangeByLex mocks base method.
func (m *MockCmdable) ZRevRangeByLex(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByLex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByLex indicates an expected call of ZRevRangeByLex.
func (mr *MockCmdableMockRecorder) ZRevRangeByLex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByLex), arg0, arg1, arg2)
}

// ZRevRangeByScore mocks base method.
func (m *MockCmdable) ZRevRangeByScore(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByScore indicates an expected call of ZRevRangeByScore.
func (mr *MockCmdableMockRecorder) ZRevRangeByScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByScore), arg0, arg1, arg2)
}

// ZRevRangeByScoreWithScores mocks base method.
func (m *MockCmdable) ZRevRangeByScoreWithScores(arg0 context.Context, arg1 string, arg2 *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScoreWithScores", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeByScoreWithScores indicates an expected call of ZRevRangeByScoreWithScores.
func (mr *MockCmdableMockRecorder) ZRevRangeByScoreWithScores(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScoreWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByScoreWithScores), arg0, arg1, arg2)
}

// ZRevRangeWithScores mocks base method.
func (m *MockCmdable) ZRevRangeWithScores(arg0 context.Context, arg1 string, arg2, arg3 int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeWithScores", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeWithScores indicates an expected call of ZRevRangeWithScores.
func (mr *MockCmdableMockRecorder) ZRevRangeWithScores(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeWithScores), arg0, arg1, arg2, arg3)
}

// ZRevRank mocks base method.
func (m *MockCmdable) ZRevRank(arg0 context.Context, arg1, arg2 string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRank", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRevRank indicates an expected call of ZRevRank.
func (mr *MockCmdableMockRecorder) ZRevRank(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRank", reflect.TypeOf((*MockCmdable)(nil).ZRevRank), arg0, arg1, arg2)
}

// ZScan mocks base method.
func (m *MockCmdable) ZScan(arg0 context.Context, arg1 string, arg2 uint64, arg3 string, arg4 int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScan", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ZScan indicates an expected call of ZScan.
func (mr *MockCmdableMockRecorder) ZScan(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScan", reflect.TypeOf((*MockCmdable)(nil).ZScan), arg0, arg1, arg2, arg3, arg4)
}

// ZScore mocks base method.
func (m *MockCmdable) ZScore(arg0 context.Context, arg1, arg2 string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZScore indicates an expected call of ZScore.
func (mr *MockCmdableMockRecorder) ZScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScore", reflect.TypeOf((*MockCmdable)(nil).ZScore), arg0, arg1, arg2)
}

// ZUnion mocks base method.
func (m *MockCmdable) ZUnion(arg0 context.Context, arg1 redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnion", arg0, arg1)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZUnion indicates an expected call of ZUnion.
func (mr *MockCmdableMockRecorder) ZUnion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnion", reflect.TypeOf((*MockCmdable)(nil).ZUnion), arg0, arg1)
}

// ZUnionStore mocks base method.
func (m *MockCmdable) ZUnionStore(arg0 context.Context, arg1 string, arg2 *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionStore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZUnionStore indicates an expected call of ZUnionStore.
func (mr *MockCmdableMockRecorder) ZUnionStore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionStore", reflect.TypeOf((*MockCmdable)(nil).ZUnionStore), arg0, arg1, arg2)
}

// ZUnionWithScores mocks base method.
func (m *MockCmdable) ZUnionWithScores(arg0 context.Context, arg1 redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionWithScores", arg0, arg1)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZUnionWithScores indicates an expected call of ZUnionWithScores.
func (mr *MockCmdableMockRecorder) ZUnionWithScores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionWithScores", reflect.TypeOf((*MockCmdable)(nil).ZUnionWithScores), arg0, arg1)
}
