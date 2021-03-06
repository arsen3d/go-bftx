// File: ./blockfreight/lib/validator/validator.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith, Gian Felipe & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright © 2017 Blockfreight, Inc. All Rights Reserved.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

// Package validator is a package that provides functions to assure the input JSON is correct.
package validator

import (
	// =======================
	// Golang Standard library
	// =======================
	"errors"  // Implements functions to manipulate errors.
	"reflect" // Implements run-time reflection, allowing a program to manipulate objects with arbitrary types.

	// ======================
	// Blockfreight™ packages
	// ======================
	"github.com/blockfreight/go-bftx/lib/app/bf_tx" // Defines the Blockfreight™ Transaction (BF_TX) transaction standard and provides some useful functions to work with the BF_TX.
)

// ValidateBFTX is a function that receives the BF_TX and return the proper message according with the result of ValidateFields function.
func ValidateBFTX(bftx bf_tx.BF_TX) (string, error) {
	var espErr error

	valid, err := ValidateFields(bftx)
	if valid {
		return "Success! [OK]", nil
	}

	if err != "" {
		espErr = errors.New(`
    Specific Error [01]:
    ` + err)
	}
	return `
    Blockfreight, Inc. © 2017. Open Source (MIT) License.

    Error [01]:

    Invalid structure in JSON provided. JSON 结构无效.
    Struttura JSON non valido. هيكل JSON صالح. 無効なJSON構造.
    Estructura inválida en el JSON dado.

    support: support@blockfreight.com`, espErr
}

// ValidateFields is a function that receives the BF_TX, validates every field in the BF_TX and return true or false, and a message if some field is wrong.
func ValidateFields(bftx bf_tx.BF_TX) (bool, string) {
	if reflect.TypeOf(bftx.Type) != reflect.TypeOf("s") {
		return false, "bftx.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.Shipper.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.Shipper.Type is not a string."
	}
	if (reflect.TypeOf(bftx.Properties.BolNum.Type) != reflect.TypeOf(1)) || bftx.Properties.BolNum.Type == 0 {
		return false, "bftx.Properties.BolNum.Type is not a number."
	}
	if (reflect.TypeOf(bftx.Properties.RefNum.Type) != reflect.TypeOf(1)) || bftx.Properties.RefNum.Type == 0 {
		return false, "bftx.Properties.RefNum.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.Consignee.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.Consignee.Type is not a string."
	}
	if (reflect.TypeOf(bftx.Properties.Vessel.Type) != reflect.TypeOf(1)) || bftx.Properties.Vessel.Type == 0 {
		return false, "bftx.Properties.Vessel.Type is not a number."
	}
	if (reflect.TypeOf(bftx.Properties.PortOfLoading.Type) != reflect.TypeOf(1)) || bftx.Properties.PortOfLoading.Type == 0 {
		return false, "bftx.Properties.PortOfLoading.Type is not a number."
	}
	if (reflect.TypeOf(bftx.Properties.PortOfDischarge.Type) != reflect.TypeOf(1)) || bftx.Properties.PortOfDischarge.Type == 0 {
		return false, "bftx.Properties.PortOfDischarge.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.NotifyAddress.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.NotifyAddress.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.DescOfGoods.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.DescOfGoods.Type is not a string."
	}
	if (reflect.TypeOf(bftx.Properties.GrossWeight.Type) != reflect.TypeOf(1)) || bftx.Properties.GrossWeight.Type == 0 {
		return false, "bftx.Properties.GrossWeight.Type is not a number."
	}
	if (reflect.TypeOf(bftx.Properties.FreightPayableAmt.Type) != reflect.TypeOf(1)) || bftx.Properties.FreightPayableAmt.Type == 0 {
		return false, "bftx.Properties.FreightPayableAmt.Type is not a number."
	}
	if (reflect.TypeOf(bftx.Properties.FreightAdvAmt.Type) != reflect.TypeOf(1)) || bftx.Properties.FreightAdvAmt.Type == 0 {
		return false, "bftx.Properties.FreightAdvAmt.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.GeneralInstructions.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.GeneralInstructions.Type is not a string."
	}
	if (reflect.TypeOf(bftx.Properties.DateShipped.Type) != reflect.TypeOf(1)) || bftx.Properties.DateShipped.Type == 0 {
		return false, "bftx.Properties.DateShipped.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.DateShipped.Format) != reflect.TypeOf("s") {
		return false, "bftx.Properties.DateShipped.Format is not a date format."
	}
	if reflect.TypeOf(bftx.Properties.IssueDetails.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.IssueDetails.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.IssueDetails.Properties.PlaceOfIssue.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.IssueDetails.Properties.PlaceOfIssue.Type is not a string."
	}
	if (reflect.TypeOf(bftx.Properties.IssueDetails.Properties.DateOfIssue.Type) != reflect.TypeOf(1)) || bftx.Properties.IssueDetails.Properties.DateOfIssue.Type == 0 {
		return false, "bftx.Properties.IssueDetails.Properties.DateOfIssue.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.IssueDetails.Properties.DateOfIssue.Format) != reflect.TypeOf("s") {
		return false, "bftx.Properties.IssueDetails.Properties.PlaceOfIssue.Format is not a date format."
	}
	if (reflect.TypeOf(bftx.Properties.NumBol.Type) != reflect.TypeOf(1)) || bftx.Properties.NumBol.Type == 0 {
		return false, "bftx.Properties.NumBol.Type is not a number."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.Properties.FirstName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.FirstName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.Properties.LastName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.LastName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.MasterInfo.Properties.Sig.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.MasterInfo.Properties.Sig.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.Properties.FirstName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.FirstName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.Properties.LastName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.LastName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForMaster.Properties.Sig.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForMaster.Properties.Sig.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Properties.FirstName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.FirstName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Properties.LastName.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.LastName.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Properties.Sig.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.Sig.Type is not a string."
	}
	if reflect.TypeOf(bftx.Properties.AgentForOwner.Properties.ConditionsForCarriage.Type) != reflect.TypeOf("s") {
		return false, "bftx.Properties.AgentForOwner.Properties.ConditionsForCarriage.Type is not a string."
	}
	return true, ""
}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================
