package session

type ViduSession struct {
	Numberofelements int `json:"numberOfElements"`
	Content          []struct {
		ID                string `json:"id"`
		Object            string `json:"object"`
		Sessionid         string `json:"sessionId"`
		Createdat         int64  `json:"createdAt"`
		Mediamode         string `json:"mediaMode"`
		Recordingmode     string `json:"recordingMode"`
		Defaultoutputmode string `json:"defaultOutputMode"`
		Customsessionid   string `json:"customSessionId"`
		Connections       struct {
			Numberofelements int `json:"numberOfElements"`
			Content          []struct {
				ID                      string      `json:"id"`
				Object                  string      `json:"object"`
				Status                  string      `json:"status"`
				Connectionid            string      `json:"connectionId"`
				Sessionid               string      `json:"sessionId"`
				Createdat               int64       `json:"createdAt"`
				Activeat                int64       `json:"activeAt"`
				Location                string      `json:"location"`
				Platform                string      `json:"platform"`
				Token                   interface{} `json:"token"`
				Type                    string      `json:"type"`
				Record                  bool        `json:"record"`
				Role                    interface{} `json:"role"`
				Kurentooptions          interface{} `json:"kurentoOptions"`
				Rtspuri                 string      `json:"rtspUri"`
				Adaptativebitrate       bool        `json:"adaptativeBitrate"`
				Onlyplaywithsubscribers bool        `json:"onlyPlayWithSubscribers"`
				Networkcache            int         `json:"networkCache"`
				Serverdata              string      `json:"serverData"`
				Clientdata              interface{} `json:"clientData"`
				Publishers              []struct {
					Createdat    int64  `json:"createdAt"`
					Streamid     string `json:"streamId"`
					Rtspuri      string `json:"rtspUri"`
					Mediaoptions struct {
						Hasaudio        bool        `json:"hasAudio"`
						Audioactive     bool        `json:"audioActive"`
						Hasvideo        bool        `json:"hasVideo"`
						Videoactive     bool        `json:"videoActive"`
						Typeofvideo     string      `json:"typeOfVideo"`
						Framerate       interface{} `json:"frameRate"`
						Videodimensions interface{} `json:"videoDimensions"`
						Filter          struct {
						} `json:"filter"`
						Adaptativebitrate       bool `json:"adaptativeBitrate"`
						Onlyplaywithsubscribers bool `json:"onlyPlayWithSubscribers"`
						Networkcache            int  `json:"networkCache"`
					} `json:"mediaOptions"`
				} `json:"publishers"`
				Subscribers []interface{} `json:"subscribers"`
			} `json:"content"`
		} `json:"connections"`
		Recording bool `json:"recording"`
	} `json:"content"`
}

var testOpenviduJson = `
{
  "numberOfElements": 54,
  "content": [
    {
      "id": "1294",
      "object": "session",
      "sessionId": "1294",
      "createdAt": 1615791529046,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1294",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_GVQJ_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_GVQJ_203_113_25_162_10306_live_0",
            "sessionId": "1294",
            "createdAt": 1615793472938,
            "activeAt": 1615793472938,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615793472943,
                "streamId": "str_IPC_Uhdl_ipc_IPCAM_rtsp_GVQJ_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1573",
      "object": "session",
      "sessionId": "1573",
      "createdAt": 1616850848032,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1573",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_2U6X_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_2U6X_203_113_25_162_10306_live_0",
            "sessionId": "1573",
            "createdAt": 1616850848063,
            "activeAt": 1616850848063,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616850848070,
                "streamId": "str_IPC_XHwp_ipc_IPCAM_rtsp_2U6X_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1416",
      "object": "session",
      "sessionId": "1416",
      "createdAt": 1616127606753,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1416",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_JBAA_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_JBAA_203_113_25_162_10306_live_0",
            "sessionId": "1416",
            "createdAt": 1616127606785,
            "activeAt": 1616127606785,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616127606793,
                "streamId": "str_IPC_SN1V_ipc_IPCAM_rtsp_JBAA_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1502",
      "object": "session",
      "sessionId": "1502",
      "createdAt": 1616423985905,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1502",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_779O_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_779O_203_113_25_162_10906_live_0",
            "sessionId": "1502",
            "createdAt": 1616424020518,
            "activeAt": 1616424020518,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616424020525,
                "streamId": "str_IPC_LW9U_ipc_IPCAM_rtsp_779O_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1130",
      "object": "session",
      "sessionId": "1130",
      "createdAt": 1615622943349,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1130",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_B3NI_192_168_9_6_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_B3NI_192_168_9_6_10106_live_0",
            "sessionId": "1130",
            "createdAt": 1615624823277,
            "activeAt": 1615624823277,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://192.168.9.6:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615624823283,
                "streamId": "str_IPC_Kz7U_ipc_IPCAM_rtsp_B3NI_192_168_9_6_10106_live_0",
                "rtspUri": "rtsp://192.168.9.6:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1492",
      "object": "session",
      "sessionId": "1492",
      "createdAt": 1616421107908,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1492",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_PUWV_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_PUWV_203_113_25_162_10906_live_0",
            "sessionId": "1492",
            "createdAt": 1616421108132,
            "activeAt": 1616421108133,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421108139,
                "streamId": "str_IPC_GATc_ipc_IPCAM_rtsp_PUWV_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1325",
      "object": "session",
      "sessionId": "1325",
      "createdAt": 1616121893022,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1325",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_1MBE_192_168_9_6_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_1MBE_192_168_9_6_10306_live_0",
            "sessionId": "1325",
            "createdAt": 1616123534088,
            "activeAt": 1616123534088,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://192.168.9.6:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616123534094,
                "streamId": "str_IPC_AdGF_ipc_IPCAM_rtsp_1MBE_192_168_9_6_10306_live_0",
                "rtspUri": "rtsp://192.168.9.6:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1503",
      "object": "session",
      "sessionId": "1503",
      "createdAt": 1616424033920,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1503",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_6O4E_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_6O4E_203_113_25_162_10906_live_0",
            "sessionId": "1503",
            "createdAt": 1616424039956,
            "activeAt": 1616424039956,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616424039962,
                "streamId": "str_IPC_DC8J_ipc_IPCAM_rtsp_6O4E_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1434",
      "object": "session",
      "sessionId": "1434",
      "createdAt": 1616227129393,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1434",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_0TTP_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_0TTP_203_113_25_162_10306_live_0",
            "sessionId": "1434",
            "createdAt": 1616227129453,
            "activeAt": 1616227129453,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616227129461,
                "streamId": "str_IPC_WYO0_ipc_IPCAM_rtsp_0TTP_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1490",
      "object": "session",
      "sessionId": "1490",
      "createdAt": 1616421059279,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1490",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_FKR6_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_FKR6_203_113_25_162_10906_live_0",
            "sessionId": "1490",
            "createdAt": 1616421059515,
            "activeAt": 1616421059515,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421059522,
                "streamId": "str_IPC_ZN5t_ipc_IPCAM_rtsp_FKR6_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1497",
      "object": "session",
      "sessionId": "1497",
      "createdAt": 1616421930595,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1497",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_2KGJ_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_2KGJ_203_113_25_162_10906_live_0",
            "sessionId": "1497",
            "createdAt": 1616421933438,
            "activeAt": 1616421933438,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421933446,
                "streamId": "str_IPC_Wwpl_ipc_IPCAM_rtsp_2KGJ_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1537",
      "object": "session",
      "sessionId": "1537",
      "createdAt": 1616597450554,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1537",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_EQAH_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_EQAH_203_113_25_162_10306_live_0",
            "sessionId": "1537",
            "createdAt": 1616597450589,
            "activeAt": 1616597450589,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616597450597,
                "streamId": "str_IPC_TYEI_ipc_IPCAM_rtsp_EQAH_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1531",
      "object": "session",
      "sessionId": "1531",
      "createdAt": 1616577767894,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1531",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_F6KR_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_F6KR_203_113_25_162_10306_live_0",
            "sessionId": "1531",
            "createdAt": 1616577767925,
            "activeAt": 1616577767926,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616577767978,
                "streamId": "str_IPC_JDAd_ipc_IPCAM_rtsp_F6KR_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1499",
      "object": "session",
      "sessionId": "1499",
      "createdAt": 1616423834440,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1499",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_EJT6_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_EJT6_203_113_25_162_10906_live_0",
            "sessionId": "1499",
            "createdAt": 1616423934482,
            "activeAt": 1616423934482,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616423934490,
                "streamId": "str_IPC_Dplg_ipc_IPCAM_rtsp_EJT6_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1518",
      "object": "session",
      "sessionId": "1518",
      "createdAt": 1616493827641,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1518",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_E1QO_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_E1QO_203_113_25_162_10306_live_0",
            "sessionId": "1518",
            "createdAt": 1616493827671,
            "activeAt": 1616493827671,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616493827678,
                "streamId": "str_IPC_UJg2_ipc_IPCAM_rtsp_E1QO_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1498",
      "object": "session",
      "sessionId": "1498",
      "createdAt": 1616423798137,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1498",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_PRRF_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_PRRF_203_113_25_162_10906_live_0",
            "sessionId": "1498",
            "createdAt": 1616423807507,
            "activeAt": 1616423807507,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616423807516,
                "streamId": "str_IPC_R5p6_ipc_IPCAM_rtsp_PRRF_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1458",
      "object": "session",
      "sessionId": "1458",
      "createdAt": 1616376897895,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1458",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_KUN4_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_KUN4_203_113_25_162_10306_live_0",
            "sessionId": "1458",
            "createdAt": 1616376897925,
            "activeAt": 1616376897926,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616376897934,
                "streamId": "str_IPC_TcUF_ipc_IPCAM_rtsp_KUN4_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1588",
      "object": "session",
      "sessionId": "1588",
      "createdAt": 1616984984204,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1588",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    },
    {
      "id": "1284",
      "object": "session",
      "sessionId": "1284",
      "createdAt": 1615780959490,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1284",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_MMAM_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_MMAM_203_113_25_162_10106_live_0",
            "sessionId": "1284",
            "createdAt": 1615783291770,
            "activeAt": 1615783291770,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615783291776,
                "streamId": "str_IPC_Bls2_ipc_IPCAM_rtsp_MMAM_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1415",
      "object": "session",
      "sessionId": "1415",
      "createdAt": 1616127461908,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1415",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_2HHG_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_2HHG_203_113_25_162_10306_live_0",
            "sessionId": "1415",
            "createdAt": 1616127461938,
            "activeAt": 1616127461938,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616127461945,
                "streamId": "str_IPC_IBs9_ipc_IPCAM_rtsp_2HHG_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1465",
      "object": "session",
      "sessionId": "1465",
      "createdAt": 1616398683287,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1465",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_DMSI_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_DMSI_203_113_25_162_10106_live_0",
            "sessionId": "1465",
            "createdAt": 1616398683317,
            "activeAt": 1616398683317,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616398683326,
                "streamId": "str_IPC_Gsa8_ipc_IPCAM_rtsp_DMSI_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1460",
      "object": "session",
      "sessionId": "1460",
      "createdAt": 1616387947730,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1460",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_XYKT_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_XYKT_203_113_25_162_10306_live_0",
            "sessionId": "1460",
            "createdAt": 1616387947762,
            "activeAt": 1616387947762,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616387947771,
                "streamId": "str_IPC_BTg7_ipc_IPCAM_rtsp_XYKT_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1493",
      "object": "session",
      "sessionId": "1493",
      "createdAt": 1616421122504,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1493",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_R9XK_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_R9XK_203_113_25_162_10906_live_0",
            "sessionId": "1493",
            "createdAt": 1616421122707,
            "activeAt": 1616421122707,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421122714,
                "streamId": "str_IPC_LdRV_ipc_IPCAM_rtsp_R9XK_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1432",
      "object": "session",
      "sessionId": "1432",
      "createdAt": 1616213891338,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1432",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_WSNB_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_WSNB_203_113_25_162_10106_live_0",
            "sessionId": "1432",
            "createdAt": 1616213891368,
            "activeAt": 1616213891369,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616213891377,
                "streamId": "str_IPC_EWxx_ipc_IPCAM_rtsp_WSNB_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1579",
      "object": "session",
      "sessionId": "1579",
      "createdAt": 1616906209833,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1579",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_CF5R_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_CF5R_203_113_25_162_10306_live_0",
            "sessionId": "1579",
            "createdAt": 1616906209865,
            "activeAt": 1616906209866,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616906209874,
                "streamId": "str_IPC_INMl_ipc_IPCAM_rtsp_CF5R_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1494",
      "object": "session",
      "sessionId": "1494",
      "createdAt": 1616421136818,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1494",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_V5TN_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_V5TN_203_113_25_162_10906_live_0",
            "sessionId": "1494",
            "createdAt": 1616421137047,
            "activeAt": 1616421137048,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421137055,
                "streamId": "str_IPC_Ge4f_ipc_IPCAM_rtsp_V5TN_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1590",
      "object": "session",
      "sessionId": "1590",
      "createdAt": 1617001864205,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1590",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": false
    },
    {
      "id": "1552",
      "object": "session",
      "sessionId": "1552",
      "createdAt": 1616653541252,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1552",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_BN3S_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_BN3S_203_113_25_162_10306_live_0",
            "sessionId": "1552",
            "createdAt": 1616653541282,
            "activeAt": 1616653541283,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616653541290,
                "streamId": "str_IPC_Kmcc_ipc_IPCAM_rtsp_BN3S_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1578",
      "object": "session",
      "sessionId": "1578",
      "createdAt": 1616904134268,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1578",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_TTJV_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_TTJV_203_113_25_162_10306_live_0",
            "sessionId": "1578",
            "createdAt": 1616904134302,
            "activeAt": 1616904134303,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616904134353,
                "streamId": "str_IPC_CjdU_ipc_IPCAM_rtsp_TTJV_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1565",
      "object": "session",
      "sessionId": "1565",
      "createdAt": 1616830431347,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1565",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    },
    {
      "id": "1583",
      "object": "session",
      "sessionId": "1583",
      "createdAt": 1616931327489,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1583",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_QJJI_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_QJJI_203_113_25_162_10106_live_0",
            "sessionId": "1583",
            "createdAt": 1616931327524,
            "activeAt": 1616931327524,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616931327531,
                "streamId": "str_IPC_QdU6_ipc_IPCAM_rtsp_QJJI_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1489",
      "object": "session",
      "sessionId": "1489",
      "createdAt": 1616420987608,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1489",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_HVHV_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_HVHV_203_113_25_162_10906_live_0",
            "sessionId": "1489",
            "createdAt": 1616420987979,
            "activeAt": 1616420987979,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616420987986,
                "streamId": "str_IPC_MxZ9_ipc_IPCAM_rtsp_HVHV_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1524",
      "object": "session",
      "sessionId": "1524",
      "createdAt": 1616557239976,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1524",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_SCZ3_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_SCZ3_203_113_25_162_10106_live_0",
            "sessionId": "1524",
            "createdAt": 1616557240008,
            "activeAt": 1616557240008,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616557240017,
                "streamId": "str_IPC_PG0J_ipc_IPCAM_rtsp_SCZ3_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1515",
      "object": "session",
      "sessionId": "1515",
      "createdAt": 1616481334319,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1515",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_KXZI_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_KXZI_203_113_25_162_10306_live_0",
            "sessionId": "1515",
            "createdAt": 1616481334350,
            "activeAt": 1616481334350,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616481334357,
                "streamId": "str_IPC_Zc6Q_ipc_IPCAM_rtsp_KXZI_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1226",
      "object": "session",
      "sessionId": "1226",
      "createdAt": 1615719466647,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1226",
      "connections": {
        "numberOfElements": 2,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_OSM5_192_168_9_6_10406_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_OSM5_192_168_9_6_10406_live_0",
            "sessionId": "1226",
            "createdAt": 1615721054407,
            "activeAt": 1615721054407,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://192.168.9.6:10406/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615721054414,
                "streamId": "str_IPC_PNYG_ipc_IPCAM_rtsp_OSM5_192_168_9_6_10406_live_0",
                "rtspUri": "rtsp://192.168.9.6:10406/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "ipc_IPCAM_rtsp_PAIY_192_168_9_6_10406_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_PAIY_192_168_9_6_10406_live_0",
            "sessionId": "1226",
            "createdAt": 1615719590479,
            "activeAt": 1615719590480,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://192.168.9.6:10406/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615719590485,
                "streamId": "str_IPC_WEcS_ipc_IPCAM_rtsp_PAIY_192_168_9_6_10406_live_0",
                "rtspUri": "rtsp://192.168.9.6:10406/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1292",
      "object": "session",
      "sessionId": "1292",
      "createdAt": 1615785365197,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1292",
      "connections": {
        "numberOfElements": 5,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_QED5_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_QED5_203_113_25_162_10106_live_0",
            "sessionId": "1292",
            "createdAt": 1615868367544,
            "activeAt": 1615868367544,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615868367551,
                "streamId": "str_IPC_S1F9_ipc_IPCAM_rtsp_QED5_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "ipc_IPCAM_rtsp_MM35_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_MM35_203_113_25_162_10106_live_0",
            "sessionId": "1292",
            "createdAt": 1615868331754,
            "activeAt": 1615868331755,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615868331761,
                "streamId": "str_IPC_DUpu_ipc_IPCAM_rtsp_MM35_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "ipc_IPCAM_rtsp_UHBT_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_UHBT_203_113_25_162_10106_live_0",
            "sessionId": "1292",
            "createdAt": 1615868383997,
            "activeAt": 1615868383997,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615868384004,
                "streamId": "str_IPC_Jc5w_ipc_IPCAM_rtsp_UHBT_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "ipc_IPCAM_rtsp_CL7R_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_CL7R_203_113_25_162_10106_live_0",
            "sessionId": "1292",
            "createdAt": 1615785407921,
            "activeAt": 1615785407921,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615785407929,
                "streamId": "str_IPC_Mn0Y_ipc_IPCAM_rtsp_CL7R_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "ipc_IPCAM_rtsp_YKUU_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_YKUU_203_113_25_162_10106_live_0",
            "sessionId": "1292",
            "createdAt": 1615868424459,
            "activeAt": 1615868424459,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1615868424466,
                "streamId": "str_IPC_WkYf_ipc_IPCAM_rtsp_YKUU_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": true
    },
    {
      "id": "1423",
      "object": "session",
      "sessionId": "1423",
      "createdAt": 1616157230523,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1423",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_VOVO_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_VOVO_203_113_25_162_10106_live_0",
            "sessionId": "1423",
            "createdAt": 1616157230554,
            "activeAt": 1616157230554,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616157230562,
                "streamId": "str_IPC_GKvA_ipc_IPCAM_rtsp_VOVO_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1338",
      "object": "session",
      "sessionId": "1338",
      "createdAt": 1615949519510,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1338",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    },
    {
      "id": "1491",
      "object": "session",
      "sessionId": "1491",
      "createdAt": 1616421097927,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1491",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_XURJ_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_XURJ_203_113_25_162_10906_live_0",
            "sessionId": "1491",
            "createdAt": 1616421098161,
            "activeAt": 1616421098162,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421098168,
                "streamId": "str_IPC_OSfQ_ipc_IPCAM_rtsp_XURJ_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1425",
      "object": "session",
      "sessionId": "1425",
      "createdAt": 1616160308828,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1425",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_JRS4_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_JRS4_203_113_25_162_10106_live_0",
            "sessionId": "1425",
            "createdAt": 1616160308858,
            "activeAt": 1616160308858,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616160308910,
                "streamId": "str_IPC_O67r_ipc_IPCAM_rtsp_JRS4_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1557",
      "object": "session",
      "sessionId": "1557",
      "createdAt": 1616678945525,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1557",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_SELB_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_SELB_203_113_25_162_10106_live_0",
            "sessionId": "1557",
            "createdAt": 1616678945562,
            "activeAt": 1616678945562,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616678945572,
                "streamId": "str_IPC_QZrL_ipc_IPCAM_rtsp_SELB_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1589",
      "object": "session",
      "sessionId": "1589",
      "createdAt": 1616987798287,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1589",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_SAVT_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_SAVT_203_113_25_162_10306_live_0",
            "sessionId": "1589",
            "createdAt": 1616987798319,
            "activeAt": 1616987798319,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616987798328,
                "streamId": "str_IPC_Ojye_ipc_IPCAM_rtsp_SAVT_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1496",
      "object": "session",
      "sessionId": "1496",
      "createdAt": 1616421711603,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1496",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_DBDE_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_DBDE_203_113_25_162_10906_live_0",
            "sessionId": "1496",
            "createdAt": 1616421719504,
            "activeAt": 1616421719505,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616421719512,
                "streamId": "str_IPC_YkXg_ipc_IPCAM_rtsp_DBDE_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1519",
      "object": "session",
      "sessionId": "1519",
      "createdAt": 1616495181691,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1519",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_6IUU_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_6IUU_203_113_25_162_10106_live_0",
            "sessionId": "1519",
            "createdAt": 1616495181723,
            "activeAt": 1616495181723,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616495181731,
                "streamId": "str_IPC_Zoxb_ipc_IPCAM_rtsp_6IUU_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1390",
      "object": "session",
      "sessionId": "1390",
      "createdAt": 1616131839325,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1390",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    },
    {
      "id": "1392",
      "object": "session",
      "sessionId": "1392",
      "createdAt": 1616036438708,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1392",
      "connections": {
        "numberOfElements": 3,
        "content": [
          {
            "id": "con_SYAKSM8Rir",
            "object": "connection",
            "status": "active",
            "connectionId": "con_SYAKSM8Rir",
            "sessionId": "1392",
            "createdAt": 1617003110842,
            "activeAt": 1617003111047,
            "location": "unknown",
            "platform": "Chrome 89.0.4389.90 on Windows 10 64-bit",
            "token": "wss://mis-vdo-03.dems1669.com?sessionId=1392&token=tok_PXb3ku7ZHDZLAA8J&role=PUBLISHER&version=2.16.0&coturnIp=203.113.25.164&turnUsername=SNTSBB&turnCredential=8etdbc",
            "type": "WEBRTC",
            "record": true,
            "role": "PUBLISHER",
            "kurentoOptions": null,
            "rtspUri": null,
            "adaptativeBitrate": null,
            "onlyPlayWithSubscribers": null,
            "networkCache": null,
            "serverData": "",
            "clientData": "{\"clientData\":\"นุชจนีย์\"}",
            "publishers": [],
            "subscribers": [
              {
                "createdAt": 1617003331856,
                "streamId": "str_IPC_Qe2B_ipc_IPCAM_rtsp_RJEO_203_113_25_162_10106_live_0"
              }
            ]
          },
          {
            "id": "ipc_IPCAM_rtsp_RJEO_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_RJEO_203_113_25_162_10106_live_0",
            "sessionId": "1392",
            "createdAt": 1616036438742,
            "activeAt": 1616036438742,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616036438790,
                "streamId": "str_IPC_Qe2B_ipc_IPCAM_rtsp_RJEO_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          },
          {
            "id": "con_JCkMUB8N4T",
            "object": "connection",
            "status": "active",
            "connectionId": "con_JCkMUB8N4T",
            "sessionId": "1392",
            "createdAt": 1617005249784,
            "activeAt": 1617005249854,
            "location": "unknown",
            "platform": "Chrome 89.0.4389.90 on Windows 10 64-bit",
            "token": "wss://mis-vdo-03.dems1669.com?sessionId=1392&token=tok_XA3gzPqhl12jv5Rx&role=PUBLISHER&version=2.16.0&coturnIp=203.113.25.164&turnUsername=Q7CSOK&turnCredential=fbm6o8",
            "type": "WEBRTC",
            "record": true,
            "role": "PUBLISHER",
            "kurentoOptions": null,
            "rtspUri": null,
            "adaptativeBitrate": null,
            "onlyPlayWithSubscribers": null,
            "networkCache": null,
            "serverData": "",
            "clientData": "{\"clientData\":\"นุชจนีย์\"}",
            "publishers": [],
            "subscribers": [
              {
                "createdAt": 1617005250012,
                "streamId": "str_IPC_Qe2B_ipc_IPCAM_rtsp_RJEO_203_113_25_162_10106_live_0"
              }
            ]
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1424",
      "object": "session",
      "sessionId": "1424",
      "createdAt": 1616159569139,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1424",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_E84U_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_E84U_203_113_25_162_10306_live_0",
            "sessionId": "1424",
            "createdAt": 1616159569168,
            "activeAt": 1616159569168,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616159569177,
                "streamId": "str_IPC_XQet_ipc_IPCAM_rtsp_E84U_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1451",
      "object": "session",
      "sessionId": "1451",
      "createdAt": 1616307247812,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1451",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_VUEV_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_VUEV_203_113_25_162_10306_live_0",
            "sessionId": "1451",
            "createdAt": 1616307247844,
            "activeAt": 1616307247844,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616307247851,
                "streamId": "str_IPC_VFCr_ipc_IPCAM_rtsp_VUEV_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1545",
      "object": "session",
      "sessionId": "1545",
      "createdAt": 1616637985496,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1545",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_2RLO_203_113_25_162_10106_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_2RLO_203_113_25_162_10106_live_0",
            "sessionId": "1545",
            "createdAt": 1616637985526,
            "activeAt": 1616637985527,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10106/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616637985537,
                "streamId": "str_IPC_GBlD_ipc_IPCAM_rtsp_2RLO_203_113_25_162_10106_live_0",
                "rtspUri": "rtsp://203.113.25.162:10106/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1417",
      "object": "session",
      "sessionId": "1417",
      "createdAt": 1616132035164,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1417",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_V7X4_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_V7X4_203_113_25_162_10306_live_0",
            "sessionId": "1417",
            "createdAt": 1616132035223,
            "activeAt": 1616132035223,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616132035229,
                "streamId": "str_IPC_PmZh_ipc_IPCAM_rtsp_V7X4_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1530",
      "object": "session",
      "sessionId": "1530",
      "createdAt": 1616673579276,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1530",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    },
    {
      "id": "1501",
      "object": "session",
      "sessionId": "1501",
      "createdAt": 1616423953094,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1501",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_YKGH_203_113_25_162_10906_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_YKGH_203_113_25_162_10906_live_0",
            "sessionId": "1501",
            "createdAt": 1616423977883,
            "activeAt": 1616423977883,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10906/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616423977934,
                "streamId": "str_IPC_U5IM_ipc_IPCAM_rtsp_YKGH_203_113_25_162_10906_live_0",
                "rtspUri": "rtsp://203.113.25.162:10906/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1420",
      "object": "session",
      "sessionId": "1420",
      "createdAt": 1616153433578,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1420",
      "connections": {
        "numberOfElements": 1,
        "content": [
          {
            "id": "ipc_IPCAM_rtsp_RAH4_203_113_25_162_10306_live_0",
            "object": "connection",
            "status": "active",
            "connectionId": "ipc_IPCAM_rtsp_RAH4_203_113_25_162_10306_live_0",
            "sessionId": "1420",
            "createdAt": 1616153433608,
            "activeAt": 1616153433608,
            "location": "unknown",
            "platform": "IPCAM",
            "token": null,
            "type": "IPCAM",
            "record": true,
            "role": null,
            "kurentoOptions": null,
            "rtspUri": "rtsp://203.113.25.162:10306/live/0",
            "adaptativeBitrate": true,
            "onlyPlayWithSubscribers": true,
            "networkCache": 2000,
            "serverData": "Security camera",
            "clientData": null,
            "publishers": [
              {
                "createdAt": 1616153433615,
                "streamId": "str_IPC_OZqk_ipc_IPCAM_rtsp_RAH4_203_113_25_162_10306_live_0",
                "rtspUri": "rtsp://203.113.25.162:10306/live/0",
                "mediaOptions": {
                  "hasAudio": true,
                  "audioActive": true,
                  "hasVideo": true,
                  "videoActive": true,
                  "typeOfVideo": "IPCAM",
                  "frameRate": null,
                  "videoDimensions": null,
                  "filter": {},
                  "adaptativeBitrate": true,
                  "onlyPlayWithSubscribers": true,
                  "networkCache": 2000
                }
              }
            ],
            "subscribers": []
          }
        ]
      },
      "recording": false
    },
    {
      "id": "1566",
      "object": "session",
      "sessionId": "1566",
      "createdAt": 1616789385197,
      "mediaMode": "ROUTED",
      "recordingMode": "ALWAYS",
      "defaultOutputMode": "INDIVIDUAL",
      "customSessionId": "1566",
      "connections": {
        "numberOfElements": 0,
        "content": []
      },
      "recording": true
    }
  ]
}
`
