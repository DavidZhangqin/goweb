package util

import (
	log "github.com/cihub/seelog"
)

var default_template string = `
<seelog type="sync" >
    <outputs>
        <console formatid="console" />
        <filter levels="error,critical" formatid="fmterror">
            <file path="/home/work/logs/applogs/goweb-errors.log"/>
        </filter>
    </outputs>
    <formats>
        <format id="fmterror" format="[%Date(2006/01/02 15:04:05 MST)] [%Level] [%FuncShort @ %File.%Line] %Msg%n"/>
        <format id="console" format="[%Date(2006/01/02 15:04:05 MST)] [%EscM(91)%Level%EscM(0)] [%File.%Line] %EscM(95)%Msg%n%EscM(0)"/>
    </formats>
</seelog>
`

func NewLogs() {
	logger, err := log.LoggerFromConfigAsBytes([]byte(default_template))
	if err != nil {
		log.Criticalf("log config err: %v", err)
	}
	log.ReplaceLogger(logger)
}
