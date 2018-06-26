package game

import "log"

// Room ... default entity to structure rooms
type Room struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Avatars []*Avatar `json:"avatars"`
}

// NewRoom ... creates and returns a new room instance
func NewRoom(id string, title string, description string) *Room {
	return &Room{
		ID:          id,
		Title:       title,
		Description: description,
		Avatars:     []*Avatar{},
	}
}

func (room *Room) removeAvatar(i int) {
	s := room.Avatars
	s = append(s[:i], s[i+1:]...)
	room.Avatars = s
}

func (room *Room) addAvatarToRoom(avatar *Avatar) {
	room.Avatars = append(room.Avatars, avatar)
}

func (room *Room) removeAvatarFromRoom(avatar *Avatar) {

	//find index
	for idx, element := range room.Avatars {
		if element == avatar {
			room.removeAvatar(idx)
		}
	}
}

// Enter ... processes every entity
func (room *Room) Enter(avatar *Avatar) {

	log.Println("avatar entered room " + avatar.CurrentUser.ID + " room: " + room.Title)

	avatar.LastKnownRoom = room
	GetInstance().OnAvatarJoinedRoom <- &AvatarJoinedRoom{
		Avatar: avatar,
		Room:   room,
	}

}

// Leave ... processes every entity
func (room *Room) Leave(avatar *Avatar) {

	/*	GetInstance().OnAvatarJoinedRoom <- &AvatarJoinedRoom{
			Avatar: avatar,
			Room:   room,
		}
	*/
}
