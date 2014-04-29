package sdp

import (
    "time"
    )

var (
    MediaTypes = []string{"audio", "video", "text", "application", "message"}
    TransportTypes =  []string{"udp", "RTP/AVP", "RTP/SAVP"}
    AttrTypes = []string{"cat", "keywds", "tool", "ptime", "maxptime", "rtpmap", "orient", "type", "charset", "framerate", "quality", "fmtp", "recvonly", "sendrecv", "sendonly", "inactive", "sdplang", "lang"}
    KeyTypes = []string{"prompt", "clear", "base64", "uri"}
    )

type SessionDescription struct {
    Version           int
    Origin            Origin
    SessionName       string
    Info              string
    Uri               string
    Emails            []Email
    Phones            []Phone
    Connection        Connection
    Bandwidths        []Bandwidth
    Times             []TimeDescription
    Key               Key
    Attributes        map[string]string
    MediaDescriptions []MediaDescription
}

type Origin struct {
    Username       string
    SessionId      string
    SessionVersion string
    NetType        string
    AddrType       string
    UnicastAddr    string
}

type Email struct {
    Address string
    Name    string
}

type Phone struct {
    Address string
    Name    string
}

type Connection struct {
    NetType  string
    AddrType string
    Address  string
}

type Bandwidth struct {
    Type string
    Bandwidth string
}

type TimeDescription struct {
    Start time.Time
    Stop time.Time
    Repeats []Repeat
    Zones []Zone
}

type Repeat struct {
    Interval time.Duration
    Active   time.Duration
    Offsets  []time.Duration
}

type Zone struct {
    Time   time.Time
    Offset time.Duration
}

type Key struct {
    Method string
    Key    string
}

type MediaDescription struct {
    Type        string
    Port        int
    NumPorts    int
    Proto       string
    Fmt         string
    Info        string
    Connections []Connection
    Bandwidths  []Bandwidth
    Key         Key
    Attributes  map[string]string
}
