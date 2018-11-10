const GreenToken = artifacts.require('GreenToken');

module.exports = async function (deployer) {
    deployer.deploy(GreenToken);
};
