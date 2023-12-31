package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
)

const rotationDeltaUnset = -999

type botState int

const (
	bstateNone botState = iota
	bstateIdle
	bstateRotateToWaypoint
	bstateFlyToWaypoint
	bstateFollow
	bstateRetreat
	bstateDefense
)

type boss1player struct {
	world  *battleState
	vessel *vesselNode
	scene  *ge.Scene

	bstate botState

	stateTicker float64

	strafeDelay     float64
	rotationDelay   float64
	rotationDelta   float64
	strafeLeftTime  float64
	strafeRightTime float64
	canStrafe       bool

	directHits      int
	followCancelled bool

	waypoint gmath.Vec
}

func newBoss1Player(world *battleState, vessel *vesselNode) *boss1player {
	return &boss1player{
		world:  world,
		vessel: vessel,
	}
}

func (p *boss1player) Init(scene *ge.Scene) {
	p.scene = scene

	p.vessel.EventOnDamage.Connect(p, func(dmg gamedata.Damage) {
		if p.bstate == bstateFollow {
			p.directHits++
			if p.directHits >= 4 && p.scene.Rand().Chance(0.85) {
				p.setState(bstateNone)
				p.followCancelled = true
				return
			}
		}

		if p.bstate == bstateIdle && p.scene.Rand().Chance(0.4) {
			p.setState(bstateNone)
			return
		}

		if p.scene.Rand().Chance(0.02) {
			p.setState(bstateNone)
			return
		}
	})
}

func (p *boss1player) GetVessel() *vesselNode { return p.vessel }

func (p *boss1player) IsDisposed() bool { return false }

func (p *boss1player) Update(delta float64) {
	if p.strafeLeftTime > 0 {
		p.vessel.orders.strafe = true
		p.vessel.orders.rotateLeft = true
	}
	if p.strafeRightTime > 0 {
		p.vessel.orders.strafe = true
		p.vessel.orders.rotateRight = true
	}

	p.stateTicker = gmath.ClampMin(p.stateTicker-delta, 0)
	p.strafeLeftTime = gmath.ClampMin(p.strafeLeftTime-delta, 0)
	p.strafeRightTime = gmath.ClampMin(p.strafeRightTime-delta, 0)

	switch p.bstate {
	case bstateNone:
		p.updateNoneState(delta)
	case bstateIdle:
		p.updateIdleState(delta)
	case bstateFollow:
		p.updateFollowState(delta)
	case bstateRotateToWaypoint:
		p.updateRotateToWaypointState(delta)
	case bstateFlyToWaypoint:
		p.updateFlyToWaypointState(delta)
	case bstateDefense:
		p.updateDefenseState(delta)
	}
}

func (p *boss1player) updateNoneState(delta float64) {
	hpPercent := p.vessel.hp / p.vessel.design.HP

	roll := p.scene.Rand().Float()
	switch {
	case roll <= 0.05: // 5%
		p.stateTicker = p.scene.Rand().FloatRange(0.5, 1)
		p.setState(bstateIdle)

	case roll <= 0.35: // 30%
		// Find a random pos and fly there.
		p.waypoint = p.posAroundPlayer(256)
		p.setState(bstateRotateToWaypoint)

	case roll <= 0.55: // 20%
		// Actively rotate and strafe.
		p.stateTicker = p.scene.Rand().FloatRange(3, 10)
		p.setState(bstateDefense)
		p.canStrafe = p.scene.Rand().Chance(0.85)

	case roll <= 0.65 && hpPercent >= 0.6: // 10%
		// Fly to the last known pos of the player.
		p.waypoint = p.vessel.enemy.pos.Add(p.scene.Rand().Offset(-40, 40))
		p.setState(bstateRotateToWaypoint)

	default: // 25%
		// Follow the player.
		maxDuration := 7.0
		if p.followCancelled {
			maxDuration = 2.5
		}
		p.stateTicker = p.scene.Rand().FloatRange(1, maxDuration)
		p.setState(bstateFollow)
		p.canStrafe = p.followCancelled || p.scene.Rand().Chance(0.3)
	}
}

func (p *boss1player) posAroundPlayer(r float64) gmath.Vec {
	offset := gmath.Vec{X: r, Y: r}
	rect := gmath.Rect{
		Min: p.vessel.enemy.pos.Sub(offset),
		Max: p.vessel.enemy.pos.Add(offset),
	}
	return randomRectPos(p.scene.Rand(), rect)
}

func (p *boss1player) setState(bstate botState) {
	p.rotationDelta = rotationDeltaUnset
	p.strafeDelay = 0
	p.directHits = 0
	p.canStrafe = false
	p.bstate = bstate
	p.followCancelled = false
}

func (p *boss1player) updateIdleState(delta float64) {
	if p.stateTicker > 0 {
		return
	}

	p.setState(bstateNone)
}

func (p *boss1player) updateDefenseState(delta float64) {
	if p.stateTicker == 0 {
		p.setState(bstateNone)
		return
	}

	v := p.vessel

	p.strafeDelay = gmath.ClampMin(p.strafeDelay-delta, 0)

	if p.canStrafe && p.strafeDelay == 0 && p.rotationDelta == rotationDeltaUnset {
		if p.scene.Rand().Bool() {
			p.strafeLeftTime = p.scene.Rand().FloatRange(0.1, 0.45)
			p.strafeDelay = p.scene.Rand().FloatRange(1.1, 4.8)
		} else {
			p.strafeRightTime = p.scene.Rand().FloatRange(0.1, 0.45)
			p.strafeDelay = p.scene.Rand().FloatRange(1.1, 4.8)
		}
	}

	if p.rotationDelta != rotationDeltaUnset {
		p.doRotation(delta)
	}

	if p.rotationDelta == rotationDeltaUnset && p.strafeLeftTime == 0 && p.strafeRightTime == 0 {
		targetAngle := v.pos.AngleToPoint(v.enemy.pos).Normalized()
		angleDelta := angleDelta(v.rotation.Normalized(), targetAngle)
		if angleDelta.Abs() > 0.3 {
			p.rotationDelta = float64(angleDelta)
		}
	}
}

func (p *boss1player) updateFollowState(delta float64) {
	if p.stateTicker == 0 {
		p.setState(bstateNone)
		return
	}

	v := p.vessel

	p.strafeDelay = gmath.ClampMin(p.strafeDelay-delta, 0)

	if p.rotationDelta != rotationDeltaUnset {
		p.doRotation(delta)
	}

	if p.rotationDelta == rotationDeltaUnset {
		p.rotationDelay = gmath.ClampMin(p.rotationDelay-delta, 0)
	}

	targetAngle := v.pos.AngleToPoint(v.enemy.pos).Normalized()
	angleDelta := angleDelta(v.rotation.Normalized(), targetAngle)
	vesselSpeed := v.currentSpeed()

	if p.canStrafe && p.strafeDelay == 0 {
		nextPos := v.pos.MoveInDirection(vesselSpeed, v.rotation)
		targetDist := nextPos.DistanceTo(v.enemy.pos)
		leftPos := v.pos.MoveInDirection(vesselSpeed, v.rotation-math.Pi/4)
		rightPos := v.pos.MoveInDirection(vesselSpeed, v.rotation+math.Pi/4)
		switch {
		case targetDist > 1.05*leftPos.DistanceTo(v.enemy.pos):
			p.strafeLeftTime = p.scene.Rand().FloatRange(0.03, 0.1)
			p.strafeDelay = p.scene.Rand().FloatRange(1.2, 6.8)
		case targetDist > 1.05*rightPos.DistanceTo(v.enemy.pos):
			p.strafeRightTime = p.scene.Rand().FloatRange(0.03, 0.1)
			p.strafeDelay = p.scene.Rand().FloatRange(1.2, 6.8)
		default:
			p.strafeDelay = p.scene.Rand().FloatRange(0.2, 2.8)
		}
	}

	if p.rotationDelay == 0 && p.rotationDelta == rotationDeltaUnset {
		if angleDelta.Abs() > 0.5 {
			p.rotationDelta = float64(angleDelta) + p.scene.Rand().FloatRange(-0.05, 0.05)
			p.rotationDelay = p.scene.Rand().FloatRange(1.1, 3.5)
		} else {
			p.rotationDelay = p.scene.Rand().FloatRange(0.5, 1.5)
		}
	}

	if angleDelta.Abs() < 0.8 {
		dist := v.pos.DistanceTo(v.enemy.pos)
		desiredSpeed := (dist / 2) + 20
		if vesselSpeed < desiredSpeed {
			v.orders.turbo = true
		}
	}
}

func (p *boss1player) updateRotateToWaypointState(delta float64) {
	v := p.vessel

	if p.rotationDelta == rotationDeltaUnset {
		targetAngle := v.pos.AngleToPoint(p.waypoint).Normalized()
		p.rotationDelta = float64(angleDelta(v.rotation.Normalized(), targetAngle))
	}

	if p.doRotation(delta) {
		dist := p.vessel.pos.DistanceTo(p.waypoint)
		p.stateTicker = dist / p.vessel.maxSpeed()
		p.setState(bstateFlyToWaypoint)
	}
}

func (p *boss1player) updateFlyToWaypointState(delta float64) {
	v := p.vessel

	if p.stateTicker == 0 || p.waypoint.DistanceTo(v.pos) <= p.vessel.currentSpeed() {
		p.setState(bstateNone)
		return
	}

	v.orders.turbo = true
}

func (p *boss1player) doRotation(delta float64) bool {
	v := p.vessel

	rotationAmount := float64(v.rotationSpeed()) * delta

	if math.Abs(p.rotationDelta) <= rotationAmount {
		p.rotationDelta = rotationDeltaUnset
		return true
	}

	if p.rotationDelta > 0 {
		p.rotationDelta -= rotationAmount
		v.orders.rotateLeft = true
	} else {
		p.rotationDelta += rotationAmount
		v.orders.rotateRight = true
	}
	return false
}
