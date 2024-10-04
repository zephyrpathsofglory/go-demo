package od

import (
	"math"
	"sort"
)

type RentingSystem struct {
	roomMap map[int]room
}

type room struct {
	id      int
	area    int
	price   int
	rooms   int
	address []int
}

func New() *RentingSystem {
	return &RentingSystem{
		roomMap: make(map[int]room),
	}
}

func (s *RentingSystem) addRoom(id, area, price, rooms int, address []int) bool {
	newRoom := room{
		id:      id,
		area:    area,
		price:   price,
		rooms:   rooms,
		address: address,
	}
	_, ok := s.roomMap[id]

	s.roomMap[id] = newRoom

	return !ok
}

func (s *RentingSystem) deleteRoom(id int) bool {
	_, ok := s.roomMap[id]
	delete(s.roomMap, id)

	return ok
}

type target struct {
	id       int
	area     int
	price    int
	rooms    int
	distance float64
}

type roomSorter struct {
	rooms   []target
	ordreBy [][]int
}

func (rs roomSorter) Len() int {
	return len(rs.rooms)
}

func (rs roomSorter) Less(x, y int) bool {
	return applyOrderBy(rs.rooms[x], rs.rooms[y], rs.ordreBy)
}

func (rs roomSorter) Swap(x, y int) {
	rs.rooms[x], rs.rooms[y] = rs.rooms[y], rs.rooms[x]
}

func applyOrderBy(x, y target, orderBy [][]int) bool {
	for _, o := range orderBy {
		switch o[0] {
		case 1:
			if x.area*o[1] == y.area*o[1] {
				continue
			}

			return x.area*o[1] < y.area*o[1]
		case 2:
			if x.price*o[1] == y.area*o[1] {
				continue
			}

			return x.price*o[1] < y.price*o[1]
		case 3:
			if x.distance*float64(o[1]) == y.distance*float64(o[1]) {
				continue
			}

			return x.distance*float64(o[1]) < y.distance*float64(o[1])
		case 4:
			return x.id < y.id
		}
	}

	return false
}

func (s *RentingSystem) queryRoom(area, price, rooms int, address []int, orderBy [][]int) []int {
	roomArr := make([]target, 0, len(s.roomMap))

	for _, r := range s.roomMap {
		if r.area >= area && r.price <= price && r.rooms == rooms {
			distance := math.Abs(float64(r.address[0]-address[0])) + math.Abs(float64(r.address[1]-address[1]))

			roomArr = append(roomArr, target{
				id:       r.id,
				area:     r.area,
				price:    r.price,
				rooms:    r.rooms,
				distance: distance,
			})
		}
	}

	sorter := roomSorter{
		rooms:   roomArr,
		ordreBy: append(orderBy, []int{4, 1}),
	}

	sort.Sort(sorter)

	res := make([]int, 0, len(sorter.rooms))
	for _, r := range sorter.rooms {
		res = append(res, r.id)
	}

	return res
}
