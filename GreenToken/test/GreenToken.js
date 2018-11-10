require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

const GreenToken = artifacts.require('GreenToken');

contract('GreenToken', function ([_, wallet1, wallet2, wallet3, wallet4, wallet5]) {
    beforeEach(async function () {
        this.token = await GreenToken.new();
    });
});
