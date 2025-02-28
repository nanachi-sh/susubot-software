package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/nanachi-sh/susubot-code/plugin/twoonone/internal/configs"
	"github.com/nanachi-sh/susubot-code/plugin/twoonone/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	once sync.Once

	jwks *jwk.Cache

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	claims = struct {
		JWKSUri string `json:"jwks_uri"`
	}{}

	xlogger = logx.WithContext(context.Background())
)

func initialize() {
	provider, err := oidc.NewProvider(context.Background(), configs.OIDC_ISSUER)
	if err != nil {
		logger.Fatalln(err)
	}
	if err := provider.Claims(&claims); err != nil {
		logger.Fatalln(err)
	}
	{
		c, err := jwk.NewCache(context.Background(), httprc.NewClient())
		if err != nil {
			logger.Fatalln(err)
		}
		jwks = c
	}
	if err := jwks.Register(context.Background(), claims.JWKSUri); err != nil {
		logger.Fatalln(err)
	}
	for {
		if jwks.Ready(context.Background(), claims.JWKSUri) {
			break
		}
	}
}

func Handle(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	once.Do(initialize)

	access_token, err := r.Cookie(types.COOKIE_KEY_access_token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !verifyToken(access_token.Value) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	next(w, r)
}

// func cutDomain(domain string) string {
// 	domain = regexp.
// 		MustCompile(`^(?:[a-z0-9-]+\.)+([a-z0-9-]+\.[a-z]{2,})$`).
// 		ReplaceAllString(domain, `$1`)
// 	if domain[0] != []byte(".")[0] {
// 		domain = "." + domain
// 	}
// 	return domain
// }

func verifyToken(token_raw string) bool {
	m := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(token_raw, m, func(t *jwt.Token) (interface{}, error) {
		set, err := jwks.Lookup(context.Background(), claims.JWKSUri)
		if err != nil {
			return nil, err
		}
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, errors.New("token kid empty or !string")
		}
		k, ok := set.LookupKeyID(kid)
		if !ok {
			return nil, errors.New("no exist token pub key")
		}
		var pk any
		if err := jwk.Export(k, &pk); err != nil {
			return nil, err
		}
		return pk, nil
	})
	if err != nil {
		xlogger.Error(err)
		return false
	}
	// 验证access_token是否包含email, name, user_id
	ok := func() bool {
		if _, ok := m["email"]; !ok {
			return false
		}
		if _, ok := m["name"]; !ok {
			return false
		}
		if mi, ok := m["federated_claims"].(map[string]any); !ok {
			return false
		} else {
			if _, ok := mi["user_id"]; !ok {
				return false
			}
		}
		return true
	}()
	if !ok {
		return false
	}
	return token.Valid
}

// func loginHandle(w http.ResponseWriter, r *http.Request) bool {
// 	// 登录验证
// 	{
// 		access_token, err := r.Cookie(types.COOKIE_KEY_access_token)
// 		if err == nil {
// 			if verifyToken(access_token.Value) {
// 				w.WriteHeader(http.StatusNotAcceptable)
// 				return false
// 			}
// 		}
// 	}
// 	//
// 	session_id, ok := getSessionId(r)
// 	if !ok {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return false
// 	}
// 	state := utils.RandomString(16, utils.Dict_Mixed)
// 	if err := configs.Redis.Hset(session_id, types.REDIS_KEY_SESSION_State, state); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		xlogger.Error(err)
// 		return false
// 	}
// 	http.Redirect(w, r, o2cfg.AuthCodeURL(state), http.StatusFound)
// 	return true
// }

// func formatCookie(key, value, domain string, expires_in int) *http.Cookie {
// 	domain = ".unturned.fun"
// 	return &http.Cookie{
// 		Name:     key,
// 		Value:    value,
// 		Path:     "/",
// 		Domain:   domain,
// 		MaxAge:   expires_in,
// 		Secure:   true,
// 		HttpOnly: true,
// 		SameSite: http.SameSiteNoneMode,
// 	}
// }
