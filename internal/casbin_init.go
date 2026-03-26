package blog_r

import (
	"blog_r/internal/global"

	"github.com/casbin/casbin/v3"
)

func InitEnforce() error {
	enforce, err := casbin.NewEnforcer("./internal/casbin/model.conf", "./internal/casbin/policy.csv")
	if err != nil {
		return err
	}
	global.Enforcer = enforce
	return nil
}
