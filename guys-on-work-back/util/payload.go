package util

import (
	"errors"
	"time"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
    ErrInvalidToken = errors.New("token is invalid")
    ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
    ID        uuid.UUID `json:"id"`
    Username  string    `json:"username"`
    IssuedAt  time.Time `json:"issued_at"`
    ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }

    payload := &Payload{
        ID:        tokenID,
        Username:  username,
        IssuedAt:  time.Now(),
        ExpiredAt: time.Now().Add(duration),
    }
    return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiredAt) {
        return ErrExpiredToken
    }
    return nil
}








// func loginUser(server *gin.Engine, ctx *gin.Context) {
//     var req Login
//     if err := ctx.ShouldBindJSON(&req); err != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     user, err := server.db.GetUser(ctx, req.Username)
//     if err != nil {
//         if err == sql.ErrNoRows {
//             ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//             return
//         }
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//     }

//     err = util.CheckPassword(req.UserSystemPassword, user.HashedPassword)
//     if err != nil {
//         ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//         return
//     }

//     accessToken, accessPayload, err := server.tokenMaker.CreateToken(
//         user.Username,
//         server.config.AccessTokenDuration,
//     )
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
//         user.Username,
//         server.config.RefreshTokenDuration,
//     )
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     session, err := server.db.CreateSession(ctx, db.CreateSessionParams{
//         ID:           refreshPayload.ID,
//         Username:     user.Username,
//         RefreshToken: refreshToken,
//         UserAgent:    ctx.Request.UserAgent(),
//         ClientIp:     ctx.ClientIP(),
//         IsBlocked:    false,
//         ExpiresAt:    refreshPayload.ExpiredAt,
//     })
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     rsp := loginUserResponse{
//         SessionID:             session.ID,
//         AccessToken:           accessToken,
//         AccessTokenExpiresAt:  accessPayload.ExpiredAt,
//         RefreshToken:          refreshToken,
//         RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
//         User:                  newUserResponse(user),
//     }
//     ctx.JSON(http.StatusOK, rsp)
// }