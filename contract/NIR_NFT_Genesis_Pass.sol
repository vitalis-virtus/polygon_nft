// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NIR_NFT is ERC721, Ownable {
    uint256 MAX_SUPPLY = 50000;
    string constant TOKEN_URI = "ipfs://bafkreigcjy6bkginsexsjthc2xl4sqsllp453c2bhenx63lsnpw2dvafcm";

    uint256 private _nextTokenId;

    constructor(address initialOwner)
        ERC721("Nirimly Labs (Nirimly Genesis Pass)", "NIRGP")
        Ownable(initialOwner)
    {}

    function tokenURI(uint256 /* tokenId */) public pure override returns (string memory) {
    return TOKEN_URI;
    }

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
        require(tokenId <= MAX_SUPPLY, "Sorry, all NFTs have been minted!");
        _safeMint(to, tokenId);
    }
}