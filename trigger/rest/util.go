package rest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/project-flogo/core/data/coerce"
)

const HeaderSetCookie = "Set-Cookie"

func getFileDetails(key string, header *multipart.FileHeader) (map[string]interface{}, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	fileDetails := map[string]interface{}{
		"key":      key,
		"fileName": header.Filename,
		"fileType": header.Header.Get("Content-Type"),
		"size":     header.Size,
		"file":     buf.Bytes(),
	}

	return fileDetails, nil
}

func addCookies(w http.ResponseWriter, cookies []interface{}) error {

	w.Header().Del(HeaderSetCookie)

	for _, cookie := range cookies {

		strCookie := ""

		if c, ok := cookie.(*http.Cookie); ok {
			strCookie = c.String()
		} else if c, ok := cookie.(string); ok {
			strCookie = c
		} else if c, ok := cookie.(map[string]interface{}); ok {
			cookie, err := mapToCookie(c)
			if err != nil {
				return err
			}
			strCookie = cookie.String()
		} else if params, ok := cookie.(map[string]string); ok {
			c, _ := coerce.ToObject(params)

			cookie, err := mapToCookie(c)
			if err != nil {
				return err
			}
			strCookie = cookie.String()
		} else {
			return fmt.Errorf("unsupported cookie format: %v", cookie)
		}

		if strCookie != "" {
			w.Header().Add(HeaderSetCookie, strCookie)
		}
	}

	return nil
}

func mapToCookie(values map[string]interface{}) (cookie *http.Cookie, err error) {

	cookie = &http.Cookie{}

	for key, value := range values {
		switch strings.ToLower(key) {
		case "name":
			cookie.Name = value.(string)
		case "value":
			cookie.Value, err = coerce.ToString(value)
			if err != nil {
				return nil, err
			}
		case "path":
			cookie.Path = value.(string)
		case "domain":
			cookie.Domain = value.(string)
		case "expires":
			if dt, ok := value.(time.Time); ok {
				cookie.Expires = dt.UTC()
			} else if dt, ok := value.(string); ok {
				exptime, err := time.Parse(time.RFC1123, dt)
				if err != nil {
					exptime, err = time.Parse("Mon, 02-Jan-2006 15:04:05 MST", dt)
					if err != nil {
						return nil, fmt.Errorf("invalid expiration: %s", dt)
					}
				}
				cookie.Expires = exptime.UTC()
			}
		case "maxage":
			cookie.MaxAge, err = coerce.ToInt(value)
			if err != nil {
				return nil, err
			}
		case "secure":
			cookie.Secure, err = coerce.ToBool(value)
			if err != nil {
				return nil, err
			}
		case "httponly":
			cookie.HttpOnly, err = coerce.ToBool(value)
			if err != nil {
				return nil, err
			}
		case "samesite":
			if ss, err := coerce.ToInt(value); err == nil {
				cookie.SameSite = http.SameSite(ss)
			} else if ss, ok := value.(string); ok {
				switch strings.ToLower(ss) {
				case "lax":
					cookie.SameSite = http.SameSiteLaxMode
				case "strict":
					cookie.SameSite = http.SameSiteStrictMode
				default:
					cookie.SameSite = http.SameSiteDefaultMode
				}
			}
		}

	}

	return cookie, nil
}
