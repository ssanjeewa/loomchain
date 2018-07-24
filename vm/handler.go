package vm

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/loomchain"
	"github.com/loomnetwork/loomchain/auth"
	"github.com/loomnetwork/loomchain/eth/utils"
	"github.com/loomnetwork/loomchain/log"
	"github.com/loomnetwork/loomchain/registry"
)

type DeployTxHandler struct {
	*Manager
}

func (h *DeployTxHandler) ProcessTx(
	state loomchain.State,
	txBytes []byte,
) (loomchain.TxHandlerResult, error) {
	var r loomchain.TxHandlerResult

	var msg MessageTx
	err := proto.Unmarshal(txBytes, &msg)
	if err != nil {
		return r, err
	}

	origin := auth.Origin(state.Context())
	caller := loom.UnmarshalAddressPB(msg.From)

	if caller.Compare(origin) != 0 {
		if origin.Local.Compare(caller.Local) != 0 {
			return r, fmt.Errorf("Origin doesn't match caller: - %v != %v", origin, caller)
		} else {
			//TODO investigate why the client is bugged
			log.Error("Local address same but chainID is wrong, allowing transaction to go through: %v != %v", origin, caller)
		}
	}

	var tx DeployTx
	err = proto.Unmarshal(msg.Data, &tx)
	if err != nil {
		return r, err
	}

	vm, err := h.Manager.InitVM(tx.VmType, state)
	if err != nil {
		return r, err
	}

	retCreate, addr, errCreate := vm.Create(origin, tx.Code)

	response, errMarshal := proto.Marshal(&DeployResponse{
		Contract: &types.Address{
			ChainId: addr.ChainID,
			Local:   addr.Local,
		},
		Output: retCreate,
	})
	if errMarshal != nil {
		if errCreate != nil {
			return r, errors.Wrapf(errCreate, "[DeployTxHandler] Error deploying EVM contract on create")
		} else {
			return r, errors.Wrapf(errMarshal, "[DeployTxHandler] Error deploying EVM contract on marshaling evm error")
		}
	}
	r.Data = append(r.Data, response...)
	if errCreate != nil {
		return r, errors.Wrapf(errCreate, "[DeployTxHandler] Error deploying EVM contract on create")
	}

	if len(tx.Name) > 0 {
		reg := &registry.StateRegistry{
			State: state,
		}
		reg.Register(tx.Name, addr, caller)
	}
	if tx.VmType == VMType_EVM {
		r.Info = utils.DeployEvm
	} else {
		r.Info = utils.DeployPlugin
	}
	return r, nil
}

type CallTxHandler struct {
	*Manager
}

func (h *CallTxHandler) ProcessTx(
	state loomchain.State,
	txBytes []byte,
) (loomchain.TxHandlerResult, error) {
	var r loomchain.TxHandlerResult

	var msg MessageTx
	err := proto.Unmarshal(txBytes, &msg)
	if err != nil {
		return r, err
	}

	origin := auth.Origin(state.Context())
	caller := loom.UnmarshalAddressPB(msg.From)
	addr := loom.UnmarshalAddressPB(msg.To)

	if caller.Compare(origin) != 0 {
		if origin.Local.Compare(caller.Local) != 0 {
			return r, fmt.Errorf("Origin doesn't match caller: - %v != %v", origin, caller)
		} else {
			//TODO investigate why the client is bugged
			log.Error("Local address same but chainID is wrong, allowing transaction to go through: %v != %v", origin, caller)
		}
	}

	var tx CallTx
	err = proto.Unmarshal(msg.Data, &tx)
	if err != nil {
		return r, err
	}

	vm, err := h.Manager.InitVM(tx.VmType, state)
	if err != nil {
		return r, err
	}

	r.Data, err = vm.Call(origin, addr, tx.Input)
	if err != nil {
		return r, err
	}
	if tx.VmType == VMType_EVM {
		r.Info = utils.CallEVM
	} else {
		r.Info = utils.CallPlugin
	}
	return r, err
}
