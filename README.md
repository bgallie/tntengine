
# tntengine

**tntengine** is a *golang* implementation of a Z-80 assembler program described in an article in Dr. Dobbs Journal Volume 9, Number 94, 1984 titled [*An Infinite Key Encryption System*](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up) by John A. Thomas and Joan Thersites.  I will not be detailing the reasoning behind the design of the system, but instead refer you to the article for those detail.  I will be describing the design of this *golang* implementation and what is different from the original code.

The first change is that program is broken out into two pieces:

* `tntengine`, a module that only contains the code necessary to support the rotors and permutator that encrypts/decrypts the plaintext.  This module is essentially a pseudo-random number generator with a very long period.
* `tnt`, The code that reads the file that to be encrypted/decrypted.

n the original program, the rotors and permutator (note that there is only one permutator that is used twice) are applied sequentially then the rotors are stepped and the permutator is cycled.  The next change is that the *golang* implementation makes use of *golang*'s built-in concurrency features to run the rotors and permutators concurrently, which means the rotors are applied and stepped individually along with two _**identical**_ permutators (instead of one) that are applied and cycled individually.

To support this, an interface called `Crypter` wad defined:
```
type Crypter interface {
	Update(*Rand)                   // function to update the rotor/permutator
	SetIndex(*big.Int)              // setter for the index value
	Index() *big.Int                // getter for the index value
	ApplyF(CipherBlock) CipherBlock // encryption function
	ApplyG(CipherBlock) CipherBlock // decryption function
}
```
There are three types, `Rotor`, `Permutator`, and `Counter` that satisfies the `Crypter` interface.  The `Rotor` and `Permutator` types provides the rotors and permutators used to encrypt/decrypt the plaintext.  `Counter`, is a type that just counts the number of 32-byte blocks that have been encrypted/decrypted.
