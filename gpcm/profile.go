package gpcm

import (
	"errors"
	"strconv"
	"wwfc/common"
	"wwfc/database"
	"wwfc/logging"
	"wwfc/qr2"

	"github.com/logrusorgru/aurora/v3"
)

func (g *GameSpySession) getProfile(command common.GameSpyCommand) {
	strProfileId := command.OtherValues["profileid"]
	profileId, err := strconv.ParseUint(strProfileId, 10, 32)
	if err != nil {
		// There was an error getting profile info.
		g.replyError(ErrGetProfile)
		return
	}

	logging.Info(g.ModuleName, "Looking up the profile of", aurora.Cyan(profileId).String())

	user := database.User{}
	locstring := ""

	mutex.Lock()
	if session, ok := sessions[uint32(profileId)]; ok && session.LoggedIn {
		locstring = session.LocString
		user = session.User
		mutex.Unlock()
	} else {
		mutex.Unlock()
		user, err = database.GetProfile(pool, ctx, uint32(profileId))
		if err != nil {
			// The profile info was requested on is invalid.
			g.replyError(ErrGetProfileBadProfile)
			return
		}
	}

	if user.ProfileId == g.User.ProfileId {
		g.WriteBuffer += common.CreateGameSpyMessage(common.GameSpyCommand{
			Command:      "pi",
			CommandValue: "",
			OtherValues: map[string]string{
				"profileid":  command.OtherValues["profileid"],
				"nick":       user.UniqueNick,
				"userid":     strconv.FormatUint(uint64(user.UserId), 10),
				"email":      user.Email,
				"sig":        common.RandomHexString(32),
				"uniquenick": user.UniqueNick,
				"firstname":  user.FirstName,
				"lastname":   user.LastName,
				"pid":        "11",
				"lon":        "0.000000",
				"lat":        "0.000000",
				"loc":        locstring,
				"id":         command.OtherValues["id"],
			},
		})
	} else {
		g.WriteBuffer += common.CreateGameSpyMessage(common.GameSpyCommand{
			Command:      "pi",
			CommandValue: "",
			OtherValues: map[string]string{
				"profileid":  command.OtherValues["profileid"],
				"nick":       "000000000" + user.GsbrCode[:4] + "0000000",
				"userid":     "0",
				"email":      "000000000" + user.GsbrCode[:4] + "0000000" + "@nds",
				"sig":        common.RandomHexString(32),
				"uniquenick": "000000000" + user.GsbrCode[:4] + "0000000",
				"firstname":  user.FirstName,
				"lastname":   "000000000" + user.GsbrCode[:4] + "0000000",
				"pid":        "11",
				"lon":        "0.000000",
				"lat":        "0.000000",
				"loc":        locstring,
				"id":         command.OtherValues["id"],
			},
		})
	}
}

func (g *GameSpySession) updateProfile(command common.GameSpyCommand) {
	if openHost, ok := command.OtherValues["wl:oh"]; ok {
		enabled := openHost != "0"
		if !g.User.OpenHost && enabled {
			g.openHostEnabled(true, true)
		} else if g.User.OpenHost && !enabled {
			g.openHostDisabled()
		}

		qr2.SetLoginOpenHost(g.User.ProfileId, enabled)
	}

	g.User.UpdateProfile(pool, ctx, command.OtherValues)
}

var ErrNoSession = errors.New("Session not found for given profileID")

func SetSessionDiscordInfo(profileId uint32, stage database.LinkStage, discordId string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if session, exists := sessions[profileId]; exists {
		session.User.LinkStage = stage
		session.User.DiscordID = discordId
		logging.Info("GPCM", "Set Discord ID for profile", aurora.Cyan(profileId), "to", aurora.Cyan(discordId), "stage", aurora.Cyan(stage))

		return nil
	} else {
		logging.Info("GPCM", "Session not found for profile ID", aurora.Cyan(profileId))

		return ErrNoSession
	}
}

func GetSessionDiscordInfo(profileId uint32) (database.LinkStage, string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if session, exists := sessions[profileId]; exists {
		return session.User.LinkStage, session.User.DiscordID, nil
	} else {
		logging.Info("GPCM", "Session not found for profile ID", aurora.Cyan(profileId))
		return 0, "", ErrNoSession
	}
}

func VerifyPlayerSearch(profileId uint32, sessionKey int32, gameName string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if session, ok := sessions[profileId]; ok && session.LoggedIn && session.SessionKey == sessionKey && session.GameName == gameName {
		return "000000000" + session.User.GsbrCode[:4] + "0000000", true
	}

	return "", false
}
