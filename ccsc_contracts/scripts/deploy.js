const { ethers } = require("hardhat");
const { networks } = require("../hardhat.config");

const ETH_RELAY_ADDRESS = "";
const BSC_RELAY_ADDRESS = "";

const ropstenProvider = ethers.getDefaultProvider(networks.ropsten.url);
const ropstenWallet = new ethers.Wallet(
  networks.ropsten.accounts[0],
  ropstenProvider
);

const bscTestnetProvider = ethers.getDefaultProvider(networks.bsc_testnet.url);
const bscTestnetWallet = new ethers.Wallet(
  networks.bsc_testnet.accounts[0],
  bscTestnetProvider
);

async function main() {
  console.log("### Deploying to BSC testnet ###");
  const storageContractFactory = await ethers.getContractFactory(
    "Storage",
    bscTestnetWallet
  );
  const storage = await storageContractFactory.deploy();

  console.log("Storage deployed to:", storage.address);

  const storageServerContractFactory = await ethers.getContractFactory(
    "StorageServer",
    bscTestnetWallet
  );
  const storageServer = await storageServerContractFactory.deploy(
    storage.address,
    ETH_RELAY_ADDRESS
  );

  console.log("StorageServer deployed to:", storageServer.address);

  console.log("### Deploying to Ropsten ###");

  const storageProxyContractFactory = await ethers.getContractFactory(
    "StorageProxy",
    ropstenWallet
  );
  const storageProxy = await storageProxyContractFactory.deploy(
    storageServer.address,
    BSC_RELAY_ADDRESS
  );
  console.log("StorageProxy deployed to:", storageProxy.address);
  await storageServer.setProxy(storageProxy.address);

  const clientContractFactory = await ethers.getContractFactory(
    "Client",
    ropstenWallet
  );
  const client = await clientContractFactory.deploy(storageProxy.address);

  console.log("Client deployed to:", client.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
