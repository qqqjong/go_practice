package main

import "fmt"

type item struct {
	name string
	price int
	amount int
}

type buyer struct {
	point int
	shoppingBucket map[string]int
}

func newBuyer() *buyer {
	d := buyer{}
	d.point = 1000000
	d.shoppingBucket = map[string]int{}
	return &d
}

func buying(itm []item, byr *buyer, itmchoice int) {
	inputamount := 0
	fmt.Print("수량을 입력하세요: ")
	fmt.Scanln(&inputamount)
	fmt.Println()

	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	if inputamount <= 0 {
		panic("올바른 수량을 입력하세요")
	}
	if byr.point < itm[itmchoice-1].price * inputamount || itm[itmchoice-1].amount < inputamount {
		panic("주문이 불가능합니다")
	} else {
		for {
			buy := 0
			fmt.Println("1. 바로주문\n2. 장바구니에 담기")
			fmt.Print("실행할 기능을 입력하시오 : ")
			fmt.Scanln(&buy)
			fmt.Println()

			if buy == 1 {
				itm[itmchoice-1].amount -= inputamount
				byr.point -= itm[itmchoice-1].price * inputamount

				fmt.Println("상품이 주문 접수 되었습니다")
				break
			} else if buy == 2 {
				checkbucket := false

				for itms := range byr.shoppingBucket {
					if itms == itm[itmchoice-1].name {
						checkbucket = true
					}
				}

				if checkbucket == true {
					byr.shoppingBucket[itm[itmchoice-1].name] += inputamount
				} else {
					byr.shoppingBucket[itm[itmchoice-1].name] = inputamount
				}
				fmt.Println("상품이 장바구니에 추가되었습니다")
				break
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
			}
		}
	}
}

func requiredPoint(itm []item, byr *buyer) (canbuy bool) {
	bucketpoint := 0
	for index, val := range byr.shoppingBucket {
		for i:=0; i<len(itm); i++ {
			if itm[i].name == index {
				bucketpoint += itm[i].price * val
			}
		}
	}
	fmt.Println("필요 마일리지 : ", bucketpoint)
	fmt.Println("보유 마일리지 : ", byr.point)
	fmt.Println()
	if byr.point < bucketpoint {
		fmt.Println("마일리지가 %d점 부족합니다", bucketpoint - byr.point)
		return false
	}
	return true
}

func emptyBucket(byr *buyer) {
	if len(byr.shoppingBucket) == 0 {
		fmt.Println("장바구니가 비었습니다")
	} else {
		for index, val := range byr.shoppingBucket {
			fmt.Printf("%s 수량: %d\n", index, val)
		}
	}
	fmt.Println()
}

func main() {
	items := make([]item, 5)
	buyer := newBuyer()

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"롱패딩", 500000, 20}
	items[2] = item{"투미 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menu := 0 // 첫 메뉴

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오 :")

		fmt.Scanln(&menu)
		fmt.Println()

		if menu == 1 { // 물건 구매
			for {
				itemchoice := 0

				for i:=0; i<len(items); i++ {
					fmt.Printf("물품%d: %s, 가격: %d, 잔여 수량: %d\n", i+1, items[i].name, items[i].price, items[i].amount)
				}
				fmt.Print("구매할 물품을 선택하세요 : ")
				fmt.Scanln(&itemchoice)
				fmt.Println()

				if itemchoice == 1 {
					buying()
					break
				} else if itemchoice == 2 {
					buying()
					break
				} else if itemchoice == 3 {
					buying()
					break
				} else if itemchoice == 4 {
					buying()
					break
				} else if itemchoice == 5 {
					buying()
					break
				} else {
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 2 { // 남은 수량 확인
			for i:=0; i<len(items); i++ {
				fmt.Printf("%s, 잔여수량: %d\n", items[i].name, items[i].amount)
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 3 { // 잔여 마일리지 확인
			fmt.Printf("현재 잔여 마일리지: %d\n", buyer.point)
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 4 { // 배송 상태 확인

			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 5 { // 장바구니 확인
			bucketmenu := 0

			for {
				emptyBucket(buyer)

				fmt.Println("1. 장바구니 상품 주문")
				fmt.Println("2. 메뉴로 돌아가기")
				fmt.Print("실행할 기능을 입력하시오")
				fmt.Scanln(&bucketmenu)
				fmt.Println()

				if bucketmenu == 1 {

				} else if bucketmenu == 2 {
					fmt.Println()
					break
				} else {
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
				}
			}
		} else if menu == 6 { // 프로그램 종료
			fmt.Println("프로그램을 종료합니다.")

			return // main함수 종료
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
		}
	}
}