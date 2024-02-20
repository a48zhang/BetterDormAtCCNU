package form

import (
	"context"
	"main/model"
)

func CreateForm(ctx context.Context, data model.Form) error {
	err := model.Insert(ctx, &data)
	return err
}

func GetUsersForm(ctx context.Context, uid string) ([]model.Form, error) {
	forms, err := model.GetMany(ctx, &model.Form{StudentID: uid})
	return forms, err
}

// CheckForm
// opt: 0 for waiting, 1 for teacher approve, -1 for teacher reject
// 2 for school approve, -2 for school reject
func CheckForm(ctx context.Context, fid string, opt int) error {
	data := model.Form{Fid: fid, Status: opt}
	err := model.Update(ctx, &data)
	return err
}
