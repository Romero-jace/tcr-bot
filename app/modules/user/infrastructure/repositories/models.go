package userdb

import (
	usertypes "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/types"
	"github.com/uptrace/bun"
)

// User represents a user in the system.
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64                  `bun:"id,pk,autoincrement" json:"id"`
	Name          string                 `bun:"name" json:"name"`
	DiscordID     usertypes.DiscordID    `bun:"discord_id,unique"`
	Role          usertypes.UserRoleEnum `bun:"role,notnull" json:"role"`
}

// Add these methods to your User struct
func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetDiscordID() usertypes.DiscordID {
	return u.DiscordID
}

func (u *User) GetRole() usertypes.UserRoleEnum {
	return u.Role
}