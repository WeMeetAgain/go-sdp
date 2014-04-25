package sdp

import (
    "fmt"
    )

func (sd *SessionDescription) Encode() (string, error) {
    str := ""
    // Version
    str += "v=" + string(sd.Version) + "\n"
    // Origin
    str += "o=" + sd.Origin.String()  + "\n"
    // Session Name
    str += "s=" + sd.SessionName  + "\n"
    // Info
    if sd.Info != "" {
        str += "i=" + sd.Info  + "\n"
    }
    // URI
    if sd.Uri != "" {
        str += "u=" + sd.Uri  + "\n"
    }
    // Emails
    for _, email := range sd.Emails {
        str += email.String() + "\n"
    }
    // Phone Numbers
    for _, phone := range sd.Phones {
        str += phone.String() + "\n"
    }
    // Connection
    if sd.Connection != nil {
        str += sd.Connection.String() + "\n"
    }
    // Bandwidths
    for _, bandwidth := range sd.Bandwidths {
        str += bandwidth.String() + "\n"
    }
    // Times
    for _, time := range sd.Times {
        str += time.String() + "\n"
    }
    // Key
    if sd.Key != nil {
        str += sd.Key.String() + "\n"
    }
    // Addributes
    for field, value := range sd.Attributes {
        str += field + ":" + attr + "\n"
    }
    // Media Descriptions
    for _, md := range sd.MediaDescription {
        str += md.String() + "\n"
    }
    return str, nil
}

func (o *Origin) String() string {
    return fmt.Sprintf("%s %s %s %s %s %s", o.Username, o.SessionId, o.SessionVersion, o.NetType, o.AddrType. o.UnicastAddr)
}

func (e *Email) String() string {
    email := e.Address
    if e.Name != "" {
        email += " (" + e.Name + ")"
    }
    return email
}

func (e *Phone) String() string {
    phone := e.Address
    if e.Name != "" {
        phone += " (" + e.Name + ")"
    }
    return phone
}