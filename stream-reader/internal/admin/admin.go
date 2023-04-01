package admin

import (
	commonsAdmin "github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/admin"
)

var kAdmin commonsAdmin.AdminInterface

func Init() {
	commonsAdmin.InitAdmin()

	kAdmin = commonsAdmin.GetAdmin()
}

func GetAdminInstance() commonsAdmin.AdminInterface {
	return kAdmin
}
