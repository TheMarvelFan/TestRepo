package diffiehellman

import (
    "crypto/rand"
    "math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
    min := big.NewInt(2)
    diff := new(big.Int).Sub(p, min)
	n, _ := rand.Int(rand.Reader, diff)
    return n.Add(n, min)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	ret := new(big.Int).Exp(big.NewInt(g), private, p)
    return ret
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
    private := PrivateKey(p)
	public := PublicKey(private, p, g)

    return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	ret := new(big.Int).Exp(public2, private1, p)
    return ret
}
