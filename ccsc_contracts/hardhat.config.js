require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-solhint");
require("hardhat-gas-reporter");
require("solidity-coverage");

task("accounts", "Prints the list of accounts", async () => {
  const accounts = await ethers.getSigners();

  for (let i = 0; i < accounts.length; i++) {
    const account = accounts[i];
    console.log(account.address);
  }
});

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {},
    ropsten: {
      url: "",
      accounts: [
        "",
      ],
    },
    bsc_testnet: {
      url: "",
      accounts: [
        "",
      ],
    },
  },
  solidity: "0.7.3",
};
