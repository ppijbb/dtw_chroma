// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	coordinatorpb "github.com/chroma-core/chroma/go/pkg/proto/coordinatorpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// SysDBClient is an autogenerated mock type for the SysDBClient type
type SysDBClient struct {
	mock.Mock
}

// CreateCollection provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) CreateCollection(ctx context.Context, in *coordinatorpb.CreateCollectionRequest, opts ...grpc.CallOption) (*coordinatorpb.CreateCollectionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCollection")
	}

	var r0 *coordinatorpb.CreateCollectionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateCollectionRequest, ...grpc.CallOption) (*coordinatorpb.CreateCollectionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateCollectionRequest, ...grpc.CallOption) *coordinatorpb.CreateCollectionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.CreateCollectionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.CreateCollectionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDatabase provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) CreateDatabase(ctx context.Context, in *coordinatorpb.CreateDatabaseRequest, opts ...grpc.CallOption) (*coordinatorpb.CreateDatabaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabase")
	}

	var r0 *coordinatorpb.CreateDatabaseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateDatabaseRequest, ...grpc.CallOption) (*coordinatorpb.CreateDatabaseResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateDatabaseRequest, ...grpc.CallOption) *coordinatorpb.CreateDatabaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.CreateDatabaseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.CreateDatabaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSegment provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) CreateSegment(ctx context.Context, in *coordinatorpb.CreateSegmentRequest, opts ...grpc.CallOption) (*coordinatorpb.CreateSegmentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSegment")
	}

	var r0 *coordinatorpb.CreateSegmentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateSegmentRequest, ...grpc.CallOption) (*coordinatorpb.CreateSegmentResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateSegmentRequest, ...grpc.CallOption) *coordinatorpb.CreateSegmentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.CreateSegmentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.CreateSegmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTenant provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) CreateTenant(ctx context.Context, in *coordinatorpb.CreateTenantRequest, opts ...grpc.CallOption) (*coordinatorpb.CreateTenantResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTenant")
	}

	var r0 *coordinatorpb.CreateTenantResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateTenantRequest, ...grpc.CallOption) (*coordinatorpb.CreateTenantResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.CreateTenantRequest, ...grpc.CallOption) *coordinatorpb.CreateTenantResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.CreateTenantResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.CreateTenantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCollection provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) DeleteCollection(ctx context.Context, in *coordinatorpb.DeleteCollectionRequest, opts ...grpc.CallOption) (*coordinatorpb.DeleteCollectionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCollection")
	}

	var r0 *coordinatorpb.DeleteCollectionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.DeleteCollectionRequest, ...grpc.CallOption) (*coordinatorpb.DeleteCollectionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.DeleteCollectionRequest, ...grpc.CallOption) *coordinatorpb.DeleteCollectionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.DeleteCollectionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.DeleteCollectionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSegment provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) DeleteSegment(ctx context.Context, in *coordinatorpb.DeleteSegmentRequest, opts ...grpc.CallOption) (*coordinatorpb.DeleteSegmentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSegment")
	}

	var r0 *coordinatorpb.DeleteSegmentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.DeleteSegmentRequest, ...grpc.CallOption) (*coordinatorpb.DeleteSegmentResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.DeleteSegmentRequest, ...grpc.CallOption) *coordinatorpb.DeleteSegmentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.DeleteSegmentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.DeleteSegmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushCollectionCompaction provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) FlushCollectionCompaction(ctx context.Context, in *coordinatorpb.FlushCollectionCompactionRequest, opts ...grpc.CallOption) (*coordinatorpb.FlushCollectionCompactionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FlushCollectionCompaction")
	}

	var r0 *coordinatorpb.FlushCollectionCompactionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.FlushCollectionCompactionRequest, ...grpc.CallOption) (*coordinatorpb.FlushCollectionCompactionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.FlushCollectionCompactionRequest, ...grpc.CallOption) *coordinatorpb.FlushCollectionCompactionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.FlushCollectionCompactionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.FlushCollectionCompactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollections provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) GetCollections(ctx context.Context, in *coordinatorpb.GetCollectionsRequest, opts ...grpc.CallOption) (*coordinatorpb.GetCollectionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCollections")
	}

	var r0 *coordinatorpb.GetCollectionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetCollectionsRequest, ...grpc.CallOption) (*coordinatorpb.GetCollectionsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetCollectionsRequest, ...grpc.CallOption) *coordinatorpb.GetCollectionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.GetCollectionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.GetCollectionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabase provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) GetDatabase(ctx context.Context, in *coordinatorpb.GetDatabaseRequest, opts ...grpc.CallOption) (*coordinatorpb.GetDatabaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabase")
	}

	var r0 *coordinatorpb.GetDatabaseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetDatabaseRequest, ...grpc.CallOption) (*coordinatorpb.GetDatabaseResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetDatabaseRequest, ...grpc.CallOption) *coordinatorpb.GetDatabaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.GetDatabaseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.GetDatabaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastCompactionTimeForTenant provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) GetLastCompactionTimeForTenant(ctx context.Context, in *coordinatorpb.GetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*coordinatorpb.GetLastCompactionTimeForTenantResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLastCompactionTimeForTenant")
	}

	var r0 *coordinatorpb.GetLastCompactionTimeForTenantResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetLastCompactionTimeForTenantRequest, ...grpc.CallOption) (*coordinatorpb.GetLastCompactionTimeForTenantResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetLastCompactionTimeForTenantRequest, ...grpc.CallOption) *coordinatorpb.GetLastCompactionTimeForTenantResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.GetLastCompactionTimeForTenantResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.GetLastCompactionTimeForTenantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSegments provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) GetSegments(ctx context.Context, in *coordinatorpb.GetSegmentsRequest, opts ...grpc.CallOption) (*coordinatorpb.GetSegmentsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSegments")
	}

	var r0 *coordinatorpb.GetSegmentsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetSegmentsRequest, ...grpc.CallOption) (*coordinatorpb.GetSegmentsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetSegmentsRequest, ...grpc.CallOption) *coordinatorpb.GetSegmentsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.GetSegmentsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.GetSegmentsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTenant provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) GetTenant(ctx context.Context, in *coordinatorpb.GetTenantRequest, opts ...grpc.CallOption) (*coordinatorpb.GetTenantResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTenant")
	}

	var r0 *coordinatorpb.GetTenantResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetTenantRequest, ...grpc.CallOption) (*coordinatorpb.GetTenantResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.GetTenantRequest, ...grpc.CallOption) *coordinatorpb.GetTenantResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.GetTenantResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.GetTenantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetState provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) ResetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*coordinatorpb.ResetStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ResetState")
	}

	var r0 *coordinatorpb.ResetStateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*coordinatorpb.ResetStateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) *coordinatorpb.ResetStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.ResetStateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLastCompactionTimeForTenant provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) SetLastCompactionTimeForTenant(ctx context.Context, in *coordinatorpb.SetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetLastCompactionTimeForTenant")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.SetLastCompactionTimeForTenantRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.SetLastCompactionTimeForTenantRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.SetLastCompactionTimeForTenantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCollection provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) UpdateCollection(ctx context.Context, in *coordinatorpb.UpdateCollectionRequest, opts ...grpc.CallOption) (*coordinatorpb.UpdateCollectionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCollection")
	}

	var r0 *coordinatorpb.UpdateCollectionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.UpdateCollectionRequest, ...grpc.CallOption) (*coordinatorpb.UpdateCollectionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.UpdateCollectionRequest, ...grpc.CallOption) *coordinatorpb.UpdateCollectionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.UpdateCollectionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.UpdateCollectionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSegment provides a mock function with given fields: ctx, in, opts
func (_m *SysDBClient) UpdateSegment(ctx context.Context, in *coordinatorpb.UpdateSegmentRequest, opts ...grpc.CallOption) (*coordinatorpb.UpdateSegmentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSegment")
	}

	var r0 *coordinatorpb.UpdateSegmentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.UpdateSegmentRequest, ...grpc.CallOption) (*coordinatorpb.UpdateSegmentResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coordinatorpb.UpdateSegmentRequest, ...grpc.CallOption) *coordinatorpb.UpdateSegmentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coordinatorpb.UpdateSegmentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coordinatorpb.UpdateSegmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSysDBClient creates a new instance of SysDBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSysDBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SysDBClient {
	mock := &SysDBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}