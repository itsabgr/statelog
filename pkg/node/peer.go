package node

type Peer struct {
	pk   []byte
	addr string
}

func (p *Peer) ID() []byte {
	return p.pk
}
