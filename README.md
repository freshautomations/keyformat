# keyformat
Utility to format YubiHSM private keys into CometBFT/Tendermint JSON format or TMKMS softsign format.

This utility allows unwrapped YubiHSM ed25519 asymmetric private keys to be converted into
CometBFT (and Tendermint) priv_validator.json format or tmkms softsign format.

The exported YubiHSM key has to be unwrapped first using the `yubihsm-unwrap` utility.

## How to install
```bash
go get github.com/freshautomations/keyformat
```

## How to use
Below is an example of exporting a key from YubiHSM for plain-text use:

```bash
# Export asymmetric ed25519 key ID 9 using wrap key ID 1. The tmkms.toml file
# defines a password that gives at least operator access to the HSM device.
tmkms yubihsm keys export -i 9 -w 1 wrapped.enc -c tmkms.toml

# Unwrap the key. The wrap key is saved in a binary file.
yubihsm-unwrap --in wrapped.enc --wrapkey wrap.key --out decrypted.key

# Format the private key for CometBFT/Tendermint
keyformat -key decrypted.key -output priv_validator.json

# Format the private key for tmkms
keyformat -key decrypted.key -output softsign.key -softsign
```

## Resources
* [tmkms source](https://github.com/iqlusioninc/tmkms)
* [yubihsm-unwrap (PR)](https://github.com/Yubico/yubihsm-shell/pull/323)
* [CometBFT](https://github.com/cometbft/cometbft)
