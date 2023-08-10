# tntengine
**tntengine** is a *golang* implementation of a Z-80 assembler program described in an article in Dr. Dobbs Journal Volume 9, Number 94, 1984 titled [*An Infinite Key Encryption System*](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up) by John A. Thomas and Joan Thersites.  I will not be detailing the reasoning behind the design of the system, but instead refer you to the article for those detail.  I will be describing the design of this _golang_ implementation and what is different from the original code.

The first change is that *__tntengine__* is a module that only contains the code necessary to support the rotors and permutator that encrypts the plaintext.  The code that reads the file that to be encrypted/decrypted is put in the *__tnt__* program.

