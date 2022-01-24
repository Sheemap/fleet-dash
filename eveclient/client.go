package eveclient

import "C"
import (
	"errors"
	"fleet-dash-core/eveclient/client"
	"fleet-dash-core/eveclient/client/fleets"
	"fleet-dash-core/utilities"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

var (
	ErrExpiredToken = errors.New("expired access token")
)

type EveClient interface {
	GetCurrentFleet() (*fleets.GetCharactersCharacterIDFleetOKBody, error)
}

type clientImpl struct {
	characterID int32
	accessToken *jwt.Token
	authenticator *authenticator
}

type authenticator struct {
	accessToken *jwt.Token
}

func NewEveClient(characterToken *jwt.Token) (EveClient, error) {
	characterID, err := utilities.GetCharacterIDFromToken(characterToken)
	if err != nil {
		return nil, err
	}

	charID, err := strconv.ParseInt(characterID, 10, 32)
	if err != nil {
		return nil, err
	}

	authy := &authenticator{
		accessToken: characterToken,
	}

	return &clientImpl{
		characterID: int32(charID),
		accessToken: characterToken,
		authenticator: authy,
	}, nil
}

func (c *clientImpl) GetCurrentFleet() (*fleets.GetCharactersCharacterIDFleetOKBody, error){
	eveClient := client.NewHTTPClient(strfmt.Default)

	params := fleets.NewGetCharactersCharacterIDFleetParams()
	params.CharacterID = c.characterID
	res, err := eveClient.Fleets.GetCharactersCharacterIDFleet(params, c.authenticator)

	switch err.(type) {
	case *fleets.GetCharactersCharacterIDFleetNotFound:
		return nil, nil
	default:
		return res.Payload, nil

	}
}



func (a *authenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error{
	err := req.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", a.accessToken.Raw))
	return err
}