

# Model
Semi-honest Security
2-PC
## Y(Yao's Garbled Circuit)
G. Asharov, Y. Lindell, T. Schneider, and M. Zohner.
More efficient oblivious transfer and extensions for
faster secure computation. In CCS, 2013.

Due to constant round, perform well in high-latency networks

SOTA: point-and-permute + free-XOR + fixedkey AES + half-gates 

store garbled table may nees lots of memory, pipeline is useful but will shift communication from setup phase to online phase.
## HE(Homomorphic Encryption)

Additively HE:
- Paillier, Public-key cryptosystems based on composite degree residuosity classes 
- DGK, Homomorphic encryption and secure comparison
- RLWE-AHE, Improved multiplication triple generation over rings via RLWEbased AHE

Full HE:
## SS(Secret Sharing)
Beaver’s Triple

data over domain $2^{l}$, variables are $l$ bit number.

### Multiplication triple
via Additively Homomorphic Encryption
使用 Paillier, 或者DGK with full decryption using the Pohlig-Hellman algorithm.
