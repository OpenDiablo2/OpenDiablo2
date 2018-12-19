namespace OpenDiablo2.Common.Enums
{
    // TODO: I don't think this needs to live in core...
    public enum eMessageFrameType
    {
        None = 0x00,
        SetSeed = 0x01,
        JoinGame = 0x02,
        LocatePlayers = 0x03,
        PlayerInfo = 0x04,
        FocusOnPlayer = 0x05,
        MoveRequest = 0x06,
        PlayerMove = 0x07,
        UpdateEquipment = 0x08,
        ChangeEquipment = 0x09,

        MAX = 0xFF, // NOTE:
        // You absolutely cannot have a higher ID than this without
        // changing the message header to multi-byte for ALL frame types!!!
    }
}
