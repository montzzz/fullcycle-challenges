package report

import (
	"bytes"
	"testing"

	"github.com/montzzzzz/challenges/stress-test/internal/dto"
	"github.com/stretchr/testify/suite"
)

type ReportTestSuite struct {
	suite.Suite
	buf *bytes.Buffer
}

func (suite *ReportTestSuite) SetupTest() {
	suite.buf = new(bytes.Buffer)
}

func (suite *ReportTestSuite) TestPrintOutput() {
	r := dto.Report{
		TotalTime: 2 * 1e9,
		Total:     5,
		Success:   3,
		ByStatus:  map[int]int{200: 3, 404: 2},
	}

	PrintTo(r, suite.buf)

	output := suite.buf.String()
	suite.Contains(output, "Total execution time")
	suite.Contains(output, "Total requests: 5")
	suite.Contains(output, "Successful (200): 3")
	suite.Contains(output, "200: 3")
	suite.Contains(output, "404: 2")
}

func TestReportTestSuite(t *testing.T) {
	suite.Run(t, new(ReportTestSuite))
}
