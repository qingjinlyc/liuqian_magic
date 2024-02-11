package main

import (
	"fmt"
	"math/rand"
	"time"
)

/***
	sex: 1 - 男， 2 - 女
 */
func transfer(cards []string, sex int, region string, nameLen int) {
	fmt.Println("cards -> ", cards)
	// 名字有多长，往下放几张
	cardsTmp := []string{}
	cardsTmp = append(cardsTmp, cards[nameLen:]...)
	cardsTmp = append(cardsTmp, cards[0 : nameLen]...)
	fmt.Printf("名字有%d个字，往下翻%d张\n", nameLen, nameLen)
	fmt.Println("cards -> ", cardsTmp)

	// 拿起最上面三张，插进中间任何一个位置
	// 生成随机位置
	randomInt := randomPos(3, 6)
	// 插入到该位置
	cardsTmp1 := []string{}
	cardsTmp1 = append(cardsTmp1, cardsTmp[3 : randomInt + 1]...)
	cardsTmp1 = append(cardsTmp1, cardsTmp[0: 3]...)
	cardsTmp1 = append(cardsTmp1, cardsTmp[randomInt + 1:]...)
	top := cardsTmp1[0]
	// 最上面的藏起来
	cardsTmp1 = cardsTmp1[1: ]
	fmt.Printf("拿起最上面三张，插进第%d张牌之后，藏起最上面的牌\n", randomInt + 1)
	fmt.Println("cards -> ", cardsTmp1)

	// 南方人拿一张，北方人拿两张，不确定拿三张
	pos := pickByRegin(region)
	randomInt = randomPos(pos + 1, 5)
	cardsTmp2 := []string{}
	// 插到该位置
	cardsTmp2 = append(cardsTmp2, cardsTmp1[pos + 1 : randomInt + 1]...)
	cardsTmp2 = append(cardsTmp2, cardsTmp1[0: pos + 1]...)
	cardsTmp2 = append(cardsTmp2, cardsTmp1[randomInt + 1:]...)
	fmt.Printf("南方人拿一张，北方人拿两张，不确定拿三张。实际拿了%d张，插到第%d张牌之后\n", pos + 1, randomInt + 1)
	fmt.Println("cards -> ", cardsTmp2)

	// 男的丢一张，女的丢两张
	cardsTmp2 = cardsTmp2[sex:]
	fmt.Printf("男的丢一张，女的丢两张。实际丢了%d张\n", sex)
	fmt.Println("cards -> ", cardsTmp2)

	// 见证奇迹的时刻
	fmt.Println("见证奇迹的时刻")
	for i := 0; i < 7; i++ {
		cardsTmp2 = append(cardsTmp2, cardsTmp2[0])
		cardsTmp2 = cardsTmp2[1:]
		fmt.Printf("第%d次\n", i + 1)
		fmt.Println("cards -> ", cardsTmp2)
	}

	// 每次拿一张下去，丢掉上面一张，直到只剩一张
	fmt.Println("每次拿一张下去，丢掉上面一张，直到只剩一张")
	for len(cardsTmp2) != 1 {
		cardsTmp2 = append(cardsTmp2, cardsTmp2[0])
		cardsTmp2 = cardsTmp2[2:]
		fmt.Println("cards -> ", cardsTmp2)
	}

	fmt.Println("最后剩下：")
	print(top, cardsTmp2[0])

}

func pickByRegin(region string) int{
	switch region {
	case "s":
		return 0
	case "n":
		return 1
	default:
		return 2

	}
}

func randomPos(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	// 生成随机位置
	randomInt := rand.Intn(max-min) + min

	return randomInt
}

func randCards() []string {
	cards := []string{"a", "b", "c", "d"}

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 打乱slice
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	res := []string{}
	for _, c := range cards {
		res = append(res, c + "1")
	}

	for _, c := range cards {
		res = append(res, c + "2")
	}

	return res
}
func main() {
	cards := randCards()
	transfer(cards, 1, "s", 3)
}
