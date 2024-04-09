package form

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"main/model"
)

func CreateForm(ctx context.Context, data model.Form) error {
	err := data.Insert(ctx)
	return err
}

func GetUsersForm(ctx context.Context, uid string) ([]model.Form, error) {
	forms, err := (&model.Form{StudentID: uid}).FindByStudentId(ctx, uid)
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
func CheckForm(ctx context.Context, fid string, opt int, advice string) error {
	data := model.Form{Fid: fid, Status: opt, TeacherAdvice: advice}
	err := data.Update(ctx)
	return err
}

func GetAssignedForms(ctx context.Context, uid string, role int) ([]model.Form, error) {

	switch role {
	case 1:
		forms, err := (&model.Form{TeacherID: uid}).FindByTeacherId(ctx, uid)
		if err != nil {
			return nil, err
		}
		ret := make([]model.Form, len(forms))
		for i, v := range forms {
			ret[i] = v.(model.Form)
		}
		return ret, err

	case 2:
		forms, err := (&model.Form{}).FindBy(ctx, bson.M{"status": 1})
		if err != nil {
			return nil, err
		}
		ret := make([]model.Form, len(forms))
		for i, v := range forms {
			ret[i] = v.(model.Form)
		}
		return ret, err
	}

	return nil, errors.New("role not satisfied")
}
