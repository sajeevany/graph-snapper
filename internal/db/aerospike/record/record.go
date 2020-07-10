package record

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/sajeevany/graph-snapper/internal/common"
	"github.com/sirupsen/logrus"
)

const (
	MetadataBinName    = "Metadata"
	AccountBinName     = "Account"
	CredentialsBinName = "Credentials"
	VersionAttrName    = "Version"

	GrafanaAPIUserNamespace            = "GrafanaAPIUser"
	ConfluenceServerBasicUserNamespace = "ConfluenceServerBasicUser"
)

const (
	VersionLevel_1 = "1"
)

type Record interface {
	//GetFields - returns logrus fields for logging
	GetFields() logrus.Fields
	//ToASBinSlice - converts record to bin map. Used to write record to db in the latest record format
	ToASBinSlice() []*aerospike.Bin
	//ToRecordViewV1 - converts to v1 record view
	ToRecordViewV1() RecordViewV1
	//AddUserCredentialsV1 - Adds input credentials to record. Does not overwrite any existing records
	AddUserCredentialsV1([]common.GrafanaUserV1, []common.ConfluenceServerUserV1)
}

//Record - Aerospike configuration + credentials data
type RecordV1 struct {
	Metadata    MetadataV1    `json:"Metadata"`
	Account     AccountV1     `json:"Account"`
	Credentials CredentialsV1 `json:"Credentials"`
}

func (r *RecordV1) ToRecordViewV1() RecordViewV1 {
	return RecordViewV1{
		Metadata:    r.Metadata.toMetadataView1(),
		Account:     r.Account.toAccountView1(),
		Credentials: r.Credentials.toCredentialsView1(),
	}
}

func (r *RecordV1) GetFields() logrus.Fields {
	return logrus.Fields{
		"MetadataV1":    r.Metadata.GetFields(),
		"AccountV1":     r.Account.GetFields(),
		"CredentialsV1": r.Credentials.GetFields(),
	}
}

//ToASBinSlice - converts to aerospike bins. Currently writes record in recordv1 format
func (r *RecordV1) ToASBinSlice() []*aerospike.Bin {
	return []*aerospike.Bin{
		r.Metadata.getMetadataBin(),
		r.Account.getAccountBin(),
		r.Credentials.getCredentialBin(),
	}
}

func (r *RecordV1) AddUserCredentialsV1(grafanaUsers []common.GrafanaUserV1, confluenceUser []common.ConfluenceServerUserV1) {

	//
	for _, gu := range grafanaUsers {
		index := getNextFreeIdx(r.Credentials.GrafanaAPIUsers, GrafanaAPIUserNamespace)
		r.Credentials.GrafanaAPIUsers[index] = gu
	}

}
