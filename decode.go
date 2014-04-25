package sdp

import (
    "bufio"
    "errors"
    "strings"
    )

const (
    badChar string = "bad character found"
    badGrammar string = "bad grammar found"
    )

func Decode(str string) (SessionDescription, error) {
    state := 0
    sd := SessionDescription{}
    scanner := bufio.NewScanner(strings.NewReader(str))
    for scanner.Scan() {
        line := scanner.Text()
        if line[1] != '=' {
            return nil, errors.New(badChar)
        switch state {
        case 0: // Version
            if line[0] == 'v' {
                sd.Version = line[2:]
                state++
            } else {
                return nil, errors.New(badChar)
            }
        case 1: // Origin
            if line[0] == 'o' {
                if o, err := parseOrigin(line[2:]); err == nil {
                    sd.Origin = o
                } else {
                    return nil, err
                }
                state++
            } else {
                return nil, errors.New(badChar)
            }
        case 2: // Session Name
            if line[0] == 's' {
                sd.SessionName = line[2:]
                state++
            } else {
                return nil, errors.New(badChar)
            }
        case 3: // Information
            if line[0] == 'i' {
                sd.Info = line[2:]
                state++
            } else {
                state++
                fallthrough
            }
        case 4: // URI
            if line[0] == 'u' {
                sd.Uri = line[2:]
                state++
            } else {
                state++
                fallthrough
            }
        case 5: // Email
            if line[0] == 'e' {
                if e, err := parseEmail(line[2:]); err == nil {
                    sd.Emails = append(sd.Emails, e)
                } else {
                    return nil, err
                }
                state++
            } else {
                state++
                fallthrough
            }
        case 6: // Phone
            if line[0] == 'p' {
                if p, err := parsePhone(line[2:]); err == nil {
                    sd.Phones = append(sd.Phones, p)
                } else {
                    return nil, err
                }
                state++
            } else {
                state++
                fallthrough
            }
        case 7: // Connection
            if line[0] == 'c' {
                if c, err := parseConnection(line[2:]); err == nil {
                    sd.Connection = c
                } else {
                    return nil, err
                }
                state++
            } else {
                state++
                fallthrough
            }
        case 8: // Time
            if line[0] == 't' {
                if t, err := parseTime(line[2:]); err == nil {
                    sd.Times = append(sd.Times, t)
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 9: // Repeat Times
            if line[0] == 'r' {
                if r, err := parseRepeat(line[2:]); err == nil {
                    sd.Times[(len(sd.Times)-1)] = r
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 10: // Zone Adjustments
            if line[0] == 'z' {
                if z, err := parseZone(line[2:]); err == nil {
                    sd.Times[(len(sd.Times)-1)] = z
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 11: // Key
            if line[0] == 'k' {
                if k, err := parseZone(line[2:]); err == nil {
                    sd.Key = k
                } else {
                    return nil, err
                }
                state++
            } else {
                state++
                fallthrough
            }
        case 12: // Attribute
            if line[0] == 'a' {
                if a, err := parseAttribute(line[2:]); err == nil {
                    sd.Attribute = a
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 13: // Media Description
            if line[0] == 'm' {
                if m, err := parseMedia(line[2:]); err == nil {
                    sd.MediaDescriptions = append(sd.MediaDescriptions, m)
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 14: // Information
            if line[0] == 'i' {
                if i, err := parseInformation(line[2:]; err == nil {
                    sd.MediaDescriptions[(len(sd.MediaDescriptions)-1)].Info = i)
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 15: // Connection
            if line[0] == 'c' {
                if c, err := parseConnection(line[2:]); err == nil {
                    sd.MediaDescriptions[(len(sd.MediaDescriptions)-1)].Connection = c
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 16: // Bandwidth
            if line[0] == 'b' {
                if b, err := parseBandwidth(line[2:]); err == nil {
                    sd.MediaDescriptions[(len(sd.MediaDescriptions)-1)].Bandwidths = append(sd.MediaDescriptions[(len(sd.Times)-1)].Bandwidths, b)
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        case 17: // Attribute
            if line[0] == 'a' {
                if a, b, err := parseAttribute(line[2:]); err == nil {
                    sd.MediaDescriptions[(len(sd.MediaDescriptions)-1)].Attributes[a] = b
                } else {
                    return nil, err
                }
            } else {
                state++
                fallthrough
            }
        }
    }
}

func contains(s []string, e string) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func parseOrigin(s string) (Origin, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 6 {
        return nil, errors.New(badGrammar)
    }
    return Origin{
        tokens[0]
        tokens[1]
        tokens[2]
        tokens[3]
        tokens[4]
        tokens[5]
    }, nil
}

func parseEmail(s string) (Email, error) {
    tokens := strings.Split(s," ")
    if len(tokens) == 1 {
        return Email{
            Address: tokens[0]
        }, nil
    } else {
        if bracket1 := strings.Index(s,"<"), bracket2 := strings.Index(s,">"); bracket1 != -1 && bracket2 != -1 {
            if (bracket1-2) > -1 {
                return Email{
                    Address: s[(bracket1+1):(bracket2-1)]
                    Name: s[0:(bracket1-2)]
                }, nil
            } else {
                return nil, errors.New(badGrammar)
            }
        } else if bracket1 != -1 || bracket2 != -1 {
            return nil, errors.New(badGrammar)
        }
        if paren1 := strings.Index(s,"("), paren2 := strings.Index(s,")"); paren1 != -1 && paren2 != -1 {
            if (paren1-2) > -1 {
                return Email{
                    Address: s[0:(paren1-2)]
                    Name: s[(paren1+1):(paren2-1)]
                }, nil
            } else {
                return nil, errors.New(badGrammar)
            }
        } else if paren1 != -1 || paren2 != -1 {
            return nil, errors.New(badGrammar)
        }
        return nil, errors.New(badGrammar)
    }
}

func parsePhone(s string) (Phone, error) {
    if bracket1 := strings.Index(s,"<"), bracket2 := strings.Index(s,">"); bracket1 != -1 && bracket2 != -1 {
        if (bracket1-2) > -1 {
            return Phone{
                Address: s[(bracket1+1):(bracket2-1)]
                Name: s[0:(bracket1-2)]
            }, nil
        } else {
            return nil, errors.New(badGrammar)
        }
    } else if bracket1 != -1 || bracket2 != -1 {
        return nil, errors.New(badGrammar)
    }
    if paren1 := strings.Index(s,"("), paren2 := strings.Index(s,")"); paren1 != -1 && paren2 != -1 {
        if (paren1-2) > -1 {
            return Phone{
                Address: s[0:(paren1-2)]
                Name: s[(paren1+1):(paren2-1)]
            }, nil
        } else {
            return nil, errors.New(badGrammar)
        }
    } else if paren1 != -1 || paren2 != -1 {
        return nil, errors.New(badGrammar)
    }
    return Phone {
        Address: tokens[0]
    }, nil
}

func parseConnection(s string) (Connection, error) {
    tokens := strings.Split(s," ")
    if len(tokens) != 3 {
        return nil, errors.New(badGrammar)
    }
    return Connection{
        tokens[0]
        tokens[1]
        tokens[2]
    }, nil
}

func parseTime(s string) (TimeField, error) {
    
}

func parseRepeat(s string) (Repeat, error) {
    
}

func parseZone(s string) (Zone, error) {
    
}

func parseKey(s string) (Key, error) {
    tokens := strings.Split(s,":")
    if !contains(KeyTypes, tokens[0]) {
        return nil, errors.New(badGrammar)
    }
    if len(tokens) == 1 {
        return tKey{tokens[0], ""}, nil
    } else if len(tokens) == 2 {
        return Key{tokens[0], tokens[1]}, nil
    } else
        return nil, errors.New(badGrammar)
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
    } else
        return "", "", errors.New(badGrammar)
    }
}

func parseMedia(s string) (MediaDescription, error) {
    tokens := strings.Split(s," ")
    m := MediaDescription{}
    if len(tokens) != 4 {
        return nil, errors.New(badGrammar)
    }
    if !contains(MediaTypes, tokens[0]) {
        return nil, errors.New(badGrammar)
    }
    if p := strings.Split(tokens[1],"/"); len(p) > 2 || len(p) < 1 {
        return nil, errors.New(badGrammar)
    } else if len(p) == 2 {
        m.NumPorts = p[1]
        m.Port = p[0]
    } else {
        m.Port = p[0]
    }
    if !contains(ProtoTypes, tokens[2]) {
        return nil, errors.New(badGrammar)
    }
    m.Type = tokens[0]
    m.Proto = tokens[2]
    m.Fmt = tokens[3]
    return m, nil
}
}
