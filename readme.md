defined in RFC <a>6350</a>.


# 1. Install vcard-wa
### 1.1- Create folder of project
### 1.2- Create mod
    $ go mod init <yourModuleNAme>
    $ go mod tidy
### 1.3- Add dependencies to your project module
    $ go get -u github.com/ecsavigne/vcard_wa
### 1.4- Test code
    package main

    import (
        "bytes"
        "fmt"
        v1 "github.com/ecsavigne/vcard_wa"
    )

    func main() {
        text := "BEGIN:VCARD\nVERSION:3.0\nN:;Contact name;;;\nFN:ContactFirstName\nX-WA-BIZ-NAME:My name\nX-WA-BIZ-DESCRIPTION:One description\nTEL;type=CELL;type=VOICE;waid=########:+## ## #####-####\nEND:VCARD"

        m := v1.ParseVCard(bytes.NewReader([]byte(text)))
        fmt.Println(m.GetPartRecord())
        fmt.Println("Tags:")
        fmt.Println(m.GetTags())
        fmt.Println("GetTagValue:")
        for _, k := range m.GetTags() {
            fmt.Printf("key: '%s' = value: '%s\n", k, m.GetTagValue(k))
        }

        fmt.Println("GetTagContentValues:")
        fmt.Printf("key: '%s' = value: '%s\n", "X-WA-BIZ-NAME", m.GetTagContentValues("X-WA-BIZ-NAME"))
        fmt.Printf("key: '%s' = value: '%s\n", "TEL", m.GetTagContentValues("TEL"))
    }   