// Time        : 2019/09/05
// Description :

package model

type Role struct {
	RoleId   string `bson:"_id"`
	Students []StudentValue
}
