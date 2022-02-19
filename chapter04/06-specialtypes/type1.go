package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	vg := &voteGame{
		students: []*student{
			&student{name: "01"},
			&student{name: "02"},
			&student{name: "03"},
			&student{name: "04"},
			&student{name: "05"},
		},
	}
	leader := vg.goRun()
	fmt.Println(*leader)
	leader.Distribute()
}

type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *Leader {
	//todo ....
	for _, item := range g.students {
		rand.Seed(time.Now().Unix())
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteD(g.students[randInt%len(g.students)])
		}
	}
	maxScore := -1
	maxScoreindex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreindex = i
		}
	}
	if maxScoreindex >= 0 {
		return (*Leader)(g.students[maxScoreindex])
	}
	return nil
}

////使用嵌套对象定义（继承）方式来定义班长
//type Leader struct {
//	student
//}
type Leader student

func (l *Leader) Distribute() {
	fmt.Println("发作业了")
}

type student struct {
	name     string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}
func (std *student) voteD(target *student) {
	target.disagree++
}
