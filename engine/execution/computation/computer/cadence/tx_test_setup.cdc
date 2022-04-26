// Setup Account

import FlowToken from 0x%s

// This transaction configures an account to store and receive tokens defined by
// the ExampleToken contract.
transaction (addr: Address) {
	prepare(acct: AuthAccount) {
		// Create a new empty Vault object
		let vaultA <- FlowToken.createEmptyVault()

		// Store the vault in the account storage
		acct.save<@FlowToken.Vault>(<-vaultA, to: /storage/flowTokenVault)

        log("Empty Vault stored")

        // Create a public Receiver capability to the Vault
		let ReceiverRef = acct.link<&FlowToken.Vault{FlowToken.Receiver, FlowToken.Balance}>(/public/flowTokenReceiver, target: /storage/flowTokenVault)

        log("References created")
	}

    post {
        // Check that the capabilities were created correctly
        getAccount(addr).getCapability<&FlowToken.Vault{FlowToken.Receiver}>(/public/flowTokenReceiver)
                        .check():
                        "Vault Receiver Reference was not created correctly"
    }
}
