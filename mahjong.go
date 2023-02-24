package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Mahjong struct {
	hand    int
	players [4]Player
	remain  []int
	sea     []int
}

type Player struct {
	hand  [17]int
	table []int
	see   [3*9 + 4 + 3]int // 數牌與字牌剩下的數量
}

func (m *Mahjong) nToChinese(n int) (s string) {
	if n < 0 || n >= 3*9*4+4*4+3*4+8 {
		return "？"
	} else if n < 9*4 { // 餅
		return string(rune('1'+n/4)) + "筒" // 1~9筒
	} else if n < 2*9*4 { // 條
		return string(rune('1'-9+n/4)) + "索" // 1~9索
	} else if n < 3*9*4 { // 萬
		return string(rune('1'-2*9+n/4)) + "萬" // 1~9萬
	} else if n < 4*4+3*9*4 { // 風
		return []string{"東", "南", "西", "北"}[-3*9+n/4] // 東南西北
	} else if n < 3*4+4*4+3*9*4 { // 龍
		return []string{"中", "發", "白"}[-4-3*9+n/4] // 中發白 "*$@"
	} else { // 花
		return []string{"春", "夏", "秋", "冬", "梅", "蘭", "菊", "竹"}[n-3*4-4*4-3*9*4] // 春夏秋冬 "uxqo", 梅蘭菊竹 "mljz"
	}
}

func (m *Mahjong) deal1() (n int) {
	tile := -1
	if len(m.remain) > 0 { // 有牌可發
		tile, m.remain = m.remain[0], m.remain[1:]
	}
	return tile
}

func (m *Mahjong) initDeal() {
	m.remain = rand.Perm(3*9*4 + 4*4 + 3*4 + 8)
	for i := 0; i < m.hand; i++ {
		for j := 0; j < 4; j++ {
			m.players[j].hand[i] = m.deal1()
		}
	}
}

func (m *Mahjong) showBonus() { // 補花
	for player := 0; player < 4; player++ {
		p := &(m.players[player])
		fmt.Printf("\n%d", player)
		for i := 0; i < m.hand; i++ {
			fmt.Printf(" %s", m.nToChinese(p.hand[i]))
			m.iShowBonus(p, i)
		}
		// m.initSee(p) // 可摸的牌不含手上的牌
	}
}

func (m *Mahjong) iShowBonus(p *Player, i int) {
	n := p.hand[i]
	for n >= 3*4+4*4+3*9*4 { // 花
		p.table = append(p.table, n)
		if n = m.deal1(); n < 0 {
			return
		}
		fmt.Printf("補 %s", m.nToChinese(n))
	}
	p.hand[i] = n
}

func (m *Mahjong) initSee(p *Player) {
	for _, t := range p.hand[:m.hand] { // 不含將打出去的牌
		p.addSee(t) // 可摸的牌不含手上的牌
	}
}

func (p *Player) addSee(t int) {
	p.see[t/4]++ // 可摸的牌不含剛打出的牌
}

func (m *Mahjong) decidePlay(p *Player) (n int) {
	return rand.Intn(m.hand + 1) // 隨機選一張牌
}

func (p *Player) play(n int, hand int) {
	p.hand[n], p.hand[hand] = p.hand[hand], p.hand[n] // 打出第n張牌, 換入摸到的牌
}

func (m *Mahjong) sort(tiles [17]int) (s []int) {
	s = append(tiles[:len(tiles)-1], tiles[len(tiles)-1])
	sort.Ints(s)
	return s
}

func (m *Mahjong) isWin(p *Player) (win bool) { //是否胡牌
	sortedHand, pairs := m.sort(p.hand), []int{}
	suited, honor := m.findPair(sortedHand, &pairs) // 找到眼及牌點與字的分佈

	for _, n := range sortedHand {
		fmt.Printf(" %s", m.nToChinese(n))
	}
	for _, n := range p.table {
		fmt.Printf("|%s", m.nToChinese(n))
	}

	for _, a := range pairs {
		n := sortedHand[a]
		if n < 3*4*9 { // 數牌
			suited[n/4] -= 2
			if isHonor(honor) && m.isSuit(suited) { // 去掉眼和字牌，找數牌胡牌型
				return true
			}
			suited[n/4] += 2
		} else { // 字牌
			honor[n/4-3*9] -= 2
			if isHonor(honor) && m.isSuit(suited) { // 去掉眼和字牌，找數牌胡牌型
				return true
			}
			honor[n/4-3*9] += 2
		}
	}
	return false
}

func isHonor(honor [4 + 3]int) bool {
	for _, count := range honor {
		if count > 0 && count != 3 {
			return false
		}
	}
	return true
}

func (m *Mahjong) isSuit(suited [3 * 9]int) bool { // 定理一 https://www.thenewslens.com/article/100657
	count := 0
	for t := 0; t < 3; t++ { // 餅, 條, 萬
		for i := t * 9; i < t*9+9; i++ {
			n := suited[i]
			count += n
			if n < 0 {
				return false
			} else if n == 0 {
				continue
			} else if n >= 3 {
				suited[i] -= 3
				return m.isSuit(suited)
			} else if i > t*9+9-3 || suited[i+1] < 1 || suited[i+2] < 1 {
				return false
			} else {
				suited[i]--
				suited[i+1]--
				suited[i+2]--
				return m.isSuit(suited)
			}
		}
	}
	if count == 0 {
		return true
	}
	return false
}

func (m *Mahjong) findPair(s []int, pairs *[]int) (suited [3 * 9]int, honor [4 + 3]int) { // 定理二 https://www.thenewslens.com/article/100657
	i, j := 0, 0
	for ; i < len(s) && s[i] < 4*9; i++ {
		suited[s[i]/4]++
	} // 餅
	m.findSuitPair(0, s[:i], pairs)
	for j = i; i < len(s) && s[i] < 2*4*9; i++ {
		suited[s[i]/4]++
	} // 條
	m.findSuitPair(j, s[j:i], pairs)
	for j = i; i < len(s) && s[i] < 3*4*9; i++ {
		suited[s[i]/4]++
	} // 萬
	m.findSuitPair(j, s[j:i], pairs)

	if i >= len(s) || s[i]/4 >= 3+4+3*9 {
		return suited, honor
	}
	j = i
	for _, n := range s[j:] {
		honor[n/4-3*9]++
	} // 字
	m.findHonorPair(j, s[j:], pairs)
	return suited, honor
}

func (m *Mahjong) findHonorPair(pad int, s []int, pairs *[]int) { // 字的眼
	for i := range s {
		m.findiPair(pad, s, i, pairs)
	}
}

func (m *Mahjong) findiPair(pad int, s []int, i int, pairs *[]int) {
	j := i + 1
	if j >= len(s) || s[i]/4 != s[j]/4 || (i > 0 && s[i-1]/4 == s[i]/4) {
		return
	} // 每張牌有4個複製, 只取第一次看到的複製
	fmt.Printf(" 眼%s", m.nToChinese(s[i]))
	*pairs = append(*pairs, pad+i)
}

func (m *Mahjong) findSuitPair(pad int, s []int, pairs *[]int) { // 順的眼 // 定理二 https://www.thenewslens.com/article/100657
	// filterMin := t * 4 * 9 // filterMax := filterMin + 4*9
	bin := [3]int{0, 0, 0}
	indexBin := map[int][]int{}
	for i, t := range s {
		bin[(t/4)%3]++
		indexBin[(t/4)%3] = append(indexBin[(t/4)%3], i)
	}
	b := 2 // 不在第0堆也不在第1堆, 那必在第2堆
	if bin[0]%3 != bin[1]%3 {
		if bin[0]%3 != bin[2]%3 { // 眼在第0堆
			b = 0
		} else { // 眼在第1堆
			b = 1
		}
	}
	for _, i := range indexBin[b] {
		m.findiPair(pad, s, i, pairs)
	}
}

func main() {
	m := Mahjong{}
	m.hand = 16 // 十六張麻將
	m.initDeal()
	m.showBonus()

	fmt.Println()

	for player := 0; len(m.remain) > 0; player = (player + 1) % 4 {
		p := &m.players[player]
		p.hand[m.hand] = m.deal1()
		fmt.Printf("\n%d摸 %s", player, m.nToChinese(p.hand[m.hand]))
		m.iShowBonus(p, m.hand)
		//p.hand = [17]int{0, 1, 2, 4, 8, 12, 16, 20, 24, 28, 32, 33, 34, 35, 36, 40, 44}
		if len(m.remain) <= 0 {
			fmt.Printf("\n和局")
			sort.Ints(m.sea)
			fmt.Println(m.sea)
			for player = 0; player < 4; player++ {
				fmt.Println(m.players[player].see)
			}
			break
		} else if m.isWin(p) {
			fmt.Printf("\n%d胡", player)
			break
		}
		p.addSee(p.hand[m.hand])        // 記錄摸到的牌
		p.play(m.decidePlay(p), m.hand) // 將打出的牌與摸到的牌交換
		fmt.Printf("\n%d打 %s_", player, m.nToChinese(p.hand[m.hand]))
		m.sea = append(m.sea, p.hand[m.hand]) // 海底加上打出的牌
		for other := 1; other < 4; other++ {  // 其他三家記錄打出的牌
			(&m).players[(player+other)%4].addSee(p.hand[m.hand])
		}
		p.hand[m.hand] = -1 // 打出的牌移出玩家
	}
}

