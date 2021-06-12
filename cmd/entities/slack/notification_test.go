package slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

/**
 * ステータスがSUCCEED場合に，PostMessageメソッドが成功することをテストします．
 */
func (suite *NotificationSuite) TestPostMessage_EventDetailSucceed_ReturnNil() {

	suite.T().Helper()

	// 外部サーバのモックを構築します．
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	}))

	defer ts.Close()

	slackMessage := NewSlackMessage(
		suite.eventDetail,
		suite.getBranchOutput.Branch,
		suite.jobStatusColor,
	)

	slackClient := NewSlackClient(
		&http.Client{},
		ts.URL, // モックサーバのURLに差し替えます．
	)

	slackNotification := NewSlackNotification(
		slackClient,
		slackMessage,
	)

	err := slackNotification.PostMessage()

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	assert.Nil(suite.T(), err)
}
