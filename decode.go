package sdp

import (
    "bufio"
    "errors"
    "strings"
    )

const (
    badChar = "bad character found"
    badGrammar = "bad grammar found"
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
                sd.Origin = parseOrigin(line[2:])
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
                sd.Emails = append(sd.Emails, parseEmail(line[2:]))
                state++
            } else {
                state++
                fallthrough
            }
        case 6: // Phone
            if line[0] == 'p' {
                sd.Phones = append(sd.Phones, parsePhone(line[2:]))
                state++
            } else {
                state++
                fallthrough
            }
        case 7: // Connection
            if line[0] == 'c' {
                sd.Connection = parseConnection(line[2:])
                state++
            } else {
                state++
                fallthrough
            }
        case 8: // Time
            if line[0] == 't' {
                sd.Times = append(sd.Times,parseTime(line[2:]))
            } else {
                state++
                fallthrough
            }
        case 9: // Repeat Times
            if line[0] == 'r' {
                sd.Times[(len(sd.Times)-1)] = parseRepeat(line[2:])
            } else {
                state++
                fallthrough
            }
        case 10: // Zone Adjustments
            if line[0] == 'z' {
                sd.Times[(len(sd.Times)-1)] = parseRepeat(line[2:])
            } else {
                state++
                fallthrough
            }
        case 11: // Key
            if line[0] == 'k' {
                sd.Key = parseKey(line[2:])
                state++
            } else {
                state++
                fallthrough
            }
        case 12: // Attribute
            if line[0] == 'k' {
                sd.Attribute = parseAttribute(line[2:])
            } else {
                state++
                fallthrough
            }
        case 13: // Media Description
            if line[0] == 'm' {
                sd.MediaDescriptions = append(sd.MediaDescriptions, parseAttribute(line[2:]))
            } else {
                state++
                fallthrough
            }
        case 14: // Information
            if line[0] == 'i' {
                sd.MediaDescriptions[(len(sd.Times)-1)].Info = parseInformation(line[2:])
            } else {
                state++
                fallthrough
            }
        case 15: // Connection
            if line[0] == 'c' {
                sd.MediaDescriptions[(len(sd.Times)-1)].Connection = parseConnection(line[2:])
            } else {
                state++
                fallthrough
            }
        case 16: // Bandwidth
            if line[0] == 'i' {
                sd.MediaDescriptions[(len(sd.Times)-1)].Bandwidths = append(sd.MediaDescriptions[(len(sd.Times)-1)].Bandwidths, parseBandwidth(line[2:]))
            } else {
                state++
                fallthrough
            }
        case 17: // Attribute
        }
    }
}

func parseOrigin(s string) (Origin, error) {
    tokens := strings.Split(s," ")
    if tokens) != 6 {
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
    if tokens) != 3 {
        return nil, errors.New(badGrammar)
    }
    return Connection{
        tokens[0]
        tokens[1]
        tokens[2]
    }, nil
}
