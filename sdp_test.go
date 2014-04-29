package sdp

import (
    "testing"
    )

var s1 = 
`v=0
o=jdoe 2890844526 2890842807 IN IP4 10.47.16.5
s=SDP Seminar
i=A Seminar on the session description protocol
u=http://www.example.com/seminars/sdp.pdf
e=j.doe@example.com (Jane Doe)
p=Jane Doe<123-456-7890>
c=IN IP4 224.2.17.12/127
t=2873397496 2873404696
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
