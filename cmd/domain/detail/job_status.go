package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobStatus struct {
	status string
}

var _ core.ValueObject = (*JobStatus)(nil)

// NewJobStatus コンストラクタ
func NewJobStatus(status string) *JobStatus {

	return &JobStatus{
		status: status,
	}
}

// Status 属性を返却します．
func (js *JobStatus) Status() string {
	return js.status
}

// Equals 等価性を検証します．
func (js *JobStatus) Equals(target core.ValueObject) bool {
	return js.status == target.(*JobStatus).Status()
}
