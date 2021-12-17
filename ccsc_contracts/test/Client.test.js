const { expect } = require("chai");
const { getProof } = require("./utils/helpers");

describe("Client", () => {
  let relayContractFactory, relay;
  let storageContractFactory, storage;
  let storageProxyContractFactory, storageProxy;
  let storageServerContractFactory, storageServer;
  let clientContractFactory, client;

  beforeEach(async () => {
    accounts = await ethers.getSigners();

    relayContractFactory = await ethers.getContractFactory("RelayMock");
    relay = await relayContractFactory.deploy();

    storageContractFactory = await ethers.getContractFactory("Storage");
    storage = await storageContractFactory.deploy();

    storageServerContractFactory = await ethers.getContractFactory(
      "StorageServer"
    );
    storageServer = await storageServerContractFactory.deploy(
      storage.address,
      relay.address
    );

    storageProxyContractFactory = await ethers.getContractFactory(
      "StorageProxy"
    );
    storageProxy = await storageProxyContractFactory.deploy(
      storageServer.address,
      relay.address
    );
    await storageServer.setProxy(storageProxy.address);

    const amount = await storageProxy.MIN_DEPOSIT_PER_CALL();
    await storageProxy.deposit({ value: amount });

    clientContractFactory = await ethers.getContractFactory("Client");
    client = await clientContractFactory.deploy(storageProxy.address);
  });

  describe("store", () => {
    it("should use the stub to store a number", async () => {
      await client.store(42);

      const calldata =
        storageServerContractFactory.interface.encodeFunctionData("store", [
          ethers.BigNumber.from("42"),
          accounts[0].address,
          client.address,
        ]);

      const blockNumber = await ethers.provider.getBlockNumber();
      let tx = await storageProxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      let proof = await getProof(tx);

      tx = await storageServer.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await storageProxy.acknowledgeCrossChainCall(proof, [
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);

      expect(await client.isStored()).to.be.equal(true);
      expect(await storage.retrieve()).to.be.equal(42);
    });
  });

  describe("retrieve", () => {
    it("should use the stub to retrieve the number", async () => {
      await storage.store(42);

      await client.retrieve();

      const calldata =
        storageServerContractFactory.interface.encodeFunctionData("retrieve", [
          accounts[0].address,
          client.address,
        ]);
      const blockNumber = await ethers.provider.getBlockNumber();
      let tx = await storageProxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "retrieveCallback",
        250000,
        blockNumber + 1000,
      ]);
      let proof = await getProof(tx);

      tx = await storageServer.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await storageProxy.acknowledgeCrossChainCall(proof, [
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "retrieveCallback",
        250000,
        blockNumber + 1000,
      ]);

      expect(await client.getRetrievedNumber()).to.be.equal(42);
    });
  });
});
