package game

import (
	"fmt"
	"time"

	"github.com/yenkeia/yams/game/cm"
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
		soulFireBall(ctx)
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
	case cm.SpellFireWall:
		fireWall(ctx)
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

// TODO 灵魂火符
func soulFireBall(ctx *magicContext) {

}

// TODO 召唤骷髅
func summonSkeleton(ctx *magicContext) {

}

// TODO 瞬息移动
func teleport(ctx *magicContext) {

}

// TODO 隐身术
func hiding(ctx *magicContext) {

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
