/*
Author : Haoyuan Liu
Time   : 2018/6/5
*/
package model

import (
	"arthur/app/model/dyn"
	"arthur/app/model/stat"
	"arthur/utils/jsonutils"
	"arthur/utils/panicutils"

	"arthur/app/info/errors"
	"arthur/app/model/center"
)

//配置中心

type ProfileCenter struct {
	Achievement            []stat.Achievement
	AchievementDesc        []stat.AchievementDesc
	AdventureBoss          []stat.AdventureBoss
	AdventureCity          []stat.AdventureCity
	AnabasisChapter        []stat.AnabasisChapter
	ArenaBuffCost          []stat.ArenaBuffCost
	ArenaFirstBuffCost     []stat.ArenaFirstBuffCost
	ArenaLevelBuff         []stat.ArenaLevelBuff
	ArenaLevelRate         []stat.ArenaLevelRate
	ArenaNpc               []stat.ArenaNpc
	ArenaNpcHero           []stat.ArenaNpcHero
	ArenaReloadCost        []stat.ArenaReloadCost
	Award                  []stat.Award
	Beauty                 []stat.Beauty
	BeautyEventAward       []stat.BeautyEventAward
	BeautyFavor            []stat.BeautyFavor
	BeautyLv               []stat.BeautyLv
	BeautyLvChild          []stat.BeautyLvChild
	BeautySkills           []stat.BeautySkills
	BeautySkillsUpgrade    []stat.BeautySkillsUpgrade
	ChatReward             []stat.ChatReward
	CollectEvent           []stat.CollectEvent
	DailyActivity          []stat.DailyActivity
	DailyTask              []stat.DailyTask
	DiplomacyAchievement   []stat.DiplomacyAchievement
	DiplomacyLv            []stat.DiplomacyLv
	DiplomacyProposal      []stat.DiplomacyProposal
	DungeonChest           []stat.DungeonChest
	DungeonSingleBattle    []stat.DungeonSingleBattle
	DungeonSingleExplore   []stat.DungeonSingleExplore
	Fashion                []stat.Fashion
	GameSystem             []stat.GameSystem
	GameSystemUnlock       []stat.GameSystemUnlock
	Hero                   []stat.Hero
	HeroFaith              []stat.HeroFaith
	HeroFaithSkill         []stat.HeroFaithSkill
	HeroFight              []stat.HeroFight
	HeroLv                 []stat.HeroLv
	HeroPeerage            []stat.HeroPeerage
	HeroProf               []stat.HeroProf
	HeroStar               []stat.HeroStar
	HeroTalent             []stat.HeroTalent
	Item                   []stat.Item
	ItemCompound           []stat.ItemCompound
	KeyValue               []stat.KeyValue
	LvUpgrade              []stat.LvUpgrade
	MainTask               []stat.MainTask
	MarriageProposalLetter []stat.MarriageProposalLetter
	Miracle                []stat.Miracle
	MonthCardAward         []stat.MonthCardAward
	Phase                  []stat.AnabasisPhase
	ResBonus               []stat.ResBonus
	RoleTitle              []stat.RoleTitle
	RoleValue              []stat.RoleValue
	RoleVip                []stat.RoleVip
	SkillInfo              []stat.SkillSolo
	StoreItem              []stat.StoreItem
	StudentAmity           []stat.StudentAmity
	StudentAvatar          []stat.StudentAvatar
	StudentCloister        []stat.StudentCloister
	StudentEncourage       []stat.StudentEncourage
	StudentExam            []stat.StudentExam
	StudentMedal           []stat.StudentMedal
	StudentName            []stat.StudentName
	StudentSubject         []stat.StudentSubject
	StudentTalent          []stat.StudentTalent
	Sword                  []stat.Sword
	TextClientHans         []stat.TextClientHans
	TextTipsHans           []stat.TextTipsHans
	ValueEvent             []stat.ValueEvent
	VisitCountry           []stat.VisitCountry
	VisitCountryLv         []stat.VisitCountryLv
	VisitCountryNpc        []stat.VisitCountryNpc
	VisitCountryNpcEvent   []stat.VisitCountryNpcEvent
	VisitEntrustEvents     []stat.VisitEntrustEvents
	VisitEntrustType       []stat.VisitEntrustType
	VisitEvent             []stat.VisitEvent
	VisitFavor             []stat.VisitFavor
	VisitMainStory         []stat.VisitMainStory
	WishAward              []stat.WishAward
	WishChoice             []stat.WishChoice
}

var ProfCenter = ProfileCenter{}

var CharacterSet = []interface{}{
	&center.AppInfo{},
	&center.LoginRecords{},
	&center.LoginRecord{},
	&center.ModuleCtrl{},
	&center.ServerInfo{},
	&center.ServerName{},
	&dyn.Achievement{},
	&dyn.AdventureNonlinear{},
	&dyn.Adventure{},
	&dyn.AnabasisPhase{},
	&dyn.ArenaChallengeRecord{},
	&dyn.ArenaFightBuff{},
	&dyn.ArenaFightMemory{},
	&dyn.ArenaFightRecord{},
	&dyn.ArenaFightRound{},
	&dyn.ArenaFight{},
	&dyn.ArenaHeroRound{},
	&dyn.ArenaPerfectRecord{},
	&dyn.ArenaReloadHeroRecord{},
	&dyn.ArenaSoloChallenge{},
	&dyn.ArenaSoloDefence{},
	&dyn.ArenaSolo{},
	&dyn.ArenaStorePurchase{},
	&dyn.ArenaStore{},
	&dyn.Backpack{},
	&dyn.BeautyChild{},
	&dyn.Beauty{},
	&dyn.ChatReport{},
	&dyn.ChatRewardGet{},
	&dyn.ChatRewardSend{},
	&dyn.ChatShare{},
	&dyn.ChatShield{},
	&dyn.DailyActivity{},
	&dyn.DailyTask{},
	&dyn.DiplomacyAchievement{},
	&dyn.DiplomacyProposal{},
	&dyn.DiplomacyRole{},
	&dyn.DungeonChest{},
	&dyn.DungeonGlobal{},
	&dyn.DungeonSingleExplore{},
	&dyn.DungeonSingle{},
	&dyn.Fashion{},
	&dyn.HeroFaith{},
	&dyn.HeroFight{},
	&dyn.HeroTalent{},
	&dyn.Hero{},
	&dyn.KeyValue{},
	&dyn.MailAward{},
	&dyn.MailTemplate{},
	&dyn.MarriageHall{},
	&dyn.MarriageProposal{},
	&dyn.MarriageSeek{},
	&dyn.MarriageTemp{},
	&dyn.Miracle{},
	&dyn.MonthCard{},
	&dyn.RankValue{},
	&dyn.RankWorship{},
	&dyn.Rank{},
	&dyn.ResBonus{},
	&dyn.RoleInfo{},
	&dyn.RoleTitle{},
	&dyn.RoleValue{},
	&dyn.StoreBuyRecords{},
	&dyn.StudentCloister{},
	&dyn.StudentExamResult{},
	&dyn.StudentExam{},
	&dyn.StudentNew{},
	&dyn.VipAward{},
	&dyn.VisitCountryNpc{},
	&dyn.VisitCountry{},
	&dyn.VisitEvent{},
	&dyn.VisitRole{},
	&dyn.WishChoice{},
	&dyn.WishTotalAward{},
	&dyn.Wish{},
}

//获取KeyValue配置表值
func GetKeyValue(key string) json.Any {
	v, ok := KeyValue[key]
	panicutils.TrueOrPanic(ok, errors.ErrNoConf)
	return json.Get([]byte(v))
}

var KeyValue = make(map[string]string)
