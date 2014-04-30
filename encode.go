package sdp

import (
    "fmt"
    "strconv"
    )

func (sd *SessionDescription) Encode() (string, error) {
    str := ""
    // Version
    str += "v=" + strconv.FormatInt(int64(sd.Version),10) + "\n"
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
    if sd.Connection.String() != new(Connection).String() {
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
    if sd.Key.String() != new(Key).String() {
        str += sd.Key.String() + "\n"
    }
    // Addributes
    for field, value := range sd.Attributes {
        str += "a=" + field + ":" + value + "\n"
    }
    // Media Descriptions
    for _, md := range sd.MediaDescriptions {
        str += md.String() + "\n"
    }
    return str, nil
}

func (o *Origin) String() string {
    return fmt.Sprintf("%s %s %s %s %s %s", o.Username, o.SessionId, o.SessionVersion, o.NetType, o.AddrType, o.UnicastAddr)
}

func (e *Email) String() string {
    email := "e=" + e.Address
    if e.Name != "" {
        email += " (" + e.Name + ")"
    }
    return email
}

func (e *Phone) String() string {
    phone := "p=" + e.Address
    if e.Name != "" {
        phone += " (" + e.Name + ")"
    }
    return phone
}

func (c *Connection) String() string {
    return "c=" + c.NetType + " " + c.AddrType + " " + c.Address
}

func (b *Bandwidth) String() string {
    return "b=" + b.Type + ":" + b.Bandwidth
}

func (t *TimeDescription) String() string {
    s := "t=" + strconv.FormatInt(t.Start.Unix(), 10)
    s += " " + strconv.FormatInt(t.Stop.Unix(), 10)
    for _,r := range t.Repeats {
        s += "\n" +  r.String()
    }
    if t.Zones != nil {
        s += "\nz="
        for _,z := range t.Zones {
            s += z.String() + " "
        }
    }
    return s
}

func (r *Repeat) String() string {
    s := "r=" + strconv.FormatInt(int64(r.Interval.Seconds()), 10)
    s += " " + strconv.FormatInt(int64(r.Active.Seconds()), 10)
    for _, o := range r.Offsets {
        s += " " + strconv.FormatInt(int64(o.Seconds()), 10)
    }
    return s
}

func (z *Zone) String() string {
    return strconv.FormatInt(z.Time.Unix() + ntpUnix, 10) + " " + strconv.FormatInt(int64(z.Offset.Seconds()), 10)
}

func (k *Key) String() string {
    if k.Key != "" {
        return "k=" + k.Method + ":" + k.Key
    }
    return "k=" + k.Method
}

func (m *MediaDescription) String() string {
    s := ""
    if m.NumPorts > 0 {
        s = "m=" + m.Type + " " + strconv.FormatInt(int64(m.Port), 10) + "/" + strconv.FormatInt(int64(m.NumPorts), 10) + " " + m.Proto + " " + m.Fmt
    } else {
        s = "m=" + m.Type + " " + strconv.FormatInt(int64(m.Port), 10) + " " + m.Proto + " " + m.Fmt
    }
    if m.Info != "" {
        s += "\ni=" + m.Info
    }
    for _, c := range m.Connections {
        s += "\n" + c.String()
    }
    for _, b := range m.Bandwidths {
        s += "\n" + b.String()
    }
    if m.Key.String() != new(Key).String() {
        s += "\n" + m.Key.String()
    }
    for field, value := range m.Attributes {
        s += "\na=" + field + ":" + value
    }
    return s
}
