package schedule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/graph-snapper/internal/grafana"
	"github.com/sajeevany/graph-snapper/internal/report"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	Group         = "/schedule"
	CheckEndpoint = "/check"
)

//@Summary Check and execute schedule
//@Description Non-authenticated endpoint which checks and runs a schedule to validate connectivity and storage behaviour by the end user
//@Produce json
//@Param schedule body CheckScheduleV1 true "Check schedule"
//@Success 200 {object} report.CheckDashboardSnapshotReportV1
//@Fail 400 {object} gin.H
//@Fail 500 {object} gin.H
//@Router /schedule/check [post]
//@Tags schedule
func CheckV1(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Debug("Starting schedule check (v1)")

		//Bind schedule object
		var schedule CheckScheduleV1
		if bErr := ctx.BindJSON(&schedule); bErr != nil {
			msg := fmt.Sprintf("Unable to bind request body to schedule object %v", bErr)
			logger.Errorf(msg)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		//Quick check if all required attributes are present
		if isValid, err := schedule.IsValid(); !isValid {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//Run snapshot and upload process
		report, err := snapshotAndUpload(schedule)
		if err != nil {
			logger.WithFields(schedule.GetFields()).Errorf("Error when running snapshotAndUpload err <%v>", err)
			return
		}

		ctx.JSON(report.GetResultCode(), report.ToCheckScheduleV1View())

	}
}

//snapshotAndUpload -
func snapshotAndUpload(logger *logrus.Logger, schedule CheckScheduleV1) (report.DashboardSnapshotReport, error) {

	snapReport := report.DashboardSnapshotReport{
		Title:     "CheckV1 schedule test",
		Timestamp: time.Now(),
	}

	//create snapshots, group, and upload as a subpage to the specified datastore(s)
	for reqKey, dashboard := range schedule.GraphDashboards.GrafanaDashboards {

		dashReport := report.GrafanaDashboardReport{
			Title:     fmt.Sprintf("Grafana dashboard <%s> snapshot panel report", dashboard.UID),
			StartTime: time.Now(),
			UID:       dashboard.UID,
			Steps: report.Steps{
				DashboardExistsCheck:  report.NewNotExecutedResult(),
				PanelSnapshot:         nil,
				BasicUILogin:          report.NewNotExecutedResult(),
				PanelSnapshotDownload: nil,
				DataStorePageCreation: report.NewNotExecutedResult(),
				UploadSnapshots:       report.NewNotExecutedResult(),
			},
		}
		snapReport.GrafanaDBReports[reqKey] = &dashReport

		//check if specified dashboard exists. Get the dashboard information so it can be reused to create the snapshot
		gdbExists, _, dashErr := grafana.DashboardExists(dashboard.UID, dashboard.Host, dashboard.Port, dashboard.User.Auth.Basic)
		if dashErr != nil {
			logger.Errorf("Internal error <%v> when checking if dashboard <%v> exists at host <%v> port <%v>", dashErr, dashboard.UID, dashboard.Host, dashboard.Port)
			dashReport.Steps.DashboardExistsCheck = report.Result{
				Result: false,
				Cause:  "Internal error when checking dashboard for existence",
			}
			continue
		}
		if !gdbExists {
			msg := fmt.Sprintf("Dashboard <%v> does not exist at host <%v> port <%v>", dashboard.UID, dashboard.Host, dashboard.Port)
			logger.Debug(msg)
			dashReport.Steps.DashboardExistsCheck = report.Result{
				Result: false,
				Cause:  msg,
			}
			continue
		}

		//for each panel create snapshot(has expiry)

		//run chromedb login

		//build url to snapshot

		//screen shot each snapshot and save to local dir

		//create page under parent page with correct file names

		//upload all attachments with name to page

		//Close off report
		dashReport.Finalize()

	}

	return snapReport, nil
}
