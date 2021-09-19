package mlib

import "errors"

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics){
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry  {
	if len(m.musics) == 0{
		return nil
	}
	for _, music := range m.musics {
		if music.Name == name{
			return &music
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry)  {
	// *type 表示这个是type的指针类型
	// & 为寻址符号 *ptr 表示查看这个指针指向的内容
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics){
		return nil
	}
	removeMusic := &m.musics[index]
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	}else if index ==0{
		m.musics = m.musics[1:]
	}else {
		m.musics = m.musics[:index]
	}
	return  removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}
