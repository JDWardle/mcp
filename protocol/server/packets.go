package server

type Play int

const (
	TeleportConfirm Play = iota
	QueryBlockNBT
	ChatMessage
	ClientStatus
	ClientSettings
	TabComplete
	ConfirmTransaction
	EnchantItem
	ClickWindow
	CloseWindow
	PluginMessage
	EditBook
	QueryEntityNBT
	UseEntity
	KeepAlive
	Player
	PlayerPosition
	PlayerPositionAndLook
	PlayerLook
	VehicleMove
	SteerBoat
	PickItem
	CraftRecipeRequest
	PlayerAbilities
	PlayerDigging
	EntityAction
	SteerVehicle
	RecipeBookData
	NameItem
	ResourcePackStatus
	AdvancementTab
	SelectTrade
	SetBeaconEffect
	HeldItemChange
	UpdateCommandBlock
	UpdateCommandBlockMinecart
	CreativeInventoryAction
	UpdateStructureBlock
	UpdateSign
	Animation
	Spectate
	PlayerBlockPlacement
	UseItem
)

type Status int

const (
	Request Status = iota
	Ping
)
