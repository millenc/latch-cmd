package session

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"text/tabwriter"
)

//Stores messages produced during the command run
type LatchCmdSession struct {
	Messages []Message
	ExitCode int
}

//Message to be printed out
type Message struct {
	Type    string
	Content string
}

//Message types
const (
	MESSAGE_TYPE_INFO    = "info"
	MESSAGE_TYPE_WARNING = "warning"
	MESSAGE_TYPE_ERROR   = "error"
	MESSAGE_TYPE_SUCCESS = "success"
)

//Adds a message
func (session *LatchCmdSession) AddMessage(Type string, Content string) {
	message := Message{
		Type:    Type,
		Content: Content,
	}

	session.Messages = append(session.Messages, message)
}

//Adds an Info message
func (session *LatchCmdSession) AddInfo(Content string) {
	session.AddMessage(MESSAGE_TYPE_INFO, Content)
}

//Adds a Warning message
func (session *LatchCmdSession) AddWarning(Content string) {
	session.AddMessage(MESSAGE_TYPE_WARNING, Content)
}

//Adds an error message
func (session *LatchCmdSession) AddError(Content string) {
	session.AddMessage(MESSAGE_TYPE_ERROR, Content)
}

//Adds a Success message
func (session *LatchCmdSession) AddSuccess(Content string) {
	session.AddMessage(MESSAGE_TYPE_SUCCESS, Content)
}

//Outputs messages
func (session *LatchCmdSession) Output() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintln(w, "")
	for _, message := range session.Messages {
		fmt.Fprintln(w, session.FormatMessageType(message.Type)+"\t"+message.Content)
	}

	w.Flush()
}

//Prints all the messages stored in the session and halts the command execution with the provided error
func (session *LatchCmdSession) Halt(err error) {
	session.AddError(err.Error())
	session.ExitCode = -1
	session.End()
}

//Outputs all messages and exits with ExitCode
func (session *LatchCmdSession) End() {
	session.Output()
	os.Exit(session.ExitCode)
}

//Colorize message types
func (session *LatchCmdSession) FormatMessageType(messageType string) string {
	switch messageType {
	case MESSAGE_TYPE_INFO:
		return color.BlueString("info")
	case MESSAGE_TYPE_WARNING:
		return color.YellowString("warning")
	case MESSAGE_TYPE_ERROR:
		return color.RedString("error")
	case MESSAGE_TYPE_SUCCESS:
		return color.GreenString("success")
	}

	return messageType
}
