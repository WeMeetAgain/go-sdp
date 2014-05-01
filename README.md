#GO-SDP - Session Description Protocol

This is a simple SDP encoding/decoding library witten in go, as defined in RFC 4566.

##Get the package

        go get github.com/WeMeetAgain/go-sdp

##Example

        import "github.com/WeMeetAgain/go-sdp"
        ...
        
        ...
        sdpString :=
        `v=0
        o=jdoe 2890844526 2890842807 IN IP4 10.47.16.5
        s=SDP Seminar
        i=A Seminar on the session description protocol
        u=http://www.example.com/seminars/sdp.pdf
        e=j.doe@example.com (Jane Doe)
        c=IN IP4 224.2.17.12/127
        t=2873397496 2873404696
        a=recvonly
        m=audio 49170 RTP/AVP 0
        m=video 51372 RTP/AVP 99
        a=rtpmap:99 h263-1998/90000
        `
        // decode string
        sessionDescription, err := sdp.Decode(sdpString)
        // access struct info
        uname := sessionDescription.Origin.Username
        email := sessionDescription.Emails[0].Address
        // change struct info
        sessionDescription.Origin.Username = "jd2014xoxo"
        // encode new SDP string
        str, err := sessionDescription.Encode()
