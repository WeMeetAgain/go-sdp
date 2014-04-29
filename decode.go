package sdp

import (
    "bufio"
    "errors"
    "strings"
    "strconv"
    "time"
    )

const (
    noLine string = "no line"
    badChar string = "bad character found"
    badGrammar string = "bad grammar found"
    )

const (
    ntpUnix int64 = 2208988800
    )
    
type SDPParser struct {
    SD SessionDescription
    Rules []func(*SDPParser,string) error
    Index int
}


func Decode(str string) (SessionDescription, error) {
    sdp := NewSDPParser()
    return sdp.Decode(str)
}

var DefaultRules = []func(*SDPParser,string) error{versionLine, originLine, sessionNameLine, infoLine, uriLine, emailLine, phoneLine, connectionLine, timeLine, repeatLine, zoneLine, keyLine, attrLine, mediaLine, mediaInfoLine, mediaConnectionLine, mediaBandwidthLine, mediaAttrLine}

func NewSDPParser() *SDPParser {
    return &SDPParser{SessionDescription{Attributes: make(map[string]string)},DefaultRules,0}
}

func (p *SDPParser) Next(s string) error {
    err := p.Rules[p.Index](p,s)
    if err != nil && err != errors.New(noLine) {
        return err
    }
    return nil
}

func (p *SDPParser) Decode(str string) (SessionDescription, error) {
    scanner := bufio.NewScanner(strings.NewReader(str))
    for scanner.Scan() {
        line := scanner.Text()
        if err := p.Next(line); err != nil {
            return p.SD, err
        }
    }
    return p.SD, nil
}

func versionLine(p *SDPParser,line string) error {
    if line[0] == 'v' {
        v, err := strconv.ParseInt(line[2:], 10, 32)
        if err != nil {
            return err
        }
        p.SD.Version = int(v)
        p.Index++
        return nil
    } else {
        return errors.New(badChar)
    }
}

func originLine(p *SDPParser,line string) error {
    if line[0] == 'o' {
        if o, err := parseOrigin(line[2:]); err == nil {
            p.SD.Origin = o
        } else {
            return err
        }
        p.Index++
        return nil
    } else {
        return errors.New(badChar)
    }
}

func sessionNameLine(p *SDPParser,line string) error {
    if line[0] == 's' {
        p.SD.SessionName = line[2:]
        p.Index++
        return nil
    } else {
        return errors.New(badChar)
    }
}

func infoLine(p *SDPParser,line string) error {
    if line[0] == 'i' {
        p.SD.Info = line[2:]
        p.Index++
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func uriLine(p *SDPParser,line string) error {
    if line[0] == 'u' {
        p.SD.Uri = line[2:]
        p.Index++
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func emailLine(p *SDPParser,line string) error {
    if line[0] == 'e' {
        if e, err := parseEmail(line[2:]); err == nil {
            p.SD.Emails = append(p.SD.Emails, e)
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func phoneLine(sdp *SDPParser,line string) error {
    if line[0] == 'p' {
        if p, err := parsePhone(line[2:]); err == nil {
            sdp.SD.Phones = append(sdp.SD.Phones, p)
        } else {
            return err
        }
    } else {
        sdp.Index++
        if err := sdp.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func connectionLine(p *SDPParser,line string) error {
    if line[0] == 'c' {
        if c, err := parseConnection(line[2:]); err == nil {
            p.SD.Connection = c
        } else {
            return err
        }
        p.Index++
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func timeLine(p *SDPParser,line string) error {
    if line[0] == 't' {
        if t, err := parseTime(line[2:]); err == nil {
            p.SD.Times = append(p.SD.Times, t)
            
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func repeatLine(p *SDPParser,line string) error {
    if line[0] == 'r' {
        if r, err := parseRepeat(line[2:]); err == nil {
            p.SD.Times[(len(p.SD.Times)-1)].Repeats = append(p.SD.Times[(len(p.SD.Times)-1)].Repeats,r)
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func zoneLine(p *SDPParser,line string) error {
    if line[0] == 'z' {
        if z, err := parseZones(line[2:]); err == nil {
            p.SD.Times[(len(p.SD.Times)-1)].Zones = z
            p.Index = p.Index-2
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func keyLine(p *SDPParser,line string) error {
    if line[0] == 'k' {
        if k, err := parseKey(line[2:]); err == nil {
            p.SD.Key = k
        } else {
            return err
        }
        p.Index++
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func attrLine(p *SDPParser,line string) error {
    if line[0] == 'a' {
        if a, b, err := parseAttribute(line[2:]); err == nil {
            p.SD.Attributes[a] = b
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func mediaLine(p *SDPParser,line string) error {
    if line[0] == 'm' {
        if m, err := parseMedia(line[2:]); err == nil {
            p.SD.MediaDescriptions = append(p.SD.MediaDescriptions, m)
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func mediaInfoLine(p *SDPParser,line string) error {
    if line[0] == 'i' {
        if i, err := parseInformation(line[2:]); err == nil {
            p.SD.MediaDescriptions[len(p.SD.MediaDescriptions)-1].Info = i
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func mediaConnectionLine(p *SDPParser,line string) error {
    if line[0] == 'c' {
        if c, err := parseConnection(line[2:]); err == nil {
            p.SD.MediaDescriptions[(len(p.SD.MediaDescriptions)-1)].Connections = append(p.SD.MediaDescriptions[(len(p.SD.MediaDescriptions)-1)].Connections,c)
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func mediaBandwidthLine(p *SDPParser,line string) error {
    if line[0] == 'b' {
        if b, err := parseBandwidth(line[2:]); err == nil {
            p.SD.MediaDescriptions[(len(p.SD.MediaDescriptions)-1)].Bandwidths = append(p.SD.MediaDescriptions[(len(p.SD.Times)-1)].Bandwidths, b)
        } else {
            return err
        }
    } else {
        p.Index++
        if err := p.Next(line); err != nil {
            return err
        }
    }
    return nil
}

func mediaAttrLine(p *SDPParser,line string) error {
    if line[0] == 'a' {
        if a, b, err := parseAttribute(line[2:]); err == nil {
            p.SD.MediaDescriptions[(len(p.SD.MediaDescriptions)-1)].Attributes[a] = b
        } else {
            return err
        }
    } else {
        p.Index = p.Index - 4
    }
    return nil
}

func contains(s []string, e string) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func parseOrigin(s string) (Origin, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 6 {
        return Origin{}, errors.New(badGrammar)
    }
    return Origin{
        tokens[0],
        tokens[1],
        tokens[2],
        tokens[3],
        tokens[4],
        tokens[5],
    }, nil
}

func parseInformation(s string) (string, error) {
    return s, nil
}

func parseEmail(s string) (Email, error) {
    tokens := strings.Split(s," ")
    if len(tokens) == 1 {
        return Email{
            Address: tokens[0],
        }, nil
    } else {
        if bracket1, bracket2 := strings.Index(s,"<"), strings.Index(s,">"); bracket1 != -1 && bracket2 != -1 {
            if (bracket1-2) > -1 {
                return Email{
                    Address: s[(bracket1+1):(bracket2)],
                    Name: s[0:(bracket1)],
                }, nil
            } else {
                return Email{}, errors.New(badGrammar)
            }
        } else if bracket1 != -1 || bracket2 != -1 {
            return Email{}, errors.New(badGrammar)
        }
        if paren1, paren2 := strings.Index(s,"("), strings.Index(s,")"); paren1 != -1 && paren2 != -1 {
            if (paren1-2) > -1 {
                return Email{
                    Address: s[0:(paren1-1)],
                    Name: s[(paren1+1):(paren2)],
                }, nil
            } else {
                return Email{}, errors.New(badGrammar)
            }
        } else if paren1 != -1 || paren2 != -1 {
            return Email{}, errors.New(badGrammar)
        }
        return Email{}, errors.New(badGrammar)
    }
}

func parsePhone(s string) (Phone, error) {
    if bracket1, bracket2 := strings.Index(s,"<"), strings.Index(s,">"); bracket1 != -1 && bracket2 != -1 {
        if (bracket1-2) > -1 {
            return Phone{
                Address: s[(bracket1+1):(bracket2)],
                Name: s[0:(bracket1)],
            }, nil
        } else {
            return Phone{}, errors.New(badGrammar)
        }
    } else if bracket1 != -1 || bracket2 != -1 {
        return Phone{}, errors.New(badGrammar)
    }
    if paren1, paren2 := strings.Index(s,"("), strings.Index(s,")"); paren1 != -1 && paren2 != -1 {
        if (paren1-2) > -1 {
            return Phone{
                Address: s[0:(paren1-1)],
                Name: s[(paren1+1):(paren2)],
            }, nil
        } else {
            return Phone{}, errors.New(badGrammar)
        }
    } else if paren1 != -1 || paren2 != -1 {
        return Phone{}, errors.New(badGrammar)
    }
    return Phone {
        Address: s,
    }, nil
}

func parseBandwidth(s string) (Bandwidth, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 2 {
        return Bandwidth{}, errors.New(badGrammar)
    }
    return Bandwidth{
        tokens[0],
        tokens[1],
    }, nil
}

func parseConnection(s string) (Connection, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 3 {
        return Connection{}, errors.New(badGrammar)
    }
    return Connection{
        tokens[0],
        tokens[1],
        tokens[2],
    }, nil
}

func parseTime(s string) (TimeDescription, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 2 {
        return TimeDescription{}, errors.New(badGrammar)
    }
    start, err := strconv.ParseInt(tokens[0], 10, 64)
    if err != nil {
        return TimeDescription{}, err
    }
    stop, err := strconv.ParseInt(tokens[1], 10, 64)
    if err != nil {
        return TimeDescription{}, err
    }
    return TimeDescription{
        Start: time.Unix(start - ntpUnix, 0),
        Stop: time.Unix(stop - ntpUnix, 0),
    }, nil
}

func parseDuration(s string) (time.Duration, error) {
    var interval time.Duration
    if s[len(s)-1] == 'd' {
        days, err := strconv.ParseInt(s[0:len(s)-2], 10, 64)
        if err != nil {
            return time.Nanosecond, err
        }
        interval, err = time.ParseDuration(strconv.FormatInt(days * 3600, 10)+"h")
        if err != nil {
            return time.Nanosecond, err
        }
    } else if s[len(s)-1] == 'h'  ||
    s[len(s)-1] == 'm'  ||
    s[len(s)-1] == 's'  {
        interval, err := time.ParseDuration(s)
        if err != nil {
            return interval, err
        }
    } else {
        seconds, err := strconv.ParseInt(s, 10, 64)
        if err != nil {
            return time.Nanosecond, err
        }
        interval, err = time.ParseDuration(strconv.FormatInt(seconds, 10)+"s")
        if err != nil {
            return time.Nanosecond, err
        }
    }
    return interval, nil
}

func parseRepeat(s string) (Repeat, error) {
    tokens := strings.Split(s," ")
    if len(tokens) < 3 {
        return Repeat{}, errors.New(badGrammar)
    }
    interval, err := parseDuration(tokens[0])
    if err != nil {
        return Repeat{}, err
    }
    active, err := parseDuration(tokens[1])
    if err != nil {
        return Repeat{}, err
    }
    var offsets []time.Duration
    for i := 2; i < len(tokens); i++ {
        o, err := parseDuration(tokens[i])
        if err != nil {
            return Repeat{}, err
        }
        offsets = append(offsets, o)
    }
    return Repeat{
        Interval: interval,
        Active: active,
        Offsets: offsets,
    }, nil
}

func parseZones(s string) ([]Zone, error) {
    tokens := strings.Split(s," ")
    if len(tokens) % 2 != 0 {
        return nil, errors.New(badGrammar)
    }
    var zones []Zone
    for i:=0; i<len(tokens); i=i+2 {
        t, err := strconv.ParseInt(tokens[i], 10, 64)
        if err != nil {
            return nil, err
        }
        offset, err := parseDuration(tokens[i+1])
        if err != nil {
            return nil, err
        }
        z := Zone{
            Time: time.Unix(t - ntpUnix, 0),
            Offset: offset,
        }
        zones = append(zones, z)
    }
    return zones, nil
}

func parseKey(s string) (Key, error) {
    tokens := strings.Split(s,":")
    if !contains(KeyTypes, tokens[0]) {
        return Key{}, errors.New(badGrammar)
    }
    if len(tokens) == 1 {
        return Key{tokens[0], ""}, nil
    } else if len(tokens) == 2 {
        return Key{tokens[0], tokens[1]}, nil
    } else {
        return Key{}, errors.New(badGrammar)
    }
}

func parseAttribute(s string) (string, string, error) {
    tokens := strings.Split(s,":")
    if !contains(AttrTypes, tokens[0]) {
        return "", "", errors.New(badGrammar)
    }
    if len(tokens) == 1 {
        return tokens[0], "", nil
    } else if len(tokens) == 2 {
        return tokens[0], tokens[1], nil
    } else {
        return "", "", errors.New(badGrammar)
    }
}

func parseMedia(s string) (MediaDescription, error) {
    tokens := strings.Split(s," ")
    m := MediaDescription{Attributes: make(map[string]string)}
    if len(tokens) != 4 {
        return MediaDescription{}, errors.New(badGrammar)
    }
    if !contains(MediaTypes, tokens[0]) {
        return MediaDescription{}, errors.New(badGrammar)
    }
    if p := strings.Split(tokens[1],"/"); len(p) > 2 || len(p) < 1 {
        return MediaDescription{}, errors.New(badGrammar)
    } else if len(p) == 2 {
        np, err := strconv.ParseInt(p[1], 10, 32)
        if err != nil {
            return m, err
        }
        m.NumPorts = int(np)
        p, err := strconv.ParseInt(p[0], 10, 32)
        if err != nil {
            return m, err
        }
        m.NumPorts = int(p)
    } else {
        p, err := strconv.ParseInt(p[0], 10, 32)
        if err != nil {
            return m, err
        }
        m.Port = int(p)
    }
    if !contains(TransportTypes, tokens[2]) {
        return m, errors.New(badGrammar)
    }
    m.Type = tokens[0]
    m.Proto = tokens[2]
    m.Fmt = tokens[3]
    return m, nil
}
