package util

import (
	"fmt"
	"github.com/millenc/golatch"
	"github.com/millenc/latch-cmd/session"
	"net/http"
	"strings"
	t "time"
)

//Replaces all characters with '*' if NoShadow is false
func FormatSecret(secret string, NoShadow bool) string {
	if NoShadow {
		return secret
	} else {
		return strings.Repeat("*", len(secret))
	}
}

//Formats usage statistics
func FormatUsageString(inUse int, limit int) string {
	return fmt.Sprintf("in use: %d, limit: %d", inUse, limit)
}

//Parses a date received from the command line
func ParseCmdDate(date string) (parsed t.Time, err error) {
	if date == "" {
		return t.Time{}, nil
	}

	parsed, err = t.Parse("02-01-2006 15:04:05", date)

	return parsed, err
}

//Formats the client versions received from the API
func FormatClientVersions(clientVersions []golatch.LatchClientVersion) (formatted string) {
	versions := []string{}
	for _, version := range clientVersions {
		versions = append(versions, version.Platform+" - "+version.App)
	}

	return strings.Join(versions, ",")
}

//Show request (used when verbose is true)
func ShowRequestInfoFn(Session *session.LatchCmdSession, NoShadow bool) func(request *golatch.LatchRequest) {
	return func(request *golatch.LatchRequest) {
		Session.AddInfo("latch:\t")
		Session.AddInfo("app/user\t" + request.AppID)
		Session.AddInfo("secret\t" + FormatSecret(request.SecretKey, NoShadow) + "\n\t\t")
		Session.AddInfo("request:\t")
		Session.AddInfo("url\t" + request.URL.String())
		Session.AddInfo("http-method\t" + request.HttpMethod)
		Session.AddInfo("date\t" + request.GetFormattedDate())
		Session.AddInfo("params\t" + request.GetSerializedParams())
		Session.AddInfo("headers\t" + request.GetSerializedHeaders())
		Session.AddInfo("signature\t" + strings.Replace(request.GetRequestSignature(), "\n", "\n\t\t", -1))
		Session.AddInfo("signature-sha1\t" + request.GetSignedRequestSignature())
		Session.AddInfo("auth-header\t" + request.GetAuthorizationHeader() + "\n\t\t")
	}
}

//Show response information (used when verbose is true)
func ShowResponseInfoFn(Session *session.LatchCmdSession) func(request *golatch.LatchRequest, response *http.Response, responseBody string) {
	return func(request *golatch.LatchRequest, response *http.Response, responseBody string) {
		Session.AddInfo("response:\t")
		Session.AddInfo(fmt.Sprintf("http-status\t%d", response.StatusCode))
		Session.AddInfo("body\t" + responseBody + "\n\t\t")
	}
}
