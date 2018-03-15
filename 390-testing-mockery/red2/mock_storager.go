// Code generated by mockery v1.0.0
package red

import mock "github.com/stretchr/testify/mock"
import redis "github.com/mediocregopher/radix.v2/redis"

// mockStorager is an autogenerated mock type for the storager type
type mockStorager struct {
	mock.Mock
}

// Cmd provides a mock function with given fields: _a0, _a1
func (_m *mockStorager) Cmd(_a0 string, _a1 ...interface{}) *redis.Resp {
	ret := _m.Called(_a0, _a1)

	var r0 *redis.Resp
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *redis.Resp); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Resp)
		}
	}

	return r0
}