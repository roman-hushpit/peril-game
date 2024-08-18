package main

func (u user) doBattles(subCh <-chan move) []piece {
	battlePieces := make([]piece, 0)
	for move := range subCh {
		userLocations := u.pieces 
		for _, userLocation := range userLocations {
			if move.piece.location == userLocation.location {
				battlePieces = append(battlePieces, userLocation)
			}
		}
	}
	return battlePieces
}

// don't touch below this line

type user struct {
	name   string
	pieces []piece
}

type move struct {
	userName string
	piece    piece
}

type piece struct {
	location string
	name     string
}

func (u user) march(p piece, publishCh chan<- move) {
	publishCh <- move{
		userName: u.name,
		piece:    p,
	}
}

func distributeBattles(publishCh <-chan move, subChans []chan move) {
	for mv := range publishCh {
		for _, subCh := range subChans {
			subCh <- mv
		}
	}
}
