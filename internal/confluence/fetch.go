package confluence

import (
	"github.com/sajeevany/graph-snapper/internal/common"
	"github.com/sirupsen/logrus"
)

func PageExistsByID(logger *logrus.Logger, pageUID, host string, port int, user common.Auth) (bool, error) {
	return true, nil
}
