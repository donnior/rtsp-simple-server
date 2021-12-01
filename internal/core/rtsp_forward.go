package core

import (
	"github.com/aler9/gortsplib"
)

type forwardRequest struct {
	source      string
	destination string
}

type forwarder struct {
	source      string
	destination string
	client      *gortsplib.Client
}

var forwarders = []*forwarder{}

func init() {
	// forwarders = append(forwarders, &forwarder{
	// 	"live.stream",
	// 	"rtsp://localhost:8554/live2.stream",
	// 	nil,
	// })
}

func regiserForward(source, destination string) {
	//should check it's already been registered
	forwarders = append(forwarders, &forwarder{
		source,
		destination,
		nil,
	})
}

func OnRTPForward(session *rtspSession, ctx *gortsplib.ServerHandlerOnPacketRTPCtx) {
	for _, f := range forwarders {
		if f.source == session.path.Name() {
			if f.client == nil {
				f.client = &gortsplib.Client{}
				err := f.client.StartPublishing(f.destination, session.stream.tracks())
				if err != nil {
					panic(err)
				}
			}
			f.client.WritePacketRTP(ctx.TrackID, ctx.Payload)
		}
	}
}
