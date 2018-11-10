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
    }

    mapping(uint256 => Data) private tokens;

    function getInfo1(
        uint256 tokenId
    )
        public
        view
        returns(
            uint256 createdDate,
            string ownerName,
            string ogrn,
            string addr,
            uint256 volume
        )
    {
        Data storage token = tokens[tokenId];
        return (
            token.createdDate,
            token.ownerName,
            token.ogrn,
            token.addr,
            token.volume
        );
    }

    function getInfo2(
        uint256 tokenId
    )
        public
        view
        returns(
            string period,
            uint256 co2Volume, // multiplied by 10**18
            string sertNumber,
            uint256 lifeTime, // in seconds
            string _type
        )
    {
        Data storage token = tokens[tokenId];
        return (
            token.period,
            token.co2Volume,
            token.sertNumber,
            token.lifeTime,
            token._type
        );
    }

    function destroy(
        uint256 tokenId
    )
        public
        onlyOwner
    {
        Data storage token = tokens[tokenId];
        require(now < token.createdDate + token.lifeTime, "Token should expire before destroying");

        _burn(ownerOf(tokenId), tokenId);
    }

    function create(
        address to,
        uint256 tokenId,
        string ownerName,
        string ogrn,
        string addr,
        uint256 volume, // multiplied by 10**18
        string period,
        uint256 co2Volume, // multiplied by 10**18
        string sertNumber,
        uint256 lifeTime, // in seconds
        string _type
    )
        public
        onlyOwner
    {
        require(tokenId != 0, "Token with the zero id can't be recreated");
        require(tokens[tokenId].tokenId == 0, "Token with the same id can't be recreated");

        _mint(to, tokenId);

        tokens[tokenId] = Data({
            tokenId: tokenId,
            createdDate: now,
            ownerName: ownerName,
            ogrn: ogrn,
            addr: addr,
            volume: volume,
            period: period,
            co2Volume: co2Volume,
            sertNumber: sertNumber,
            lifeTime: lifeTime,
            _type: _type
        });
    }
}
