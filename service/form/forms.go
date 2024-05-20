package form

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dao"
)

func CreateForm(ctx context.Context, data dao.Form) error {
	err := data.Insert(ctx)
	return err
}

func GetUsersForm(ctx context.Context, uid string) ([]dao.Form, error) {
	forms, err := (&dao.Form{StudentID: uid}).FindByStudentId(ctx, uid)
	if err != nil {
		return nil, err
	}
	ret := make([]dao.Form, len(forms))
	for i, v := range forms {
		ret[i] = v.(dao.Form)
	}
	return ret, err
}

func GetOneForm(ctx context.Context, fid string) (dao.Form, error) {
	// convert fid to []byte
	fidByte, err := primitive.ObjectIDFromHex(fid)
	data := &dao.Form{MID: fidByte}
	res, err := data.FindBy(ctx, bson.M{"_id": fidByte})
	if err != nil {
		return dao.Form{}, err
	}
	*data = res[0].(dao.Form)
	return *data, err
}

// CheckForm
// opt: 0 for waiting, 1 for teacher approve, -1 for teacher reject
// 2 for school approve, -2 for school reject
func CheckForm(ctx context.Context, fid string, opt int, advice string) error {
	fidByte, err := primitive.ObjectIDFromHex(fid)
	data := dao.Form{MID: fidByte, Status: opt, TeacherAdvice: advice}
	err = data.Update(ctx)
	return err
}

func GetAssignedForms(ctx context.Context, uid string, role int) ([]dao.Form, error) {
	if role > 0 {
		if role == 1 {
			forms, err := (&dao.Form{TeacherID: uid}).FindByTeacherId(ctx, uid)
			if err != nil {
				return nil, err
			}
			ret := make([]dao.Form, len(forms))
			for i, v := range forms {
				ret[i] = v.(dao.Form)
			}
			return ret, err
		} else {
			forms, err := (&dao.Form{}).FindBy(ctx, bson.M{"status": 1})
			if err != nil {
				return nil, err
			}
			ret := make([]dao.Form, len(forms))
			for i, v := range forms {
				ret[i] = v.(dao.Form)
			}
			return ret, err
		}
	}
	return nil, errors.New("role not satisfied")
}
