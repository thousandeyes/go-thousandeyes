package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAgents(t *testing.T) {
	out := `{"agents":[{"agentId": 1, "enabled": 1}, {"agentId": 2, "enabled": 0}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := Agents{
		Agent{
			AgentID: 1,
			Enabled: 1,
		},
		Agent{
			AgentID: 2,
			Enabled: 0,
		},
	}
	res, err := client.GetAgents()
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgent(t *testing.T) {
	out := `{"agents":[{"agentId": 1, "enabled": 1}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	expected := Agent{AgentID: 1, Enabled: 1}
	res, err := client.GetAgent(1)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetAgentsError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetAgents()
	teardown()
	assert.Error(t, err)
}

func TestClient_GetAgentsJsonError(t *testing.T) {
	out := `{"agents":[{"agentId":4492,agentName":"Dallas, TX (Trial)","agentType":"Cloud","countryId":"US","ipAddresses":["104.130.154.136","104.130.156.108","104.130.141.203","104.130.155.161"],"location":"Dallas Area"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgents()
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}

func TestClient_GetAgentJsonError(t *testing.T) {
	out := `{"agents":[{"agentId":4492,agentName":"Dallas, TX (Trial)","agentType":"Cloud","countryId":"US","ipAddresses":["104.130.154.136","104.130.156.108","104.130.141.203","104.130.155.161"],"location":"Dallas Area"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetAgent(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'a' looking for beginning of object key string")
}

func TestClient_GetAgentsStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetAgents()
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_GetAgentStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetAgent(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
