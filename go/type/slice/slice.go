// Time        : 2019/01/18
// Description :

package slice

import (
	"fmt"
	"sort"
)

func sliceBase() {
	s := make([]int, 5, 10)
	s1 := s[5:10]
	s = append(s, 20, 21, 22)
	s1[0] = 10
	s1[1] = 11
	fmt.Printf("s  = %v,len(s)  = %d,cap(s)  = %d,ptr(s)  = %p\n", s, len(s), cap(s), s)
	fmt.Printf("s1 = %v,len(s1) = %d,cap(s1) = %d,ptr(s1) = %p\n", s1, len(s1), cap(s1), s1)
}

func sliceAppend() {
	sBase := make([]int, 0, 5)
	printInfo(sBase)
	fmt.Println("append 3 integers:1-3")
	sBase = append(sBase, 1, 2, 3)
	printInfo(sBase)
	fmt.Println("append 3 integers:4-6")
	sBase = append(sBase, 4, 5, 6)
	printInfo(sBase)
	fmt.Println("append 4 integers:7-10")
	sBase = append(sBase, 7, 8, 9, 10)
	printInfo(sBase)
	fmt.Println("append 13 integers:11-23")
	sBase = append(sBase, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23)
	printInfo(sBase)
}

func printInfo(s []int) {
	fmt.Printf("%v\tlen:%d\tcap:%d\tptr:%p\n", s, len(s), cap(s), s)
}

func insert(list []int, value, pos int) []int {
	list = append(list[:pos-1], append([]int{value}, list[pos-1:]...)...)
	return list
}

func sortSlice() {
	sorter := &RankSort{list: list}
	sort.Sort(sorter)
	list = sorter.list
}

type RankInfo struct {
	Id         string // 联盟id
	ShowId     int    // showId
	Ranking    int    // 排名
	FlagNo     int16  // 旗帜
	Name       string // 名字
	Lv         int16  // 等级
	LeaderName string // 盟主
	Count      int16  // 人数
	Power      int64  // 国力
	HasApplied bool   // 是否申请过
}

type RankSort struct {
	list []RankInfo
}

func (r *RankSort) Len() int {
	return len(r.list)
}

func (r *RankSort) Less(i, j int) bool {
	return r.list[i].Power > r.list[j].Power ||
		r.list[i].Power == r.list[j].Power && r.list[i].Lv > r.list[j].Lv
}

func (r *RankSort) Swap(i, j int) {
	r.list[i], r.list[j] = r.list[j], r.list[i]
	r.list[i].Ranking, r.list[j].Ranking = r.list[j].Ranking, r.list[i].Ranking
}

var list = []RankInfo{
	{"w548yvcyksqr", 8783834, 1, 6, "啦啦德玛西亚", 10, "", 3, 1788157228457, false},
	{"11xoiqf3p0z4", 2532116, 2, 8, "谢广坤粉丝团", 10, "", 28, 27985967041, false},
	{"w4u4lej6a6au", 6241789, 3, 8, "联盟联盟", 1, "", 1, 4866851770, false},
	{"w5wqwrvegsyn", 5562813, 4, 8, "Memory", 10, "", 5, 3318434376, false},
	{"w54bhbjj1rho", 8704615, 5, 6, "MWMWMWMWMWMWMWMW", 1, "", 1, 750289723, false},
	{"w5b3vb0k8ui8", 9500053, 6, 8, "848484", 1, "", 1, 732999956, false},
	{"2fr8keinlu44", 7765443, 7, 1, "111111111", 1, "", 1, 346076635, false},
	{"w53qmx4hvlmu", 3043990, 8, 8, "斋藤家的环奈", 1, "", 4, 275150624, false},
	{"w4pieesy52xd", 9806812, 9, 8, "99999", 2, "", 1, 247838048, false},
	{"w44ip4kdnmrn", 9362715, 10, 8, "蔡徐坤律师团", 1, "", 1, 204810234, false},
	{"2p4evzre35hr", 7318071, 11, 6, "1111111111111111", 5, "", 22, 129484911, false},
	{"7qbvc1wtisso", 7698959, 12, 8, "gdfg", 2, "", 10, 110026235, false},
	{"w8b1kpnpiobk", 7391680, 13, 7, "hei", 1, "", 1, 75189254, false},
	{"w6prwxw0caii", 8340388, 14, 7, "dsafad", 1, "", 1, 60035755, false},
	{"w4j6e5rl9ts2", 9616055, 15, 5, "谢广坤律师团", 1, "", 1, 15060528, false},
	{"w6w2glqinwvh", 4375668, 16, 7, "zero", 1, "", 1, 7439044, false},
	{"w6wuj4rtul8l", 3134058, 17, 7, "SKY", 10, "", 1, 5015824, false},
	{"w6lcr69wvbie", 8003034, 18, 7, " 联盟22233333", 10, "", 1, 4963040, false},
	{"w8lfqtjpmmfi", 3485346, 19, 7, " 1111", 1, "", 1, 3399325, false},
	{"w874tjw0u0xo", 5205439, 20, 7, " 你好阿", 1, "", 1, 2742954, false},
	{"g1lkg5wg3n1", 6313394, 21, 8, " 1234561z", 2, "", 2, 1334892, false},
	{"w5x3pgjtar3z", 4324170, 22, 8, " 955959", 1, "", 1, 1027843, false},
	{"w4trbrfi8cpe", 4399594, 23, 8, " 命啊", 1, "", 1, 775917, false},
	{"w4t3u1etalfv", 4064225, 24, 8, " 大", 1, "", 1, 673337, false},
	{"w6prwl52zvgo", 9715584, 25, 6, " 测试联盟		", 1, "", 1, 548608, false},
	{"w6oxd94w6rdb", 3886789, 26, 7, " 111111111", 1, "", 1, 383247, false},
	{"16w9q1inrwxd", 9102810, 27, 1, " 111 		", 1, "", 1, 347159, false},
	{"3583mkndkd6o", 3531324, 28, 8, " 112 		", 2, "", 1, 325495, false},
	{"w548yvd12p0e", 6310582, 29, 8, " 4399联盟 	", 1, "", 3, 320654, false},
	{"w31c3t1pe8ph", 1400419, 30, 8, " zoeunion1", 1, "", 1, 303215, false},
	{"w4mfz4tca3gh", 3811212, 31, 2, " abc 		", 1, "", 1, 276156, false},
	{"401mmbs6lgqi", 8529558, 32, 8, " 999999 	", 2, "", 1, 210405, false},
	{"w5boussol43l", 5791522, 33, 1, " MWWMWWMWW", 1, "", 1, 156921, false},
	{"w4mygqmr29ls", 3999992, 34, 7, " emmmmm 	", 1, "", 2, 102485, false},
	{"2clvw9qi2len", 6359932, 35, 8, " 银河护卫 	", 1, "", 1, 57724, false},
	{"w6ichb9y34nc", 2574528, 36, 7, " aaaaaaaaa", 1, "", 1, 38648, false},
	{"w28sq0vv33q0", 6399135, 37, 8, " zpyAllian", 1, "", 10, 33864, false},
	{"w8axxrsp6mh1", 3724352, 38, 7, " 艾欧尼亚 	", 1, "", 1, 33172, false},
	{"w5anx4fp4w76", 7772617, 39, 8, " advance 	", 1, "", 1, 24521, false},
	{"w6opb5pud5pb", 8492965, 40, 7, " aaaassd 	", 1, "", 1, 18203, false},
	{"w5izpu9u26mu", 3203378, 41, 8, " 咸鱼666 	", 1, "", 2, 11667, false},
	{"w54eqnsedj5x", 5107062, 42, 8, " WWMWWMWWM", 1, "", 1, 9088, false},
	{"w5tptop5wvf4", 1037258, 43, 8, " AAA11 	", 1, "", 1, 5853, false},
	{"w5emxl2v6zgo", 3758017, 44, 8, " 55555 	", 1, "", 1, 5597, false},
	{"w4qj8buk3z5l", 9338872, 45, 8, " 前往前往 	", 1, "", 2, 5006, false},
	{"w54bhbjgjv2e", 6187515, 46, 8, " 咸鱼联盟 	", 1, "", 1, 4516, false},
	{"w7wgzs52644i", 8836349, 47, 7, " 鞍山所所所所所所 ", 1, "", 1, 3712, false},
	{"4py91uu9esov", 6016923, 48, 8, " 123123123", 1, "", 2, 1728, false},
	{"6r35ihre5116", 2944401, 49, 8, " 啾咪啾咪 	", 1, "", 1, 1288, false},
	{"8f4fdchqj0nr", 2214926, 50, 8, " 阿达 		", 2, "", 1, 1226, false},
	{"w5mgu95u5vfy", 3528677, 51, 8, " 9595959 	", 1, "", 1, 1050, false},
	{"w6pa4xlaiz3x", 9018114, 52, 7, " s 		", 1, "", 1, 960, false},
	{"w64h410zfsbw", 1461039, 53, 7, " 0213 	", 1, "", 1, 928, false},
	{"w4swh5oynv9q", 8165463, 54, 3, " 46516518 ", 1, "", 2, 800, false},
	{"4mzmy6muozp5", 9750998, 55, 8, " 321 		", 1, "", 1, 764, false},
	{"8m6xlk2f33mj", 8089997, 56, 8, " 21 		", 1, "", 1, 730, false},
	{"w5b3vnc87llt", 5445899, 57, 8, " 121 		", 1, "", 1, 610, false},
	{"w4t8vjlssj76", 8304117, 58, 8, " 123456 	", 1, "", 1, 610, false},
	{"w63jog067s3l", 8941384, 59, 7, " linfeng66", 1, "", 1, 580, false},
	{"w7m14b16zeva", 3229624, 60, 7, " 345 		", 1, "", 1, 510, false},
	{"w5i1khboquc2", 9409347, 61, 8, " 123456啊啊 ", 1, "", 1, 400, false},
	{"w54md357mus2", 6457578, 62, 8, " demo111 	", 1, "", 1, 400, false},
	{"w54mcov7j1cn", 1051173, 63, 8, " 4399联盟2部 ", 1, "", 1, 400, false},
	{"w3uc5v4gpzjb", 1100896, 64, 8, " 1111111 	", 2, "", 1, 400, false},
	{"nszabqkzssf", 3769774, 65, 2, " AAA 		", 2, "", 1, 200, false},
	{"5zc14j8nd7ob", 5771804, 66, 8, " 反牟秋波联盟 	", 1, "", 1, 0, false},
}
