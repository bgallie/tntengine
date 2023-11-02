
# tntengine

**tntengine** is a *golang* implementation of a Z-80 assembler program described in an article in Dr. Dobbs Journal Volume 9, Number 94, 1984 titled [*An Infinite Key Encryption System*](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up) by John A. Thomas and Joan Thersites.  I will not be detailing the reasoning behind the design of the system, but instead refer you to the article for those detail.  I will be describing the design of this *golang* implementation and what is different from the original code.

The first change is that program is broken out into two pieces:

* `tntengine`, a module that only contains the code necessary to support the rotors and permutator that encrypts/decrypts the plaintext.  it contains the code to initialize the rotors and permutators based on a supplied secret key.  This module is essentially a pseudo-random number generator with a very long period.
* `tnt`, the code that obtains the secret key, the file that to be encrypted/decrypted, the starting block number, and saves the next starting block number after the encryption is done.

In the original program, the rotors and permutator (note that there is only one permutator that is used twice) are applied sequentially then the rotors are stepped and the permutator is cycled.  The next change is that the *golang* implementation makes use of *golang*'s built-in concurrency features to run the rotors and permutators concurrently, which means the rotors are applied and stepped individually along with two _**identical**_ permutators (instead of one) that are applied and cycled individually.  This simulates the original code where one permutator is used twice before it is cycled.

To support this, an interface called `Crypter` wad defined:
```go
type Crypter interface {
	Update(*Rand)                   // function to update the Crypter (rotor or permutator)
	SetIndex(*big.Int)              // setter for the index value
	Index() *big.Int                // getter for the index value
	ApplyF(CipherBlock) CipherBlock // encryption function
	ApplyG(CipherBlock) CipherBlock // decryption function
}
````
There are three types, `Rotor`, `Permutator`, and `Counter` that satisfies the `Crypter` interface.  The `Rotor` and `Permutator` types provides the rotors and permutators used to encrypt/decrypt the plaintext.  `Counter`, is a type that just counts the number of 32-byte blocks that have been encrypted/decrypted.

The `Rotor` and `Permutator` objects are wrapped in a go function that reads a `CypherBlock` from the input channel, calls the `Applyf` or `Applyg` function depending on whether the file is being encrypted or decrypted, and then sends processed `CiperBlock` to the output channel.  The wrapped `Rotor` and `Permutator` objects are then chained together by connecting the ouput channel of one wrapped object to the input channel of the following object.  The input of the first object in the chain is feed the data to be encrypted/decrypted and the output of the last object is the encrypted/decrypted data.
To initialize tntengine, the (hard-coded) proforma rotors and permutator are used to build the proforma machine using the sequece of rotor1, rotor2, permutator1, rotor3, rotor4, permutator1, rotor5 rotor6

![proforma Encryption Machine][def]

Once the TntEngine.engine is set up with the proforma machine, an encryption machine is created by linking the rotors and permutatures togeher by wrapping them in a go function that reads a `CypherBlock` from the input channel, calls the `Applyf` function then sends processed `CiperBlock` to the output channel.  Now a random number generator is created so that we can update the rotors and permutator with new (psudo) random data in a very non-linear manner.  The rotors are updated `CipherBlockBytes` bytes at a time to add additional complexity since the rotors are being modified as they are bing used.

Once the permutator and rotors have been updated, a new TntEngine.engine is created using the rotors in a random order and the rotors and permutator is used to create a new encryption machine.  In the diagram below, the random order is [Rotor 3, Rotor 5, Rotor 1, Rotor 6, Rotor 4, Rotor 2].

![Updated Encryption Machine][def2]

To create a decryption machine, the rotors are taken in reveres order.

![Updated Decryption Machine][def3]

[def]: assets/images/proformaEncryption.png
[def2]: assets/images/updatedEncryption.png
[def3]: assets/images/updatedDecryption.png