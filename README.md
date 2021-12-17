# A Framework for Asynchronous Cross-Blockchain Smart Contract Calls

This project contains a prototypical implementation of our framework to enable smart contract interoperability across different blockchains. The framework allows smart contracts residing on two different EVM-based blockchains to call each other. The framework provides the opportunity to create stub contracts which serve as an abstraction layer to hide the background details of cross-blockchain communication. For that, the stub contracts use blockchain relays to transfer the information to another blockchain. A  client  contract  can conveniently  call  a  stub  contract  to  call  a  smart  contract  on another  blockchain,  as  the  transaction  sender  facilitates  the communication between the stubs by forwarding the necessary information with the respective proof data.

### Prerequisites

* [Hardhat](https://hardhat.org/)
* [Golang](https://go.dev/)
* [Solidity](https://docs.soliditylang.org/en/v0.8.10/)

## Example

In `./ccsc_contracts/contracts/ccsc`, you can find the general  client  and  server stub, which can be used to create the stubs for specific smart contracts. We used it for a simple storage contract, which can be found in `./ccsc_contracts/contracts/examples/`. Further, we added a simple cross-blockchain caller in `ccsc_go` which executes some cross-blockchain calls with the example contracts and captures different metrics. To install the example, follow these steps:

### Installing the example contracts

1. Deploy [ETH-Relay](https://github.com/pantos-io/ethrelay) and [BSC-Relay](https://github.com/soberm/bsc_relay).
2. Change into the contract directory: `cd ccsc_contracts/`
3. Install all dependencies: `npm install`
4. Change the network configuration in `hardhat.config.js`
5. Update the addresses for ETH-Relay and BSC-Relay in `./scripts/deploy.js`
6. Run `hardhat run ./scripts/deploy.js` to deploy the example on the configured networks

### Installing the example caller

1. Change into the contract directory: `cd ccsc_go/`
2. Build the example caller with `go build -o example_caller ./cmd/main.go`
3. Adapt the configuration in `./configs/config.json`
4. Run the example caller with `./example_caller`

## Contributing

This is a research prototype. We welcome anyone to contribute. File a bug report or submit feature requests through the issue tracker. If you want to contribute feel free to submit a pull request.

## Acknowledgement

The financial support by the Austrian Federal Ministry for Digital and Economic Affairs, the National Foundation for Research, Technology and Development as well as the Christian Doppler Research Association is gratefully acknowledged.

## License

This project is licensed under the [MIT License](LICENSE).
