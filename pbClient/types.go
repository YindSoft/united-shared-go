package pbClient

type HasMsgTypeId interface{ MsgTypeId() uint32 }

func (*Talk) MsgTypeId() uint32 { return 1 }
func (*Walk) MsgTypeId() uint32 { return 2 }
