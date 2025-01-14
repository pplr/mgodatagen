package generators

import (
	"errors"
	"fmt"
)

type stringFromPartGenerator struct {
	base
	parts []Generator
}

func newStringFromPartsGenerator(config *Config, base base, ci *CollInfo, buffer *DocBuffer) (Generator, error) {
	if len(config.Parts) == 0 {
		return nil, errors.New("'parts' can't be null or empty")
	}

	parts := make([]Generator, 0, len(config.Parts))
	for _, c := range config.Parts {

		// if thoses attributes are set, the parts becom a 'fromArray' generator
		// with pre-computed BSON values that we can't convert back as string
		//
		// TODO: find a way to make it work
		if c.Unique || c.MaxDistinctValue > 0 {
			return nil, errors.New("parts generator can't have 'unique' or 'maxDistinctValue' attributes")
		}

		g, err := ci.newGenerator(buffer, "", &c)
		if err != nil {
			return nil, fmt.Errorf("invalid parts generator: %v", err)
		}
		parts = append(parts, g)
	}

	return &stringFromPartGenerator{
		base:  base,
		parts: parts,
	}, nil
}

func (g *stringFromPartGenerator) EncodeValue() {

	sizePos := g.buffer.Len()
	g.buffer.Reserve()
	start := g.buffer.Len()

	for _, p := range g.parts {
		p.EncodeValueAsString()
	}
	g.buffer.WriteSingleByte(0)
	g.buffer.WriteAt(sizePos, int32Bytes(int32(g.buffer.Len()-start)))
}
