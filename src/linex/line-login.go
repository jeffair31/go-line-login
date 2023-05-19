package linex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// Constants
const (
	ClientID     = "1661154522"
	ClientSecret = "8600b46336182bcdad79cfc8f84887f7"
	RedirectURI  = "http://localhost:8080/callback"
	AuthURL      = "https://access.line.me/oauth2/v2.1/authorize"
	TokenURL     = "https://api.line.me/oauth2/v2.1/token"
)

func CallbackHandler(c *gin.Context) {

	code := c.Query("code")

	// Exchange the authorization code for an access token
	params := url.Values{
		"grant_type":    []string{"authorization_code"},
		"code":          []string{code},
		"client_id":     []string{os.Getenv("CLIENT_ID")},
		"client_secret": []string{os.Getenv("CLIENT_SECRET")},
		"redirect_uri":  []string{RedirectURI},
	}

	// dd := io.Reader{}

	// http.Post(TokenURL, "", "")

	//
	resp, err := http.PostForm(TokenURL, params)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	// Handle the response
	// You can parse the JSON response to extract the access token
	// and use it to make API calls on behalf of the user
	// For simplicity, we'll just print the response body here

	// Log the request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var rx interface{}
	if err := json.Unmarshal(body, &rx); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//
	c.JSON(http.StatusOK, rx)
}
