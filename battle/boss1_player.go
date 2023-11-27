package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
)

type boss1player struct {
	world  *battleState
	vessel *vesselNode
	scene  *ge.Scene

	shiftDelay        float64
	shiftRight        float64
	shiftLeft         float64
	rotationDelay     float64
	waypointThreshold float64
	waypoint          gmath.Vec
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
		if p.scene.Rand().Chance(0.2) {
			if p.scene.Rand().Chance(0.6) {
				p.SetWaypoint(p.world.human.vessel.pos.Add(p.scene.Rand().Offset(-96, 96)))
			} else {
				p.SetWaypoint(gmath.Vec{})
			}
		}
	})
}

func (p *boss1player) IsDisposed() bool { return false }

func (p *boss1player) Update(delta float64) {
	p.shiftLeft = gmath.ClampMin(p.shiftLeft-delta, 0)
	p.shiftRight = gmath.ClampMin(p.shiftRight-delta, 0)
	p.shiftDelay = gmath.ClampMin(p.shiftDelay-delta, 0)
	p.rotationDelay = gmath.ClampMin(p.rotationDelay-delta, 0)

	if !p.waypoint.IsZero() {
		p.followWaypoint()
	} else {
		p.findWaypoint()
	}
}

func (p *boss1player) findWaypoint() {
	p.SetWaypoint(p.world.human.vessel.pos)
}

func (p *boss1player) SetWaypoint(pos gmath.Vec) {
	p.waypointThreshold = 50
	p.waypoint = pos
}

func (p *boss1player) followWaypoint() {
	dist := p.waypoint.DistanceTo(p.vessel.pos)

	if dist < p.waypointThreshold {
		p.waypoint = gmath.Vec{}
		return
	}

	if p.shiftLeft > 0 {
		p.vessel.orders.strafe = true
		p.vessel.orders.rotateLeft = true
		return
	}
	if p.shiftRight > 0 {
		p.vessel.orders.strafe = true
		p.vessel.orders.rotateRight = true
		return
	}

	vesselSpeed := p.vessel.currentSpeed()

	if p.shiftDelay == 0 && dist <= 128 {
		nextPos := p.vessel.pos.MoveInDirection(vesselSpeed, p.vessel.rotation)
		nextDist := nextPos.DistanceTo(p.waypoint)
		leftPos := p.vessel.pos.MoveInDirection(vesselSpeed, p.vessel.rotation-math.Pi/4)
		rightPos := p.vessel.pos.MoveInDirection(vesselSpeed, p.vessel.rotation+math.Pi/4)
		switch {
		case nextDist > 1.05*leftPos.DistanceTo(p.waypoint):
			p.shiftLeft = p.scene.Rand().FloatRange(0.03, 0.06)
			p.shiftDelay = p.scene.Rand().FloatRange(1.2, 4.8)
			p.waypointThreshold += 12
		case nextDist > 1.05*rightPos.DistanceTo(p.waypoint):
			p.shiftRight = p.scene.Rand().FloatRange(0.03, 0.06)
			p.shiftDelay = p.scene.Rand().FloatRange(1.2, 4.8)
			p.waypointThreshold += 12
		default:
			p.shiftDelay = p.scene.Rand().FloatRange(0.2, 1.2)
		}
	}

	if p.rotationDelay == 0 {
		targetAngle := p.vessel.pos.AngleToPoint(p.waypoint).Normalized()
		angleDelta := targetAngle.AngleDelta(p.vessel.rotation.Normalized())
		if angleDelta.Abs() > 0.1 {
			if angleDelta > 0 {
				p.vessel.orders.rotateLeft = true
			} else {
				p.vessel.orders.rotateRight = true
			}
		} else {
			p.waypointThreshold += 12
			p.rotationDelay = p.scene.Rand().FloatRange(0.5, 1.5)
		}
	}

	desiredSpeed := (dist / 2) + 20
	if vesselSpeed < desiredSpeed {
		p.vessel.orders.turbo = true
	}
}
