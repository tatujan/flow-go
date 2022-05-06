import FungibleToken from 0x%s
import FlowToken from 0x%s

transaction(amount: UFix64, recipient: Address) {
	let sentVault: @FungibleToken.Vault
// These are transaction field variables
//	var preBalance: UFix64
//	var postBalance: UFix64

	prepare(signer: AuthAccount) {
		let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
			?? panic("failed to borrow reference to sender vault")
		self.sentVault <- vaultRef.withdraw(amount: amount)
	}
	execute {
		let receiverRef =  getAccount(recipient)
		    .getCapability(/public/flowTokenReceiver)
		    .borrow<&{FungibleToken.Receiver}>()
			    ?? panic("failed to borrow reference to recipient vault receiver")

		let balanceRef =  getAccount(recipient)
            .getCapability(/public/flowTokenBalance)
        	.borrow<&{FungibleToken.Balance}>()
        		?? panic("failed to borrow reference to recipient vault balance")

        // the preBalance/postBalance transfer check should live in the post conditions (failure in the post conditions
        // will reverse the transaction), not sure why it doesn't work here
        let preBalance = balanceRef.balance
        // Deposit the tokens in the recipient's receiver
        receiverRef.deposit(from: <-self.sentVault)
        let postBalance = balanceRef.balance

        if postBalance != preBalance + amount {
            panic("Account was not funded correctly.")
        }

        // This should work, not sure what's wrong with the syntax
        // save pre balance
        //self.preBalance = balanceRef.balance
        // deposit amount
		//receiverRef.deposit(from: <-self.sentVault)
        // save post balance
        //self.postBalance = balanceRef.balance
	}
	post {
	    // This should work, not sure why
	    // self.postBalance == self.preBalance + amount : "Balance was not transferred correctly."
	}
}
