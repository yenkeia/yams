package game

import (
	"fmt"
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/proto/server"
)

type magicContext struct {
	spell       cm.Spell
	target      attackTarget
	player      *player
	targetPoint cm.Point
}

func startMagic(ctx *magicContext) (targetID int, err error) {
	if ctx.target != nil {
		targetID = 0
		if ctx.target.isAttackTarget(ctx.player) {
			targetID = ctx.target.getObjectID()
		}
	}
	switch ctx.spell {
	case cm.SpellFireBall, // 火球术
		cm.SpellGreatFireBall, // 大火球
		cm.SpellFrostCrunch:   // 寒冰掌
		fireBall(ctx)
	case cm.SpellHealing: // 治愈术
		healing(ctx)
	case cm.SpellRepulsion, // 抗拒火环
		cm.SpellEnergyRepulsor: // 气功波
		repulsion(ctx)
	case cm.SpellElectricShock: // 诱惑之光
		electricShock(ctx)
	case cm.SpellPoisoning: // 施毒术
		poisoning(ctx)
	case cm.SpellHellFire: // 地狱火
		hellFire(ctx)
	case cm.SpellThunderBolt: // 雷电术
		thunderBolt(ctx)
	case cm.SpellSoulFireBall: //灵魂火符
		ok := soulFireBall(ctx)
		if !ok {
			targetID = 0
		}
	case cm.SpellSummonSkeleton: //召唤骷髅
		summonSkeleton(ctx)
	case cm.SpellTeleport: // 瞬息移动
		teleport(ctx)
	case cm.SpellHiding: // 隐身术
		hiding(ctx)
	case cm.SpellFury: // 血龙剑法
		fury(ctx)
	case cm.SpellFireBang, // 爆裂火焰
		cm.SpellIceStorm: // 冰咆哮
		fireBang(ctx)
	case cm.SpellMassHiding: // 集体隐身术
		massHiding(ctx)
	case cm.SpellSoulShield, // 幽灵盾
		cm.SpellBlessedArmour: // 神圣战甲术
		soulShield(ctx)
	case cm.SpellFireWall: // 火墙
		fireWall(ctx)
	case cm.SpellLightning: // 疾光电影
		lightning(ctx)
	case cm.SpellMassHealing: // 群体治疗术
		massHealing(ctx)
	case cm.SpellShoulderDash: // 野蛮冲撞
		shoulderDash(ctx)
	case cm.SpellThunderStorm, cm.SpellFlameField: // 地狱雷光/火龙气焰
		thunderStorm(ctx)
	case cm.SpellMagicShield: // 魔法盾
		magicShield(ctx)
	case cm.SpellFlameDisruptor: // 火龙术
		flameDisruptor(ctx)
	case cm.SpellTurnUndead: // 圣言术
		turnUndead(ctx)
	case cm.SpellMagicBooster: // 深延术
		magicBooster(ctx)
	case cm.SpellVampirism: // 嗜血术
		vampirism(ctx)
	case cm.SpellSummonShinsu: // 召唤神兽
		summonShinsu(ctx)
	case cm.SpellPurification: // 净化术
		purification(ctx)
	case cm.SpellLionRoar: // 狮子吼
		lionRoar(ctx)
	case cm.SpellRevelation: // 心灵启示
		revelation(ctx)
	case cm.SpellPoisonCloud: // 毒云
		poisonCloud(ctx)
	case cm.SpellEntrapment: // 捕绳剑
		entrapment(ctx)
	case cm.SpellBladeAvalanche: // 攻破斩
		bladeAvalanche(ctx)
	case cm.SpellSlashingBurst: // 日闪
		slashingBurst(ctx)
	case cm.SpellRage: // 剑气爆
		rage(ctx)
	case cm.SpellMirroring: // 分身术
		mirroring(ctx)
	case cm.SpellBlizzard: // 天霜冰环
		blizzard(ctx)
	case cm.SpellMeteorStrike: // 流星火雨
		meteorStrike(ctx)
	case cm.SpellIceThrust: // 冰焰术
		iceThrust(ctx)
	case cm.SpellProtectionField: // 护身气幕
		protectionField(ctx)
	case cm.SpellPetEnhancer: // 血龙水
		petEnhancer(ctx)
	case cm.SpellTrapHexagon: // 困魔咒
		trapHexagon(ctx)
	case cm.SpellReincarnation: // 复活术
		reincarnation(ctx)
	case cm.SpellCurse: // 诅咒术
		curse(ctx)
	case cm.SpellSummonHolyDeva: // 召唤月灵
		summonHolyDeva(ctx)
	case cm.SpellHallucination: // 迷魂术
		hallucination(ctx)
	case cm.SpellEnergyShield: // 阴阳盾
		energyShield(ctx)
	case cm.SpellUltimateEnhancer: // 无极真气
		ultimateEnhancer(ctx)
	case cm.SpellPlague: // 瘟疫
		plague(ctx)
	default:
		return 0, fmt.Errorf("技能暂未实现")
	}
	return
}

// 火球/大火球/寒冰掌
func fireBall(ctx *magicContext) {
	p := ctx.player
	t := ctx.target
	m := ctx.player.magics[ctx.spell]
	if t == nil || !t.isAttackTarget(p) {
		return
	}
	damage := m.getDamage(p.getAttackPower(p.minMC, p.maxMC))
	delay := cm.MaxDistance(p.location, t.getPosition())*50 + 500
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(delay)*time.Millisecond, func() {
		t.attacked(p, damage, cm.DefenceTypeMAC, false)
	})
}

// 治愈术
func healing(ctx *magicContext) {
	p := ctx.player
	t := ctx.target
	m := p.magics[ctx.spell]
	if t == nil || !t.isFriendlyTarget(p) {
		return
	}
	value := m.getDamage(p.getAttackPower(p.minSC, p.maxSC)*2) + p.level
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(500*time.Millisecond), func() {
		t.changeHP(value)
	})
}

// TODO 抗拒火环/气功波
func repulsion(ctx *magicContext) {

}

// TODO 诱惑之光
func electricShock(ctx *magicContext) {

}

// 施毒术
func poisoning(ctx *magicContext) {
	p := ctx.player
	t := ctx.target
	m := ctx.player.magics[ctx.spell]
	if t == nil || !t.isAttackTarget(p) {
		return
	}
	item := p.getPoison(1)
	if item == nil {
		return
	}
	value := m.getDamage(p.getAttackPower(p.minSC, p.maxSC))
	p.consumeItem(item, 1)
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(500*time.Millisecond), func() {
		tickSec := (value * 2) + ((m.level + 1) * 7) // 总共持续多少秒
		switch item.info.Shape {
		case 1:
			t.applyPoison(newPoison(
				time.Duration(2000*time.Millisecond), // 时间间隔2秒
				tickSec/2,                            // 持续多少秒 / 时间间隔2秒 = 要跳多少次
				p.objectID,
				cm.PoisonTypeGreen,
				value/15+m.level+1+cm.RandomInt(0, p.poisonAttack-1),
			), p)
		case 2:
			t.applyPoison(newPoison(
				time.Duration(2000*time.Millisecond), // 时间间隔2秒
				tickSec/2,                            // 持续多少秒 / 时间间隔2秒 = 要跳多少次
				p.objectID,
				cm.PoisonTypeRed,
				0,
			), p)
		}
	})
}

// 地狱火
func hellFire(ctx *magicContext) {
	p := ctx.player
	m := p.magics[ctx.spell]
	mp := env.maps[p.mapID]
	loc := p.location
	value := m.getDamage(p.getAttackPower(p.minMC, p.maxMC))
	mp.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(500*time.Millisecond), func() {
		dir := p.direction
		points := []cm.Point{}
		for i := 0; i < 4; i++ {
			point := cm.PointMove(loc, dir, i+1)
			points = append(points, point)
		}
		for _, point := range points {
			if !mp.validPoint(point) {
				return
			}
			c := mp.getCell(point)
			for it := c.objects.Front(); it != nil; it = it.Next() {
				if t, ok := it.Value.(attackTarget); ok {
					t.attacked(p, value, cm.DefenceTypeMAC, false)
				}
			}
		}
	})
}

// 雷电术
func thunderBolt(ctx *magicContext) {
	p := ctx.player
	t := ctx.target
	m := p.magics[ctx.spell]
	if t == nil || !t.isAttackTarget(p) {
		return
	}
	value := m.getDamage(p.getAttackPower(p.minMC, p.maxMC))
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(time.Millisecond*500), func() {
		t.attacked(p, value, cm.DefenceTypeMAC, false)
	})
}

// 灵魂火符
func soulFireBall(ctx *magicContext) bool {
	p := ctx.player
	t := ctx.target
	m := p.magics[ctx.spell]
	item := p.getAmulet(1)
	if item == nil {
		return false
	}
	if t == nil || !t.isAttackTarget(p) {
		return false
	}
	value := m.getDamage(p.getAttackPower(p.minSC, p.maxSC))
	delay := cm.MaxDistance(p.location, t.getPosition())*50 + 500
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(delay)*time.Millisecond, func() {
		t.attacked(p, value, cm.DefenceTypeMAC, false)
	})
	p.consumeItem(item, 1)
	return true
}

// TODO 召唤骷髅
func summonSkeleton(ctx *magicContext) {

}

// 瞬息移动
func teleport(ctx *magicContext) {
	p := ctx.player
	m := p.magics[ctx.spell]
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(200)*time.Millisecond, func() {
		bindMap := env.maps[p.bindMapID]
		bindLocation := p.bindLocation
		currentMap := env.maps[p.mapID]
		mapSizeX := bindMap.width / (m.level + 1)
		mapSizeY := bindMap.height / (m.level + 1)
		if currentMap.info.NoTeleport {
			p.receiveChat("无法在这里传送。", cm.ChatTypeSystem)
			return
		}
		for i := 0; i < 200; i++ {
			newLocation := cm.NewPoint(
				int(bindLocation.X)+cm.RandomInt(-mapSizeX, mapSizeX-1),
				int(bindLocation.Y)+cm.RandomInt(-mapSizeY, mapSizeY-1),
			)
			if p.teleport(bindMap, newLocation) {
				break
			}
		}
	})
}

// 隐身术
func hiding(ctx *magicContext) {
	p := ctx.player
	m := p.magics[ctx.spell]
	item := p.getAmulet(1)
	if item == nil {
		return
	}
	p.consumeItem(item, 1)
	value := p.getAttackPower(p.minSC, p.maxSC) + (m.level+1)*5
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(500*time.Millisecond), func() {
		if p.buffs.has(cm.BuffTypeHiding) {
			return
		}
		mp := env.maps[p.mapID]
		p.addBuff(newBuff(cm.BuffTypeHiding, p.objectID, mp.now.Add(time.Duration(value*1000)*time.Millisecond), []int{}))
		// LevelMagic(magic);
	})
}

// TODO 血龙剑法
func fury(ctx *magicContext) {

}

// TODO 爆裂火焰/冰咆哮
func fireBang(ctx *magicContext) {

}

// TODO 集体隐身术
func massHiding(ctx *magicContext) {

}

// TODO 幽灵盾/神圣战甲术
func soulShield(ctx *magicContext) {

}

// 火墙
func fireWall(ctx *magicContext) {
	p := ctx.player
	m := p.magics[ctx.spell]
	location := ctx.targetPoint
	value := m.getDamage(p.getAttackPower(p.minMC, p.maxMC))
	mp := env.maps[p.mapID]
	mp.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Millisecond*time.Duration(500), func() {
		dir := cm.MirDirectionUp
		points := []cm.Point{}
		for i := 0; i < 4; i++ {
			point := cm.PointMove(location, dir, 1)
			dir = dir.NextDirection().NextDirection()
			points = append(points, point)
		}
		points = append(points, location)
		for _, point := range points {
			if mp.validPoint(point) {
				c := mp.getCell(point)
				cast := true
				for it := c.objects.Front(); it != nil; it = it.Next() {
					if obj, ok := it.Value.(*spell); ok {
						if obj.spell == cm.SpellFireWall {
							cast = false
							break
						}
					}
				}
				if cast {
					expireTime := mp.now.Add(time.Millisecond * time.Duration((10+value/2)*1000))
					tickSpeed := time.Duration(2000 * time.Millisecond)
					s := newSpell(p.objectID, ctx.spell, value, mp.info.ID, point, tickSpeed, expireTime)
					mp.addObject(s)
					s.broadcastInfo()
				}
			}
		}
	})
}

// TODO 疾光电影
func lightning(ctx *magicContext) {}

// TODO 疾光电影
func massHealing(ctx *magicContext) {}

// TODO 野蛮冲撞
func shoulderDash(ctx *magicContext) {}

// TODO 地狱雷光/火龙气焰
func thunderStorm(ctx *magicContext) {
	// if (spell == Spell.FlameField)
	//     SpellTime = Envir.Time + 2499; //Spell Delay
	// if (spell == Spell.StormEscape)
	//     //Start teleport.
	//     ActionList.Add(new DelayedAction(DelayedType.Magic, Envir.Time + 749, magic, location));
	// break;
}

// 魔法盾
func magicShield(ctx *magicContext) {
	p := ctx.player
	m := p.magics[ctx.spell]
	value := m.getPower(p.getAttackPower(p.minMC, p.maxMC) + 15)
	p.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(500*time.Millisecond), func() {
		if p.hasMagicShield {
			return
		}
		p.hasMagicShield = true
		mp := env.maps[p.mapID]
		p.magicShieldExpireTime = mp.now.Add(time.Duration(value*1000) * time.Millisecond)
		p.magicShieldLevel = m.level
		mp.broadcast(p.location, &server.ObjectEffect{ObjectID: uint32(p.objectID), Effect: cm.SpellEffectMagicShieldUp}, 0)
		p.buffs.addBuff(newBuff(cm.BuffTypeMagicShield, p.objectID, p.magicShieldExpireTime, []int{p.magicShieldLevel}))
		// LevelMagic(magic);
	})
}

// TODO 火龙术
func flameDisruptor(ctx *magicContext) {}

// TODO 圣言术
func turnUndead(ctx *magicContext) {}

// TODO 深延术
func magicBooster(ctx *magicContext) {}

// TODO 嗜血术
func vampirism(ctx *magicContext) {}

// TODO 召唤神兽
func summonShinsu(ctx *magicContext) {}

// TODO 净化术
func purification(ctx *magicContext) {
	//     if (target == null)
	//     {
	//         target = this;
	//         targetID = ObjectID;
	//     }
}

// TODO 狮子吼
func lionRoar(ctx *magicContext) {
	//     CurrentMap.ActionList.Add(new DelayedAction(DelayedType.Magic, Envir.Time + 500, this, magic, CurrentLocation));
}

// TODO 心灵启示
func revelation(ctx *magicContext) {}

// TODO 毒云
func poisonCloud(ctx *magicContext) {}

// TODO 捕绳剑
func entrapment(ctx *magicContext) {}

// TODO 攻破斩
func bladeAvalanche(ctx *magicContext) {}

// TODO 日闪
func slashingBurst(ctx *magicContext) {}

// TODO 剑气爆
func rage(ctx *magicContext) {}

// TODO 分身术
func mirroring(ctx *magicContext) {}

// TODO 天霜冰环
func blizzard(ctx *magicContext) {}

// TODO 流星火雨
func meteorStrike(ctx *magicContext) {}

// TODO 冰焰术
func iceThrust(ctx *magicContext) {}

// TODO 护身气幕
func protectionField(ctx *magicContext) {}

// TODO 血龙水
func petEnhancer(ctx *magicContext) {}

// TODO 困魔咒
func trapHexagon(ctx *magicContext) {}

// TODO 复活术
func reincarnation(ctx *magicContext) {}

// TODO 诅咒术
func curse(ctx *magicContext) {}

// TODO 召唤月灵
func summonHolyDeva(ctx *magicContext) {}

// TODO 迷魂术
func hallucination(ctx *magicContext) {}

// TODO 阴阳盾
func energyShield(ctx *magicContext) {}

// TODO 无极真气
func ultimateEnhancer(ctx *magicContext) {}

// TODO 瘟疫
func plague(ctx *magicContext) {}
