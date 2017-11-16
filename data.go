package serializer

type Data struct {
	ID      TypeID `json:"id"`
	Payload []byte `json:"payload"`
}
