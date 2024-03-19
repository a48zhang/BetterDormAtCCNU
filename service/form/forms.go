package form

import (
	"context"
	"main/model"
)

func CreateForm(ctx context.Context, data model.Form) error {
	err := data.Insert(ctx)
	return err
}

func GetUsersForm(ctx context.Context, uid string) ([]model.Form, error) {
	forms, err := (&model.Form{StudentID: uid}).FindAll(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]model.Form, len(forms))
	for i, v := range forms {
		ret[i] = v.(model.Form)
	}
	return ret, err
}

func GetOneForm(ctx context.Context, fid string) (model.Form, error) {
	data := &model.Form{Fid: fid}
	err := data.Find(ctx)
	return *data, err
}

// CheckForm
// opt: 0 for waiting, 1 for teacher approve, -1 for teacher reject
// 2 for school approve, -2 for school reject
func CheckForm(ctx context.Context, fid string, opt int) error {
	data := model.Form{Fid: fid, Status: opt}
	err := data.Update(ctx)
	return err
}
