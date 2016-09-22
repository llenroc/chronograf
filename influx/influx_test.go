package influx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/influxdata/mrfusion"
	"github.com/influxdata/mrfusion/influx"
	"golang.org/x/net/context"
)

func Test_Influx_MakesRequestsToQueryEndpoint(t *testing.T) {
	t.Parallel()
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{}`))
		called = true
		if path := r.URL.Path; path != "/query" {
			t.Error("Expected the path to contain `/query` but was", path)
		}
	}))
	defer ts.Close()

	var series mrfusion.TimeSeries
	series, err := influx.NewClient(ts.URL)
	if err != nil {
		t.Fatal("Unexpected error initializing client: err:", err)
	}

	query := mrfusion.Query{
		Command: "show databases",
	}
	_, err = series.Query(context.Background(), query)
	if err != nil {
		t.Fatal("Expected no error but was", err)
	}

	if called == false {
		t.Error("Expected http request to Influx but there was none")
	}
}

func Test_Influx_CancelsInFlightRequests(t *testing.T) {
	t.Parallel()

	started := false
	finished := false
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		started = true
		time.Sleep(2 * time.Second)
		finished = true
	}))
	defer func() {
		ts.CloseClientConnections()
		ts.Close()
	}()

	series, _ := influx.NewClient(ts.URL)
	ctx, cancel := context.WithCancel(context.Background())

	errs := make(chan (error))
	go func() {
		query := mrfusion.Query{
			Command: "show databases",
		}

		_, err := series.Query(ctx, query)
		errs <- err
	}()

	cancel()

	if started != true && finished != false {
		t.Errorf("Expected cancellation during request processing. Started: %t. Finished: %t", started, finished)
	}

	err := <-errs
	if err != mrfusion.ErrUpstreamTimeout {
		t.Error("Expected timeout error but wasn't. err was", err)
	}
}
