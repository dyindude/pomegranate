package main

import (
    "os"
    "os/user"
    "os/exec"
    "bufio"
    "encoding/json"
    "fmt"
    "time"
    "strings"
    "io/ioutil"
)

func Pom(topic string, minutes time.Duration) {
    ticker := time.NewTicker(time.Millisecond * 1000)
    start := time.Now()
    go func() {
        for range ticker.C {
            fmt.Printf("\r%s %s/%s", topic, time.Since(start).String(), (time.Minute * minutes).String())
        }
    }()
    time.Sleep(time.Minute * minutes)
    ticker.Stop()
    exec.Command("say", fmt.Sprintf("Your %s has expired.", topic)).Output()
}

type Topic struct {
    Duration    time.Duration
    Cycles      int64
}

func ConfigFile() string {
    user,_ := user.Current()
    configfile := user.HomeDir + "/.pomegranate.json"
    return configfile
}

func InitConfig() map[string]*Topic {
    var jsonString = "{\"Break\":{\"Duration\":5,\"Cycles\":0},\"Focus\":{\"Duration\":25,\"Cycles\":0}}"
    var m = make(map[string]*Topic)
    json.Unmarshal([]byte(jsonString),&m)
    WriteConfig(m)
    return m
}
func WriteConfig(m map[string]*Topic) {
    configfile := ConfigFile()
    jsonConf, _ := json.MarshalIndent(m, "", "    ")
    e := ioutil.WriteFile(configfile, []byte(jsonConf), 0644)
    if e != nil {
        panic(e)
    }
}
func LoadConfig() map[string]*Topic {
    configfile := ConfigFile()
    file, e := ioutil.ReadFile(configfile)
    if e != nil {
        return InitConfig()
    }
    var m = make(map[string]*Topic)
    json.Unmarshal(file, &m)
    return m
}

func main() {
    Topics := LoadConfig()

    for {
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
        jsonString, _ := json.Marshal(Topics)
        var TopicsCopy = make(map[string]*Topic)
        json.Unmarshal(jsonString,&TopicsCopy)
        for key, value := range Topics {
            fmt.Printf("%-8s%-8s%-8d%-8s\n", key,value.Duration * time.Minute,value.Cycles,time.Duration(value.Duration * time.Duration(value.Cycles) * time.Minute).String())
        }
        
        reader := bufio.NewReader(os.Stdin)
        time.Sleep(time.Second * 1)
        fmt.Printf("\r%-25s", " ")
        fmt.Printf("\rEnter text: ")
        key, _ := reader.ReadString('\n')
        skey := strings.TrimRight(key,"\n")
        if _, ok := Topics[skey]; ok {
                Pom(skey,Topics[skey].Duration)
                Topics[skey].Cycles += 1
                WriteConfig(Topics)
        } else {
            fmt.Printf("\rInvalid entry. Enter text: ")
        }
    }
}
