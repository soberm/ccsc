const { expect } = require("chai");
const { getProof } = require("./utils/helpers");

describe("Proxy", () => {
  let accounts;
  let relayContractFactory, relay;
  let proxyContractFactory, proxy;
  let serverContractFactory, server;
  let clientContractFactory, client;
  let callOpts, callData;

  beforeEach(async () => {
    accounts = await ethers.getSigners();

    relayContractFactory = await ethers.getContractFactory("RelayMock");
    relay = await relayContractFactory.deploy();

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

  describe("deposit", () => {
    it("should deposit specified amount", async () => {
      const amount = await proxy.MIN_DEPOSIT_PER_CALL();
      expect(await proxy.deposit({ value: amount }))
        .to.emit(proxy, "Deposited")
        .withArgs(accounts[0].address, amount);
    });
  });

  describe("withdraw", () => {
    it("should withdraw all funds", async () => {
      const amount = await proxy.MIN_DEPOSIT_PER_CALL();
      expect(await proxy.deposit({ value: amount }))
        .to.emit(proxy, "Deposited")
        .withArgs(accounts[0].address, amount);
    });
    it("should revert if calls pending", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];

      await proxy.initiateCrossChainCall(callData, callOpts);
      await expect(proxy.withdraw()).to.be.revertedWith("calls pending");
    });
  });

  describe("initiateCrossChainCall", () => {
    it("should emit cross chain call initiated", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];

      const tx = await proxy.initiateCrossChainCall(callData, callOpts);
      const blockNumber = await ethers.provider.getBlockNumber();
      expect(tx)
        .to.emit(proxy, "CrossChainCallInitiated")
        .withArgs(
          1,
          accounts[0].address,
          client.address,
          callData,
          250000,
          "storeCallback",
          250000,
          blockNumber + 1000
        );
    });
    it("should revert if relayer deposit too low", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);
      await expect(
        proxy.initiateCrossChainCall(callData, callOpts)
      ).to.be.revertedWith("deposit too low");
    });
  });

  describe("prepareCrossChainCall", () => {
    it("should emit cross chain call prepared", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      expect(
        await proxy.prepareCrossChainCall([
          1,
          accounts[0].address,
          client.address,
          callData,
          250000,
          "storeCallback",
          250000,
          blockNumber + 1000,
        ])
      )
        .to.emit(proxy, "CrossChainCallPrepared")
        .withArgs(1, accounts[0].address, 250000, callData);
    });
    it("should revert if call does not exist", async () => {
      callData = [];
      await expect(
        proxy.prepareCrossChainCall([
          1,
          accounts[0].address,
          client.address,
          callData,
          250000,
          "storeCallback",
          250000,
          1000,
        ])
      ).to.be.revertedWith("call does not exist");
    });
  });

  describe("acknowledgeCrossChainCall", () => {
    it("should acknowledge cross chain call", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await proxy.acknowledgeCrossChainCall(proof, crossChainRequest);
    });
    it("should acknowledge cross chain call and slash", async () => {
      callData = [];
      callOpts = [250000, 0, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await proxy
        .connect(accounts[1])
        .acknowledgeCrossChainCall(proof, crossChainRequest);
    });
    it("should revert if sender is invalid", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      const otherServer = await serverContractFactory.deploy(relay.address);
      await otherServer.setProxy(proxy.address);

      tx = await otherServer.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await expect(
        proxy.acknowledgeCrossChainCall(proof, crossChainRequest)
      ).to.be.revertedWith("invalid sender");
    });
    it("should revert if sender is not origin before timeout", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await expect(
        proxy
          .connect(accounts[1])
          .acknowledgeCrossChainCall(proof, crossChainRequest)
      ).to.be.revertedWith("invalid sender");
    });
    it("should revert if already acknowledged", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await proxy.acknowledgeCrossChainCall(proof, crossChainRequest);
      await expect(
        proxy.acknowledgeCrossChainCall(proof, crossChainRequest)
      ).to.be.revertedWith("call does not exist");
    });
    it("should revert if transaction not verified", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await relay.setVerifyTransactionResult(1);

      await expect(
        proxy.acknowledgeCrossChainCall(proof, crossChainRequest)
      ).to.be.revertedWith("transaction not verified");
    });
    it("should revert if transaction receipt not verified", async () => {
      callData = [];
      callOpts = [250000, 1000, client.address, "storeCallback", 250000];
      await proxy.initiateCrossChainCall(callData, callOpts);

      const blockNumber = await ethers.provider.getBlockNumber();
      crossChainRequest = [
        1,
        accounts[0].address,
        client.address,
        callData,
        250000,
        "storeCallback",
        250000,
        blockNumber + 1000,
      ];

      let tx = await proxy.prepareCrossChainCall(crossChainRequest);
      let proof = await getProof(tx);

      tx = await server.executeCrossChainCall(proof, { value: 2 });
      proof = await getProof(tx);

      await relay.setVerifyReceiptResult(1);

      await expect(
        proxy.acknowledgeCrossChainCall(proof, crossChainRequest)
      ).to.be.revertedWith("transaction receipt not verified");
    });
  });
});
