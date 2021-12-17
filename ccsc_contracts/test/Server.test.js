const { expect } = require("chai");
const { getProof } = require("./utils/helpers");

describe("Server", () => {
  let accounts;
  let relayContractFactory, relay;
  let storageContractFactory;
  let serverContractFactory, server;
  let proxyContractFactory, proxy;
  let clientContractFactory, client;

  beforeEach(async () => {
    accounts = await ethers.getSigners();

    relayContractFactory = await ethers.getContractFactory("RelayMock");
    relay = await relayContractFactory.deploy();

    storageContractFactory = await ethers.getContractFactory("Storage");
    storage = await storageContractFactory.deploy();

    serverContractFactory = await ethers.getContractFactory("Server");
    server = await serverContractFactory.deploy(relay.address);

    proxyContractFactory = await ethers.getContractFactory("Proxy");
    proxy = await proxyContractFactory.deploy(server.address, relay.address);
    await server.setProxy(proxy.address);

    const amount = await proxy.MIN_DEPOSIT_PER_CALL();
    await proxy.deposit({ value: amount });

    clientContractFactory = await ethers.getContractFactory("Client");
    client = await clientContractFactory.deploy(proxy.address);
  });

  describe("executeCrossChainCall", () => {
    it("should revert with invalid sender", async () => {
      const otherProxy = await proxyContractFactory.deploy(
        server.address,
        relay.address
      );
      const amount = await otherProxy.MIN_DEPOSIT_PER_CALL();
      await otherProxy.deposit({ value: amount });

      const calldata = storageContractFactory.interface.encodeFunctionData(
        "store",
        [ethers.BigNumber.from("42")]
      );

      await proxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);
      await otherProxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);

      const blockNumber = await ethers.provider.getBlockNumber();
      const tx = await otherProxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      const proof = await getProof(tx);

      await expect(
        server.executeCrossChainCall(proof, { value: 2 })
      ).to.be.revertedWith("invalid sender");
    });
    it("should revert if transaction not verified", async () => {
      const calldata = storageContractFactory.interface.encodeFunctionData(
        "store",
        [ethers.BigNumber.from("42")]
      );
      await proxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);
      const blockNumber = await ethers.provider.getBlockNumber();
      const tx = await proxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      await relay.setVerifyTransactionResult(1);
      const proof = await getProof(tx);

      await expect(
        server.executeCrossChainCall(proof, { value: 2 })
      ).to.be.revertedWith("transaction not verified");
    });
    it("should revert if transaction receipt not verified", async () => {
      const calldata = storageContractFactory.interface.encodeFunctionData(
        "store",
        [ethers.BigNumber.from("42")]
      );
      await proxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);
      const blockNumber = await ethers.provider.getBlockNumber();
      const tx = await proxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      await relay.setVerifyReceiptResult(1);
      const proof = await getProof(tx);

      await expect(
        server.executeCrossChainCall(proof, { value: 2 })
      ).to.be.revertedWith("transaction receipt not verified");
    });
    it("should revert if call already executed", async () => {
      const calldata = storageContractFactory.interface.encodeFunctionData(
        "store",
        [ethers.BigNumber.from("42")]
      );
      await proxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);
      const blockNumber = await ethers.provider.getBlockNumber();
      const tx = await proxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      const proof = await getProof(tx);
      await server.executeCrossChainCall(proof, { value: 2 });
      await expect(
        server.executeCrossChainCall(proof, { value: 2 })
      ).to.be.revertedWith("already executed");
    });
    it("should execute cross chain call", async () => {
      const calldata = storageContractFactory.interface.encodeFunctionData(
        "store",
        [ethers.BigNumber.from("42")]
      );
      await proxy.initiateCrossChainCall(calldata, [
        250000,
        1000,
        client.address,
        "storeCallback",
        250000,
      ]);
      const blockNumber = await ethers.provider.getBlockNumber();
      const tx = await proxy.prepareCrossChainCall([
        1,
        accounts[0].address,
        client.address,
        calldata,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ]);
      const proof = await getProof(tx);

      expect(await server.executeCrossChainCall(proof, { value: 2 }))
        .to.emit(server, "CrossChainCallExecuted")
        .withArgs(1, false, "0x");
    });
  });
});
