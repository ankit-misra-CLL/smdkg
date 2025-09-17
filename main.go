package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	//"math/big"

	"github.com/ankit-misra-CLL/smdkg/p256keyring"
	"github.com/ankit-misra-CLL/smdkg/utils/crypto/dkgtypes"

	//"github.com/ankit-misra-CLL/smdkg/utils/crypto/math"
	"github.com/ankit-misra-CLL/smdkg/utils/crypto/p256keyringshim"
	//"github.com/ankit-misra-CLL/smdkg/utils/crypto/vess"
)

func main() {
	// curve := math.CurveByName("P256")
	// iid := "dummy_iid"
	// tag := "dummy_tag"
	n := 2
	// t := 2

	fmt.Println("generating recipient key pairs")
	krs := make([]dkgtypes.P256Keyring, n)
	pks := make([]dkgtypes.P256PublicKey, n)
	for i := 0; i < n; i += 1 {
		kr, err := p256keyring.New(rand.Reader)
		if err != nil {
			log.Fatalf("error generating keyring for party %d: %v", i, err)
		}
		krInternal, err := p256keyringshim.New(kr)
		if err != nil {
			log.Fatalf("error generating keyring for party %d: %v", i, err)
		}
		krs[i] = krInternal
		pks[i] = krInternal.PublicKey()

		kr_mar, _ := kr.MarshalBinary()
		kr_mar_hex := hex.EncodeToString(kr_mar)
		fmt.Printf("keyring %d: %s\n", i, kr_mar_hex)

		kr_mar2, _ := hex.DecodeString(kr_mar_hex)
		kr2, _ := p256keyring.New(rand.Reader)
		_ = kr2.UnmarshalBinary(kr_mar2)

		if kr.String() != kr2.String() {
			fmt.Println("noooooo")
		}

		fmt.Printf("public key %d: %s\n", i, hex.EncodeToString(pks[i].Bytes()))
	}

	// fmt.Println("generating new vess instance")
	// v, err := vess.NewVESS(curve, dkgtypes.InstanceID(iid), tag, n, t, pks)
	// if err != nil {
	// 	log.Fatalf("error generating new vess instance: %v", err)
	// }

	// fmt.Println("generating secret")
	// modulus := curve.GroupOrder()
	// s := math.NewScalarFromString("165186230875895249184305411371791848164", modulus)

	// fmt.Println("generating dealing")
	// dealing, err := v.Deal(s, []byte{}, rand.Reader)
	// if err != nil {
	// 	log.Fatalf("error generating new dealing: %v", err)
	// }

	// fmt.Println("verifying dealing")
	// ver_dealing, err := v.VerifyDealing(&dealing.UnverifiedDealing, []byte{})
	// if err != nil {
	// 	log.Fatalf("error verifying the dealing: %v", err)
	// }

	// fmt.Println("decrypting dealing")
	// dec_values := make([]math.Scalar, n)
	// for i := 0; i < n; i += 1 {
	// 	dec_val, err := v.Decrypt(i, krs[i], ver_dealing, []byte{})
	// 	if err != nil {
	// 		log.Fatalf("error decrypting the dealing for party %d: %v", i, err)
	// 	}
	// 	dec_values[i] = dec_val
	// }

	// fmt.Println("verifying decrypted shares")
	// for i := 0; i < n; i += 1 {
	// 	err := v.VerifyShare(dec_values[i], ver_dealing, i)
	// 	if err != nil {
	// 		log.Fatalf("failed to verify share of party %d: %v", i, err)
	// 	}
	// }

	// hexNums := make([]string, n)
	// fmt.Println("printing shares")
	// for i := 0; i < n; i += 1 {
	// 	share_str := hex.EncodeToString(dec_values[i].Bytes())
	// 	hexNums[i] = "0x" + share_str
	// 	fmt.Printf("Party %d: %s\n", i, hexNums[i])
	// }

	// // num1 := new(big.Int)
	// // num2 := new(big.Int)
	// // _, ok1 := num1.SetString(hexNums[0], 0)
	// // if !ok1 {
	// // 	fmt.Println("Error parsing hexNum1")
	// // 	return
	// // }
	// // _, ok2 := num2.SetString(hexNums[1], 0)
	// // if !ok2 {
	// // 	fmt.Println("Error parsing hexNum2")
	// // 	return
	// // }
	// // sum1 := new(big.Int)
	// // sum2 := new(big.Int)
	// // sum1.Add(num1, num1)
	// // sum2.Sub(sum1, num2)
	// // fmt.Printf("Number 1: %s\n", hexNums[0])
	// // fmt.Printf("Number 2: %s\n", hexNums[1])
	// // fmt.Printf("Sum(hex): 0x%x\n", sum2)

	// // modulusHex := "0x" + hex.EncodeToString(modulus.Bytes())
	// // mod := new(big.Int)
	// // _, ok3 := mod.SetString(modulusHex, 0)
	// // if !ok3 {
	// // 	fmt.Println("Error parsing modulusHex")
	// // 	return
	// // }
	// // ans := new(big.Int)
	// // ans.Mod(sum2, mod)
	// // fmt.Printf("Sum(hex): 0x%x\n", ans)



	// // aes_key_hex := "7c45b5f9e6a0e2acddd68e98768492e4"
	// // nn := new(big.Int)
	// // _, success := nn.SetString(aes_key_hex, 16)
	// // if !success {
	// // 	fmt.Println("Error setting string")
	// // }
	// // aes_key_dec := nn.String()
	// // fmt.Printf("key: %s\n", aes_key_dec)

	// var nums [4]*big.Int
	// for i := 0; i < n; i += 1 {
	// 	nums[i] = new(big.Int)
	// 	_, _ = nums[i].SetString(hexNums[i], 0)
	// }
	// modulusHex := "0x" + hex.EncodeToString(modulus.Bytes())
	// mod := new(big.Int)
	// _, _ = mod.SetString(modulusHex, 0)
	// for i := 0; i < n; i += 1 {
	// 	for j := 0; j < n; j += 1 {
	// 		if i >= j {
	// 			continue
	// 		}

	// 		ydiff := new(big.Int)
	// 		ydiff.Sub(nums[j], nums[i])
	// 		xdiff := new(big.Int).SetInt64(int64(j-i))

	// 		xdiffinv := new(big.Int)
	// 		xdiffinv.ModInverse(xdiff, mod)

	// 		m := new(big.Int)
	// 		m.Mul(ydiff, xdiffinv)
	// 		m.Mod(m, mod)

	// 		fmt.Printf("Slope for parties %d and %d: %x\n", i+1, j+1, m)

	// 		c := new(big.Int)
	// 		xi := new(big.Int).SetInt64(int64(i+1))
	// 		c.Mul(m, xi)
	// 		c.Sub(nums[i], c)
	// 		c.Mod(c, mod)
	// 		fmt.Printf("secret: %x\n", c)
	// 	}
	// }


}