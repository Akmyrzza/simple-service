package usecases

import "github.com/Akmyrzza/simple-service/internal/domain/core"

type St struct {
	cr *core.St
}

func New() *St {
	u := &St{}
	return u
}
func (u *St) SetCore(core *core.St) {
	u.cr = core
}
