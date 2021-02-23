package main

import "fmt"

type Command interface {
	Execute(*Calculator)
	Undo(*Calculator)
}

type PlusCommand struct {
	beforeVal int
	num       int
}

func (a *PlusCommand) Execute(calculator *Calculator) {
	a.beforeVal = calculator.val
	calculator.Add(a.num)
	fmt.Println(calculator.val)
}

func (a *PlusCommand) Undo(calculator *Calculator) {
	calculator.val = a.beforeVal
	fmt.Println(calculator.val)
}

type MinusCommand struct {
	beforeVal int
	num       int
}

func (m *MinusCommand) Execute(calculator *Calculator) {
	m.beforeVal = calculator.val
	calculator.Minus(m.num)
	fmt.Println(calculator.val)
}

func (m *MinusCommand) Undo(calculator *Calculator) {
	calculator.val = m.beforeVal
	fmt.Println(calculator.val)
}

type Calculator struct {
	val int
}
func (c *Calculator) Add(num int) {
	c.val += num
}
func (c *Calculator) Minus(num int) {
	c.val -= num
}

type CommandQueue struct {
	queue []Command
	actor *Calculator
}

func (p *CommandQueue) AddCommand(c Command) {
	// 명령을 큐에 저장하고
	p.queue = append(p.queue, c)
	// 실행합니다
	c.Execute(p.actor)
	// 길이가 10이 될경우 작업내역을 지웁니다
	if len(p.queue) == 10 {
		p.queue = make([]Command, 10)
	}
}

func (p *CommandQueue) RemoveCommand() {
	// 마지막 명령을 꺼내서 Undo를 호출합니다.
	lastCommand := p.queue[len(p.queue)- 1]
	lastCommand.Undo(p.actor)
	// 마지막 명령을 큐슬라이스에서 제거합니다.
	p.queue = p.queue[:len(p.queue)-1]
}

func main() {
	calculator := &Calculator{val:0} // 0으로 시작하는 계산기
	queue := CommandQueue{actor: calculator} // 계산기를 액터로 지정후 커맨드큐생성

	queue.AddCommand(CreatePlusCommand(3)) // +3
	queue.AddCommand(CreatePlusCommand(3)) // +3
	queue.AddCommand(CreatePlusCommand(3)) // +3
	queue.AddCommand(CreatePlusCommand(30)) // +3
	//queue.AddCommand(CreateMinusCommand(3)) // -3
	//queue.AddCommand(CreateMinusCommand(3)) // -3
	//queue.RemoveCommand() // 돌려돌려 되돌려줘
	//queue.RemoveCommand() // 돌려돌려 되돌려줘...
}

func CreatePlusCommand(num int) Command {
	plus := new(PlusCommand)
	plus.beforeVal = plus.num
	plus.num = num
	return plus
}
