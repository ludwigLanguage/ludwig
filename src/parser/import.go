package parser

import (
	"ludwig/src/ast"
	"ludwig/src/lexer"
	"ludwig/src/source"
	"ludwig/src/tokens"
	"path/filepath"
)

/* Syntax:
 * import <path>
 */

func (p *Parser) parseImport() ast.Node {
	p.lxr.MoveUp() //Move over 'import' token

	path := p.getImportPath()
	pkg := p.getPkgFrom(path)

	assignedPkg := p.assignPackageToGivenId(pkg)

	return assignedPkg
}

func (p *Parser) getImportPath() string {
	if p.lxr.CurTok.Alias != tokens.STR {
		p.raiseError("Syntax", "Expected string after 'import,' got: "+p.lxr.CurTok.Value)
	}

	path, err := filepath.Abs(p.lxr.CurTok.Value)
	if err != nil {
		p.raiseError("File", "Could not find file '"+p.lxr.CurTok.Value+"'")
	}

	p.lxr.MoveUp() //Move over filepath string

	return path
}

func (p *Parser) getPkgFrom(path string) ast.Node {
	src := source.New(path)
	lxr := lexer.New(src)
	prs := New(lxr)
	return prs.ParsePackage()
}

func (p *Parser) assignPackageToGivenId(pkg ast.Node) ast.Node {
	id := pkg.(*ast.Package).Id
	assignment := &ast.InfixExpr{id, pkg, "=", pkg.GetTok()}

	return assignment
}
