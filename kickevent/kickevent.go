package kickevent

type KickEvent int

const (
    ChatMessage KickEvent = iota
    UserBanned
    MessageDeleted
)