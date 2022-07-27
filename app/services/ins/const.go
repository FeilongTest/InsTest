package ins

import "github.com/levigross/grequests"

const (
	UserAgent    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36 Edg/103.0.1264.49"
	ShareDataUrl = "https://www.instagram.com/data/shared_data/"
	LoginUrl     = "https://www.instagram.com/accounts/login/ajax/"
)

type Client struct {
	Session   *grequests.Session
	Header    map[string]string
	UserAgent string
}
