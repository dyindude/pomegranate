# pomegranate
a pomodoro timer written in go, that keeps track of total cycle count, time spent, etc.

This is my first golang project, please be kind :)

example output (I used this to keep myself on track while packing up my apartment to move out):

```
pack    25m0s   16      6h40m0s 
break   5m0s    14      1h10m0s 
lbreak  30m0s   5       2h30m0s 
Enter text:              
```

Configuration is stored in ~/.pomegranate.json. A default config will be written if it does not yet exist.

```json
{
    "Break": {
        "Duration": 5, 
        "Cycles": 0
    },   
    "Focus": {
        "Duration": 25,
        "Cycles": 0
    }    
}
```
Enter the name of a topic (currently case-sensitive) to start a cycle. The cycle count is written to the config every time a cycle completes.

Currently, the notification when a cycle completes is issuing an OSX "say" command to use the TTS engine for the announcement. I've also used the following for using mplayer to play a sound on linux:

```go
exec.Command("mplayer", "/path/to/filename.ogg").Output()
```
