transaction(publicKey: [UInt8]) {
	prepare(signer: AuthAccount) {
		let acct = AuthAccount(payer: signer)
		acct.addPublicKey(publicKey)
	}
}
