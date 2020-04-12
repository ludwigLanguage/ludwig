package ast

/*
	Example:
		for = macro
		syntax: for <ident> in <node> <node>
		do: func(pool) do
			utilFn = func(iter) do
				if !(iter > len(quote(pool[1])) do
					block = quote(do nil end)
					body = 
						if pool[2].Type() == "<block>"
							pool[2].GetBody()
						else
							quote(do nil end).GetBody()

					block.SetBody([ quote( unquote(pool[0]) = unquote(pool[1][iter]) ) ] + body)
					unquote(block)
				end
			end
			utilFn()
		end

		@for i in [1, 2, 3] do
			println(i)
		end

	Notes:
		We can just parse the do function like a normal function
		The syntax must be parsed with a speacial procedure. It will
		generate a list outlining the syntax (ie: [for, identNode, in, node, node])
*/


// <node_type>
type CarrotNode struct {
	NodeType string
}

type MacroDef struct {
	Syntax []Node
	Function 
}

