# Web3 Tool Suite

```shell
go mod download
```

## mos server
```shell
go run main.go server
```

## mos cli

### btc utils
```shell
# generate bitcoin accounts from mnemonic
go run main.go btc hdkey -m "test test test test test test test test test test test junk"
```

### eth utils
```shell
# generate 10 eth accounts from mnemonic
go run main.go eth hdkey -m "test test test test test test test test test test test junk"

m/44'/60'/0'/0/0   0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
m/44'/60'/0'/0/1   0x70997970C51812dc3A010C7d01b50e0d17dc79C8 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
m/44'/60'/0'/0/2   0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC 5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
m/44'/60'/0'/0/3   0x90F79bf6EB2c4f870365E785982E1f101E93b906 7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
m/44'/60'/0'/0/4   0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
m/44'/60'/0'/0/5   0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc 8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
m/44'/60'/0'/0/6   0x976EA74026E726554dB657fA54763abd0C3a0aa9 92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
m/44'/60'/0'/0/7   0x14dC79964da2C08b23698B3D3cc7Ca32193d9955 4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
m/44'/60'/0'/0/8   0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
m/44'/60'/0'/0/9   0xa0Ee7A142d267C1f36714E4a8F75612F20a79720 2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
m/44'/60'/0'/0/10  0xBcd4042DE499D14e55001CcbB24a551F3b954096 f214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897

```

```shell
# generate 2 eth accounts from mnemonic
go run main.go eth hdkey -m "test test test test test test test test test test test junk" -N 2

m/44'/60'/0'/0/0   0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
m/44'/60'/0'/0/1   0x70997970C51812dc3A010C7d01b50e0d17dc79C8 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
```

```shell
# generate address privateKey publicKey from mnemonic
go run main.go eth hdkey -m "test test test test test test test test test test test junk" -H "m/44'/60'/0'/0/0"

hdPath:  m/44'/60'/0'/0/0
address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
privKey: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
pubKey:  8318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5
```

### cosmos utils
```shell

```

### p2p utils
```shell
# generate p2p node id from node privateKey
go run main.go p2p gen-node-id -i ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
16Uiu2HAmMUjGmiUhJeiZgu6ZZnLRkE2VViR2JgjqtW9aTZnHQqgg

# generate p2p node id from node publicKey
go run main.go p2p gen-node-id -u 8318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5
16Uiu2HAmMUjGmiUhJeiZgu6ZZnLRkE2VViR2JgjqtW9aTZnHQqgg
```
