package report

import (
	"fmt"
	"sort"
	"time"
)

type Rows []Row
type Row struct {
	Time time.Time
	In   bool
	User string
}

type dailyHour struct {
	in  time.Time
	out time.Time
	hours float64
}

type dailyReport struct {
	day        string
	userReport map[string]*dailyHour
}

type report map[string]dailyReport

func (rows Rows) NewReport() {
	report := make(report)
	var absDate string
	for _, row := range rows {
		absDate = row.Time.Format("02/01/2006")
		if dReport, ok := report[absDate]; ok {
			if uReport, ok := dReport.userReport[row.User]; ok {
				uReport.compareDate(row)
			} else {
				dReport.userReport[row.User] = row.newDailyHour()
			}
		} else {
			report[absDate] = dailyReport{
				day:        absDate,
				userReport: make(map[string]*dailyHour),
			}
		}
	}
	report.print()
}

func (r Row) newDailyHour() *dailyHour {
	if r.In {
		return &dailyHour{
			in:  r.Time,
			out: time.Date(1970,1,1,0,0,0,0,time.UTC),
		}
	}else {
		return &dailyHour{
			in:  time.Time{},
			out: r.Time,
		}
	}
}

func (dh *dailyHour) compareDate(row Row){
	if row.In && dh.in.After(row.Time) {
		dh.in = row.Time
	} else if !row.In && dh.out.Before(row.Time){
		dh.out = row.Time
	}

	dh.hours = dh.out.Sub(dh.in).Hours()
}

func (r report) print(){
	dates := make([]string, 0, len(r))
	for date := range r {
		dates = append(dates, date)
	}
	sort.Strings(dates) //sort by key
	for _, date := range dates {
		fmt.Println(r[date].day)
		for user,day := range r[date].userReport {
			fmt.Println(fmt.Sprintf("\t%s : %f",user,day.hours))
		}
	}
}
