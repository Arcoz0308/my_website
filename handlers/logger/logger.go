package logger

import (
	"fmt"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/jwalton/gchalk"
	"io"
	"os"
	"time"
)

var (
	output      io.Writer = os.Stdout
	cyan                  = gchalk.Ansi256(14) // if possible, same cyan in each terminal
	fatalLevel            = gchalk.BgAnsi256(88)
	fatalText             = gchalk.Ansi256(124)
	errorLevel            = gchalk.Ansi256(160)
	errorText             = gchalk.Ansi256(196)
	warnLevel             = gchalk.Ansi256(220)
	warnText              = gchalk.Ansi256(220)
	infoLevel             = gchalk.Ansi256(40)
	noticeLevel           = gchalk.Ansi256(83)
	debugLevel            = gchalk.Ansi256(251)
	debugText             = gchalk.Ansi256(251)
	app                   = gchalk.Ansi256(21)
)

func Info(a ...any) (n int, err error) {
	return AppInfo("main", a...)
}
func Infof(format string, a ...any) (n int, err error) {
	return AppInfof("main", format, a...)
}
func Infoln(a ...any) (n int, err error) {
	return AppInfoln("main", a...)
}
func AppInfo(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), infoLevel("[INFO]"), fmt.Sprint(a...))
}
func AppInfof(appName string, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), infoLevel("[INFO]"), fmt.Sprintf(format, a...))
}
func AppInfoln(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), infoLevel("[INFO]"), fmt.Sprintln(a...))
}

func Notice(a ...any) (n int, err error) {
	return AppNotice("main", a...)
}
func Noticef(format string, a ...any) (n int, err error) {
	return AppNoticef("main", format, a...)
}
func Noticeln(a ...any) (n int, err error) {
	return AppNoticeln("main", a...)
}
func AppNotice(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), noticeLevel("[NOTICE]"), fmt.Sprint(a...))
}
func AppNoticef(appName string, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), noticeLevel("[NOTICE]"), fmt.Sprintf(format, a...))
}
func AppNoticeln(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), noticeLevel("[NOTICE]"), fmt.Sprintln(a...))
}

func Debug(a ...any) (n int, err error) {
	return AppDebug("main", a...)
}
func Debugf(format string, a ...any) (n int, err error) {
	return AppDebugf("main", format, a...)
}
func Debugln(a ...any) (n int, err error) {
	return AppDebugln("main", a...)
}
func AppDebug(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), debugLevel("[DEBUG]"), debugText(fmt.Sprint(a...)))
}
func AppDebugf(appName string, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), debugLevel("[DEBUG]"), debugText(fmt.Sprintf(format, a...)))
}
func AppDebugln(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), debugLevel("[DEBUG]"), debugText(fmt.Sprintln(a...)))
}

func Warn(a ...any) (n int, err error) {
	return AppWarn("main", a...)
}
func Warnf(format string, a ...any) (n int, err error) {
	return AppWarnf("main", format, a...)
}
func Warnln(a ...any) (n int, err error) {
	return AppWarnln("main", a...)
}
func AppWarn(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), warnLevel("[WARN]"), warnText(fmt.Sprint(a...)))
}
func AppWarnf(appName string, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), warnLevel("[WARN]"), warnText(fmt.Sprintf(format, a...)))
}
func AppWarnln(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), warnLevel("[WARN]"), warnText(fmt.Sprintln(a...)))
}

func Error(a ...any) (n int, err error) {
	return AppError("main", a...)
}
func Errorf(format string, a ...any) (n int, err error) {
	return AppErrorf("main", format, a...)
}
func Errorln(a ...any) (n int, err error) {
	return AppErrorln("main", a...)
}
func AppError(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), errorLevel("[ERROR]"), errorText(fmt.Sprint(a...)))
}
func AppErrorf(appName string, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), errorLevel("[ERROR]"), errorText(fmt.Sprintf(format, a...)))
}
func AppErrorln(appName string, a ...any) (n int, err error) {
	return fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), errorLevel("[ERROR]"), errorText(fmt.Sprintln(a...)))
}

func Fatal(exit bool, a ...any) {
	AppFatal(exit, "main", a...)
}
func Fatalf(exit bool, format string, a ...any) {
	AppFatalf(exit, "main", format, a...)
}
func Fatalln(exit bool, a ...any) {
	AppFatalln(exit, "main", a...)
}
func AppFatal(exit bool, appName string, a ...any) {
	fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), fatalLevel("[FATAL]"), fatalText(fmt.Sprint(a...)))
	if exit {
		utils.CloseServ()
	}
}
func AppFatalf(exit bool, appName string, format string, a ...any) {
	fmt.Fprintf(output, "%s %s %s %s\n", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), fatalLevel("[FATAL]"), fatalText(fmt.Sprintf(format, a...)))
	if exit {
		utils.CloseServ()
	}
}
func AppFatalln(exit bool, appName string, a ...any) {
	fmt.Fprintf(output, "%s %s %s %s", cyan(time.Now().Format("[2006-01-02 15:04:05]")), app("["+appName+"]"), fatalLevel("[FATAL]"), fatalText(fmt.Sprintln(a...)))
	if exit {
		utils.CloseServ()
	}
}
