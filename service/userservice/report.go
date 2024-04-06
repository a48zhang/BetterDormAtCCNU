package userservice

import (
	"context"
	"main/model"
)

func NewReport(ctx context.Context, uid string, info string, contact string) error {
	data := model.Report{
		Uid:     uid,
		Info:    info,
		Contact: contact,
		Comment: "No Comment",
	}
	return data.Insert(ctx)
}

func MyReport(ctx context.Context, uid string) ([]model.Report, error) {
	data := model.Report{
		Uid: uid,
	}
	res, err := data.FindByUid(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]model.Report, 0)
	for _, v := range res {
		ret = append(ret, v.(model.Report))
	}
	return ret, nil
}
