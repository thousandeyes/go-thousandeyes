package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_CreatePageLoad(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 19:15:36","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":1226422,"testName":"test1","type":"page-load","interval":300,"httpInterval":300,"url":"https://test.com","protocol":"TCP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBgp":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"pageLoadTimeLimit":10,"pageLoadTargetTime":6,"followRedirects":1,"includeHeaders":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":62,"ipAddress":"2001:1890:111d:1::63","countryId":"US","monitorName":"New York, NY-6","network":"AT&T Services, Inc. (AS 7018)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/page-load/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"}],"sslVersion":"Auto"}]}`
	setup()
	defer teardown()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := PageLoad{
		CreatedDate:           "2020-02-06 19:15:36",
		CreatedBy:             "William Fleming (wfleming@grumpysysadm.com)",
		Enabled:               1,
		SavedEvent:            0,
		TestId:                1226422,
		TestName:              "test1",
		Type:                  "page-load",
		Interval:              300,
		HttpInterval:          300,
		Url:                   "https://test.com",
		Protocol:              "TCP",
		NetworkMeasurements:   1,
		MtuMeasurements:       1,
		BandwidthMeasurements: 0,
		BgpMeasurements:       1,
		AlertsEnabled:         1,
		LiveShare:             0,
		HttpTimeLimit:         5,
		HttpTargetTime:        1000,
		HttpVersion:           2,
		PageLoadTimeLimit:     10,
		PageLoadTargetTime:    6,
		IncludeHeaders:        1,
		SslVersionId:          0,
		VerifyCertificate:     1,
		UseNtlm:               0,
		AuthType:              "NONE",
		ProbeMode:             "AUTO",
		Agents: []Agent{
			{
				AgentId:     48620,
				AgentName:   "Seattle, WA (Trial) - IPv6",
				CountryId:   "US",
				IpAddresses: []string{"135.84.184.153"},
				Location:    "Seattle Area",
				Network:     "Astute Hosting Inc. (AS 54527)",
				Prefix:      "135.84.184.0/22",
			},
		},
		SharedWithAccounts: []AccountGroup{
			{
				Aid:  176592,
				Name: "Cloudreach",
			},
		},
		BgpMonitors: []Monitor{
			{
				MonitorId:   62,
				IpAddress:   "2001:1890:111d:1::63",
				CountryId:   "US",
				MonitorName: "New York, NY-6",
				Network:     "AT&T Services, Inc. (AS 7018)",
				MonitorType: "Public",
			},
		},
		NumPathTraces: 3,
		ApiLinks: []ApiLink{
			{
				Rel:  "self",
				Href: "https://api.thousandeyes.com/v6/tests/1226422",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/web/http-server/1226422",
			}, {
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/web/page-load/1226422"},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226422",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226422",
			}, {
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226422",
			},
		},
		SslVersion: "Auto",
	}
	create := PageLoad{
		TestName:     "test1",
		Url:          "https://test.com",
		Interval:     300,
		HttpInterval: 300,
	}
	res, err := client.CreatePageLoad(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetPageLoad(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 19:15:36","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":1226422,"testName":"test1","type":"page-load","interval":300,"httpInterval":300,"url":"https://test.com","protocol":"TCP","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBgp":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"pageLoadTimeLimit":10,"pageLoadTargetTime":6,"followRedirects":1,"includeHeaders":1,"sslVersionId":0,"verifyCertificate":1,"useNtlm":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":62,"ipAddress":"2001:1890:111d:1::63","countryId":"US","monitorName":"New York, NY-6","network":"AT&T Services, Inc. (AS 7018)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/page-load/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226422"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226422"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1226422.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := PageLoad{
		CreatedDate:           "2020-02-06 19:15:36",
		CreatedBy:             "William Fleming (wfleming@grumpysysadm.com)",
		Enabled:               1,
		SavedEvent:            0,
		TestId:                1226422,
		TestName:              "test1",
		Type:                  "page-load",
		Interval:              300,
		HttpInterval:          300,
		Url:                   "https://test.com",
		Protocol:              "TCP",
		NetworkMeasurements:   1,
		MtuMeasurements:       1,
		BandwidthMeasurements: 0,
		BgpMeasurements:       1,
		AlertsEnabled:         1,
		LiveShare:             0,
		HttpTimeLimit:         5,
		HttpTargetTime:        1000,
		HttpVersion:           2,
		PageLoadTimeLimit:     10,
		PageLoadTargetTime:    6,
		IncludeHeaders:        1,
		SslVersionId:          0,
		VerifyCertificate:     1,
		UseNtlm:               0,
		AuthType:              "NONE",
		ProbeMode:             "AUTO",
		Agents: []Agent{
			{
				AgentId:     48620,
				AgentName:   "Seattle, WA (Trial) - IPv6",
				CountryId:   "US",
				IpAddresses: []string{"135.84.184.153"},
				Location:    "Seattle Area",
				Network:     "Astute Hosting Inc. (AS 54527)",
				Prefix:      "135.84.184.0/22",
			},
		},
		SharedWithAccounts: []AccountGroup{
			{
				Aid:  176592,
				Name: "Cloudreach",
			},
		},
		BgpMonitors: []Monitor{
			{
				MonitorId:   62,
				IpAddress:   "2001:1890:111d:1::63",
				CountryId:   "US",
				MonitorName: "New York, NY-6",
				Network:     "AT&T Services, Inc. (AS 7018)",
				MonitorType: "Public",
			},
		},
		NumPathTraces: 3,
		ApiLinks: []ApiLink{
			{
				Rel:  "self",
				Href: "https://api.thousandeyes.com/v6/tests/1226422",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/web/http-server/1226422",
			}, {
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/web/page-load/1226422"},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226422",
			},
			{
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226422",
			}, {
				Rel:  "data",
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226422",
			},
		},
		SslVersion: "Auto",
	}

	res, err := client.GetPageLoad(1226422)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeletePageLoad(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/tests/page-load/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeletePageLoad(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdatePageLoad(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"page-load","url":"https://test.com"}]}`
	mux.HandleFunc("/tests/page-load/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	httpS := PageLoad{Url: "https://test.com"}
	res, err := client.UpdatePageLoad(id, httpS)
	if err != nil {
		t.Fatal(err)
	}
	expected := PageLoad{TestId: 1, TestName: "test123", Type: "page-load", Url: "https://test.com"}
	assert.Equal(t, &expected, res)

}

func TestPageLoad_AddAgent(t *testing.T) {
	test := PageLoad{TestName: "test", Agents: Agents{}}
	expected := PageLoad{TestName: "test", Agents: []Agent{{AgentId: 1}}}
	test.AddAgent(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetPageLoadError(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_PageLoadJsonError(t *testing.T) {
	out := `{"test": [test]}`
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetPageLoad(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'e' in literal true (expecting 'r')")
}

func TestClient_GetPageLoadStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetPageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreatePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreatePageLoad(PageLoad{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdatePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdatePageLoad(1, PageLoad{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeletePageLoadStatusCode(t *testing.T) {
	setup()
	var client = &Client{ApiEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/page-load/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeletePageLoad(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
