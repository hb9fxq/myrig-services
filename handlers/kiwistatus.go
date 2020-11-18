package handlers

import (
	"encoding/json"
	"github.com/hb9fxq/myrig-services/globals"
	"net/http"
	"strconv"
)

type Kiwistatus struct {
	HeaderText string
}

func KiwistatusHandler(route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(http.StatusOK)

		kiwistatus := &Kiwistatus{}
		kiwistatus.HeaderText = ""

		strRotorOpti := strconv.Itoa(globals.GlobalAppCtx.RotorStatus.Deg)
		strRotorLoop := strconv.Itoa(globals.GlobalAppCtx.LoopRotorStatus.Deg)

		switch globals.GlobalAppCtx.AntennaStatus.Ant { // TODO: add antenna map to app config....
		case "1B":
			kiwistatus.HeaderText += "Optibeam@" + strRotorOpti + "°"
		case "2B":
			kiwistatus.HeaderText += "Longwire"
		case "3B":
			kiwistatus.HeaderText += "MagneticLoop@" + strRotorLoop + "°"
		default:
			kiwistatus.HeaderText = "FLEXRADIO is in use. KIWI SDR disabled, come back later :-)"
		}

		if err := json.NewEncoder(w).Encode(kiwistatus); err != nil {
			panic(err)
		}
	}
}
