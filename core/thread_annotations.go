package core

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/textileio/textile-go/pb"
	"github.com/textileio/textile-go/repo"
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
)

// AddComment adds an outgoing comment block
func (t *Thread) AddComment(dataId string, body string) (mh.Multihash, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	// build block
	return t.addAnnotation(&pb.ThreadAnnotation{
		Type:    pb.ThreadAnnotation_COMMENT,
		Data:    dataId,
		Caption: body,
	})
}

// AddLike adds an outgoing like block
func (t *Thread) AddLike(dataId string) (mh.Multihash, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	// build block
	return t.addAnnotation(&pb.ThreadAnnotation{
		Type: pb.ThreadAnnotation_LIKE,
		Data: dataId,
	})
}

// addAnnotation adds an outgoing comment or like block
func (t *Thread) addAnnotation(msg *pb.ThreadAnnotation) (mh.Multihash, error) {
	// commit to ipfs
	res, err := t.commitBlock(msg, pb.ThreadBlock_ANNOTATION, nil)
	if err != nil {
		return nil, err
	}

	// index it locally
	dconf := &repo.DataBlockConfig{
		DataId:      msg.Data,
		DataCaption: msg.Caption,
	}
	var btype repo.BlockType
	switch msg.Type {
	case pb.ThreadAnnotation_COMMENT:
		btype = repo.CommentBlock
	case pb.ThreadAnnotation_LIKE:
		btype = repo.LikeBlock
	default:
		return nil, errors.New("invalid annotation")
	}
	if err := t.indexBlock(res, btype, dconf); err != nil {
		return nil, err
	}

	// update head
	if err := t.updateHead(res.hash); err != nil {
		return nil, err
	}

	// post it
	if err := t.post(res, t.Peers()); err != nil {
		return nil, err
	}

	log.Debugf("added ANNOTATION to %s: %s", t.Id, res.hash.B58String())

	// all done
	return res.hash, nil
}

// handleAnnotationBlock handles an incoming data block
func (t *Thread) handleAnnotationBlock(hash mh.Multihash, block *pb.ThreadBlock) (*pb.ThreadAnnotation, error) {
	msg := new(pb.ThreadAnnotation)
	if err := ptypes.UnmarshalAny(block.Payload, msg); err != nil {
		return nil, err
	}

	// index it locally
	dconf := &repo.DataBlockConfig{
		DataId: msg.Data,
	}
	var atype repo.BlockType
	switch msg.Type {
	case pb.ThreadAnnotation_COMMENT:
		atype = repo.CommentBlock
		dconf.DataCaption = msg.Caption
	case pb.ThreadAnnotation_LIKE:
		atype = repo.LikeBlock
	}
	if err := t.indexBlock(&commitResult{hash: hash, header: block.Header}, atype, dconf); err != nil {
		return nil, err
	}
	return msg, nil
}
