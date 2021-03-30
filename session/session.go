package session

import jsoniter "github.com/json-iterator/go"

func GetSessionActive(js []byte) int {
	var session1 ViduSession
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(js, &session1)
	return session1.Numberofelements
}

func GetConnectionActive(js []byte) int {
	var session1 ViduSession
	var connections int = 0
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(js, &session1)
	for i := 0; i < session1.Numberofelements; i++ {
		connections += session1.Content[i].Connections.Numberofelements
	}
	return connections
}

func GetConnectionInSession(js []byte, index int) int {
	var session1 ViduSession
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(js, &session1)
	return session1.Content[index].Connections.Numberofelements
}
