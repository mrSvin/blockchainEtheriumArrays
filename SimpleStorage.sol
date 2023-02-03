// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract SimpleStorage {

    struct Wallet {
        string walletName;
        uint balance;
    }

    mapping(string => Wallet) public wallets;
    string nameForWallet;
    uint balanceWallet;

    function setWallet(string memory nameWallet, uint setBalance) public {
        wallets[nameWallet].walletName = nameWallet;
        wallets[nameWallet].balance = setBalance;
    }

    function getWallet(string memory nameWallet) external view returns (Wallet memory) {
        return wallets[nameWallet];
    }

    function sendMoney(string memory nameWalletSender, string memory nameWalletRecipient, uint money) public {
            wallets[nameWalletSender].balance = wallets[nameWalletSender].balance - money;
            wallets[nameWalletRecipient].balance = wallets[nameWalletRecipient].balance + money;
    }

    function sendMoneyTenThousandTransactions(string[] memory senders, string[] memory recipients, uint money) public {
        for (uint i = 0; i < 9999; i++) {
            wallets[senders[i]].balance = wallets[senders[i]].balance - money;
            wallets[recipients[i]].balance = wallets[recipients[i]].balance + money;
        }
    }

    function sendMoneyOneThousandTransactions(string[] memory senders, string[] memory recipients, uint money) public {
        for (uint i = 0; i < 999; i++) {
            wallets[senders[i]].balance = wallets[senders[i]].balance - money;
            wallets[recipients[i]].balance = wallets[recipients[i]].balance + money;
        }
    }

    function sendMoneyOneHundredTransactions(string[] memory senders, string[] memory recipients, uint money) public {
        for (uint i = 0; i < 99; i++) {
            wallets[senders[i]].balance = wallets[senders[i]].balance - money;
            wallets[recipients[i]].balance = wallets[recipients[i]].balance + money;
        }
    }

    function sendMoneyTenTransactions(string[] memory senders, string[] memory recipients, uint money) public {
        for (uint i = 0; i < 9; i++) {
            wallets[senders[i]].balance = wallets[senders[i]].balance - money;
            wallets[recipients[i]].balance = wallets[recipients[i]].balance + money;
        }
    }

}