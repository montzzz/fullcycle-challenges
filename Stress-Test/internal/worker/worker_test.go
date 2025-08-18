package worker

import (
	"net/http"
	"testing"
	"time"

	"github.com/montzzzzz/challenges/stress-test/internal/config"
	"github.com/montzzzzz/challenges/stress-test/internal/dto"
	"github.com/stretchr/testify/suite"
)

type RunnerTestSuite struct {
	suite.Suite
}

func (suite *RunnerTestSuite) TestAggregate() {
	results := make(chan dto.Result, 4)
	results <- dto.Result{StatusCode: 200}
	results <- dto.Result{StatusCode: 200}
	results <- dto.Result{StatusCode: 404}
	results <- dto.Result{StatusCode: 500}
	close(results)

	report := aggregate(results, 1*time.Second)

	suite.Equal(4, report.Total)
	suite.Equal(2, report.Success)
	suite.Equal(1, report.ByStatus[404])
	suite.Equal(1, report.ByStatus[500])
}

func (suite *RunnerTestSuite) TestRunTestWithMock() {
	http.DefaultClient.Transport = &mockRoundTripper{}

	cfg := config.Config{
		URL:         "https://example.com",
		Requests:    10,
		Concurrency: 3,
	}

	report := RunTest(cfg)

	suite.Equal(cfg.Requests, report.Total)
	suite.Equal(cfg.Requests, report.Success)
}

func TestRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(RunnerTestSuite))
}

type mockRoundTripper struct{}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       http.NoBody,
	}, nil
}
