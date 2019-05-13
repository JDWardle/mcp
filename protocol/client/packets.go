package client

type Play int

const (
	SpawnObject Play = iota
	SpawnExperienceOrb
	SpawnGlobalEntity
	SpawnMob
	SpawnPainting
	SpawnPlayer
	Animation
	Statistics
	BlockBreakAnimation
	UpdateBlockEntity
	BlockAction
	BlockChange
	BossBar
	ServerDifficulty
	ChatMessage
	MultiBlockChange
	TabComplete
	DeclareCommands
	ConfirmTransaction
	CloseWindow
	OpenWindow
	WindowItems
	WindowProperty
	SetSlot
	SetCooldown
	PluginMessage
	NamedSoundEffect
	Disconnect
	EntityStatus
	NBTQueryResponse
	Explosion
	UnloadChunk
	ChangeGameState
	KeepAlive
	ChunkData
	Effect
	Particle
	JoinGame
	MapData
	Entity
	EntityRelativeMove
	EntityLookAndRelativeMove
	EntityLook
	VehicleMove
	OpenSignEditor
	CraftRecipeResponse
	PlayerAbilities
	CombatEvent
	PlayerInfo
	FacePlayer
	PlayerPositionAndLook
	UseBed
	UnlockRecipes
	DestroyEntities
	RemoveEntityEffect
	ResourcePackSend
	Respawn
	EntityHeadLook
	SelectAdvancementTab
	WorldBorder
	Camera
	HeldItemChange
	DisplayScoreboard
	EntityMetadata
	AttachEntity
	EntityVelocity
	EntityEquipment
	SetExperience
	UpdateHealth
	ScoreboardObjective
	SetPassengers
	Teams
	UpdateScore
	SpawnPosition
	TimeUpdate
	Title
	StopSound
	SoundEffect
	PlayerListHeaderAndFooter
	CollectItem
	EntityTeleport
	Advancements
	EntityProperties
	EntityEffect
	DeclareRecipes
	Tags
)

type Status int

const (
	Response = iota
	Pong
)

type Login int

const (
	Disconnect Login = iota
	EncryptionRequest
	LoginSuccess
	SetCompression
	LoginPluginRequest
)
