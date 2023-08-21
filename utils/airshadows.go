package utils

type StateStrut struct {
	Desired struct {
		Cmd string `json:"cmd"`
	} `json:"desired"`
	Reported struct {
		Message string `json:"message"`
	} `json:"reported"`
}

type Desired struct {
	Cmd string `json:"cmd"`
}
type Reported struct {
	Message string `json:"message"`
}
type Delta struct {
	Cmd string `json:"cmd"`
}

type ShadowAcceptStrut struct {
	State struct {
		Desired struct {
			Cmd string `json:"cmd"`
		} `json:"desired"`
		Reported struct {
			Message string `json:"message"`
		} `json:"reported"`
		Delta struct {
			Cmd string `json:"cmd"`
		} `json:"delta"`
	} `json:"state"`
	Metadata struct {
		Desired struct {
			Cmd struct {
				Timestamp int `json:"timestamp"`
			} `json:"cmd"`
		} `json:"desired"`
		Reported struct {
			Message struct {
				Timestamp int `json:"timestamp"`
			} `json:"message"`
		} `json:"reported"`
	} `json:"metadata"`
	Version   int `json:"version"`
	Timestamp int `json:"timestamp"`
}
type ShadowDocumentStrut struct {
	Previous struct {
		State struct {
			Desired struct {
				Cmd string `json:"cmd"`
			} `json:"desired"`
			Reported struct {
				Message string `json:"message"`
			} `json:"reported"`
		} `json:"state"`
		Metadata struct {
			Desired struct {
				Cmd struct {
					Timestamp int `json:"timestamp"`
				} `json:"cmd"`
			} `json:"desired"`
			Reported struct {
				Message struct {
					Timestamp int `json:"timestamp"`
				} `json:"message"`
			} `json:"reported"`
		} `json:"metadata"`
		Version int `json:"version"`
	} `json:"previous"`
	Current struct {
		State struct {
			Desired struct {
				Cmd string `json:"cmd"`
			} `json:"desired"`
			Reported struct {
				Message string `json:"message"`
			} `json:"reported"`
		} `json:"state"`
		Metadata struct {
			Desired struct {
				Cmd struct {
					Timestamp int `json:"timestamp"`
				} `json:"cmd"`
			} `json:"desired"`
			Reported struct {
				Message struct {
					Timestamp int `json:"timestamp"`
				} `json:"message"`
			} `json:"reported"`
		} `json:"metadata"`
		Version int `json:"version"`
	} `json:"current"`
	Timestamp int `json:"timestamp"`
}
type ShadowGetAcceptStrut struct {
	State struct {
		Desired struct {
			Cmd string `json:"cmd"`
		} `json:"desired"`
		Reported struct {
			Message string `json:"message"`
		} `json:"reported"`
		Delta struct {
			Cmd string `json:"cmd"`
		} `json:"delta"`
	} `json:"state"`
	Metadata struct {
		Desired struct {
			Cmd struct {
				Timestamp int `json:"timestamp"`
			} `json:"cmd"`
		} `json:"desired"`
		Reported struct {
			Message struct {
				Timestamp int `json:"timestamp"`
			} `json:"message"`
		} `json:"reported"`
	} `json:"metadata"`
	Version   int `json:"version"`
	Timestamp int `json:"timestamp"`
}

type AirDeviceStruct struct {
	Wifi struct {
		Ssid        string `json:"ssid"`
		AirName     string `json:"airName"`
		AirPassword string `json:"airPassword"`
		MacAddress  string `json:"macAddress"`
		IPAddress   string `json:"ipAddress"`
		Version     string `json:"version"`
	} `json:"wifi"`
	Meta struct {
		Errors       []int `json:"errors"`
		InConnected  bool  `json:"inConnected"`
		OutConnected bool  `json:"outConnected"`
	} `json:"meta"`
	Data struct {
		Reg1000 string `json:"reg1000"`
		Reg2000 string `json:"reg2000"`
		Reg3000 string `json:"reg3000"`
		Reg4000 string `json:"reg4000"`
	} `json:"data"`
}
