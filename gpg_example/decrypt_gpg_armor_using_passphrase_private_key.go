package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/openpgp"
)

func main() {
	// Read armored private key into type EntityList
	// An EntityList contains one or more Entities.
	// This assumes there is only one Entity involved
	/*
		entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(privateKey))
		if err != nil {
			log.Fatal(err)
		}
		entity := entitylist[0]
		fmt.Printf("Found %d entities\n", len(entitylist))
	*/

	// generate a private key
	ent, err := openpgp.NewEntity("1", "", "", nil)
	if err != nil {
		panic(err)
	}

	// encode
	buf := bytes.NewBuffer(nil)

	err = ent.PrivateKey.Serialize(buf)
	if err != nil {
		panic(err)
	}

	privkey := hex.EncodeToString(buf.Bytes())

	fmt.Println("Serialized Private Key:", privkey, len(privkey))

	/*
		entity, err := openpgp.NewEntity("1", "2", "", nil)
		if err != nil || entity.PrivateKey == nil {
			log.Fatal(err)
		}

		pkBuf := new(bytes.Buffer)
		err = entity.PrivateKey.Serialize(privBuf)
		if err != nil {
			log.Fatal(err)
		}

		// base64 encode
		str := base64.StdEncoding.EncodeToString(privBuf.Bytes())
		fmt.Println("Encoded Private Key:", str, len(str))

		str2 := privBuf.String()

		fmt.Println("Serialized Prvate Key:", str2, len(str2))
	*/
	/*
		// Decrypt private key using passphrase
		passphrase := []byte("golang")
		if entity.PrivateKey != nil && entity.PrivateKey.Encrypted {
			fmt.Println("Decrypting private key using passphrase")
			err := entity.PrivateKey.Decrypt(passphrase)
			if err != nil {
				fmt.Println("failed to decrypt key")
			}
		}
		for _, subkey := range entity.Subkeys {
			if subkey.PrivateKey != nil && subkey.PrivateKey.Encrypted {
				fmt.Printf("%#v\n", *subkey.PrivateKey)
				err := subkey.PrivateKey.Decrypt(passphrase)
				if err != nil {
					fmt.Println("failed to decrypt subkey")
				}
			}
		}

		// Decrypt armor encrypted message using decrypted private key
		decbuf := bytes.NewBuffer([]byte(encryptedMessage))
		result, err := armor.Decode(decbuf)
		if err != nil {
			log.Fatal(err)
		}

		md, err := openpgp.ReadMessage(result.Body, entitylist, nil , nil)
		if err != nil {
			fmt.Println("error reading message", err)
		}

		mdBody, err := ioutil.ReadAll(md.UnverifiedBody)
		fmt.Println("md:", string(mdBody))

		// Encrypt it again

		buf := new(bytes.Buffer)
		w, err := openpgp.Encrypt(buf, entitylist, nil, nil, nil)
		if err != nil {
		}

		// Print encrypted message
		const markMessage = "Test message encrypted using public key"
		_, err = w.Write([]byte(markMessage))
		if err != nil {
		}
		err = w.Close()
		if err != nil {
		}

		// Output as base64 encoded string
		bufBytes, err := ioutil.ReadAll(buf)
		str := base64.StdEncoding.EncodeToString(bufBytes)
		fmt.Println("Public key encrypted message (base64 encoded):", str)

		// Output as armored string
		outBuf := new(bytes.Buffer)
		pgpWriter, err := armor.Encode(outBuf, "NEW PGP ENCRYPTED MESSAGE", nil)
		pgpWriter.Write(bufBytes)
		pgpWriter.Close()
		fmt.Println(outBuf.String())
	*/
}

// pub   1024R/7F98BBCE 2014-01-04
// uid                  Golang Test (Private key password is 'golang') <golangtest@test.com>
// sub   1024R/5F34A320 2014-01-04
const privateKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1

lQH+BFLHbYYBBADCjgKHmPmwBxI3c3DPVoSdu0+EJl/EsS2HEaN63dnLkGsMAs+4
32wsywmMrzKqCL40sbhJVYBcfe0chL+cry4O54DX7+gA0ZSVzFUN2EGocnkaHzyS
fuUtBdCTmoWZZAGFiBwlIS7aE/86SOyHksFo8LRC9W/GIWQS2PbcadvUywARAQAB
/gMDApJxOwcsfChBYCCmhOAvotKdYcy7nuG7dyGDBlpclLJtH/PaakKSE33NtEj4
1fyixQOdwApxvuQ2P0VX3pie/De1KpbeqXfnPLsmsXQwrRPOo38T5zeJ5ToWUGDC
Oia69ep3kmHbAW41EBH/uk/nMM91QUdl4mkYsc3dhVOXbmf0xyRoP/Afqha4UhdZ
0XKlIZP1a5+3NF/Q6dAVG0+FlO5Hcai8n98jW0id8Yf6zI+1gFGvYYKhlifkdJeK
Nf4YEvOXALEvaQqkcJOxEca+BmqsgCIFctJe9Bahx97Ep5hP7AH0aBmtZfmGmZwB
GYoevUtKa4ASVmK8RaddBvIjcrWsoAsYMpDGYaE0fcdtxsBf3uT1Q8IMsT+ZRjjV
TfvJ8aW14ZrLI98KdtXaOPZs91mML+3iw1c/1O/IEJfwxrUni2p/fDmCYU9eHR3u
Q0PwVR0MCUHI1fGuUoetW2gYIxfklvBtEFWW1BD6fCpCtERHb2xhbmcgVGVzdCAo
UHJpdmF0ZSBrZXkgcGFzc3dvcmQgaXMgJ2dvbGFuZycpIDxnb2xhbmd0ZXN0QHRl
c3QuY29tPoi4BBMBAgAiBQJSx22GAhsDBgsJCAcDAgYVCAIJCgsEFgIDAQIeAQIX
gAAKCRBVSiCHf5i7zqKJA/sFUM2TfL2VZKWC7E1N1wwZctB9Bf77SeAPSVpGCZ0c
iUYIFdwwGowKtjoDrsbYgPp+UGOyYMD6tGzWKaJrQQoDyaQqVVRhbNXB7Jz7JT2a
qKHD1t7cx5FfUzDMBNou3TOWHomDXyQGDAULAZnjaOj8/pDe6poxyBluSjMJUzfD
pp0B/gRSx22GAQQArUMDqkGng9Cppk73UBWBd7jhhbtk0eaRQh/goUHhKJerZ4LM
Q21IKyIX+GQbscDpccpXMI6eThXxrL+D8G4cNb4ewvT0zc20+T91ztgT9A/4Vifc
EPQCErTqY/oZphAzZM1p6sRenc22e42iT0Iibd5gCs2wnSNeUzybDcuQi2EAEQEA
Af4DAwKScTsHLHwoQWCYayWqio8purPTonYogZSN3QwaheS2Y0NE7skdLOvP97vi
Rh7BktS6Dkgu0T3D39+q0O6ZO7XErvTVoas1F0HXzId4tiIicmx4tYNyWI4NrSO7
6TQPz/bQe8ZN+plG5cgZowts6g6RSfQxoW21LrP8Lh+OEdcYwWf7BTukAYmD3oq9
RxdfYI7hnbVGFdOqQUQNcxZkbdrsF9ITjQb/KRln5/99E1Kp1D45VpPOs7NT3orA
mnfSslJXVNm1uK6FDBX2iUe3JaAmgh+RLGXQXRZKJW4DGDTyYdwR4hO8cYix2+8z
+XuwdVDPKBnzKn190m6xpdLyvKfj1BQhX14NShPQZ3QJiMU0k4Js23XSsWs9NSxI
FjjE9/mOFVUH25KN+X7rzBPo2S0pMQLqyQxSLIdI2LPDxzlknctT6OoBPKPJjb7S
Lt5GhIA5Cz+cohfX6LePG4FkvwU32tTRBz5YNhFBizmS+YifBBgBAgAJBQJSx22G
AhsMAAoJEFVKIId/mLvOulED/2uUh/qjOT468XoK6Xt837w45JQPpLqiGH9KJgqF
rUxJMw1bIE2G606OY6hCgeE+YC8qny29hQtXhKIquUI/0A1qK3aCZhwqyqT+QjvF
6Xi0i/HrgQwCyBopY3uGndMbvthxU0KO0d6seMZltHDr8YaU1JvDwNFDQVuw+Rqy
57ET
=nvLl
-----END PGP PRIVATE KEY BLOCK-----`

// Encrypted by public key message for
// 1024R/5F34A320 2014-01-04 "Golang Test (Private key password is 'golang') <golangtest@test.com>"
const encryptedMessage = `-----BEGIN PGP MESSAGE-----                  
Version: GnuPG v1

hIwDBZMeL180oyABA/9DnQDx2QSk5k7f1JsqmK1kgoeEzBdNJ30DYlF0CBtHi5pQ
yx0Y8EOaCXC7GLKyJXqTf+3KeVpXhC3YhVjHB4W5mDv9ifM91TRFyLm8OCmYLqV1
U9OANW6l3aQPG4FH/b/JnS74NCu2uEmWuBve8Qb808KSiDMclRHsAomV6XmNfdJh
Af37lgu/6AFbgVMBmpfkPLtn2BYLitoeKh5IbF+Fzz5dk38Ij6P94nvPN3chaq8f
/GsRWGttUkTsCE1YMZgPrr/uSGpxlNrdAY7KEhFHovLtsdJvMth5/n9IrDUVagwt
nQ==
=lNST
-----END PGP MESSAGE-----`
