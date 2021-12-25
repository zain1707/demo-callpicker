package usecase

import (
	"context"
	"github.com/gomarkho/demo-callpicker/inbound"
	"github.com/gomarkho/demo-callpicker/model"
	"strconv"
	"strings"
)

type InboundUseCase struct {
	InboundRepo inbound.Repository
}

func NewInboundUseCase(ir inbound.Repository) inbound.UseCase {
	return &InboundUseCase{ir}
}

func (u *InboundUseCase) InsertDialPlan(ctx context.Context, req *inbound.InsertDialPlan) (*model.DialPlan, int, error) {
	var dialplan model.DialPlan

	req.Source = validateNumber(req.Source)
	req.Destination = validateNumber(req.Destination)

	flag, obj := u.InboundRepo.Exists(ctx, req.Destination, req.Source)
	if flag {
		return obj, model.CodeAlreadyExists, model.ErrAlreadyExists
	}

	dialplan.Source = req.Source
	dialplan.Destination = req.Destination
	dialplan.Say = req.Say
	dialplan.Play = req.Play
	dialplan.Record = req.Record
	dialplan.Dial = req.Dial
	for i, _ := range req.Record {
		req.Record[i].Method = strings.ToUpper(req.Record[i].Method)
	}
	for i, _ := range req.Dial {
		if req.Dial[i].CallerID == "" {
			dialplan.Dial[i].CallerID = req.Source
		}
		req.Dial[i].Method = strings.ToUpper(req.Dial[i].Method)
	}
	flag = checkSteps(&dialplan)
	if !flag {
		return nil, model.CodeSteps, model.ErrSteps
	}

	if flag {
		_, err := u.InboundRepo.Store(ctx, &dialplan)
		if err != nil {
			return nil, model.CodeDatabase, model.ErrDatabase
		}
	}

	return &dialplan, 0, nil
}

func (u *InboundUseCase) DeleteDialPlan(ctx context.Context, req inbound.DeleteDialPlan) (int, error) {

	req.Source = validateNumber(req.Source)
	req.Destination = validateNumber(req.Destination)

	flag, _ := u.InboundRepo.Exists(ctx, req.Destination, req.Source)
	if !flag {
		return model.CodeNotExists, model.ErrNotExists
	}

	err := u.InboundRepo.DeleteByNumber(ctx, req)
	if err != nil {
		return model.CodeDatabase, model.ErrDatabase
	}

	return model.CodeSuccess, nil
}

func (u *InboundUseCase) GetDialPlan(ctx context.Context, req inbound.GetDialPlan) (*model.DialPlan, int, error) {

	req.Source = validateNumber(req.Source)
	req.Destination = validateNumber(req.Destination)

	flag, _ := u.InboundRepo.Exists(ctx, req.Destination, req.Source)
	if !flag {
		return nil, model.CodeNotExists, model.ErrNotExists
	}

	obj, err := u.InboundRepo.GetDialplanByNumber(ctx, req)
	if err != nil {
		return nil, model.CodeDatabase, model.ErrDatabase
	}

	return obj, model.CodeSuccess, nil
}

func (u *InboundUseCase) GetXml(ctx context.Context, req *inbound.XMLOptions) (*model.XMLResponse, error) {

	req.Source = validateNumber(req.Source)
	req.Destination = validateNumber(req.Destination)

	flag, _ := u.InboundRepo.Exists(ctx, req.Destination, req.Source)
	if !flag {
		return nil, model.ErrNotExists
	}

	dialplan, err := u.InboundRepo.GetByNumber(ctx, *req)
	if err != nil {
		return nil, model.ErrDatabase
	}
	arrModules, err := sortDialplan(dialplan)
	if err != nil {
		return nil, err
	}

	var response model.XMLResponse
	for _, v := range arrModules {
		sayCount, playCount, dialCount, recordCount := 0, 0, 0, 0
		switch v {
		case "say":
			response.Elements = append(response.Elements, dialplan.Say[sayCount])
			sayCount++
		case "play":
			response.Elements = append(response.Elements, dialplan.Play[playCount])
			playCount++
		case "dial":
			response.Elements = append(response.Elements, dialplan.Dial[dialCount])
			dialCount++
		case "record":
			response.Elements = append(response.Elements, dialplan.Record[recordCount])
			recordCount++
		default:
			//Do Nothing
		}
	}
	return &response, nil

}

func sortDialplan(dialplan *model.DialPlan) ([]string, error) {

	var arrModules []string
	totalSteps := len(dialplan.Say) + len(dialplan.Play) + len(dialplan.Record) + len(dialplan.Dial)

	for i := 0; i < totalSteps; i++ {

		if dialplan.Say != nil {
			for _, v := range dialplan.Say {
				if v.Step == strconv.Itoa(i+1) {
					arrModules = append(arrModules, "say")
				}
			}
		}
		if dialplan.Play != nil {
			for _, v := range dialplan.Play {
				if v.Step == strconv.Itoa(i+1) {
					arrModules = append(arrModules, "play")
				}
			}
		}

		if dialplan.Record != nil {
			for _, v := range dialplan.Record {
				if v.Step == strconv.Itoa(i+1) {
					arrModules = append(arrModules, "record")
				}
			}
		}
		if dialplan.Dial != nil {
			for _, v := range dialplan.Dial {
				if v.Step == strconv.Itoa(i+1) {
					arrModules = append(arrModules, "dial")
				}
			}

		}
	}

	return arrModules, nil

}

func checkSteps(dialplan *model.DialPlan) bool {

	var arrsteps []string
	totalSteps := len(dialplan.Say) + len(dialplan.Play) + len(dialplan.Record) + len(dialplan.Dial)

	for i := 0; i < totalSteps; i++ {

		if dialplan.Say != nil {
			for _, v := range dialplan.Say {
				if v.Step == strconv.Itoa(i+1) {
					arrsteps = append(arrsteps, v.Step)
				}
			}
		}
		if dialplan.Play != nil {
			for _, v := range dialplan.Play {
				if v.Step == strconv.Itoa(i+1) {
					arrsteps = append(arrsteps, v.Step)
				}
			}
		}

		if dialplan.Record != nil {
			for _, v := range dialplan.Record {
				if v.Step == strconv.Itoa(i+1) {
					arrsteps = append(arrsteps, v.Step)
				}
			}
		}
		if dialplan.Dial != nil {
			for _, v := range dialplan.Dial {
				if v.Step == strconv.Itoa(i+1) {
					arrsteps = append(arrsteps, v.Step)

				}
			}

		}
	}

	if duplicateInArray(arrsteps){
		return false
	}
	if len(arrsteps) != totalSteps {
		return false
	}
	return true

}

func validateNumber(number string) string {

	if len(number) == 11 {
		number = "+" + number
	}
	if !strings.Contains(number, "+1") {
		number = "+1" + number
	}

	return number

}


func duplicateInArray(arr []string) bool{
	visited := make(map[string]bool, 0)
	for i:=0; i<len(arr); i++{
		if visited[arr[i]] == true{
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
