package main;

import (
	"fmt"
	"time"
	"math/rand"
)

//
// using cgo magic for the keyboard io

// #include <conio.h>
import "C"

func konsole() int {
    var key int = 0
    if int(C.kbhit()) > 0 {
    	key = int(C.getch())
    	for ; int(C.kbhit()) > 0; {
    		C.getch()
    	}
    }
    return key
}


const (
	NBAR = 7
	NROT = 4
	NSIZ = 3
)

var _bar = [NBAR][NROT][NSIZ*NSIZ]int{
	{{1,1,1, 0,0,0, 0,0,0},
	 {0,1,0, 0,1,0, 0,1,0},
	 {1,1,1, 0,0,0, 0,0,0},
	 {0,1,0, 0,1,0, 0,1,0}},
	{{2,2,0, 2,2,0, 0,0,0},
	 {2,2,0, 2,2,0, 0,0,0},
	 {2,2,0, 2,2,0, 0,0,0},
	 {2,2,0, 2,2,0, 0,0,0},},
	{{3,3,0, 0,3,3, 0,0,0},
	 {0,0,3, 0,3,3, 0,3,0},
	 {3,3,0, 0,3,3, 0,0,0},
	 {0,0,3, 0,3,3, 0,3,0}},
	{{0,4,4, 4,4,0, 0,0,0},
	 {4,0,0, 4,4,0, 0,4,0},
	 {0,4,4, 4,4,0, 0,0,0},
	 {4,0,0, 4,4,0, 0,4,0}},
	{{0,5,0, 5,5,5, 0,0,0},
	 {0,5,0, 0,5,5, 0,5,0},
	 {0,0,0, 5,5,5, 0,5,0},
	 {0,5,0, 5,5,0, 0,5,0}},
	{{0,6,0, 0,6,0, 0,6,6},
	 {0,0,0, 6,6,6, 0,0,6},
	 {6,6,0, 0,6,0, 0,6,0},
	 {6,0,0, 6,6,6, 0,0,0}},
	{{0,7,0, 0,7,0, 7,7,0},
	 {7,0,0, 7,7,7, 0,0,0},
	 {0,7,7, 0,7,0, 0,7,0},
	 {0,0,0, 7,7,7, 0,0,7}},
}

type Bar struct {
	mask [NROT][NSIZ*NSIZ]int
	state int
	x int
	y int
}
func (me Bar) get(y,x int) int {
	return me.mask[me.state][y*NSIZ+x]
}


type Tetris struct {
	score int
	t int
	w int
	h int
	v []int
	bar Bar
}

func (me *Tetris) newPiece() {
	piece := rand.Intn(NBAR)
	me.bar.mask = _bar[piece]
	me.bar.x = (me.w/2-2)
	me.bar.y = 0
	me.bar.state += 1
	me.bar.state %= 4
	me.score += 1+piece/2
}

func (me *Tetris) remove_bar() {
	for i:=0; i<NSIZ; i++ {
		for j:=0; j<NSIZ; j++ {
			y:= i + me.bar.y
			x:= j + me.bar.x
			if y>= me.h-1 || x>= me.w || x < 0 {
				continue
			}
			v:= me.bar.get(i,j)
			o:= me.get(y,x)
			if o>0 && o!=v {
				continue
			}
			me.set(y,x,0);
		}
	}
}


func (me *Tetris) collapse() {
	collapsed := false
	for y:=me.h-2; y>=1; y-- {
		if collapsed == false {
			s := 0
			for x:=0; x<me.w; x++ {
				v := me.get(y,x)
				if v > 0 { s++ }
			}
			collapsed = (s == me.w)
		}
		if collapsed { // all "on", so reduce
			for x:=0; x<me.w; x++ {
				cu := me.get(y,x)
				up := me.get(y-1,x)
				me.set(y,x,up)
				me.score += cu
			}
		}
	}
}

func (me *Tetris) eval(rot,mov int) bool {
	if me.t == 0 {
		me.newPiece()
	}

	// clear prev. position
	me.remove_bar()

	// possibly reduce all filled lines:
	me.collapse()

	// translation
	if mov != 0 {
		me.bar.x += mov
		if (me.can_move() == false) {
			me.bar.x -= mov
		}
	}
	// rotation
	me.bar.state += NROT + rot
	me.bar.state %= NROT
	can_move_down := me.can_move()
	if can_move_down {
		me.bar.y += 1
	}

	// render piece
	for i:=0; i<NSIZ; i++ {
		for j:=0; j<NSIZ; j++ {
			y:= i + me.bar.y
			x:= j + me.bar.x
			if y>= me.h-1 || x>= me.w {
				continue
			}
			o:= me.get(y,x)
			if o>0 {
				continue
			}
			v:= me.bar.get(i,j)
			me.set(y,x,v);
		}
	}

	// next state
	if can_move_down == false {
		if me.bar.y<1 {
			return false // game over
		}
		me.newPiece()
	}
	return true
}

// collision detection for *next* position(down)
func (me *Tetris) can_move() bool {
	move := false
	for i:=0; i<NSIZ; i++ {
		for j:=0; j<NSIZ; j++ {
			y:= i + me.bar.y
			x:= j + me.bar.x
			if y>= me.h-1 || x>= me.w || x < 0 {
				continue
			}
			v:= me.bar.get(i,j)
			if me.get(y+1,x) != 0 && v != 0 {
				return false
			}
			move = true
		}
	}
	return move
}


// leave out 1 pixel border
func (me *Tetris) render() string {
	s := ""
	for si:=0; si<me.h-1; si++ {
		for sj:=1; sj<me.w-1; sj++ {
			v := me.get(si, sj)
			if v > 0 {
				s += fmt.Sprintf("%d", v)
			} else {
				s += "_"
			}
		}
		s += "\n"
	}
	return s;
}

func (me *Tetris) get(y,x int) int {
	return me.v[y * me.w + x]
}
func (me *Tetris) set(y,x,v int) {
	me.v[y * me.w + x] = v
}

func newGame(w,h int) (*Tetris) {
	t:= new(Tetris)
	t.w = w
	t.h = h
	t.v = make([]int, w * h)
	// we need (invisible) borders
	// for the collision detection later
	for i:=0; i<w; i++ {
		t.v[(h-1)*w + i] = 9
	}
	for j:=0; j<h; j++ {
		t.v[j*w        ] = 9
		t.v[j*w + (w-1)] = 9
	}
	return t
}

func dtime(d time.Duration, f float64) (t time.Duration){
	t = (time.Duration)((float64)(d) * f)
	return
}

func main() {
	rand.Seed(time.Now().Unix())
	tim := time.Second
	tetris := newGame(12,13) // +2 horz, +1 vert border
	for tetris.t=0; tetris.t<1000; tetris.t++ {
	    var key int = konsole()
		mov := 0
		rot := 0
		if key == 97 { mov = -1 } // 's'
		if key == 115 { mov = 1 } // 'a'
		if key == 119 { rot = 1 } // 'w'
		if key == 113 || key == 27 { return  } // 'q' or  'esc'
		fmt.Printf("tetris %v %v\n", tetris.t, tetris.score)

		run := tetris.eval(rot, mov)
		fmt.Println(tetris.render())
		if (run == false) {
			return
		}
		if (tetris.t % 50) == 49 {
			tim = dtime(tim, 0.9)
		}
		if key == 32 { // space
    		time.Sleep(dtime(tim, 0.25))
		} else {
		    time.Sleep(tim)
		}
	}
}
