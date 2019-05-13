package server

import (
	"bufio"
	"sync"

	"github.com/JDWardle/gocraft/protocol"
)

type HandlerFunc func(c *Client, r *bufio.Reader) error

type Mux struct {
	m  map[protocol.ClientState]map[int32]HandlerFunc
	mu sync.RWMutex
}

func (m Mux) GetHandler(clientState protocol.ClientState, id int32) (bool, HandlerFunc) {
	m.mu.RLock()

	if h, ok := m.m[clientState][id]; ok {
		m.mu.RUnlock()
		return ok, h
	}

	m.mu.RUnlock()
	return false, nil
}

var DefaultHandlers = &Mux{
	m: map[protocol.ClientState]map[int32]HandlerFunc{
		protocol.ClientStateHandshaking: {
			0x00: HandshakeHandler,
			0xFE: LegacyServerListPingHandler,
		},
		protocol.ClientStateStatus: {
			0x00: StatusRequestHandler,
			0x01: PingHandler,
		},
		protocol.ClientStateLogin: {
			0x00: LoginStartHandler,
			0x01: EncryptionResponseHandler,
			0x02: LoginPluginResponseHandler,
		},
		protocol.ClientStatePlay: {
			0x00: TeleportConfirmHandler,
			0x01: QueryBlockNBTHandler,
			0x02: ChatMessageHandler,
			0x03: ClientStatusHandler,
			0x04: ClientSettingsHandler,
			0x05: TabCompleteHandler,
			0x06: ConfirmTransactionHandler,
			0x07: EnchantItemHandler,
			0x08: ClickWindowHandler,
			0x09: CloseWindowHandler,
			0x0A: PluginMessageHandler,
			0x0B: EditBookHandler,
			0x0C: QueryEntityNBTHandler,
			0x0D: UseEntityHandler,
			0x0E: KeepAliveHandler,
			0x0F: PlayerHandler,
			0x10: PlayerPositionHandler,
			0x11: PlayerPositionAndLookHandler,
			0x12: PlayerLookHandler,
			0x13: VehicleMoveHandler,
			0x14: SteerBoatHandler,
			0x15: PickItemHandler,
			0x16: CraftRecipeRequestHandler,
			0x17: PlayerAbilitiesHandler,
			0x18: PlayerDiggingHandler,
			0x19: EntityActionHandler,
			0x1A: SteerVehicleHandler,
			0x1B: RecipeBookDataHandler,
			0x1C: NameItemHandler,
			0x1D: ResourcePackStatusHandler,
			0x1E: AdvancementTabHandler,
			0x1F: SelectTradeHandler,
			0x20: SetBeaconEffectHandler,
			0x21: HeldItemChangeHandler,
			0x22: UpdateCommandBlockHandler,
			0x23: UpdateCommandBlockMinecartHandler,
			0x24: CreativeInventoryActionHandler,
			0x25: UpdateStructureBlockHandler,
			0x26: UpdateSignHandler,
			0x27: AnimationHandler,
			0x28: SpectateHandler,
			0x29: PlayerBlockPlacementHandler,
			0x2A: UseItemHandler,
		},
	},
}
