package credentials

import (
	"github.com/sajeevany/graph-snapper/internal/common"
	"github.com/sirupsen/logrus"
)

//AddCredentialsV1 - Request to add users to account with specified ID
type AddCredentialsV1 struct {
	AccountID             string
	GrafanaReadUsers      []common.GrafanaUserV1          `json:"GrafanaReadUsers"`
	ConfluenceServerUsers []common.ConfluenceServerUserV1 `json:"ConfluenceServerUsers"`
}

func (a AddCredentialsV1) GetFields() logrus.Fields {

	//Convert grafana read users to loggable format
	gru := make([]logrus.Fields, len(a.GrafanaReadUsers))
	for i, v := range a.GrafanaReadUsers {
		gru[i] = v.GetFields()
	}

	//Convert confluence-server write users to loggable format
	csu := make([]logrus.Fields, len(a.GrafanaReadUsers))
	for i, v := range a.ConfluenceServerUsers {
		csu[i] = v.GetFields()
	}

	return logrus.Fields{
		"AccountID":             a.AccountID,
		"GrafanaReadUsers":      gru,
		"ConfluenceServerUsers": csu,
	}
}

//HasNoUsers - returns true if both grafana and confluence user arrays are empty
func (req AddCredentialsV1) HasNoUsers() bool {
	return len(req.GrafanaReadUsers) == 0 && len(req.ConfluenceServerUsers) == 0
}

//AddedCredentials - Users added to specified account
type AddedCredentialsV1 struct {
	GrafanaReadUsers      map[string]common.GrafanaUserV1          `json:"GrafanaReadUsers"`
	ConfluenceServerUsers map[string]common.ConfluenceServerUserV1 `json:"ConfluenceServerUsers"`
}
