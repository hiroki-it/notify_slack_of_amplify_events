package detail

type Detail struct {
	AppId         *AppId
	BranchName    *BranchName
	JobId         *JobId
	JobStatusType JobStatusType
}

// NewDetail コンストラクタ
func NewDetail(appId *AppId, branchName *BranchName, jobId *JobId, jobStatusType JobStatusType) *Detail {

	return &Detail{
		AppId:         appId,
		BranchName:    branchName,
		JobId:         jobId,
		JobStatusType: jobStatusType,
	}
}

// GetJobStatusType JobStatusTypeを返却します.
func (d *Detail) GetJobStatusType() JobStatusType {
	return d.JobStatusType
}
