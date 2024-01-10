## 安装
首先下载编译，确保本地安装有Golang环境

```xshell
git clone git@github.com:yangchen5710/fish_client.git
cd fish_client
go build fish_client
```

接下来在终端下执行可执行文件

* windows

  ```xshell
  client.exe 47.101.212.202 8096
  ```

* linux/macos

  ```xshell
  ./client 47.101.212.202 8096
  ```
  
## 玩法
![demo](http://img.goatup.cn/brand/img/demo.gif)

## 出牌规则
所有牌型：
```
┌──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐──┐
│3 |4 |5 |6 |7 |8 |9 |10|J |Q |K |A |2 |S |X |
│♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |♦ |  |  |
└──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘──┘
```

- 王炸： ``sx``
- 顺子： ``34567``
- 三带一： ``3334``
- 三带二： ``33344``
- 飞机： ``333444a2``
- 单张10： ``t``
- 单张A： ``a``
- 封顶顺子： ``3456789tjqka``
- 不想出牌： ``pass``

## 算法实现
首先我们先定义相关结构和接口以及相关常量
```go
 const (
	Rocket            = 12
	Bomb              = 11
	Pair              = 10
	Three             = 9
	ThreeWithOne      = 8
	ThreeWithTwo      = 7
	SingleStraight    = 6
	DoubleStraight    = 5
	ThreeStraight     = 4
	AircraftWithWings = 3
	FourWithTwo       = 2
	Single            = 1
)
 
 
 type Poker struct {
	Level string //牌的数值
	Type  string //牌的类型
}

type PokerMath struct {
	PokerType   int
	MinValue    int
	PokerLength int
}

type PokerAlgorithm interface {
	IsRocket(Poker, Poker) (bool, PokerMath)       //火箭 即大小王
	IsBomb([]Poker) (bool, PokerMath)              //炸弹 四张同数值牌
	IsPair([]Poker) (bool, PokerMath)              //对牌 数值相同的两张牌
	IsThree([]Poker) (bool, PokerMath)             //三张牌 数值相同的三张牌
	IsThreeWithOne([]Poker) (bool, PokerMath)      //三带一 数值相同的三张牌 + 一张单牌
	IsThreeWithTwo([]Poker) (bool, PokerMath)      //三带二 数值相同的三张牌 + 一张对牌
	IsSingleStraight([]Poker) (bool, PokerMath)    //单顺 五张或更多的连续单牌（如：45678 或 789tJQK）。不包括 2 点和双王
	IsDoubleStraight([]Poker) (bool, PokerMath)    //双顺 三对或更多的连续对牌（如：334455 、77 88 99 1010 JJ）。不包括 2 点和双王
	IsThreeStraight([]Poker) (bool, PokerMath)     //三顺：二个或更多的连续三张牌（如：333444 、 555 666 777 888）。不包括 2 点和双王
	IsAircraftWithWings([]Poker) (bool, PokerMath) //飞机带翅膀：三顺+同数量的单牌（或同数量的对牌）
	IsFourWithTwo([]Poker) (bool, PokerMath)       //四带二：四张牌+两手牌（注意：四带二不是炸弹）。
	IsSingle(Poker) (bool, PokerMath)              //单牌
	ComparePokers(PokerMath, PokerMath) bool       //比较牌型和大小
}
```
然后去实现相关的接口算法
```go
func (pm PokerMath) IsRocket(poker1, poker2 Poker) (boolean bool, pokerMath PokerMath) {
	boolean = (poker1.Level == "LEVEL_SMALL_KING" && poker2.Level == "LEVEL_BIG_KING") ||
		(poker1.Level == "LEVEL_BIG_KING" && poker2.Level == "LEVEL_SMALL_KING")
	if boolean {
		pokerMath = PokerMath{PokerType: Rocket, MinValue: 0, PokerLength: 2}
	}
	return
}

func (pm PokerMath) IsBomb(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 4 {
		return
	}

	valueMap := make(map[string]int)
	for _, poker := range pokers {
		valueMap[poker.Level]++
	}

	for pokerLevel, count := range valueMap {
		if count == 4 {
			boolean = true
			pokerMath = PokerMath{PokerType: Bomb, MinValue: convertToValueInit(pokerLevel), PokerLength: 4}
			return
		}
	}
	return
}

func (pm PokerMath) IsPair(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 2 {
		return
	}

	valueMap := make(map[string]int)
	for _, poker := range pokers {
		valueMap[poker.Level]++
	}

	for pokerLevel, count := range valueMap {
		if count == 2 {
			boolean = true
			pokerMath = PokerMath{PokerType: Pair, MinValue: convertToValueInit(pokerLevel), PokerLength: 2}
			return
		}
	}

	return
}

func (pm PokerMath) IsThree(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 3 {
		return
	}

	valueMap := make(map[string]int)
	for _, poker := range pokers {
		valueMap[poker.Level]++
	}

	for pokerLevel, count := range valueMap {
		if count == 3 {
			boolean = true
			pokerMath = PokerMath{PokerType: Three, MinValue: convertToValueInit(pokerLevel), PokerLength: 3}
			return
		}
	}

	return
}

func (pm PokerMath) IsThreeWithOne(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 4 {
		return
	}

	sort.Slice(pokers, func(i, j int) bool {
		return pokers[i].Level < pokers[j].Level
	})

	// 判断是否是三带一，即第一个三张牌点数相同，且第四张牌与其中一张牌的点数不同
	if pokers[0].Level == pokers[1].Level && pokers[1].Level == pokers[2].Level &&
		pokers[2].Level != pokers[3].Level {
		boolean = true
		pokerMath = PokerMath{PokerType: ThreeWithOne, MinValue: convertToValueInit(pokers[0].Level), PokerLength: 4}
		return
	}

	// 也可以判断最后三张牌是否相同点数
	if pokers[1].Level == pokers[2].Level && pokers[2].Level == pokers[3].Level &&
		pokers[0].Level != pokers[1].Level {
		boolean = true
		pokerMath = PokerMath{PokerType: ThreeWithOne, MinValue: convertToValueInit(pokers[1].Level), PokerLength: 4}
		return
	}

	return
}

func (pm PokerMath) IsThreeWithTwo(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 5 {
		return
	}

	sort.Slice(pokers, func(i, j int) bool {
		return pokers[i].Level < pokers[j].Level
	})

	// 判断是否是三带二
	if pokers[0].Level == pokers[1].Level && pokers[1].Level == pokers[2].Level &&
		pokers[3].Level == pokers[4].Level && pokers[2].Level != pokers[3].Level {
		boolean = true
		pokerMath = PokerMath{PokerType: ThreeWithTwo, MinValue: convertToValueInit(pokers[0].Level), PokerLength: 5}
		return
	}

	// 也可以判断三张和两张的位置对调
	if pokers[0].Level == pokers[1].Level && pokers[2].Level == pokers[3].Level &&
		pokers[3].Level == pokers[4].Level && pokers[1].Level != pokers[2].Level {
		boolean = true
		pokerMath = PokerMath{PokerType: ThreeWithTwo, MinValue: convertToValueInit(pokers[2].Level), PokerLength: 5}
		return
	}

	return
}

func (pm PokerMath) IsSingleStraight(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) < 5 {
		return
	}

	sort.Slice(pokers, func(i, j int) bool {
		// 将牌的值转换为对应的数值进行比较
		val1 := convertToValue(pokers[i].Level)
		val2 := convertToValue(pokers[j].Level)
		return val1 < val2
	})

	// 检查牌是否连续
	for i := 0; i < len(pokers)-1; i++ {
		if convertToValue(pokers[i+1].Level)-convertToValue(pokers[i].Level) != 1 {
			return // 如果有间断，不是单顺
		}
	}
	boolean = true
	pokerMath = PokerMath{PokerType: SingleStraight, MinValue: convertToValueInit(pokers[0].Level), PokerLength: len(pokers)}
	return
}

func (pm PokerMath) IsDoubleStraight(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers)%2 != 0 || len(pokers) < 6 {
		return // 牌数必须是偶数且至少6张才可能是双顺
	}

	// 对牌进行排序
	sort.Slice(pokers, func(i, j int) bool {
		// 将牌的值转换为对应的数值进行比较
		val1 := convertToValue(pokers[i].Level)
		val2 := convertToValue(pokers[j].Level)
		return val1 < val2
	})

	fmt.Println(pokers)

	// 检查牌是否连续成对
	for i := 0; i < len(pokers)-1; i += 2 {
		if convertToValue(pokers[i].Level) != convertToValue(pokers[i+1].Level) ||
			(i > 0 && convertToValue(pokers[i].Level)-convertToValue(pokers[i-1].Level) != 1) {
			return // 如果不是成对连续，不是双顺
		}
	}

	boolean = true
	pokerMath = PokerMath{PokerType: DoubleStraight, MinValue: convertToValueInit(pokers[0].Level), PokerLength: len(pokers)}
	return
}

func (pm PokerMath) IsThreeStraight(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers)%3 != 0 || len(pokers) < 6 {
		return // 牌数必须是3的倍数且至少6张才可能是三顺
	}

	// 对牌进行排序
	sort.Slice(pokers, func(i, j int) bool {
		// 将牌的值转换为对应的数值进行比较
		val1 := convertToValue(pokers[i].Level)
		val2 := convertToValue(pokers[j].Level)
		return val1 < val2
	})

	// 检查牌是否连续三张
	for i := 0; i < len(pokers)-2; i += 3 {
		if convertToValue(pokers[i].Level) != convertToValue(pokers[i+1].Level) ||
			convertToValue(pokers[i].Level) != convertToValue(pokers[i+2].Level) ||
			(i > 0 && convertToValue(pokers[i].Level)-convertToValue(pokers[i-1].Level) != 1) {
			return
		}
	}

	boolean = true
	pokerMath = PokerMath{PokerType: ThreeStraight, MinValue: convertToValueInit(pokers[0].Level), PokerLength: len(pokers)}
	return
}

func (pm PokerMath) IsAircraftWithWings(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers)%4 != 0 || len(pokers) < 8 {
		return // 牌数必须是4的倍数且至少8张才可能是飞机带翅膀
	}

	// 对牌进行排序
	sort.Slice(pokers, func(i, j int) bool {
		// 将牌的值转换为对应的数值进行比较
		val1 := convertToValue(pokers[i].Level)
		val2 := convertToValue(pokers[j].Level)
		return val1 < val2
	})

	// 计算连续三张的数量
	tripleCount := 0
	minValue := ""
	for i := 0; i < len(pokers)-2; i++ {
		if convertToValue(pokers[i].Level) == convertToValue(pokers[i+1].Level) &&
			convertToValue(pokers[i].Level) == convertToValue(pokers[i+2].Level) {
			tripleCount++
			if minValue == "" {
				minValue = pokers[i].Level
			}
		}
	}

	// 检查是否有连续三张的数量大于等于2
	if tripleCount < 2 {
		return // 如果连续三张的数量不足2组，不是飞机带翅膀
	}

	boolean = true
	pokerMath = PokerMath{PokerType: AircraftWithWings, MinValue: convertToValueInit(minValue), PokerLength: len(pokers)}
	return
}

func (pm PokerMath) IsFourWithTwo(pokers []Poker) (boolean bool, pokerMath PokerMath) {
	boolean = false
	if len(pokers) != 6 {
		return // 必须是六张牌才可能是四带二
	}

	// 对牌进行排序
	sort.Slice(pokers, func(i, j int) bool {
		// 将牌的值转换为对应的数值进行比较
		val1 := convertToValue(pokers[i].Level)
		val2 := convertToValue(pokers[j].Level)
		return val1 < val2
	})

	// 计算四张牌的数量
	fourCount := 0
	minValue := ""
	for i := 0; i < len(pokers)-3; i++ {
		if convertToValue(pokers[i].Level) == convertToValue(pokers[i+1].Level) &&
			convertToValue(pokers[i].Level) == convertToValue(pokers[i+2].Level) &&
			convertToValue(pokers[i].Level) == convertToValue(pokers[i+3].Level) {
			fourCount++
			if minValue == "" {
				minValue = pokers[i].Level
			}
		}
	}

	// 检查是否有四张牌
	if fourCount != 1 {
		return // 如果不是四张牌，不是四带二
	}

	// 计算两张其他牌的数量
	otherCount := 0
	for i := 0; i < len(pokers)-1; i++ {
		if convertToValue(pokers[i].Level) != convertToValue(pokers[i+1].Level) {
			otherCount++
		}
	}
	if otherCount != 1 {
		return // 如果不是两张其他牌，不是四带二
	}

	boolean = true
	pokerMath = PokerMath{PokerType: FourWithTwo, MinValue: convertToValueInit(minValue), PokerLength: len(pokers)}
	return
}

func (pm PokerMath) IsSingle(poker Poker) (boolean bool, pokerMath PokerMath) {
	boolean = true
	pokerMath = PokerMath{PokerType: Single, MinValue: convertToValueInit(poker.Level), PokerLength: 1}
	return
}

func (pm PokerMath) ComparePokers(roomPokerMath, playerPokerMath PokerMath) bool {
	roomPokerType := roomPokerMath.PokerType
	roomPokerLength := roomPokerMath.PokerLength
	roomMinValue := roomPokerMath.MinValue

	if roomPokerType == playerPokerMath.PokerType &&
		roomPokerLength == playerPokerMath.PokerLength &&
		roomMinValue < playerPokerMath.MinValue {
		return true
	}

	if roomPokerType != playerPokerMath.PokerType {
		switch roomPokerType {
		case Bomb:
			if playerPokerMath.PokerType == Rocket {
				return true
			}
		case Pair, Three, ThreeWithOne, ThreeWithTwo, SingleStraight, DoubleStraight, ThreeStraight, AircraftWithWings, FourWithTwo, Single:
			if playerPokerMath.PokerType == Rocket || playerPokerMath.PokerType == Bomb {
				return true
			}
		}
	}

	return false
}

func (pm PokerMath) ConvertToValueString(cardValue string) string {
	valueMap := map[string]string{
		"3": "LEVEL_3", "4": "LEVEL_4", "5": "LEVEL_5", "6": "LEVEL_6",
		"7": "LEVEL_7", "8": "LEVEL_8", "9": "LEVEL_9",
		"0": "LEVEL_10", "t": "LEVEL_10", "T": "LEVEL_10",
		"j": "LEVEL_J", "J": "LEVEL_J",
		"q": "LEVEL_Q", "Q": "LEVEL_Q",
		"k": "LEVEL_K", "K": "LEVEL_K",
		"a": "LEVEL_A", "A": "LEVEL_A", "1": "LEVEL_A",
		"2": "LEVEL_2",
		"s": "LEVEL_SMALL_KING", "S": "LEVEL_SMALL_KING",
		"x": "LEVEL_BIG_KING", "X": "LEVEL_BIG_KING",
	}

	return valueMap[cardValue]
}

// 辅助函数：将牌的值转换为数值
func convertToValueInit(cardValue string) int {
	valueMap := map[string]int{
		"LEVEL_3": 3, "LEVEL_4": 4, "LEVEL_5": 5, "LEVEL_6": 6, "LEVEL_7": 7,
		"LEVEL_8": 8, "LEVEL_9": 9, "LEVEL_10": 10, "LEVEL_J": 11, "LEVEL_Q": 12,
		"LEVEL_K": 13, "LEVEL_A": 14, "LEVEL_2": 15, "LEVEL_SMALL_KING": 16, "LEVEL_BIG_KING": 17,
	}

	return valueMap[cardValue]
}

func convertToValue(cardValue string) int {
	value := convertToValueInit(cardValue)
	if value == 15 || value == 16 || value == 17 {
		value = -1
	}
	return value
}
```
## 计划
* 支持PVE
* 支持个人得分