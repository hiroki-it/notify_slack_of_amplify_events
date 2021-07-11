package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobId struct {
	id string
}

var _ core.ID = &JobId{}

// NewJobId コンストラクタ
func NewJobId(id string) *JobId {

	return &JobId{
		id: id,
	}
}

// Id 属性を返却します．
func (ji *JobId) Id() string {
	return ji.id
}

// Equals 等価性を検証します．
func (ji *JobId) Equals(target core.ID) bool {
	return ji.id == target.(*JobId).Id()
}
