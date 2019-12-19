package models

import (
	"time"
)

type Payment struct {
	Worker Worker  `json:"worker"`
	TaskWorkers []TaskWorker  `json:"task_workers"`
	TotalAll float64 `json:"total_all"`
	Total float64 `json:"total"`
	Paid bool `json:"paid"`

}

func GetPayments(start time.Time,end time.Time) []Payment {

	var ps []Payment
	tws:=GetTaskWorkersBetweenDates(start,end)

	var ws []Worker
	for _,tw:=range tws{
		w:= GetWorkerById(tw.WorkerId)
		if w != nil{
			ws=appendIfMissing(ws,*w)
		}
	}

	for _,w:=range ws{
		p:=Payment{}
		p.Worker = w
		for _,tw:=range tws{
			if tw.WorkerId == w.Id{
				p.TaskWorkers = append(p.TaskWorkers,tw)
			}
		}
		p.TotalAll,p.Total,p.Paid = totalPaid(p.TaskWorkers)
		ps= append(ps,p)
	}

	return  ps

}

func GetDetailPayments(workerId int64,start time.Time,end time.Time) []TaskWorker {
	return GetTaskWorkersBetweenDatesByWorkerId(workerId,start,end)
}

func appendIfMissing(slice []Worker, i Worker) []Worker {
	for _, ele := range slice {
		if ele.Id == i.Id {
			return slice
		}
	}
	return append(slice, i)
}

func totalPaid(tws []TaskWorker) (float64,float64,bool){

	var t float64
	var tAll float64
	p:= false
	existFalsePaid:= false
	for _,tw:=range tws{
		tAll= tAll + tw.TotalPrice
		if !tw.Paid {
			existFalsePaid = true
			t= t + tw.TotalPrice
		}
	}

	p = !existFalsePaid

	return tAll,t,p
}