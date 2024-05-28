package platform

type Platform int

const (
    Internal Platform = iota
    Kick
    Discord
)

const Max Platform = Discord