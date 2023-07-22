package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// InsuranceContract is the chaincode definition
type InsuranceContract struct {
	contractapi.Contract
}

// Policy defines the structure of an insurance policy
type Policy struct {
	ID        string  `json:"id"`
	Holder    string  `json:"holder"`
	Type      string  `json:"type"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Premium   float64 `json:"premium"`
	Status    string  `json:"status"`
}

// Claim defines the structure of an insurance claim
type Claim struct {
	ID         string  `json:"id"`
	PolicyID   string  `json:"policyId"`
	Claimer    string  `json:"claimer"`
	ClaimDate  string  `json:"claimDate"`
	Amount     float64 `json:"amount"`
	IsApproved bool    `json:"isApproved"`
}

// Init initializes chaincode
func (ic *InsuranceContract) Init(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("Insurance chaincode initialized")
	return nil
}

// CreatePolicy creates a new insurance policy
func (ic *InsuranceContract) CreatePolicy(ctx contractapi.TransactionContextInterface, policyJSON string) error {
	var policy Policy
	err := json.Unmarshal([]byte(policyJSON), &policy)
	if err != nil {
		return fmt.Errorf("failed to unmarshal policy JSON: %v", err)
	}

	exists, err := ic.policyExists(ctx, policy.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("policy with ID %s already exists", policy.ID)
	}

	policyBytes, err := json.Marshal(policy)
	if err != nil {
		return fmt.Errorf("failed to marshal policy: %v", err)
	}

	err = ctx.GetStub().PutState(policy.ID, policyBytes)
	if err != nil {
		return fmt.Errorf("failed to put policy state: %v", err)
	}

	return nil
}

// GetPolicy returns the details of an insurance policy
func (ic *InsuranceContract) GetPolicy(ctx contractapi.TransactionContextInterface, policyID string) (*Policy, error) {
	policyBytes, err := ctx.GetStub().GetState(policyID)
	if err != nil {
		return nil, fmt.Errorf("failed to read policy state: %v", err)
	}
	if policyBytes == nil {
		return nil, fmt.Errorf("policy with ID %s does not exist", policyID)
	}

	var policy Policy
	err = json.Unmarshal(policyBytes, &policy)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal policy: %v", err)
	}

	return &policy, nil
}

// ClaimPolicy claims an insurance policy for a particular user
func (ic *InsuranceContract) ClaimPolicy(ctx contractapi.TransactionContextInterface, policyID string, claimJSON string) error {
	policy, err := ic.GetPolicy(ctx, policyID)
	if err != nil {
		return err
	}

	var claim Claim
	err = json.Unmarshal([]byte(claimJSON), &claim)
	if err != nil {
		return fmt.Errorf("failed to unmarshal claim JSON: %v", err)
	}

	if claim.PolicyID != policy.ID {
		return fmt.Errorf("claim policy ID does not match the provided policy ID")
	}

	claimBytes, err := json.Marshal(claim)
	if err != nil {
		return fmt.Errorf("failed to marshal claim: %v", err)
	}

	err = ctx.GetStub().PutState(claim.ID, claimBytes)
	if err != nil {
		return fmt.Errorf("failed to put claim state: %v", err)
	}

	return nil
}

// policyExists checks if a policy with a given ID exists
func (ic *InsuranceContract) policyExists(ctx contractapi.TransactionContextInterface, policyID string) (bool, error) {
	policyBytes, err := ctx.GetStub().GetState(policyID)
	if err != nil {
		return false, fmt.Errorf("failed to read policy state: %v", err)
	}

	return policyBytes != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&InsuranceContract{})
	if err != nil {
		fmt.Printf("Error creating insurance chaincode: %v\n", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting insurance chaincode: %v\n", err)
	}
}
