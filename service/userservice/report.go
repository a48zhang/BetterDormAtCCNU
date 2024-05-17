package userservice

import (
	"context"
	"main/dao"
)

func NewReport(ctx context.Context, uid string, info string, contact string) error {
	data := dao.Report{
		Uid:     uid,
		Info:    info,
		Contact: contact,
		Comment: "No Comment",
	}
	return data.Insert(ctx)
}

func MyReport(ctx context.Context, uid string) ([]dao.Report, error) {
	data := dao.Report{
		Uid: uid,
	}
	res, err := data.FindByUid(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]dao.Report, 0)
	for _, v := range res {
		ret = append(ret, v.(dao.Report))
	}
	return ret, nil
}
