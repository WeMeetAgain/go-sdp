package sdp

import (
    "time"
    "testing"
    )

var s1 = 
`v=0
o=jdoe 2890844526 2890842807 IN IP4 10.47.16.5
s=SDP Seminar
i=A Seminar on the session description protocol
u=http://www.example.com/seminars/sdp.pdf
e=j.doe@example.com (Jane Doe)
p=Jane Doe <123-456-7890>
c=IN IP4 224.2.17.12/127
t=2873397496 2873404696
r=7d 1h 0 25h
r=7d 1h 0 25h
z=2882844526 -1h 2898848070 0
a=recvonly
m=audio 49170 RTP/AVP 0
m=video 51372 RTP/AVP 99
a=rtpmap:99 h263-1998/90000
`

func TestDecode(t *testing.T) {
    sd, err := Decode(s1)
    if err != nil {
        t.Error(err)
    }
    if sd.Version != 0 {
        t.Errorf("Wrong Version: %s", sd.Version)
    }
    if sd.Origin.Username != "jdoe" {
        t.Errorf("Wrong Username: %s", sd.Origin.Username)
    }
    if sd.Origin.SessionId != "2890844526" {
        t.Errorf("Wrong Session Id: %s", sd.Origin.SessionId)
    }
    if sd.Origin.SessionVersion != "2890842807" {
        t.Errorf("Wrong Session Version: %s", sd.Origin.SessionVersion)
    }
    if sd.Origin.NetType != "IN" {
        t.Errorf("Wrong Net Type: %s", sd.Origin.NetType)
    }
    if sd.Origin.AddrType != "IP4" {
        t.Errorf("Wrong Address Type: %s", sd.Origin.AddrType)
    }
    if sd.Origin.UnicastAddr != "10.47.16.5" {
        t.Errorf("Wrong Unicast Address: %s", sd.Origin.UnicastAddr)
    }
    if sd.Info != "A Seminar on the session description protocol" {
        t.Errorf("Wrong Info: %s", sd.Info)
    }
    if sd.Uri != "http://www.example.com/seminars/sdp.pdf" {
        t.Errorf("Wrong Uri: %s", sd.Uri)
    }
    if sd.Emails[0].Address != "j.doe@example.com" {
        t.Errorf("Wrong Email Address: %s", sd.Emails[0].Address)
    }
    if sd.Emails[0].Name != "Jane Doe" {
        t.Errorf("Wrong Email Name: %s", sd.Emails[0].Name)
    }
    if sd.Phones[0].Address != "123-456-7890" {
        t.Errorf("Wrong Phone Address: %s", sd.Phones[0].Address)
    }
    if sd.Phones[0].Name != "Jane Doe" {
        t.Errorf("Wrong Phone Name: %s", sd.Phones[0].Name)
    }
    if sd.Connection.NetType != "IN" {
        t.Errorf("Wrong Connection Net Type: %s", sd.Connection.NetType)
    }
    if sd.Connection.AddrType != "IP4" {
        t.Errorf("Wrong Connection Address Type: %s", sd.Connection.AddrType)
    }
    if sd.Connection.Address != "224.2.17.12/127" {
        t.Errorf("Wrong Connection Address: %s", sd.Connection.Address)
    }
    if sd.Times[0].Start.Unix() != 2873397496-ntpUnix {
        t.Errorf("Wrong Phone Address: %s", sd.Times[0].Start)
    }
    if sd.Times[0].Stop.Unix() != 2873404696-ntpUnix {
        t.Errorf("Wrong Phone Name: %s", sd.Times[0].Stop)
    }
    if sd.Times[0].Repeats[0].Interval != (time.Hour * 24 * 7) {
        t.Errorf("Wrong Repeat Interval: %s", sd.Times[0].Repeats[0].Interval)
    }
    if sd.Times[0].Repeats[0].Active != time.Hour {
        t.Errorf("Wrong Repeat Active: %s", sd.Times[0].Repeats[0].Active)
    }
    if sd.Times[0].Repeats[0].Offsets[0] != 0 {
        t.Errorf("Wrong Repeat Offset: %s", sd.Times[0].Repeats[0].Offsets[0])
    }
    if sd.Times[0].Repeats[0].Offsets[1] != (time.Hour * 25) {
        t.Errorf("Wrong Repeat Offset: %s", sd.Times[0].Repeats[0].Offsets[1])
    }
    if sd.Times[0].Repeats[1].Interval != (time.Hour * 24 * 7) {
        t.Errorf("Wrong Repeat Interval: %s", sd.Times[0].Repeats[1].Interval)
    }
    if sd.Times[0].Repeats[1].Active != time.Hour {
        t.Errorf("Wrong Repeat Active: %s", sd.Times[0].Repeats[1].Active)
    }
    if sd.Times[0].Repeats[1].Offsets[0] != 0 {
        t.Errorf("Wrong Repeat Offset: %s", sd.Times[0].Repeats[1].Offsets[0])
    }
    if sd.Times[0].Repeats[1].Offsets[1] != (time.Hour * 25) {
        t.Errorf("Wrong Repeat Offset: %s", sd.Times[0].Repeats[1].Offsets[1])
    }
    if sd.Times[0].Zones[0].Time.Unix() != 2882844526-ntpUnix {
        t.Errorf("Wrong Zone Time: %s", sd.Times[0].Zones[0].Time.Unix())
    }
    if sd.Times[0].Zones[0].Offset != (time.Hour * -1) {
        t.Errorf("Wrong Zone Offset: %s", sd.Times[0].Zones[0].Offset)
    }
    if sd.Times[0].Zones[1].Time.Unix() != 2898848070-ntpUnix {
        t.Errorf("Wrong Zone Time: %s", sd.Times[0].Zones[1].Time.Unix())
    }
    if sd.Times[0].Zones[1].Offset != 0 {
        t.Errorf("Wrong Zone Offset: %s", sd.Times[0].Zones[1].Offset)
    }
    if _,ok := sd.Attributes["recvonly"]; !ok {
        t.Errorf("Attribute recvonly not found")
    }
    if sd.MediaDescriptions[0].Type != "audio" {
        t.Errorf("Wrong Media Description Type: %s", sd.MediaDescriptions[0].Type)
    }
    if sd.MediaDescriptions[0].Port != 49170 {
        t.Errorf("Wrong Media Description Port: %s", sd.MediaDescriptions[0].Port)
    }
    if sd.MediaDescriptions[0].NumPorts != 0 {
        t.Errorf("Wrong Media Connection NumPorts: %s", sd.MediaDescriptions[0].NumPorts)
    }
    if sd.MediaDescriptions[0].Proto != "RTP/AVP" {
        t.Errorf("Wrong Media Description Protocol: %s", sd.MediaDescriptions[0].Proto)
    }
    if sd.MediaDescriptions[0].Fmt != "0" {
        t.Errorf("Wrong Connection Address: %s", sd.MediaDescriptions[1].Fmt)
    }
    if sd.MediaDescriptions[1].Type != "video" {
        t.Errorf("Wrong Media Description Type: %s", sd.MediaDescriptions[1].Type)
    }
    if sd.MediaDescriptions[1].Port != 51372 {
        t.Errorf("Wrong Media Description Port: %s", sd.MediaDescriptions[1].Port)
    }
    if sd.MediaDescriptions[1].NumPorts != 0 {
        t.Errorf("Wrong Media Connection NumPorts: %s", sd.MediaDescriptions[1].NumPorts)
    }
    if sd.MediaDescriptions[1].Proto != "RTP/AVP" {
        t.Errorf("Wrong Media Description Protocol: %s", sd.MediaDescriptions[1].Proto)
    }
    if sd.MediaDescriptions[1].Fmt != "99" {
        t.Errorf("Wrong Connection Address: %s", sd.MediaDescriptions[1].Fmt)
    }
    if sd.MediaDescriptions[1].Attributes["rtpmap"] != "99 h263-1998/90000" {
        t.Errorf("Wrong Media Description Attribute: %s", sd.MediaDescriptions[1].Attributes["rtpmap"])
    }
}

var s2 = 
`v=0
o=testy 99201111 123 IN IP4 192.168.1.14
s=SessionName~
i=A test session
u=http://example.com
e=example@web.com (Testy)
p=123-123-3321 (Testy)
c=IN IP4 131.134.44.12
b=CT:128
k=base64:lol
m=video 49170/2 RTP/AVP 31
`

var sd = SessionDescription{
    Version: 0,
    Origin: Origin{"testy","99201111","123","IN","IP4","192.168.1.14"},
    SessionName: "SessionName~",
    Info: "A test session",
    Uri: "http://example.com",
    Emails: []Email{Email{"example@web.com","Testy"}},
    Phones: []Phone{Phone{"123-123-3321","Testy"}},
    Connection: Connection{"IN","IP4","131.134.44.12"},
    Bandwidths: []Bandwidth{Bandwidth{"CT","128"}},
    Key: Key{"base64","lol"},
    Attributes: make(map[string]string),
    MediaDescriptions: []MediaDescription{MediaDescription{"video",49170,2,"RTP/AVP","31","",nil,nil,Key{},make(map[string]string)}},
}

func TestEncode(t *testing.T) {
    sdp, err := sd.Encode()
    if err != nil {
        t.Error(err)
    }
    if sdp != s2 {
        t.Errorf("wrong SDP:\n%s",sdp)
    }
}

func BenchmarkDecode(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _,_ = Decode(s1)
    }
}

func BenchmarkEncode(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _,_ = sd.Encode()
    }
}
