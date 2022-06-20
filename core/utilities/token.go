package utilities

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"math/big"
	"net/http"
	"strings"
	"time"
)

var (
	ErrUnableToParseCharacterID = errors.New("unable to parse character ID")
	ErrNotAValidCharacterToken  = errors.New("token is not for a character")
)

type jwks struct {
	Keys []jwk `json:"keys"`
	SkipUnresolvableJsonWebKeys bool `json:"SkipUnresolvedJsonWebKeys"`
}

type jwk struct {
	Algorithm string `json:"alg"`
	E 		string `json:"e"`
	N 		string `json:"n"`
	Use string `json:"use"`
	Kty string `json:"kty"`
	Kid string `json:"kid"`
}

type JwtValidator struct{
	jwks *jwks
	lastFetchedJwks time.Time
}

func NewValidator() JwtValidator{
	return JwtValidator{}
}


func (jv *JwtValidator) Parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jv.constructPublicKey(token)
	})
}

func (jv *JwtValidator) fetchJwks() (*jwks, error) {
	if jv.jwks != nil && time.Now().Sub(jv.lastFetchedJwks) < time.Hour * 24 {
		return jv.jwks, nil
	}

	resp, err := http.Get("https://login.eveonline.com/oauth/jwks")
	if err != nil {
		return nil, err
	}

	keySet := &jwks{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&keySet)
	if err != nil {
		return nil, err
	}

	return keySet, nil
}

func (jv *JwtValidator) constructPublicKey(token *jwt.Token) (*rsa.PublicKey, error){
	keySet, err := jv.fetchJwks()
	if err != nil {
		return nil, err
	}

	var signingKey jwk
	kid := token.Header["kid"].(string)
	for _, key := range keySet.Keys {
		if key.Kid == kid {
			signingKey = key
		}
	}

	nString := signingKey.N
	decN, err := base64.RawURLEncoding.DecodeString(nString)
	if err != nil {
		return nil, err
	}
	n := big.NewInt(0)
	n.SetBytes(decN)

	eString := signingKey.E
	decE, err := base64.StdEncoding.DecodeString(eString)
	if err != nil {
		return nil, err
	}
	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}
	eReader := bytes.NewReader(eBytes)
	var e uint64
	err = binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		return nil, err
	}

	return &rsa.PublicKey{
		N: n,
		E: int(e),
	}, nil
}

func GetCharacterIDFromToken(token *jwt.Token) (string, error) {
	claims := token.Claims.(jwt.MapClaims)

	subject := claims["sub"].(string)
	subjectSlice := strings.Split(subject, ":")
	if len(subjectSlice) != 3 {
		return "", ErrUnableToParseCharacterID
	}

	if subjectSlice[0] != "CHARACTER" {
		return "", ErrNotAValidCharacterToken
	}

	return subjectSlice[2], nil
}

func GetCharacterNameFromToken(token *jwt.Token) (string, error) {
	claims := token.Claims.(jwt.MapClaims)

	subject := claims["name"].(string)
	return subject, nil
}