// This transaction is a template for a transaction that
// could be used by anyone to send tokens to another account
// that has been set up to receive tokens.
//
// The withdraw amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x%s
import FlowToken from 0x%s

transaction (amount: UFix64, to: Address, senderAddress: Address) {

	// The Vault resource that holds the tokens that are being transferred
	let sentVault: @FungibleToken.Vault

	prepare(signer: AuthAccount) {

		// Get a reference to the signer's stored vault
		let vaultRef = signer.borrow<&AnyResource{FungibleToken.Provider}>(from: /storage/flowTokenVault)
			?? panic("Could not borrow reference to the owner's Vault!")

		// Withdraw tokens from the signer's stored vault
		self.sentVault <- vaultRef.withdraw(amount: amount)
	}

	execute {

		// Get the recipient's public account object
		let recipient = getAccount(to)
		// let sender = getAccount(senderAddress)

		// Get a reference to the recipient's Receiver
		let receiverRef = recipient.getCapability(/public/flowTokenReceiver)
			.borrow<&{FungibleToken.Receiver}>()
			?? panic("Could not borrow receiver reference to the recipient's Vault")

        let balanceRefTo =  recipient.getCapability(/public/flowTokenBalance)
            .borrow<&{FungibleToken.Balance}>()
                ?? panic("failed to borrow reference to recipient vault balance")
        // let balanceRefSender =  sender.getCapability(/public/flowTokenBalance)
            //.borrow<&{FungibleToken.Balance}>()
                //?? panic("failed to borrow reference to sender vault balance")


        let preBalanceTo = balanceRefTo.balance
        //let preBalanceSender = balanceRefSender.balance
		// Deposit the withdrawn tokens in the recipient's receiver
		receiverRef.deposit(from: <-self.sentVault)
	    let postBalanceTo = balanceRefTo.balance
	    //let postBalanceSender = balanceRefSender.balance

        if postBalanceTo != preBalanceTo + amount {
            panic("Receiver balance was not transferred correctly.")
        }
	}
}
