package coinbase

type Error struct {
  Type string `json:"type"`
  Message string `json:"message"`
}
