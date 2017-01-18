package mocks

import mock "github.com/stretchr/testify/mock"
import persistence "code.uber.internal/devexp/minions/common/persistence"

// ExecutionManager is an autogenerated mock type for the ExecutionManager type
type ExecutionManager struct {
	mock.Mock
}

// CompleteTransferTask provides a mock function with given fields: request
func (_m *ExecutionManager) CompleteTransferTask(request *persistence.CompleteTransferTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CompleteTransferTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShard provides a mock function with given fields: request
func (_m *ExecutionManager) CreateShard(request *persistence.CreateShardRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CreateShardRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) CreateWorkflowExecution(request *persistence.CreateWorkflowExecutionRequest) (*persistence.CreateWorkflowExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.CreateWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.CreateWorkflowExecutionRequest) *persistence.CreateWorkflowExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.CreateWorkflowExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) DeleteWorkflowExecution(request *persistence.DeleteWorkflowExecutionRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.DeleteWorkflowExecutionRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetShard provides a mock function with given fields: request
func (_m *ExecutionManager) GetShard(request *persistence.GetShardRequest) (*persistence.GetShardResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetShardResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetShardRequest) *persistence.GetShardResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetShardRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTimerIndexTasks provides a mock function with given fields: request
func (_m *ExecutionManager) GetTimerIndexTasks(request *persistence.GetTimerIndexTasksRequest) (*persistence.GetTimerIndexTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetTimerIndexTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetTimerIndexTasksRequest) *persistence.GetTimerIndexTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTimerIndexTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetTimerIndexTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransferTasks provides a mock function with given fields: request
func (_m *ExecutionManager) GetTransferTasks(request *persistence.GetTransferTasksRequest) (*persistence.GetTransferTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetTransferTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetTransferTasksRequest) *persistence.GetTransferTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTransferTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetTransferTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) GetWorkflowExecution(request *persistence.GetWorkflowExecutionRequest) (*persistence.GetWorkflowExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetWorkflowExecutionRequest) *persistence.GetWorkflowExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetWorkflowExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateShard provides a mock function with given fields: request
func (_m *ExecutionManager) UpdateShard(request *persistence.UpdateShardRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.UpdateShardRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) UpdateWorkflowExecution(request *persistence.UpdateWorkflowExecutionRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.UpdateWorkflowExecutionRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ persistence.ExecutionManager = (*ExecutionManager)(nil)