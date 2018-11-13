package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"html/template"
	//"io/ioutil"
	"log"
	//"net"
	"net/http"
	//"net/url"

	"math"

	//"strconv"
	//"github.com/gin-gonic/gin"
	"os"

	//"minlite"
	"github.com/markchuang56/minlite"
	//"go-heroku/test-personal/vendor/minlite"
	"time"
	//"github.com/markchuang56/minlite"
	//"spx"
	//"go-heroku/test-personal/minlite"
	//"github.com/garyburd/go-oauth/examples/session"
	//"github.com/garyburd/go-oauth/oauth"
)

var homeTmpl = template.Must(template.New("home").ParseFiles("templates/epochx.html"))
var homeLoggedOutTmpl = template.Must(template.New("loggedout").ParseFiles("templates/loggedout.html"))

// serveAuthorize gets the OAuth temp credentials and redirects the user to the
// Twitter's authorization page.
func userServeAuthorize(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  USER SERVE AUTHORIZE  *****")
	err := minlite.ApiAuthorize(w, r)

	if err != nil {
		fmt.Println(" ERROR HAPPEN AT AUTHORIZE ")
		return
	}
}

// serveOAuthCallback handles callbacks from the OAuth server.
func demoServeOAuthCallback(w http.ResponseWriter, r *http.Request) {
	minlite.ApiOAuthCallback(w, r)
}

// serveLogout clears the authentication cookie.
func ApiLogout() {

}

func userServeLogout(w http.ResponseWriter, r *http.Request) {

}

/*
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
*/

// authHandler reads the auth cookie and invokes a handler with the result.
type demoAuthHandler struct {
	handler  func(w http.ResponseWriter, r *http.Request, c *minlite.ApiCredentials)
	optional bool
}

func (h *demoAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  DEMO SERVE HTTP  *****")

	cred := minlite.ApiServeHTTP(w, r)

	if cred == nil && !h.optional {
		//if cred == nil && !optional {
		http.Error(w, "Not logged in.", 403)
	}
	fmt.Println(cred)
	h.handler(w, r, cred)
	//h.handler(w, r, nil)
}

// response responds to a request by executing the html remplate t with data.
func respond(w http.ResponseWriter, t *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Println(" *****  RESPOND 1 *****")
	if err := t.Execute(w, data); err != nil {
		fmt.Println("*** ERROR ***")
		fmt.Println(err)
		//log.Print(err)
	}
}

func demoServeHome(w http.ResponseWriter, r *http.Request, cred *minlite.ApiCredentials) {
	fmt.Println(" *****  DEMO SERVE HOME  *****")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println(cred)
	fmt.Println("HOME")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if cred == nil {
		if err := homeLoggedOutTmpl.ExecuteTemplate(w, "loggedout.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := homeTmpl.ExecuteTemplate(w, "epochx.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func demoServeGetDailies(w http.ResponseWriter, r *http.Request, cred *minlite.ApiCredentials) {
	fmt.Println("==== SERVE GET DAILIES ====")
	//tmStart, tmEnd := timeStampProcess(r)
	//spx.ApiGetDailies(w, r, cred, tmStart, tmEnd)
	minlite.ApiGetDailies(w, r, cred, " ", " ")
	fmt.Println("==== SERVE GET DAILIES ==== OK ")
}

func demoServeGetEpochs(w http.ResponseWriter, r *http.Request, cred *minlite.ApiCredentials) {
	fmt.Println("==== SERVE GET EPOCHS ====")
	//tmStart, tmEnd := timeStampProcess(r)
	//epochs := spx.ApiGetEpochs(w, r, cred, tmStart, tmEnd)
	epochs := minlite.ApiGetEpochs(w, r, cred, " ", " ")
	//var epochs []map[string]interface{}
	fmt.Println("==== SERVE GET EPOCHS ==== OK ")
	//fmt.Println(epochs)
	//spx.ApiParseEpochs(epochs)
	clientParseEpochs(w, epochs)
	fmt.Println("WHAT'S UP ??")
	//fmt.Fprintf(w, "HOW DO YOU TURN THIS ON!!")
}

func demoServeGetActivities(w http.ResponseWriter, r *http.Request, cred *minlite.ApiCredentials) {
	fmt.Println("==== SERVE GET ACTIVITIES ====")
	//tmStart, tmEnd := timeStampProcess(r)
	//spx.ApiGetActivities(w, r, cred, tmStart, tmEnd)
	minlite.ApiGetActivities(w, r, cred, " ", " ")
	fmt.Println("==== SERVE GET ACTIVITIES ==== OK ")
}

func demoServePingEpochs(w http.ResponseWriter, r *http.Request, cred *minlite.ApiCredentials) {
	fmt.Println("==== SERVE PING EPOCHS ====")
	epochs := minlite.ApiPostEpochs(w, r, cred)
	//var epochs []map[string]interface{}
	//fmt.Println("==== SERVE GET EPOCHS ==== OK ")
	fmt.Println(epochs)
	//fmt.Println("WHAT'S UP ??")
	fmt.Fprintf(w, "DEMO PING EPOCHS!!")
}

var httpAddr = flag.String("addr", ":8080", "HTTP server address")

func main() {
	fmt.Println("===== GARMIN GO =====")
	flag.Parse()

	//if err := readCredentials(); err != nil {
	if err := minlite.ApiReadCredentials(); err != nil {
		log.Fatalf("Error reading configuration, %v", err)
		fmt.Println("===  ERROR 0  ===")
	}

	// added
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
	}
	fmt.Println(port)

	//xrouter := gin.New()
	//fmt.Println(xrouter)

	//u, _ := url.Parse("https://www.examplecasserver.com")
	// Use a different auth URL for "Sign in with Twitter."
	//
	//signinOAuthClient = spxClient
	//signinOAuthClient.ResourceOwnerAuthorizationURI = "https://connect.garmin.com/oauthConfirm"

	//http.Handle("/", &demoAuthHandler{handler: demoServeHome, optional: true})

	http.HandleFunc("/authorize", userServeAuthorize)

	http.HandleFunc("/logout", userServeLogout)
	http.HandleFunc("/callback", demoServeOAuthCallback)

	http.Handle("/dailiesPath", &demoAuthHandler{handler: demoServeGetDailies})
	http.Handle("/epochsPath", &demoAuthHandler{handler: demoServeGetEpochs})
	http.Handle("/activitiesPath", &demoAuthHandler{handler: demoServeGetActivities})

	//http.Handle("/ping_epochs", &demoAuthHandler{handler: demoServePingEpochs})

	fmt.Println("+++ HANDLE FUNC 4 +++")

	//router.Run(":" + port)

	http.Handle("/", &demoAuthHandler{handler: demoServeHome, optional: true})
	//http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

	/*
		if err := http.ListenAndServe(*httpAddr, nil); err != nil {
			fmt.Println("+++ LISTEN ERROR +++")
			log.Fatalf("Error listening, %v", err)
		}
	*/

	fmt.Println("===== MAIN END =====")
	/*
		router := gin.Default()

		s := &http.Server{
			Addr:           ":8080",
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	*/
	/*
		portx := os.Getenv("PORT")

		fmt.Println("What's the PORT ??")
		fmt.Println(port)

		if port == "" {
			log.Fatal("$PORT must be set")
		}

		router := gin.New()
		fmt.Println(router)

		router.Use(gin.Logger())
		router.LoadHTMLGlob("templates/*.tmpl.html")
		router.Static("/static", "static")

		router.GET("/", func(c *gin.Context) {
			//c.HTML(http.StatusOK, "index.tmpl.html", nil)
			c.HTML(http.StatusOK, "loggedout.tmpl.html", nil)
		})

		//router.POST("/authorize", userServeAuthorize)

		router.GET("/post", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
			//id := c.Query("id")
			//page := c.DefaultQuery("page", "0")
			//name := c.PostForm("name")
			//message := c.PostForm("message")
			fmt.Printf("What's this ??")
			//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		})

		router.Run(":" + portx)
	*/
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}

func clientParseEpochs(w http.ResponseWriter, epochs []map[string]interface{}) {
	fmt.Println("==== CLIENT PARSE EPOCHS ====")

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
			fmt.Fprintf(w, "\n")
			switch k {
			case "summaryId":
			case "activityType":
			case "intensity":
				//fmt.Printf("id : %d, %s -> %s\n", i, k, v)
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
				//fmt.Println(time.Unix(a, 0))
				//fmt.Fprintf(w, "%s", a)
				fmt.Fprintf(w, "%v", time.Unix(a, 0))
				break

			case "steps":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows
				//fmt.Printf("步數 -> %d 步\n", a)
				fmt.Fprintf(w, "步數 -> %d 步\n", a)
				break
			case "distanceInMeters":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}
				a := int64(value) // <-- never overflows
				//fmt.Printf("距離 -> %d 米\n", a)
				fmt.Fprintf(w, "距離 -> %d 米\n", a)
				break

			case "durationInSeconds":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows

				//if k == "durationInSeconds" {
				//fmt.Printf("期間 -> %d 秒\n", a)
				//}
				fmt.Fprintf(w, "期間 -> %d 秒\n", a)
				break

			case "activeTimeInSeconds":
				value, _ := v.(float64)
				if value >= math.MaxInt64 || value <= math.MinInt64 { // <-- this works !
					fmt.Println("Conversion impossible: value is out of int64 range.")
					return
				}

				a := int64(value) // <-- never overflows

				//if k == "activeTimeInSeconds" {
				//	fmt.Printf("活動時間 -> %d 秒\n", a)
				//}
				fmt.Fprintf(w, "活動時間 -> %d 秒\n", a)
				break

			default:
				value, _ := v.(float64)
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
func timeStampProcess(r *http.Request) (string, string) {

	urlString := r.URL.String()
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	//fmt.Fprintf(w, "\n")
	//fmt.Fprintf(w, m["year"][0])
	//fmt.Fprintf(w, "\n")

	//xdate := m["year"][0] + "-" + m["mon"][0] + "-" + m["day"][0]
	//xtime := m["hr"][0] + ":" + m["min"][0] + ":" + m["sec"][0] + " GMT"
	//xdatetime := xdate + "T" + xtime
	//fmt.Fprintf(w, xdatetime)
	//fmt.Fprintf(w, "\n")

	// TIME Format
	loc, _ := time.LoadLocation("Asia/Taipei")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	//tx, _ := time.ParseInLocation(longForm, "Nov 1, 2018 at 09:26pm (CEST)", loc)
	//tx, _ := time.ParseInLocation(longForm, "Nov 1, 2018 at 09:26pm (TST)", loc)
	//fmt.Println(tx)

	monNr, err := strconv.Atoi(m["mon"][0])
	monShort := monthx(monNr)
	hrNr, err := strconv.Atoi(m["hr"][0])
	//var tState string
	//var tState string
	tState := "am"
	if hrNr >= 12 {
		hrNr = (hrNr - 12)
		tState = "pm"
	}
	hrStr := strconv.Itoa(hrNr)
	xdate := monShort + " " + m["day"][0] + ", " + m["year"][0]
	xtime := hrStr + ":" + m["min"][0] + tState + " (TST)" //+ ":" + m["sec"][0] + " GMT"
	xdatetime := xdate + " at " + xtime
	fmt.Println(xdatetime)
	timeStartTmp, _ := time.ParseInLocation(longForm, xdatetime, loc)
	fmt.Println(timeStartTmp)
	timeStartSecs := strconv.FormatInt(timeStartTmp.Unix(), 10)

	timeEndTmp := time.Now()

	timeEndSecs := strconv.FormatInt(timeEndTmp.Unix(), 10)
	fmt.Println(timeStartSecs)
	fmt.Println(timeEndSecs)

	timeStartSecsNr, err := strconv.Atoi(timeStartSecs)
	timeEndSecsNr, err := strconv.Atoi(timeEndSecs)
	if timeEndSecsNr > timeStartSecsNr {
		timeDiff := timeEndSecsNr - timeStartSecsNr
		if timeDiff >= 86400 {
			timeEndSecsNr = timeStartSecsNr + 86400
			timeEndSecs = strconv.Itoa(timeEndSecsNr)
		}
	} else {

	}

	fmt.Println(timeStartSecsNr)
	fmt.Println(timeEndSecsNr)
	return timeStartSecs, timeEndSecs

}

func monthx(mon int) string {
	month := "XXX"
	switch mon {
	case 1:
		month = time.January.String()
		break
	case 2:
		month = time.February.String()
		break
	case 3:
		month = time.March.String()
		break
	case 4:
		month = time.April.String()
		break
	case 5:
		month = time.May.String()
		break
	case 6:
		month = time.June.String()
		break
	case 7:
		month = time.July.String()
		break
	case 8:
		month = time.August.String()
		break
	case 9:
		month = time.September.String()
		break
	case 10:
		month = time.October.String()
		break
	case 11:
		month = time.November.String()
		break
	case 12:
		month = time.December.String()
		break
	default:
		fmt.Println("==error==")
		break
	}
	result := month[0:3]
	fmt.Println("MONTH")
	fmt.Println(result)
	return result
}
*/

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
*/

/*
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
*/

/*
func serveAuthorize(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" *****  SERVE AUTHORIZE  *****")
	fmt.Println(spxClient)
	callback := "http://" + r.Host + "/callback"
	tempCred, err := spxClient.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return
	}
	fmt.Println("=== CRED 1 ===")
	fmt.Println(tempCred)
	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session , "+err.Error(), 500)
		return
	}
	fmt.Println("==== REQUEST TOKEN ====")
	fmt.Println(callback)
	fmt.Println(tempCred)
	redir := spxClient.AuthorizationURL(tempCred, nil)
	fmt.Println("=== TOKEN 1 ===")
	fmt.Println(redir)
	fmt.Println("=== TOKEN 2 ===")

	//fmt.Println(w)
	//fmt.Println()
	//fmt.Println(r)
	http.Redirect(w, r, redir, 302)

	//http.Redirect(w, r, spxClient.AuthorizationURL(tempCred, nil), 302)
	//http.Redirect(w, r, spxClient.AuthorizationURL(tempCred, url.Values{"oauth_callback": {callback}}), 302)
}
*/

/*
var beginSecs string
var endSecs string

func demoTimeStamp() {
	const longForm = "Mon, 02 Jan 2006 15:04:05 MST"

	xbegin, _ := time.Parse(longForm, "Sun, 14 Oct 2018 07:18:40 MST")
	xend, _ := time.Parse(longForm, "Mon, 15 Oct 2018 07:18:40 MST")

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

/*
// Session state keys.
const (
	tempCredKey  = "tempCred"
	tokenCredKey = "tokenCred"
)

var spxClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://connectapi.garmin.com/oauth-service/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://connect.garmin.com/oauthConfirm",
	TokenRequestURI:               "https://connectapi.garmin.com/oauth-service/oauth/access_token",
}

var signinOAuthClient oauth.Client

var credPath = flag.String("t-config", "tcon.json", "Path to configuration file containing the application's credentials.")



var localCredentials oauth.Credentials
*/
/*
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
*/

// serveSignin gets the OAuth temp credentials and redirects the user to the
// Twitter's authentication page.
/*
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
*/

/*
var (
	homeLoggedOutTmpl = template.Must(template.New("loggedout").Parse(
		`<html>
<head>
</head>

<body>

<a href="/authorize">Authorize</a> or
<a href="/timeline"><img src="https://www.health2sync.com/v4/images/main/cover_orange.png"></a>

<!--
<p>年 月 日 </p>
<form oninput="x.value=parseInt(a.value)">2010
<input type="range" min="2010" max="2030" id="a" value="2020">2030
<br>
= <output name="x" for="a"></output>
</form>

<form oninput="x.value=parseInt(a.value)">1
<input type="range" min="1" max="12" id="a" value="6">12
<br>
= <output name="x" for="a"></output>
</form>

<form oninput="x.value=parseInt(a.value)">1
<input type="range" min="1" max="31" id="a" value="15">31
<br>
= <output name="x" for="a"></output>
</form>



<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">2010
<input type="range" id="a" value="2010">2030
+<input type="number" id="b" value="2010">
=<output name="x" for="a b"></output>
</form>


<form oninput="x.value=parseInt(a.value)">2010 <br>
<br>
<input type="range" min="0" max="11" id="a" value="5">30
<br>
=<output name="x" for="a"></output>
</form>

<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">2010
<input type="range" id="a" value="50">30
+<input type="number" id="b" value="2010">
=<output name="x" for="a b"></output>
</form>

-->


<!--
<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">0
<input type="range" id="a" value="50">100
+<input type="number" id="b" value="50">
=<output name="x" for="a b"></output>
</form>

<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">0
<input type="range" id="a" value="50">100
+<input type="number" id="b" value="50">
=<output name="x" for="a b"></output>
</form>

<form oninput="x.value=parseInt(a.value)+parseInt(b.value)">0
<input type="range" id="a" value="50">100
+<input type="number" id="b" value="50">
=<output name="x" for="a b"></output>
</form>

-->



</body>
</html>`))
)

*/

/*
<!--
	homeTmpl = template.Must(template.New("home").Parse(
		`<html>
<head>
</head>
<body>
<p><a href="/dailies">GET DAILIES</a>
<p><a href="/epochs">GET EPOCHS</a>
<p><a href="/activities">GET ACTIVITIES</a>

<br><br>
<p><a href="ping_dailies">PING DAILIES</a>
<p><a href="ping_epochs">PING EPOCHS</a>
<p><a href="ping_activities">PING ACTIVITIES</a>


<p><a href="/logout">logout</a>
</body></heml>`))

-->
)
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
