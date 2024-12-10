package assembler

type Parser struct {
	Tokens []Token
	Index  int
}

func (p Parser) Next() bool {
	return p.Index+1 < len(p.Tokens)
}

func (p *Parser) Token() Token {
	token := p.Tokens[p.Index]
	p.Index += 1
	return token
}

func (p Parser) Look(where int) *Token {
	index := p.Index + where
	if index < len(p.Tokens) {
		return &p.Tokens[index]
	}
	return nil
}

func (p *Parser) Consume(amount int) {
	p.Index += amount
}

func (p Parser) Test(where int, content string) bool {
	if token := p.Look(where); token != nil {
		return token.Token == content
	}
	return false
}

func (p *Parser) ParseLabel() AST {
	if p.Test(1, ":") && p.Test(2, "\n") {
		tokens := p.Tokens[p.Index : p.Index+2]
		p.Consume(3)

		children := []AST{}
		for p.Next() {
			children = append(children, p.ParseInstruction())
		}

		return AST{
			Type:     Label,
			Tokens:   tokens,
			Children: children,
		}
	}

	return AST{}
}

func (p *Parser) ParseInstruction() AST {
	token := p.Token()
	children := []AST{}

	for p.Next() {
		next := p.Token()
		if p.Test(0, ",") {
			p.Consume(1)
			children = append(children, AST{
				Type:     Operand,
				Tokens:   []Token{next},
				Children: []AST{},
			})
		} else if p.Test(0, "\n") {
			p.Consume(1)
			children = append(children, AST{
				Type:     Operand,
				Tokens:   []Token{next},
				Children: []AST{},
			})
			break
		}
	}

	return AST{
		Type:     Label,
		Tokens:   []Token{token},
		Children: children,
	}
}

func Parse(tokens []Token) AST {

	parser := Parser{
		Tokens: tokens,
		Index:  0,
	}

	tree := AST{
		Type:     Root,
		Children: []AST{},
	}

	for parser.Next() {
		tree.Children = append(tree.Children, parser.ParseLabel())
	}

	return tree

}
