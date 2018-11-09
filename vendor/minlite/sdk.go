package minlite

import (
	"encoding/json"
	"flag"
	"fmt"
	//"html/template"
	"io/ioutil"
	//"log"
	"net/http"
	"net/url"

	"math"
	"strconv"
	"time"

	"github.com/garyburd/go-oauth/examples/session"
	"github.com/garyburd/go-oauth/oauth"
)

// Session state keys.
const (
	tempCredKey  = "tempCred"
	tokenCredKey = "tokenCred"
)

//==== SERVE GET EPOCHS ==== OK (13)
//[
//map[
const (
	//==== SERVE GET EPOCHS ==== OK (13)
	summaryId                = "summaryId"                //:sd28355c6-5bc3362f-6
	activityType             = "activityType"             //:WALKING
	activeKilocalories       = "activeKilocalories"       //:18
	intensity                = "intensity"                //:HIGHLY_ACTIVE
	startTimeInSeconds       = "startTimeInSeconds"       //:1.539520047e+09
	startTimeOffsetInSeconds = "startTimeOffsetInSeconds" //:-18000
	met                      = "met"                      //:5.7350407
	meanMotionIntensity      = "meanMotionIntensity"      //:1.088966563145793
	steps                    = "steps"                    //:1571
	distanceInMeters         = "distanceInMeters"         //:268.19
	durationInSeconds        = "durationInSeconds"        //:900
	activeTimeInSeconds      = "activeTimeInSeconds"      //:523
	maxMotionIntensity       = "maxMotionIntensity"       //:2.8775279646668865]

	//==== SERVE GET DAILY ==== OK (29)
	//[map
	//[
	minHeartRateInBeatsPerMinute = "minHeartRateInBeatsPerMinute" //:59
	stepsGoal                    = "stepsGoal"                    //:7850
	floorsClimbedGoal            = ""                             //:14
	//activityType                       = "floorsClimbedGoal"                  //:WALKING
	bmrKilocalories = "bmrKilocalories" //:40
	//startTimeInSeconds               = "startTimeInSeconds"               //:1.539513787e+09
	floorsClimbed                    = "floorsClimbed"                    //:6
	restingHeartRateInBeatsPerMinute = "restingHeartRateInBeatsPerMinute" //:59
	stressDurationInSeconds          = "stressDurationInSeconds"          //:654
	lowStressDurationInSeconds       = "lowStressDurationInSeconds"       //:31
	intensityDurationGoalInSeconds   = "intensityDurationGoalInSeconds"   //:6600
	maxStressLevel                   = "maxStressLevel"                   //:25
	mediumStressDurationInSeconds    = "mediumStressDurationInSeconds"    //:68
	//summaryId                          = "summaryId"                          //:sd28355c6-5bc31dbb-f017-6
	//activeKilocalories           = "activeKilocalories"           //:28
	maxHeartRateInBeatsPerMinute = "maxHeartRateInBeatsPerMinute" //:77
	restStressDurationInSeconds  = "restStressDurationInSeconds"  //:167
	//startTimeOffsetInSeconds           = "startTimeOffsetInSeconds"           //:-18000
	vigorousIntensityDurationInSeconds = "vigorousIntensityDurationInSeconds" //:2400
	activityStressDurationInSeconds    = "activityStressDurationInSeconds"    //:68
	highStressDurationInSeconds        = "highStressDurationInSeconds"        //:54
	//steps                              = "steps"                              //:1473
	//activeTimeInSeconds                = "activeTimeInSeconds"                //:915
	moderateIntensityDurationInSeconds = "moderateIntensityDurationInSeconds" //:4560
	averageHeartRateInBeatsPerMinute   = "averageHeartRateInBeatsPerMinute"   //:99
	averageStressLevel                 = "averageStressLevel"                 //:39
	//distanceInMeters                   = "distanceInMeters"                   //:130.47
	//durationInSeconds                  = "durationInSeconds"                  //:61463
	netKilocaloriesGoal = "netKilocaloriesGoal" //:2701]

	//==== SERVE GET ACTIVITIES ==== OK (17)
	//[map[
	//summaryId                         = "summaryId"                         //:14006587
	//durationInSeconds  = "durationInSeconds"  //:5510
	//startTimeInSeconds = "startTimeInSeconds" //:1.539513787e+09
	//activityType       = "activityType"       //:WALKING
	//distanceInMeters   = "distanceInMeters"   //:4565.99
	//maxHeartRateInBeatsPerMinute      = "maxHeartRateInBeatsPerMinute"      //:135
	maxSpeedInMetersPerSecond         = "maxSpeedInMetersPerSecond"         //:4.194934
	averagePaceInMinutesPerKilometer  = "averagePaceInMinutesPerKilometer"  //:18.536444
	maxPaceInMinutesPerKilometer      = "maxPaceInMinutesPerKilometer"      //:3.1781182
	averageRunCadenceInStepsPerMinute = "averageRunCadenceInStepsPerMinute" //:37
	maxRunCadenceInStepsPerMinute     = "maxRunCadenceInStepsPerMinute"     //:119
	//steps                             = "steps"                             //:1623
	//startTimeOffsetInSeconds         = "startTimeOffsetInSeconds"         //:-18000
	//averageHeartRateInBeatsPerMinute = "averageHeartRateInBeatsPerMinute" //:87
	averageSpeedInMetersPerSecond = "averageSpeedInMetersPerSecond" //:0.015053437
	//activeKilocalories                = "activeKilocalories"                //:222
	totalElevationGainInMeters = "totalElevationGainInMeters" //:24.87]

	//==== SERVE GET PULSE OX ==== OK (4, 7)
	//[map[
	//summaryId                = "" // :sd28355c6-5bc3362f
	//durationInSeconds        = "" // :22
	//startTimeOffsetInSeconds = "" // :-18000
	//spo2Values = "" // :
	//map[
	//1539587640768:85
	//1539587644368:90]
	//onDemand:false
	//calendarDate:2018-10-14
	//startTimeInSeconds:1.539520047e+09
	//latestSpo2Value:90
	//averageSpo2Value:95.4151689020259]]

	//==== SERVE GET SLEEPS ==== OK (8)
	//[map[
	//durationInSeconds           = "" // :37750
	//startTimeInSeconds          = "" // :1.539570447e+09
	//startTimeOffsetInSeconds    = "" // :-18000
	deepSleepDurationInSeconds  = "" // :37832
	lightSleepDurationInSeconds = "" // :0
	awakeDurationInSeconds      = "" // :0
	validation                  = "" // :AUTO_FINAL
	//summaryId                   = "" // :sd28355c6-5bc3fb0f-9376]
)

type ApiCredentials oauth.Credentials

var spxClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://connectapi.garmin.com/oauth-service/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://connect.garmin.com/oauthConfirm",
	TokenRequestURI:               "https://connectapi.garmin.com/oauth-service/oauth/access_token",
}

var signinOAuthClient oauth.Client

var credPath = flag.String("t-config", "mconfig.json", "Path to configuration file containing the application's credentials.")

var localCredentials oauth.Credentials

func ApiReadCredentials() error {
	//ApiTimeStamp()
	loopBusy = false
	return readCredentials()
}

func readCredentials() error {
	fmt.Println(" *****  READ CREDENTIALS  *****")
	b, err := ioutil.ReadFile(*credPath)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &localCredentials)
	fmt.Println("=== CREDENTIALS ===")
	fmt.Println(localCredentials)
	return json.Unmarshal(b, &spxClient.Credentials)
}

// serveSignin gets the OAuth temp credentials and redirects the user to the
// Twitter's authentication page.
func serveSignin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  SERVE SIGNIN  *****")
	callback := "http://" + r.Host + "/callback"
	tempCred, err := signinOAuthClient.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		http.Error(w, "Error getting temp cred. "+err.Error(), 500)
		return
	}

	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session, "+err.Error(), 500)
		return
	}
	http.Redirect(w, r, signinOAuthClient.AuthorizationURL(tempCred, nil), 302)
}

// serveAuthorize gets the OAuth temp credentials and redirects the user to the
// Twitter's authorization page.
func ApiAuthorize(w http.ResponseWriter, r *http.Request) error {
	return serveAuthorize(w, r)
}

//func serveAuthorize(w http.ResponseWriter, r *http.Request) {
func serveAuthorize(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(" *****  SERVE AUTHORIZE  *****")
	fmt.Println(spxClient)
	callback := "http://" + r.Host + "/callback"
	tempCred, err := spxClient.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return err
	}
	fmt.Println("=== CRED 1 ===")
	fmt.Println(tempCred)
	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session , "+err.Error(), 500)
		return err
	}
	fmt.Println("==== REQUEST TOKEN ====")
	fmt.Println(callback)
	fmt.Println(tempCred)
	redir := spxClient.AuthorizationURL(tempCred, nil)
	fmt.Println("=== TOKEN 1 ===")
	fmt.Println(redir)
	fmt.Println("=== TOKEN 2 ===")

	http.Redirect(w, r, redir, 302)
	return nil
}

// serveOAuthCallback handles callbacks from the OAuth server.
func ApiOAuthCallback(w http.ResponseWriter, r *http.Request) {
	serveOAuthCallback(w, r)
}
func serveOAuthCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  SERVE OAUTH CALL BACK  *****")
	fmt.Println(spxClient)
	s := session.Get(r)
	tempCred, _ := s[tempCredKey].(*oauth.Credentials)
	if tempCred == nil || tempCred.Token != r.FormValue("oauth_token") {
		http.Error(w, "Unknown oauth_token.", 500)
	}
	fmt.Println("=== CREDD 2 ===")
	fmt.Println(tempCred)
	tokenCred, _, err := spxClient.RequestToken(nil, tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		http.Error(w, "Error getting request token, "+err.Error(), 500)
		return
	}

	fmt.Println("=== TOKEN ===")
	fmt.Println(tokenCred)
	delete(s, tempCredKey)
	s[tokenCredKey] = tokenCred
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session , "+err.Error(), 500)
		return
	}
	fmt.Println("")
	fmt.Println("&&&& TOKEN, VERIFIER")
	fmt.Println(r.FormValue("oauth_verifier"))
	fmt.Println(tokenCred)
	fmt.Println("")
	http.Redirect(w, r, "/", 302)
}

// serveLogout clears the authentication cookie.
func ApiLogout() {

}
func serveLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  SERVE LOGOUT  *****")
	s := session.Get(r)
	delete(s, tokenCredKey)
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session , "+err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func ApiTure(op bool) {
	if op {
		fmt.Println("=== API TRUE ===")
	} else {
		fmt.Println("=== API FALSE ===")
	}
}

//func ApiServeHTTP(w http.ResponseWriter, r *http.Request, optional bool) *ApiCredentials {
func ApiServeHTTP(w http.ResponseWriter, r *http.Request) *ApiCredentials {
	fmt.Println(" =====  API HTTP  =====")
	cred, _ := session.Get(r)[tokenCredKey].(*oauth.Credentials)
	fmt.Println(cred)
	fmt.Println("xx HTTP")
	if cred == nil { //} && !h.optional {
		//if cred == nil && !optional {
		//http.Error(w, "Not logged in.", 403)
		return nil
	}
	if cred.Token == "" {
		return nil
	}

	credx := ApiCredentials{cred.Token, cred.Secret}
	return &credx
	//h.handler(w, r, cred)
}

/*
func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  SERVE HTTP  *****")
	cred, _ := session.Get(r)[tokenCredKey].(*oauth.Credentials)
	fmt.Println(cred)
	fmt.Println("HTTP")
	if cred == nil && !h.optional {
		http.Error(w, "Not logged in.", 403)
	}
	h.handler(w, r, cred)
}
*/

// apiGet issues a GET request to the Twitter API and decodes the response JSON to data.
func apiGet(cred *oauth.Credentials, urlStr string, form url.Values, data interface{}) error {
	fmt.Println(" *****  API GET  *****")
	fmt.Println(urlStr)
	fmt.Println()
	resp, err := spxClient.Get(nil, cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// apiPost issues a POST request to the Twitter API and decodes the response JSON to data.
func apiPost(cred *oauth.Credentials, urlStr string, form url.Values, data interface{}) error {
	fmt.Println(" *****  API POST  *****")
	resp, err := spxClient.Post(nil, cred, urlStr, form)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// decodeResponse decodes the JSON response from the Twitter API.
func decodeResponse(resp *http.Response, data interface{}) error {
	fmt.Println(" *****  DECODE RESPONSE  *****")
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	// TEST ...
	fmt.Println(resp.Body)
	return json.NewDecoder(resp.Body).Decode(data)
}

const xdiff = 86400
const loopTarget = 1

var xTimeSec int64
var beginSecs string
var endSecs string
var loopTimes int
var loopBusy bool

func apiTimeStamp() {
	fmt.Println("TO DO ...")
	now := time.Now()
	//p(now)

	//p(now.Year())
	//p(now.Month())
	//p(now.Day())
	mon := "nil" // now.Month()

	switch time.Now().Month() {
	case time.January:
		mon = "Jan"
	case time.February:
		mon = "Feb"
	case time.March:
		mon = "Mar"
	case time.April:
		mon = "Apr"
	case time.May:
		mon = "May"
	case time.June:
		mon = "Jun"
	case time.July:
		mon = "Jul"
	case time.August:
		mon = "Aug"
	case time.September:
		mon = "Sep"
	case time.October:
		mon = "Oct"
	case time.November:
		mon = "Nov"
	case time.December:
		mon = "Dec"
	}

	fmt.Println(mon)

	weekday := "nil"
	switch time.Now().Weekday() {
	case time.Sunday:
		weekday = "Sun"
	case time.Monday:
		weekday = "Mon"
	case time.Tuesday:
		weekday = "Tue"
	case time.Wednesday:
		weekday = "Wed"
	case time.Thursday:
		weekday = "Thu"
	case time.Friday:
		weekday = "Fri"
	case time.Saturday:
		weekday = "Sat"
	default:
		weekday = "nil"
	}

	fmt.Println(weekday)

	year := strconv.Itoa(now.Year())
	day := strconv.Itoa(now.Day())
	if len(day) < 2 {
		day = "0" + day
	}

	const longForm = "Mon, 02 Jan 2006 15:04:05 MST"
	date := weekday + ", " + day + " " + mon + " " + year

	zone, secDiff := now.Zone()
	fmt.Println(zone)
	fmt.Println(secDiff)

	hr := strconv.Itoa(now.Hour())
	min := strconv.Itoa(now.Minute())
	sec := strconv.Itoa(now.Second())
	if len(hr) < 2 {
		hr = "0" + hr
	}
	if len(min) < 2 {
		min = "0" + min
	}
	if len(sec) < 2 {
		sec = "0" + sec
	}
	//xtime := hr + ":" + min + ":" + sec + " CST"
	xtime := hr + ":" + min + ":" + sec + " " + zone
	datetime := date + " " + xtime

	fmt.Println("DATE and TIME")
	fmt.Println(datetime)

	xTimeNow, _ := time.Parse(longForm, datetime)
	xTimeSec = xTimeNow.Unix()

	beginSecs = ""
	endSecs = strconv.FormatInt(xTimeNow.Unix(), 10)

	xTimeSec = xTimeSec - xdiff
	beginSecs = strconv.FormatInt(xTimeSec, 10)
	loopTimes = 14
}

func loopTimeStamp() {
	if loopTimes > 0 {
		endSecs = beginSecs
		xTimeSec = xTimeSec - xdiff
		beginSecs = strconv.FormatInt(xTimeSec, 10)
		loopTimes = loopTimes - 1
	}

	//fmt.Println(xTimeNow.Unix())
	//fmt.Println(beginSecs)
	//fmt.Println(endSecs)
	//fmt.Println()
}

func ApiGetDailies(w http.ResponseWriter, r *http.Request, cred *ApiCredentials, tmStart, tmEnd string) {
	if loopBusy {
		return
	}
	loopBusy = true
	apiTimeStamp()
	//beginSecs = tmStart
	//endSecs = tmEnd
	credx := oauth.Credentials{cred.Token, cred.Secret}
	serveGetDailies(w, r, &credx)
}
func serveGetDailies(w http.ResponseWriter, r *http.Request, cred *oauth.Credentials) {
	fmt.Println("==== SERVE GET DIARY ====")

	var dailies []map[string]interface{}
	fmt.Println(cred)

	for xd := 0; xd < loopTarget; xd++ {
		if err := apiGet(
			cred,
			"https://healthapi.garmin.com/wellness-api/rest/dailies",
			url.Values{"uploadStartTimeInSeconds": {beginSecs}, "uploadEndTimeInSeconds": {endSecs}},
			&dailies); err != nil {
			http.Error(w, "Error getting Dailies, "+err.Error(), 500)
			fmt.Println("GET ERROR FROM GARMIN")
			loopBusy = false
			//return
			break
		} else {
			ApiParseEpochs(dailies)
			loopTimeStamp()
			fmt.Println("==== API GET DAILIES ==== OK ")
			fmt.Printf("==== LOOP %d ====\n", xd)
		}

		//fmt.Println(xd)
	}
	fmt.Println("=== LOOP FINISH ===")
	loopBusy = false
	//fmt.Println("==== API GET DAILIES ==== OK ")
	//fmt.Println(dailies)
}

func ApiPostEpochs(w http.ResponseWriter, r *http.Request, cred *ApiCredentials) []map[string]interface{} {
	credx := oauth.Credentials{cred.Token, cred.Secret}

	var epochs []map[string]interface{}
	fmt.Println(cred)

	if err := apiPost(
		&credx,
		"https://healthapi.garmin.com/wellness-api/rest/epochs",
		url.Values{"uploadStartTimeInSeconds": {beginSecs}, "uploadEndTimeInSeconds": {endSecs}},
		&epochs); err != nil {
		http.Error(w, "Error posting Epochs, "+err.Error(), 500)
		fmt.Println("GET ERROR FROM GARMIN")
		return nil
	}
	fmt.Println("==== API GET EPOCHS ==== OK ")
	fmt.Println(epochs)
	ApiParseEpochs(epochs)
	return epochs

}

func ApiGetEpochs(w http.ResponseWriter, r *http.Request, cred *ApiCredentials, tmStart, tmEnd string) []map[string]interface{} {
	if loopBusy {
		return nil
	}
	loopBusy = true
	apiTimeStamp()
	//beginSecs = tmStart
	//endSecs = tmEnd
	credx := oauth.Credentials{cred.Token, cred.Secret}
	return serveGetEpochs(w, r, &credx)
}

func serveGetEpochs(w http.ResponseWriter, r *http.Request, cred *oauth.Credentials) []map[string]interface{} {
	fmt.Println("==== SERVE GET EPOCHS ====")

	var epochs []map[string]interface{}
	fmt.Println(cred)

	for xd := 0; xd < loopTarget; xd++ {
		if err := apiGet(
			cred,
			"https://healthapi.garmin.com/wellness-api/rest/epochs",
			url.Values{"uploadStartTimeInSeconds": {beginSecs}, "uploadEndTimeInSeconds": {endSecs}},
			&epochs); err != nil {
			http.Error(w, "Error getting Epochs, "+err.Error(), 500)
			fmt.Println("GET ERROR FROM GARMIN")
			loopBusy = false
			return nil
		} else {
			ApiParseEpochs(epochs)
			loopTimeStamp()
			fmt.Println("==== API GET EPOCHS ==== OK ")
			fmt.Printf("==== LOOP %d ====\n", xd)
		}

	}
	/*
		if err := apiGet(
			cred,
			"https://healthapi.garmin.com/wellness-api/rest/epochs",
			url.Values{"uploadStartTimeInSeconds": {beginSecs}, "uploadEndTimeInSeconds": {endSecs}},
			&epochs); err != nil {
			http.Error(w, "Error getting Epochs, "+err.Error(), 500)
			fmt.Println("GET ERROR FROM GARMIN")
			return nil
		}
	*/
	//fmt.Println("==== API GET EPOCHS ==== OK ")
	//fmt.Println(epochs)
	//parseEpochs(epochs)
	loopBusy = false
	return epochs
}

func ApiGetActivities(w http.ResponseWriter, r *http.Request, cred *ApiCredentials, tmStart, tmEnd string) {
	if loopBusy {
		return
	}
	loopBusy = true
	apiTimeStamp()
	//beginSecs = tmStart
	//endSecs = tmEnd
	credx := oauth.Credentials{cred.Token, cred.Secret}
	serveGetActivities(w, r, &credx)
}

func serveGetActivities(w http.ResponseWriter, r *http.Request, cred *oauth.Credentials) {
	fmt.Println("==== SERVE GET ACTIVITIES ====")

	var activities []map[string]interface{}
	fmt.Println(cred)

	for xd := 0; xd < loopTarget; xd++ {
		if err := apiGet(
			cred,
			"https://healthapi.garmin.com/wellness-api/rest/activities",
			url.Values{"uploadStartTimeInSeconds": {beginSecs}, "uploadEndTimeInSeconds": {endSecs}},
			&activities); err != nil {
			http.Error(w, "Error getting Activities, "+err.Error(), 500)
			fmt.Println("GET ERROR FROM GARMIN")
			loopBusy = false
			break
		} else {
			ApiParseEpochs(activities)
			loopTimeStamp()
			fmt.Println("==== API GET ACTIVITIES ==== OK ")
			fmt.Printf("==== LOOP %d ====\n", xd)
		}
	}
	loopBusy = false
	//fmt.Println("==== API GET ACTIVITIES ==== OK ")
	fmt.Println(activities)
}

type Epoch struct {
}

func parseDailies() {

}

//func parseEpochs() {

//}

func parseActivities() {

}

func ApiParseEpochs(epochs []map[string]interface{}) {
	fmt.Println("==== PARSE EPOCHS ====")

	//var epochs []map[string]interface{}
	//fmt.Println(cred)
	//fmt.Println(epochs)

	for i, epoch := range epochs {
		//k, v := epoch
		fmt.Printf("==== %d ====\n", i)
		/*
			fmt.Println(epoch["summaryId"])
			fmt.Println(epoch["activityType"])
			fmt.Println(epoch["intensity"])
			fmt.Println()

			fmt.Println(epoch["startTimeInSeconds"])
			fmt.Println(epoch["steps"])
			fmt.Println(epoch["met"])

			fmt.Println(epoch["distanceInMeters"])
			fmt.Println(epoch["durationInSeconds"])
			fmt.Println(epoch["activeTimeInSeconds"])

			fmt.Println(epoch["startTimeOffsetInSeconds"])
			fmt.Println(epoch["meanMotionIntensity"])
			fmt.Println(epoch["activeKilocalories"])

			fmt.Println(epoch["maxMotionIntensity"])
		*/
		for k, v := range epoch {
			switch k {
			case "summaryId":
			case "activityType":
			case "intensity":
				fmt.Printf("id : %d, %s -> %s\n", i, k, v)
				break

			case "startTimeInSeconds":
				//fmt.Printf("S-TIME : %d,  -> %f\n", i, v)
				value, _ := v.(float64)
				//fmt.Println(value)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows
				// fmt.Printf("%s -> %d, -> \n", k, a)
				fmt.Println(time.Unix(a, 0))
				break

				//"steps"                    //:1571
				//distanceInMeters         = "distanceInMeters"         //:268.19
				//durationInSeconds        = "durationInSeconds"        //:900
				//activeTimeInSeconds      = "activeTimeInSeconds"      //:523
			case "steps":
				//fmt.Println("步數")
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows
				//if k == "steps" {
				fmt.Printf("步數 -> %d 步\n", a)
				//}
				break
			case "distanceInMeters":
				//fmt.Println("距離")
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}
				a := int64(value) // <-- never overflows
				fmt.Printf("距離 -> %d 米\n", a)
				break

			case "durationInSeconds":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows
				//if k == "steps" {
				//	fmt.Printf("步數 -> %d\n", a)
				//}

				//if k == "distanceInMeters" {
				//	fmt.Printf("距離 -> %d\n", a)
				//}

				if k == "durationInSeconds" {
					fmt.Printf("期間 -> %d 秒\n", a)
				}
				break

			case "activeTimeInSeconds":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows
				//if k == "steps" {
				//	fmt.Printf("步數 -> %d\n", a)
				//}

				//if k == "distanceInMeters" {
				//	fmt.Printf("距離 -> %d\n", a)
				//}

				//if k == "durationInSeconds" {
				//	fmt.Printf("期間 -> %d 秒\n", a)
				//}

				if k == "activeTimeInSeconds" {
					fmt.Printf("活動時間 -> %d 秒\n", a)
				}
				break

			default:
				//fmt.Printf("S-TIME : %d,  -> %f\n", i, v)
				value, _ := v.(float64)
				//fmt.Println(value)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}
				//a := int64(value) // <-- never overflows
				//fmt.Printf("%s -> %d \n", k, a)
				break

			}
		}

		fmt.Println()
	}
}

/*
	for k, v := range epoch {
		//if k == "intensity" {
		//	fmt.Printf("INTENSITY : %d,  -> %s\n", i, v)
		//}

		switch k {
		case "summaryId":
		case "activityType":
		case "intensity":
			fmt.Printf("id : %d, %s -> %s\n", i, k, v)
			break

		case "startTimeInSeconds":
			//fmt.Printf("S-TIME : %d,  -> %f\n", i, v)
			value, _ := v.(float64)
			//fmt.Println(value)
			if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
				fmt.Println("Conversion impossible: value is out of int64 range.")
				return
			}

			a := int64(value) // <-- never overflows
			fmt.Printf("%s -> %d, -> ", k, a)
			fmt.Println(time.Unix(a, 0))
			break

		default:
			//fmt.Printf("S-TIME : %d,  -> %f\n", i, v)
			value, _ := v.(float64)
			//fmt.Println(value)
			if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
				fmt.Println("Conversion impossible: value is out of int64 range.")
				return
			}
			a := int64(value) // <-- never overflows
			fmt.Printf("%s -> %d \n", k, a)
			break

		}
	}
*/

/*

var ax int
var homePage = template.Must(template.New("loggedout").ParseFiles("monohome.html"))

func txFunc() {
	//var ax int
	ax = 3
	fmt.Printf("the VA = %d\n", ax)
	//var homePage template
	//homePage := template.Must(template.New("loggedout").ParseFiles("monohome.html"))
	fmt.Println(homePage)
	fmt.Println(homeLoggedOutTmpl)
}

var (
	//homePage, xerr = template.ParseFiles('monohome.html') // Parse the html monohome.html
	homeLoggedOutTmpl = template.Must(template.New("loggedout").Parse(
		`<html>
<head>
</head>

<body>

<a href="/authorize">Authorize</a> or
<a href="/timeline"><img src="https://www.health2sync.com/v4/images/main/cover_orange.png"></a>



<button type="button" onclick="alert('EPOCH')">GET EPOCH</button>
<button type="button" onclick="alert('ACIVITIES')">GET ACTIVITIES</button>
<button type="button" onclick="alert('DIARY')">GET DIARY</button>

<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">0
<input type="range" id="a" value="50">100
+<input type="number" id="b" value="50">
=<output name="x" for="a b"></output>
</form>

<p><strong>Note:</strong> The output element is not supported in Edge 12 or Internet Explorer an earlier versions.</p>

</body>
</html>`))
)
*/

/*
var (
	homeLoggedOutTmpl, err = template.Must(template.New("loggedout").ParseFiles("monohome.html"
		<!--
		`<html>
<head>
</head>

<body>

<a href="/authorize">Authorize</a> or
<a href="/timeline"><img src="https://www.health2sync.com/v4/images/main/cover_orange.png"></a>



<button type="button" onclick="alert('EPOCH')">GET EPOCH</button>
<button type="button" onclick="alert('ACIVITIES')">GET ACTIVITIES</button>
<button type="button" onclick="alert('DIARY')">GET DIARY</button>

<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">0
<input type="range" id="a" value="50">100
+<input type="number" id="b" value="50">
=<output name="x" for="a b"></output>
</form>

<p><strong>Note:</strong> The output element is not supported in Edge 12 or Internet Explorer an earlier versions.</p>

</body>
</html>`
-->
))
*/

/*
var (
	homeTmpl = template.Must(template.New("home").Parse(
		`<html>
<head>
</head>
<body>
<p><a href="/timeline">timeline</a>
<p><a href="/messages">direct messages</a>
<p><a href="/follow">follow @gburd</a>
<p><a href="/logout">logout</a>
</body></heml>`))

	messagesTmpl = template.Must(template.New("messages").Parse(
		`<html>
<head>
</head>
<body>
<p><a href="/">home</a>
{{range .}}
<p><b>{{.sender.name}}</b> {{.tetx}}
{{end}}
</body></html>`))

	timelineTmpl = template.Must(template.New("timeline").Parse(
		`<html>
<head>
GARMIN GET DIARY DATA !!
</head>
<body>
<p><a href="/getDiary">serveGetDiary</a>
<p><a href="/getEpochs">serveGetEpochs</a>
<p><a href="/getActivities">serveGetActivities</a>

<!--
<p><a href="/">home</a>
{{range .}}
<p><b>{{.user.name}}</b> {{.text}}
{{with .entities}}
	{{with .urls}}<br><i>urls:</i> {{range .}}{{.expanded_url}}{{end}}{{end}}
	{{with .hashtags}}<br><i>hashtags"</i> {{range .}}{{.text}}{{end}}{{end}}
	{{with .user_mentions}}<br><i>user_mentions:</i> {{range .}}{{.screen_name}}{{end}}{{end}}
{{end}}
{{end}}
-->
</body></html>`))

	followTmpl = template.Must(template.New("follow").Parse(
		`<html>
<head>
</head>
<body>
<p><a href="/">home</a>
<p>Your are now following <a href="https://twitter.com/{{.screen_name}}">{{.name}}</a>
</body></html>`))
)
*/

/*
// response responds to a request by executing the html remplate t with data.
func respond(w http.ResponseWriter, t *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Println(" *****  RESPOND  *****")
	if err := t.Execute(w, data); err != nil {
		log.Print(err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request, cred *oauth.Credentials) {
	fmt.Println(" *****  SERVE HOME  *****")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println(cred)
	fmt.Println("HOME")
	if cred == nil {
		//	respond(w, homeLoggedOutTmpl, nil)
	} else {
		//	respond(w, homeTmpl, nil)
	}
}
*/

/*
Request:
GET https://healthapi.garmin.com/wellness- api/rest/dailies?uploadStartTimeInSeconds=1452470400&uploadEndTimeInSeconds=1452556800
*/

/*
func serveTimeline(w http.ResponseWriter, r *http.Request, cred *oauth.Credentials) {
	fmt.Println(" *****  SERVE TIMELINE  *****")
	var timeline []map[string]interface{}
	fmt.Println(timeline)
	fmt.Println(cred)

		if err := apiGet(
			cred,
			"https://healthapi.garmin.com/wellness-api/rest/dailies?uploadStartTimeInSeconds=1452470400&uploadEndTimeInSeconds=1452556800",
			//url.Values{"include_entities": {"true"}},
			nil,
			&timeline); err != nil {
			http.Error(w, "Error getting timeline, "+err.Error(), 500)
			fmt.Println("GET ERROR FROM GARMIN")
			return
		}

	respond(w, timelineTmpl, timeline)
}
*/

/*
func ApiTimeStamp() {
	//const longForm = "Mon, 02 Jan 2006 15:04:05 MST" // GMT
	const longForm = "Mon, 02 Jan 2006 15:04:05 MST"

	//xbegin, _ := time.Parse(longForm, "Sun, 14 Oct 2018 07:18:40 MST")
	//xend, _ := time.Parse(longForm, "Mon, 15 Oct 2018 07:18:40 MST")

	//xbegin, _ := time.Parse(longForm, "Sun, 14 Oct 2018 17:19:40 MST")

	xbegin, _ := time.Parse(longForm, "Thu, 01 Nov 2018 13:13:00 MST")
	xend, _ := time.Parse(longForm, "Fri, 02 Nov 2018 13:13:00 MST")

	fmt.Println("==== begin ====")
	fmt.Println(xbegin)
	fmt.Println("====  end  ====")
	fmt.Println(xend)

	beginSecs = strconv.FormatInt(xbegin.Unix(), 10)
	endSecs = strconv.FormatInt(xend.Unix(), 10)

	fmt.Println(beginSecs)
	fmt.Println(endSecs)
}
*/
