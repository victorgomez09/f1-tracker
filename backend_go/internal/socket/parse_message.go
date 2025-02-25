package internal

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"log"

	f1Model "github.com/victorgomez09/f1-tracker.git/internal/model"
)

// func parseCompressed(data string) {
// 	return JSON.parse(
// 	  new TextDecoder().decode(inflateRawSync(Buffer.from(data, "base64")))
// 	)
//   }

func Decompress(input []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(input))
	decompressed, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return decompressed, nil
}

func UpdateState(state f1Model.F1State, data f1Model.SocketData) f1Model.F1State {
	if len(data.M) > 0 {
		for _, message := range data.M {
			if message.M != "feed" {
				continue
			}

			var test = message.A
			fmt.Println("test", test)
		}
	}

	fmt.Println("if cond", data.R.CarData)
	if (data.R.CarData != "") && (data.I == "1") {
		// carData, err := Decompress([]byte(data.R.CarData))
		// if err != nil {
		// 	log.Fatalln("Error decompresing", err)
		// }
		// var parsedData = f1Model.ParsedRecap{
		// 	CarData: carData,
		// }
		buffer := bytes.NewBuffer([]byte(data.R.CarData))
		r, err := zlib.NewReader(buffer)
		if err != nil {
			log.Fatalln("Error decompresing", err)
		}
		fmt.Println("r", r)
	}

	// 	if (data.R && data.I === "1") {
	// 	  const parsedData: ParsedRecap = {
	// 		...(data.R["CarData.z"] && {
	// 		  CarData: parseCompressed(data.R["CarData.z"]),
	// 		}),
	// 		...(data.R["Position.z"] && {
	// 		  Position: parseCompressed(data.R["Position.z"]),
	// 		}),
	// 	  };

	// 	  const {
	// 		"CarData.z": z1,
	// 		"Position.z": z2,
	// 		...newState
	// 	  } = { ...data.R, ...parsedData };

	// 	  return merge(state, newState) ?? state;
	// 	}

	return state
}

// export const updateState = (state: F1State, data: SocketData): F1State => {
// 	if (data.M) {
// 	  for (const message of data.M) {
// 		if (message.M !== "feed") continue;

// 		let [cat, msg] = message.A;

// 		let parsedMsg: null | F1CarData | F1Position = null;
// 		let parsedCat: null | string = null;

// 		if (
// 		  (cat === "CarData.z" || cat === "Position.z") &&
// 		  typeof msg === "string"
// 		) {
// 		  parsedCat = cat.split(".")[0];
// 		  parsedMsg = parseCompressed<F1CarData | F1Position>(msg);
// 		}

// 		state = merge(state, { [parsedCat ?? cat]: parsedMsg ?? msg }) ?? state;
// 	  }

// 	  return state;
// 	}

// 	if (data.R && data.I === "1") {
// 	  const parsedData: ParsedRecap = {
// 		...(data.R["CarData.z"] && {
// 		  CarData: parseCompressed(data.R["CarData.z"]),
// 		}),
// 		...(data.R["Position.z"] && {
// 		  Position: parseCompressed(data.R["Position.z"]),
// 		}),
// 	  };

// 	  const {
// 		"CarData.z": z1,
// 		"Position.z": z2,
// 		...newState
// 	  } = { ...data.R, ...parsedData };

// 	  return merge(state, newState) ?? state;
// 	}

// 	return state;
//   };
