pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";


contract GreenToken is Ownable, ERC721 {
    struct Data {
        uint256 tokenId;
        uint256 createdDate;
        string ownerName;
        string ogrn;
        string addr;
        uint256 volume; // multiplied by 10**18
        string period;
        uint256 co2Volume; // multiplied by 10**18
        string sertNumber;
        uint256 lifeTime; // in seconds
        string _type;
        string wallet;
    }

    mapping(uint256 => Data) public tokens;

    function destroy(
        uint256 tokenId
    )
        public
        onlyOwner
    {
        _burn(ownerOf(tokenId), tokenId);
    }

    function create(
        address to,
        uint256 tokenId,
        uint256 createdDate,
        string ownerName,
        string ogrn,
        string addr,
        uint256 volume, // multiplied by 10**18
        string period,
        uint256 co2Volume, // multiplied by 10**18
        string sertNumber,
        uint256 lifeTime, // in seconds
        string _type,
        string wallet
    )
        public
        onlyOwner
    {
        require(tokenId != 0, "Token with the zero id can't be recreated");
        require(tokens[tokenId].tokenId == 0, "Token with the same id can't be recreated");

        _mint(to, tokenId);

        tokens[tokenId] = Data({
            tokenId: tokenId,
            createdDate: createdDate,
            ownerName: ownerName,
            ogrn: ogrn,
            addr: addr,
            volume: volume, // multiplied by 10**18
            period: period,
            co2Volume: co2Volume, // multiplied by 10**18
            sertNumber: sertNumber,
            lifeTime: lifeTime, // in seconds
            _type: _type,
            wallet: wallet
        });
    }
}
